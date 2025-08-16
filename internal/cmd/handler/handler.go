package handler

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"sync"
	"time"

	"github.com/fsnotify/fsnotify"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"github.com/spencer-weaver/goweb/frontend/editor"
	"github.com/spencer-weaver/goweb/frontend/site"
	"github.com/spencer-weaver/goweb/internal/template/html"
	"github.com/yosssi/gohtml"
)

type BuildConfig struct {
	Minify               bool   `json:"minify"`
	Testing              bool   `json:"testing"`
	ConfigDir            string `json:"configDir"`
	OutputDir            string `json:"outputDir"`
	SiteDir              string `json:"siteDir"`
	TemplatesDir         string `json:"templatesDir"`
	SharedTemplateFile   string `json:"sharedTemplateFile"`
	ElementTemplatesFile string `json:"elementTemplatesFile"`
}

type TestConfig struct {
	ServeDir    string `json:"serveDir"`
	WatchDir    string `json:"watchDir"`
	TemplateDir string `json:"templateDir"`
	StylesDir   string `json:"stylesDir"`
	MarkdownDir string `json:"markdownDir"`
	ServePort   string `json:"servePort"`
}

type MarkdownRequest struct {
	Content string `json:"content" binding:"required"`
}

var BuildSettings *BuildConfig
var TestSettings *TestConfig
var HtmlSettings *html.HTMLConfig

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool { return true },
}

var clients = make(map[*websocket.Conn]bool)
var clientsMu sync.Mutex

func BuildHandler() error {

	if err := os.MkdirAll(BuildSettings.OutputDir, 0755); err != nil {
		return fmt.Errorf("failed to create output dir: %w", err)
	}

	for name, builder := range site.Pages {

		// inject live reload
		if BuildSettings.Testing {
			script := html.NewScript()
			script.InnerHTML = "(function() { var ws = new WebSocket(\"ws://\" + location.host + \"/live-reload\"); ws.onmessage = function(event) { if (event.data === \"reload\") { window.location.reload(); } }; ws.onclose = function() { console.log(\"live reload disconnected\"); }; })();"
			builder.AddChild(script)
		}

		// render document to string
		document := html.RenderComponent(builder)

		// format document
		formattedDocument := gohtml.Format(document)
		if BuildSettings.Minify {
			formattedDocument = html.MinifyHTML(formattedDocument)
		}

		filename := filepath.Join(BuildSettings.OutputDir, strings.ToLower(name)+".html")
		if err := os.WriteFile(filename, []byte(formattedDocument), 0644); err != nil {
			return fmt.Errorf("failed to write file %s: %w", filename, err)
		}
		fmt.Printf("built %s\n", filename)
	}
	return nil
}

func onChange() error {
	cmd := exec.Command("go", "run", "./", "build")

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err := cmd.Start()
	if err != nil {
		return err
	}

	go func() {
		err := cmd.Wait()
		if err != nil {
			log.Printf("build command error: %v", err)
		} else {
			log.Printf("build command finished successfully")
			broadcastReload()
		}
	}()
	return nil
}

func broadcastReload() {
	clientsMu.Lock()
	defer clientsMu.Unlock()

	for conn := range clients {
		err := conn.WriteMessage(websocket.TextMessage, []byte("reload"))
		if err != nil {
			fmt.Println("WebSocket error:", err)
			conn.Close()
			delete(clients, conn)
		}
	}
}

func TestHandler() error {

	router := gin.Default()

	// live reload handling
	fmt.Print("crating handler\n\n")
	router.GET("/live-reload", func(c *gin.Context) {
		conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
		if err != nil {
			fmt.Println("WebSocket upgrade error:", err)
			return
		}
		defer conn.Close()

		clientsMu.Lock()
		clients[conn] = true
		clientsMu.Unlock()

		for {
			_, _, err := conn.ReadMessage()
			if err != nil {
				break
			}
		}

		clientsMu.Lock()
		delete(clients, conn)
		clientsMu.Unlock()
	})
	router.GET("/markdown/*filepath", func(c *gin.Context) {
		filePath := c.Param("filepath")
		markdownDir := strings.TrimPrefix(TestSettings.MarkdownDir, "./")

		path := markdownDir + filePath
		if _, err := os.Stat(path); os.IsNotExist(err) {
			file, err := os.Create(path)
			if err != nil {
				fmt.Println("Error creating file:", err)
				return
			}
			file.Close()
		}

		// render document to string
		document := html.RenderComponent(editor.MarkdownEditor(path))

		// format document
		formattedDocument := gohtml.Format(document)

		c.Data(200, "text/html; charset=utf-8", []byte(formattedDocument))
	})
	// router.POST("/save-markdown/*filepath", func(c *gin.Context) {
	// 	filePath := c.Param("filepath")
	// 	var req MarkdownRequest

	// 	if err := c.ShouldBindJSON(&req); err != nil {
	// 		c.JSON(http.StatusBadRequest, gin.H{"error": "content field is required"})
	// 		return
	// 	}

	// 	// Save to a timestamped file (or customize filename logic)
	// 	filename := time.Now().Format("20060102_150405") + ".md"

	// 	if err := ioutil.WriteFile("markdown/"+filename, []byte(req.Content), 0644); err != nil {
	// 		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to write file"})
	// 		return
	// 	}

	// 	c.JSON(http.StatusOK, gin.H{"message": "File saved", "filename": filename})
	// })
	router.Static("/static", TestSettings.ServeDir)

	go func() {
		fmt.Printf("serving %s at http://%s\n", TestSettings.ServeDir, TestSettings.ServePort)
		if err := router.Run(TestSettings.ServePort); err != nil {
			fmt.Println("error starting server:", err)
			os.Exit(1)
		}
	}()

	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		fmt.Println("failed to create watcher:", err)
		os.Exit(1)
	}
	defer watcher.Close()

	err = filepath.Walk(TestSettings.WatchDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.IsDir() {
			return watcher.Add(path)
		}
		return nil
	})
	if err != nil {
		fmt.Println("failed to watch directory:", err)
		os.Exit(1)
	}
	err = filepath.Walk(TestSettings.TemplateDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.IsDir() {
			return watcher.Add(path)
		}
		return nil
	})
	if err != nil {
		fmt.Println("failed to watch directory:", err)
		os.Exit(1)
	}
	err = filepath.Walk(TestSettings.StylesDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.IsDir() {
			return watcher.Add(path)
		}
		return nil
	})
	if err != nil {
		fmt.Println("failed to watch directory:", err)
		os.Exit(1)
	}
	err = filepath.Walk(TestSettings.MarkdownDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.IsDir() {
			return watcher.Add(path)
		}
		return nil
	})
	if err != nil {
		fmt.Println("failed to watch directory:", err)
		os.Exit(1)
	}

	debounce := time.Now()

	for {
		select {
		case event, ok := <-watcher.Events:
			if !ok {
				return err
			}
			if time.Since(debounce) < 300*time.Millisecond {
				continue // debounce
			}
			if event.Op&(fsnotify.Write|fsnotify.Create|fsnotify.Remove) != 0 {
				fmt.Println("change detected:", event)
				debounce = time.Now()
				onChange()
			}

		case err, ok := <-watcher.Errors:
			if !ok {
				return err
			}
			fmt.Println("watcher error:", err)
		}
	}
}

func HtmlHandler(updateHtml, updateGo bool) error {

	// load html templates
	// this defines templates that can be loaded for each element included in the update
	// you can define what status of elements you would like to include in goweb.config
	if updateHtml {
		fmt.Printf("updating html\r\n")
		err := html.RenderHTMLElementTemplates(&HtmlSettings.Render)
		if err != nil {
			fmt.Printf("failed to render html element templates: %v", err)
			return err
		}
	}

	// load go data templates
	// this generates common node structures for each html element directly accessible in
	// your go code, like javascript document methods
	if updateGo {
		fmt.Printf("updating go\r\n")
		err := html.RenderGoDataTemplates(&HtmlSettings.GoConfig, HtmlSettings.Render.OutputFile)
		if err != nil {
			fmt.Printf("failed to render go data templates: %v", err)
			return err
		}
	}
	return nil
}
