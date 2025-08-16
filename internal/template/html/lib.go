package html

func NewRtc() *Component {
	return &Component{
		Node: Node{
			Name: "rtc",
		},
		Render: RenderSpec{
			Type:        "html",
			Template:    "rtc-html",
			SelfClosing: IsSelfClosing("rtc"),
		},
	}
}

func (component *Component) AddRtc() *Component {
	newRtc := NewRtc()
	return component.AddChild(newRtc)
}

func NewTd() *Component {
	return &Component{
		Node: Node{
			Name: "td",
		},
		Render: RenderSpec{
			Type:        "html",
			Template:    "td-html",
			SelfClosing: IsSelfClosing("td"),
		},
	}
}

func (component *Component) AddTd() *Component {
	newTd := NewTd()
	return component.AddChild(newTd)
}

func NewDatalist() *Component {
	return &Component{
		Node: Node{
			Name: "datalist",
		},
		Render: RenderSpec{
			Type:        "html",
			Template:    "datalist-html",
			SelfClosing: IsSelfClosing("datalist"),
		},
	}
}

func (component *Component) AddDatalist() *Component {
	newDatalist := NewDatalist()
	return component.AddChild(newDatalist)
}

func NewDir() *Component {
	return &Component{
		Node: Node{
			Name: "dir",
		},
		Render: RenderSpec{
			Type:        "html",
			Template:    "dir-html",
			SelfClosing: IsSelfClosing("dir"),
		},
	}
}

func (component *Component) AddDir() *Component {
	newDir := NewDir()
	return component.AddChild(newDir)
}

func NewHeader() *Component {
	return &Component{
		Node: Node{
			Name: "header",
		},
		Render: RenderSpec{
			Type:        "html",
			Template:    "header-html",
			SelfClosing: IsSelfClosing("header"),
		},
	}
}

func (component *Component) AddHeader() *Component {
	newHeader := NewHeader()
	return component.AddChild(newHeader)
}

func NewTime() *Component {
	return &Component{
		Node: Node{
			Name: "time",
		},
		Render: RenderSpec{
			Type:        "html",
			Template:    "time-html",
			SelfClosing: IsSelfClosing("time"),
		},
	}
}

func (component *Component) AddTime() *Component {
	newTime := NewTime()
	return component.AddChild(newTime)
}

func NewDt() *Component {
	return &Component{
		Node: Node{
			Name: "dt",
		},
		Render: RenderSpec{
			Type:        "html",
			Template:    "dt-html",
			SelfClosing: IsSelfClosing("dt"),
		},
	}
}

func (component *Component) AddDt() *Component {
	newDt := NewDt()
	return component.AddChild(newDt)
}

func NewU() *Component {
	return &Component{
		Node: Node{
			Name: "u",
		},
		Render: RenderSpec{
			Type:        "html",
			Template:    "u-html",
			SelfClosing: IsSelfClosing("u"),
		},
	}
}

func (component *Component) AddU() *Component {
	newU := NewU()
	return component.AddChild(newU)
}

func NewTemplate() *Component {
	return &Component{
		Node: Node{
			Name: "template",
		},
		Render: RenderSpec{
			Type:        "html",
			Template:    "template-html",
			SelfClosing: IsSelfClosing("template"),
		},
	}
}

func (component *Component) AddTemplate() *Component {
	newTemplate := NewTemplate()
	return component.AddChild(newTemplate)
}

func NewArea() *Component {
	return &Component{
		Node: Node{
			Name: "area",
		},
		Render: RenderSpec{
			Type:        "html",
			Template:    "area-html",
			SelfClosing: IsSelfClosing("area"),
		},
	}
}

func (component *Component) AddArea() *Component {
	newArea := NewArea()
	return component.AddChild(newArea)
}

func NewSection() *Component {
	return &Component{
		Node: Node{
			Name: "section",
		},
		Render: RenderSpec{
			Type:        "html",
			Template:    "section-html",
			SelfClosing: IsSelfClosing("section"),
		},
	}
}

func (component *Component) AddSection() *Component {
	newSection := NewSection()
	return component.AddChild(newSection)
}

func NewSup() *Component {
	return &Component{
		Node: Node{
			Name: "sup",
		},
		Render: RenderSpec{
			Type:        "html",
			Template:    "sup-html",
			SelfClosing: IsSelfClosing("sup"),
		},
	}
}

func (component *Component) AddSup() *Component {
	newSup := NewSup()
	return component.AddChild(newSup)
}

func NewLink() *Component {
	return &Component{
		Node: Node{
			Name: "link",
		},
		Render: RenderSpec{
			Type:        "html",
			Template:    "link-html",
			SelfClosing: IsSelfClosing("link"),
		},
	}
}

func (component *Component) AddLink() *Component {
	newLink := NewLink()
	return component.AddChild(newLink)
}

func NewStrong() *Component {
	return &Component{
		Node: Node{
			Name: "strong",
		},
		Render: RenderSpec{
			Type:        "html",
			Template:    "strong-html",
			SelfClosing: IsSelfClosing("strong"),
		},
	}
}

func (component *Component) AddStrong() *Component {
	newStrong := NewStrong()
	return component.AddChild(newStrong)
}

func NewIframe() *Component {
	return &Component{
		Node: Node{
			Name: "iframe",
		},
		Render: RenderSpec{
			Type:        "html",
			Template:    "iframe-html",
			SelfClosing: IsSelfClosing("iframe"),
		},
	}
}

func (component *Component) AddIframe() *Component {
	newIframe := NewIframe()
	return component.AddChild(newIframe)
}

func NewTitle() *Component {
	return &Component{
		Node: Node{
			Name: "title",
		},
		Render: RenderSpec{
			Type:        "html",
			Template:    "title-html",
			SelfClosing: IsSelfClosing("title"),
		},
	}
}

func (component *Component) AddTitle() *Component {
	newTitle := NewTitle()
	return component.AddChild(newTitle)
}

func NewPlaintext() *Component {
	return &Component{
		Node: Node{
			Name: "plaintext",
		},
		Render: RenderSpec{
			Type:        "html",
			Template:    "plaintext-html",
			SelfClosing: IsSelfClosing("plaintext"),
		},
	}
}

func (component *Component) AddPlaintext() *Component {
	newPlaintext := NewPlaintext()
	return component.AddChild(newPlaintext)
}

func NewTrack() *Component {
	return &Component{
		Node: Node{
			Name: "track",
		},
		Render: RenderSpec{
			Type:        "html",
			Template:    "track-html",
			SelfClosing: IsSelfClosing("track"),
		},
	}
}

func (component *Component) AddTrack() *Component {
	newTrack := NewTrack()
	return component.AddChild(newTrack)
}

func NewI() *Component {
	return &Component{
		Node: Node{
			Name: "i",
		},
		Render: RenderSpec{
			Type:        "html",
			Template:    "i-html",
			SelfClosing: IsSelfClosing("i"),
		},
	}
}

func (component *Component) AddI() *Component {
	newI := NewI()
	return component.AddChild(newI)
}

func NewMeter() *Component {
	return &Component{
		Node: Node{
			Name: "meter",
		},
		Render: RenderSpec{
			Type:        "html",
			Template:    "meter-html",
			SelfClosing: IsSelfClosing("meter"),
		},
	}
}

func (component *Component) AddMeter() *Component {
	newMeter := NewMeter()
	return component.AddChild(newMeter)
}

func NewH3() *Component {
	return &Component{
		Node: Node{
			Name: "h3",
		},
		Render: RenderSpec{
			Type:        "html",
			Template:    "h3-html",
			SelfClosing: IsSelfClosing("h3"),
		},
	}
}

func (component *Component) AddH3() *Component {
	newH3 := NewH3()
	return component.AddChild(newH3)
}

func NewRb() *Component {
	return &Component{
		Node: Node{
			Name: "rb",
		},
		Render: RenderSpec{
			Type:        "html",
			Template:    "rb-html",
			SelfClosing: IsSelfClosing("rb"),
		},
	}
}

func (component *Component) AddRb() *Component {
	newRb := NewRb()
	return component.AddChild(newRb)
}

func NewEmbed() *Component {
	return &Component{
		Node: Node{
			Name: "embed",
		},
		Render: RenderSpec{
			Type:        "html",
			Template:    "embed-html",
			SelfClosing: IsSelfClosing("embed"),
		},
	}
}

func (component *Component) AddEmbed() *Component {
	newEmbed := NewEmbed()
	return component.AddChild(newEmbed)
}

func NewTt() *Component {
	return &Component{
		Node: Node{
			Name: "tt",
		},
		Render: RenderSpec{
			Type:        "html",
			Template:    "tt-html",
			SelfClosing: IsSelfClosing("tt"),
		},
	}
}

func (component *Component) AddTt() *Component {
	newTt := NewTt()
	return component.AddChild(newTt)
}

func NewHtml() *Component {
	return &Component{
		Node: Node{
			Name: "html",
		},
		Render: RenderSpec{
			Type:        "html",
			Template:    "html-html",
			SelfClosing: IsSelfClosing("html"),
		},
	}
}

func (component *Component) AddHtml() *Component {
	newHtml := NewHtml()
	return component.AddChild(newHtml)
}

func NewFrameset() *Component {
	return &Component{
		Node: Node{
			Name: "frameset",
		},
		Render: RenderSpec{
			Type:        "html",
			Template:    "frameset-html",
			SelfClosing: IsSelfClosing("frameset"),
		},
	}
}

func (component *Component) AddFrameset() *Component {
	newFrameset := NewFrameset()
	return component.AddChild(newFrameset)
}

func NewThead() *Component {
	return &Component{
		Node: Node{
			Name: "thead",
		},
		Render: RenderSpec{
			Type:        "html",
			Template:    "thead-html",
			SelfClosing: IsSelfClosing("thead"),
		},
	}
}

func (component *Component) AddThead() *Component {
	newThead := NewThead()
	return component.AddChild(newThead)
}

func NewFont() *Component {
	return &Component{
		Node: Node{
			Name: "font",
		},
		Render: RenderSpec{
			Type:        "html",
			Template:    "font-html",
			SelfClosing: IsSelfClosing("font"),
		},
	}
}

func (component *Component) AddFont() *Component {
	newFont := NewFont()
	return component.AddChild(newFont)
}

func NewSamp() *Component {
	return &Component{
		Node: Node{
			Name: "samp",
		},
		Render: RenderSpec{
			Type:        "html",
			Template:    "samp-html",
			SelfClosing: IsSelfClosing("samp"),
		},
	}
}

func (component *Component) AddSamp() *Component {
	newSamp := NewSamp()
	return component.AddChild(newSamp)
}

func NewSource() *Component {
	return &Component{
		Node: Node{
			Name: "source",
		},
		Render: RenderSpec{
			Type:        "html",
			Template:    "source-html",
			SelfClosing: IsSelfClosing("source"),
		},
	}
}

func (component *Component) AddSource() *Component {
	newSource := NewSource()
	return component.AddChild(newSource)
}

func NewButton() *Component {
	return &Component{
		Node: Node{
			Name: "button",
		},
		Render: RenderSpec{
			Type:        "html",
			Template:    "button-html",
			SelfClosing: IsSelfClosing("button"),
		},
	}
}

func (component *Component) AddButton() *Component {
	newButton := NewButton()
	return component.AddChild(newButton)
}

func NewBase() *Component {
	return &Component{
		Node: Node{
			Name: "base",
		},
		Render: RenderSpec{
			Type:        "html",
			Template:    "base-html",
			SelfClosing: IsSelfClosing("base"),
		},
	}
}

func (component *Component) AddBase() *Component {
	newBase := NewBase()
	return component.AddChild(newBase)
}

func NewMark() *Component {
	return &Component{
		Node: Node{
			Name: "mark",
		},
		Render: RenderSpec{
			Type:        "html",
			Template:    "mark-html",
			SelfClosing: IsSelfClosing("mark"),
		},
	}
}

func (component *Component) AddMark() *Component {
	newMark := NewMark()
	return component.AddChild(newMark)
}

func NewTr() *Component {
	return &Component{
		Node: Node{
			Name: "tr",
		},
		Render: RenderSpec{
			Type:        "html",
			Template:    "tr-html",
			SelfClosing: IsSelfClosing("tr"),
		},
	}
}

func (component *Component) AddTr() *Component {
	newTr := NewTr()
	return component.AddChild(newTr)
}

func NewOptgroup() *Component {
	return &Component{
		Node: Node{
			Name: "optgroup",
		},
		Render: RenderSpec{
			Type:        "html",
			Template:    "optgroup-html",
			SelfClosing: IsSelfClosing("optgroup"),
		},
	}
}

func (component *Component) AddOptgroup() *Component {
	newOptgroup := NewOptgroup()
	return component.AddChild(newOptgroup)
}

func NewUl() *Component {
	return &Component{
		Node: Node{
			Name: "ul",
		},
		Render: RenderSpec{
			Type:        "html",
			Template:    "ul-html",
			SelfClosing: IsSelfClosing("ul"),
		},
	}
}

func (component *Component) AddUl() *Component {
	newUl := NewUl()
	return component.AddChild(newUl)
}

func NewArticle() *Component {
	return &Component{
		Node: Node{
			Name: "article",
		},
		Render: RenderSpec{
			Type:        "html",
			Template:    "article-html",
			SelfClosing: IsSelfClosing("article"),
		},
	}
}

func (component *Component) AddArticle() *Component {
	newArticle := NewArticle()
	return component.AddChild(newArticle)
}

func NewH1() *Component {
	return &Component{
		Node: Node{
			Name: "h1",
		},
		Render: RenderSpec{
			Type:        "html",
			Template:    "h1-html",
			SelfClosing: IsSelfClosing("h1"),
		},
	}
}

func (component *Component) AddH1() *Component {
	newH1 := NewH1()
	return component.AddChild(newH1)
}

func NewOutput() *Component {
	return &Component{
		Node: Node{
			Name: "output",
		},
		Render: RenderSpec{
			Type:        "html",
			Template:    "output-html",
			SelfClosing: IsSelfClosing("output"),
		},
	}
}

func (component *Component) AddOutput() *Component {
	newOutput := NewOutput()
	return component.AddChild(newOutput)
}

func NewScript() *Component {
	return &Component{
		Node: Node{
			Name: "script",
		},
		Render: RenderSpec{
			Type:        "html",
			Template:    "script-html",
			SelfClosing: IsSelfClosing("script"),
		},
	}
}

func (component *Component) AddScript() *Component {
	newScript := NewScript()
	return component.AddChild(newScript)
}

func NewOl() *Component {
	return &Component{
		Node: Node{
			Name: "ol",
		},
		Render: RenderSpec{
			Type:        "html",
			Template:    "ol-html",
			SelfClosing: IsSelfClosing("ol"),
		},
	}
}

func (component *Component) AddOl() *Component {
	newOl := NewOl()
	return component.AddChild(newOl)
}

func NewSearch() *Component {
	return &Component{
		Node: Node{
			Name: "search",
		},
		Render: RenderSpec{
			Type:        "html",
			Template:    "search-html",
			SelfClosing: IsSelfClosing("search"),
		},
	}
}

func (component *Component) AddSearch() *Component {
	newSearch := NewSearch()
	return component.AddChild(newSearch)
}

func NewSpan() *Component {
	return &Component{
		Node: Node{
			Name: "span",
		},
		Render: RenderSpec{
			Type:        "html",
			Template:    "span-html",
			SelfClosing: IsSelfClosing("span"),
		},
	}
}

func (component *Component) AddSpan() *Component {
	newSpan := NewSpan()
	return component.AddChild(newSpan)
}

func NewWbr() *Component {
	return &Component{
		Node: Node{
			Name: "wbr",
		},
		Render: RenderSpec{
			Type:        "html",
			Template:    "wbr-html",
			SelfClosing: IsSelfClosing("wbr"),
		},
	}
}

func (component *Component) AddWbr() *Component {
	newWbr := NewWbr()
	return component.AddChild(newWbr)
}

func NewFrame() *Component {
	return &Component{
		Node: Node{
			Name: "frame",
		},
		Render: RenderSpec{
			Type:        "html",
			Template:    "frame-html",
			SelfClosing: IsSelfClosing("frame"),
		},
	}
}

func (component *Component) AddFrame() *Component {
	newFrame := NewFrame()
	return component.AddChild(newFrame)
}

func NewDocument() *Component {
	return &Component{
		Node: Node{
			Name: "document",
		},
		Render: RenderSpec{
			Type:        "html",
			Template:    "document-html",
			SelfClosing: IsSelfClosing("document"),
		},
	}
}

func (component *Component) AddDocument() *Component {
	newDocument := NewDocument()
	return component.AddChild(newDocument)
}

func NewBdi() *Component {
	return &Component{
		Node: Node{
			Name: "bdi",
		},
		Render: RenderSpec{
			Type:        "html",
			Template:    "bdi-html",
			SelfClosing: IsSelfClosing("bdi"),
		},
	}
}

func (component *Component) AddBdi() *Component {
	newBdi := NewBdi()
	return component.AddChild(newBdi)
}

func NewInput() *Component {
	return &Component{
		Node: Node{
			Name: "input",
		},
		Render: RenderSpec{
			Type:        "html",
			Template:    "input-html",
			SelfClosing: IsSelfClosing("input"),
		},
	}
}

func (component *Component) AddInput() *Component {
	newInput := NewInput()
	return component.AddChild(newInput)
}

func NewNoembed() *Component {
	return &Component{
		Node: Node{
			Name: "noembed",
		},
		Render: RenderSpec{
			Type:        "html",
			Template:    "noembed-html",
			SelfClosing: IsSelfClosing("noembed"),
		},
	}
}

func (component *Component) AddNoembed() *Component {
	newNoembed := NewNoembed()
	return component.AddChild(newNoembed)
}

func NewH2() *Component {
	return &Component{
		Node: Node{
			Name: "h2",
		},
		Render: RenderSpec{
			Type:        "html",
			Template:    "h2-html",
			SelfClosing: IsSelfClosing("h2"),
		},
	}
}

func (component *Component) AddH2() *Component {
	newH2 := NewH2()
	return component.AddChild(newH2)
}

func NewSelect() *Component {
	return &Component{
		Node: Node{
			Name: "select",
		},
		Render: RenderSpec{
			Type:        "html",
			Template:    "select-html",
			SelfClosing: IsSelfClosing("select"),
		},
	}
}

func (component *Component) AddSelect() *Component {
	newSelect := NewSelect()
	return component.AddChild(newSelect)
}

func NewParam() *Component {
	return &Component{
		Node: Node{
			Name: "param",
		},
		Render: RenderSpec{
			Type:        "html",
			Template:    "param-html",
			SelfClosing: IsSelfClosing("param"),
		},
	}
}

func (component *Component) AddParam() *Component {
	newParam := NewParam()
	return component.AddChild(newParam)
}

func NewAcronym() *Component {
	return &Component{
		Node: Node{
			Name: "acronym",
		},
		Render: RenderSpec{
			Type:        "html",
			Template:    "acronym-html",
			SelfClosing: IsSelfClosing("acronym"),
		},
	}
}

func (component *Component) AddAcronym() *Component {
	newAcronym := NewAcronym()
	return component.AddChild(newAcronym)
}

func NewH4() *Component {
	return &Component{
		Node: Node{
			Name: "h4",
		},
		Render: RenderSpec{
			Type:        "html",
			Template:    "h4-html",
			SelfClosing: IsSelfClosing("h4"),
		},
	}
}

func (component *Component) AddH4() *Component {
	newH4 := NewH4()
	return component.AddChild(newH4)
}

func NewTh() *Component {
	return &Component{
		Node: Node{
			Name: "th",
		},
		Render: RenderSpec{
			Type:        "html",
			Template:    "th-html",
			SelfClosing: IsSelfClosing("th"),
		},
	}
}

func (component *Component) AddTh() *Component {
	newTh := NewTh()
	return component.AddChild(newTh)
}

func NewTextarea() *Component {
	return &Component{
		Node: Node{
			Name: "textarea",
		},
		Render: RenderSpec{
			Type:        "html",
			Template:    "textarea-html",
			SelfClosing: IsSelfClosing("textarea"),
		},
	}
}

func (component *Component) AddTextarea() *Component {
	newTextarea := NewTextarea()
	return component.AddChild(newTextarea)
}

func NewTbody() *Component {
	return &Component{
		Node: Node{
			Name: "tbody",
		},
		Render: RenderSpec{
			Type:        "html",
			Template:    "tbody-html",
			SelfClosing: IsSelfClosing("tbody"),
		},
	}
}

func (component *Component) AddTbody() *Component {
	newTbody := NewTbody()
	return component.AddChild(newTbody)
}

func NewData() *Component {
	return &Component{
		Node: Node{
			Name: "data",
		},
		Render: RenderSpec{
			Type:        "html",
			Template:    "data-html",
			SelfClosing: IsSelfClosing("data"),
		},
	}
}

func (component *Component) AddData() *Component {
	newData := NewData()
	return component.AddChild(newData)
}

func NewBdo() *Component {
	return &Component{
		Node: Node{
			Name: "bdo",
		},
		Render: RenderSpec{
			Type:        "html",
			Template:    "bdo-html",
			SelfClosing: IsSelfClosing("bdo"),
		},
	}
}

func (component *Component) AddBdo() *Component {
	newBdo := NewBdo()
	return component.AddChild(newBdo)
}

func NewMain() *Component {
	return &Component{
		Node: Node{
			Name: "main",
		},
		Render: RenderSpec{
			Type:        "html",
			Template:    "main-html",
			SelfClosing: IsSelfClosing("main"),
		},
	}
}

func (component *Component) AddMain() *Component {
	newMain := NewMain()
	return component.AddChild(newMain)
}

func NewAside() *Component {
	return &Component{
		Node: Node{
			Name: "aside",
		},
		Render: RenderSpec{
			Type:        "html",
			Template:    "aside-html",
			SelfClosing: IsSelfClosing("aside"),
		},
	}
}

func (component *Component) AddAside() *Component {
	newAside := NewAside()
	return component.AddChild(newAside)
}

func NewHead() *Component {
	return &Component{
		Node: Node{
			Name: "head",
		},
		Render: RenderSpec{
			Type:        "html",
			Template:    "head-html",
			SelfClosing: IsSelfClosing("head"),
		},
	}
}

func (component *Component) AddHead() *Component {
	newHead := NewHead()
	return component.AddChild(newHead)
}

func NewPicture() *Component {
	return &Component{
		Node: Node{
			Name: "picture",
		},
		Render: RenderSpec{
			Type:        "html",
			Template:    "picture-html",
			SelfClosing: IsSelfClosing("picture"),
		},
	}
}

func (component *Component) AddPicture() *Component {
	newPicture := NewPicture()
	return component.AddChild(newPicture)
}

func NewFooter() *Component {
	return &Component{
		Node: Node{
			Name: "footer",
		},
		Render: RenderSpec{
			Type:        "html",
			Template:    "footer-html",
			SelfClosing: IsSelfClosing("footer"),
		},
	}
}

func (component *Component) AddFooter() *Component {
	newFooter := NewFooter()
	return component.AddChild(newFooter)
}

func NewBlockquote() *Component {
	return &Component{
		Node: Node{
			Name: "blockquote",
		},
		Render: RenderSpec{
			Type:        "html",
			Template:    "blockquote-html",
			SelfClosing: IsSelfClosing("blockquote"),
		},
	}
}

func (component *Component) AddBlockquote() *Component {
	newBlockquote := NewBlockquote()
	return component.AddChild(newBlockquote)
}

func NewKbd() *Component {
	return &Component{
		Node: Node{
			Name: "kbd",
		},
		Render: RenderSpec{
			Type:        "html",
			Template:    "kbd-html",
			SelfClosing: IsSelfClosing("kbd"),
		},
	}
}

func (component *Component) AddKbd() *Component {
	newKbd := NewKbd()
	return component.AddChild(newKbd)
}

func NewH6() *Component {
	return &Component{
		Node: Node{
			Name: "h6",
		},
		Render: RenderSpec{
			Type:        "html",
			Template:    "h6-html",
			SelfClosing: IsSelfClosing("h6"),
		},
	}
}

func (component *Component) AddH6() *Component {
	newH6 := NewH6()
	return component.AddChild(newH6)
}

func NewDel() *Component {
	return &Component{
		Node: Node{
			Name: "del",
		},
		Render: RenderSpec{
			Type:        "html",
			Template:    "del-html",
			SelfClosing: IsSelfClosing("del"),
		},
	}
}

func (component *Component) AddDel() *Component {
	newDel := NewDel()
	return component.AddChild(newDel)
}

func NewA() *Component {
	return &Component{
		Node: Node{
			Name: "a",
		},
		Render: RenderSpec{
			Type:        "html",
			Template:    "a-html",
			SelfClosing: IsSelfClosing("a"),
		},
	}
}

func (component *Component) AddA() *Component {
	newA := NewA()
	return component.AddChild(newA)
}

func NewTfoot() *Component {
	return &Component{
		Node: Node{
			Name: "tfoot",
		},
		Render: RenderSpec{
			Type:        "html",
			Template:    "tfoot-html",
			SelfClosing: IsSelfClosing("tfoot"),
		},
	}
}

func (component *Component) AddTfoot() *Component {
	newTfoot := NewTfoot()
	return component.AddChild(newTfoot)
}

func NewStyle() *Component {
	return &Component{
		Node: Node{
			Name: "style",
		},
		Render: RenderSpec{
			Type:        "html",
			Template:    "style-html",
			SelfClosing: IsSelfClosing("style"),
		},
	}
}

func (component *Component) AddStyle() *Component {
	newStyle := NewStyle()
	return component.AddChild(newStyle)
}

func NewDiv() *Component {
	return &Component{
		Node: Node{
			Name: "div",
		},
		Render: RenderSpec{
			Type:        "html",
			Template:    "div-html",
			SelfClosing: IsSelfClosing("div"),
		},
	}
}

func (component *Component) AddDiv() *Component {
	newDiv := NewDiv()
	return component.AddChild(newDiv)
}

func NewIns() *Component {
	return &Component{
		Node: Node{
			Name: "ins",
		},
		Render: RenderSpec{
			Type:        "html",
			Template:    "ins-html",
			SelfClosing: IsSelfClosing("ins"),
		},
	}
}

func (component *Component) AddIns() *Component {
	newIns := NewIns()
	return component.AddChild(newIns)
}

func NewDfn() *Component {
	return &Component{
		Node: Node{
			Name: "dfn",
		},
		Render: RenderSpec{
			Type:        "html",
			Template:    "dfn-html",
			SelfClosing: IsSelfClosing("dfn"),
		},
	}
}

func (component *Component) AddDfn() *Component {
	newDfn := NewDfn()
	return component.AddChild(newDfn)
}

func NewMeta() *Component {
	return &Component{
		Node: Node{
			Name: "meta",
		},
		Render: RenderSpec{
			Type:        "html",
			Template:    "meta-html",
			SelfClosing: IsSelfClosing("meta"),
		},
	}
}

func (component *Component) AddMeta() *Component {
	newMeta := NewMeta()
	return component.AddChild(newMeta)
}

func NewProgress() *Component {
	return &Component{
		Node: Node{
			Name: "progress",
		},
		Render: RenderSpec{
			Type:        "html",
			Template:    "progress-html",
			SelfClosing: IsSelfClosing("progress"),
		},
	}
}

func (component *Component) AddProgress() *Component {
	newProgress := NewProgress()
	return component.AddChild(newProgress)
}

func NewBody() *Component {
	return &Component{
		Node: Node{
			Name: "body",
		},
		Render: RenderSpec{
			Type:        "html",
			Template:    "body-html",
			SelfClosing: IsSelfClosing("body"),
		},
	}
}

func (component *Component) AddBody() *Component {
	newBody := NewBody()
	return component.AddChild(newBody)
}

func NewDoctype() *Component {
	return &Component{
		Node: Node{
			Name: "doctype",
		},
		Render: RenderSpec{
			Type:        "html",
			Template:    "doctype-html",
			SelfClosing: IsSelfClosing("doctype"),
		},
	}
}

func (component *Component) AddDoctype() *Component {
	newDoctype := NewDoctype()
	return component.AddChild(newDoctype)
}

func NewNoframes() *Component {
	return &Component{
		Node: Node{
			Name: "noframes",
		},
		Render: RenderSpec{
			Type:        "html",
			Template:    "noframes-html",
			SelfClosing: IsSelfClosing("noframes"),
		},
	}
}

func (component *Component) AddNoframes() *Component {
	newNoframes := NewNoframes()
	return component.AddChild(newNoframes)
}

func NewDd() *Component {
	return &Component{
		Node: Node{
			Name: "dd",
		},
		Render: RenderSpec{
			Type:        "html",
			Template:    "dd-html",
			SelfClosing: IsSelfClosing("dd"),
		},
	}
}

func (component *Component) AddDd() *Component {
	newDd := NewDd()
	return component.AddChild(newDd)
}

func NewDetails() *Component {
	return &Component{
		Node: Node{
			Name: "details",
		},
		Render: RenderSpec{
			Type:        "html",
			Template:    "details-html",
			SelfClosing: IsSelfClosing("details"),
		},
	}
}

func (component *Component) AddDetails() *Component {
	newDetails := NewDetails()
	return component.AddChild(newDetails)
}

func NewH5() *Component {
	return &Component{
		Node: Node{
			Name: "h5",
		},
		Render: RenderSpec{
			Type:        "html",
			Template:    "h5-html",
			SelfClosing: IsSelfClosing("h5"),
		},
	}
}

func (component *Component) AddH5() *Component {
	newH5 := NewH5()
	return component.AddChild(newH5)
}

func NewSummary() *Component {
	return &Component{
		Node: Node{
			Name: "summary",
		},
		Render: RenderSpec{
			Type:        "html",
			Template:    "summary-html",
			SelfClosing: IsSelfClosing("summary"),
		},
	}
}

func (component *Component) AddSummary() *Component {
	newSummary := NewSummary()
	return component.AddChild(newSummary)
}

func NewFieldset() *Component {
	return &Component{
		Node: Node{
			Name: "fieldset",
		},
		Render: RenderSpec{
			Type:        "html",
			Template:    "fieldset-html",
			SelfClosing: IsSelfClosing("fieldset"),
		},
	}
}

func (component *Component) AddFieldset() *Component {
	newFieldset := NewFieldset()
	return component.AddChild(newFieldset)
}

func NewHgroup() *Component {
	return &Component{
		Node: Node{
			Name: "hgroup",
		},
		Render: RenderSpec{
			Type:        "html",
			Template:    "hgroup-html",
			SelfClosing: IsSelfClosing("hgroup"),
		},
	}
}

func (component *Component) AddHgroup() *Component {
	newHgroup := NewHgroup()
	return component.AddChild(newHgroup)
}

func NewDl() *Component {
	return &Component{
		Node: Node{
			Name: "dl",
		},
		Render: RenderSpec{
			Type:        "html",
			Template:    "dl-html",
			SelfClosing: IsSelfClosing("dl"),
		},
	}
}

func (component *Component) AddDl() *Component {
	newDl := NewDl()
	return component.AddChild(newDl)
}

func NewXmp() *Component {
	return &Component{
		Node: Node{
			Name: "xmp",
		},
		Render: RenderSpec{
			Type:        "html",
			Template:    "xmp-html",
			SelfClosing: IsSelfClosing("xmp"),
		},
	}
}

func (component *Component) AddXmp() *Component {
	newXmp := NewXmp()
	return component.AddChild(newXmp)
}

func NewSvg() *Component {
	return &Component{
		Node: Node{
			Name: "svg",
		},
		Render: RenderSpec{
			Type:        "html",
			Template:    "svg-html",
			SelfClosing: IsSelfClosing("svg"),
		},
	}
}

func (component *Component) AddSvg() *Component {
	newSvg := NewSvg()
	return component.AddChild(newSvg)
}

func NewLabel() *Component {
	return &Component{
		Node: Node{
			Name: "label",
		},
		Render: RenderSpec{
			Type:        "html",
			Template:    "label-html",
			SelfClosing: IsSelfClosing("label"),
		},
	}
}

func (component *Component) AddLabel() *Component {
	newLabel := NewLabel()
	return component.AddChild(newLabel)
}

func NewPath() *Component {
	return &Component{
		Node: Node{
			Name: "path",
		},
		Render: RenderSpec{
			Type:        "html",
			Template:    "path-html",
			SelfClosing: IsSelfClosing("path"),
		},
	}
}

func (component *Component) AddPath() *Component {
	newPath := NewPath()
	return component.AddChild(newPath)
}

func NewCenter() *Component {
	return &Component{
		Node: Node{
			Name: "center",
		},
		Render: RenderSpec{
			Type:        "html",
			Template:    "center-html",
			SelfClosing: IsSelfClosing("center"),
		},
	}
}

func (component *Component) AddCenter() *Component {
	newCenter := NewCenter()
	return component.AddChild(newCenter)
}

func NewCol() *Component {
	return &Component{
		Node: Node{
			Name: "col",
		},
		Render: RenderSpec{
			Type:        "html",
			Template:    "col-html",
			SelfClosing: IsSelfClosing("col"),
		},
	}
}

func (component *Component) AddCol() *Component {
	newCol := NewCol()
	return component.AddChild(newCol)
}

func NewAbbr() *Component {
	return &Component{
		Node: Node{
			Name: "abbr",
		},
		Render: RenderSpec{
			Type:        "html",
			Template:    "abbr-html",
			SelfClosing: IsSelfClosing("abbr"),
		},
	}
}

func (component *Component) AddAbbr() *Component {
	newAbbr := NewAbbr()
	return component.AddChild(newAbbr)
}

func NewHr() *Component {
	return &Component{
		Node: Node{
			Name: "hr",
		},
		Render: RenderSpec{
			Type:        "html",
			Template:    "hr-html",
			SelfClosing: IsSelfClosing("hr"),
		},
	}
}

func (component *Component) AddHr() *Component {
	newHr := NewHr()
	return component.AddChild(newHr)
}

func NewPre() *Component {
	return &Component{
		Node: Node{
			Name: "pre",
		},
		Render: RenderSpec{
			Type:        "html",
			Template:    "pre-html",
			SelfClosing: IsSelfClosing("pre"),
		},
	}
}

func (component *Component) AddPre() *Component {
	newPre := NewPre()
	return component.AddChild(newPre)
}

func NewStrike() *Component {
	return &Component{
		Node: Node{
			Name: "strike",
		},
		Render: RenderSpec{
			Type:        "html",
			Template:    "strike-html",
			SelfClosing: IsSelfClosing("strike"),
		},
	}
}

func (component *Component) AddStrike() *Component {
	newStrike := NewStrike()
	return component.AddChild(newStrike)
}

func NewFencedframe() *Component {
	return &Component{
		Node: Node{
			Name: "fencedframe",
		},
		Render: RenderSpec{
			Type:        "html",
			Template:    "fencedframe-html",
			SelfClosing: IsSelfClosing("fencedframe"),
		},
	}
}

func (component *Component) AddFencedframe() *Component {
	newFencedframe := NewFencedframe()
	return component.AddChild(newFencedframe)
}

func NewCode() *Component {
	return &Component{
		Node: Node{
			Name: "code",
		},
		Render: RenderSpec{
			Type:        "html",
			Template:    "code-html",
			SelfClosing: IsSelfClosing("code"),
		},
	}
}

func (component *Component) AddCode() *Component {
	newCode := NewCode()
	return component.AddChild(newCode)
}

func NewLegend() *Component {
	return &Component{
		Node: Node{
			Name: "legend",
		},
		Render: RenderSpec{
			Type:        "html",
			Template:    "legend-html",
			SelfClosing: IsSelfClosing("legend"),
		},
	}
}

func (component *Component) AddLegend() *Component {
	newLegend := NewLegend()
	return component.AddChild(newLegend)
}

func NewS() *Component {
	return &Component{
		Node: Node{
			Name: "s",
		},
		Render: RenderSpec{
			Type:        "html",
			Template:    "s-html",
			SelfClosing: IsSelfClosing("s"),
		},
	}
}

func (component *Component) AddS() *Component {
	newS := NewS()
	return component.AddChild(newS)
}

func NewDialog() *Component {
	return &Component{
		Node: Node{
			Name: "dialog",
		},
		Render: RenderSpec{
			Type:        "html",
			Template:    "dialog-html",
			SelfClosing: IsSelfClosing("dialog"),
		},
	}
}

func (component *Component) AddDialog() *Component {
	newDialog := NewDialog()
	return component.AddChild(newDialog)
}

func NewCanvas() *Component {
	return &Component{
		Node: Node{
			Name: "canvas",
		},
		Render: RenderSpec{
			Type:        "html",
			Template:    "canvas-html",
			SelfClosing: IsSelfClosing("canvas"),
		},
	}
}

func (component *Component) AddCanvas() *Component {
	newCanvas := NewCanvas()
	return component.AddChild(newCanvas)
}

func NewObject() *Component {
	return &Component{
		Node: Node{
			Name: "object",
		},
		Render: RenderSpec{
			Type:        "html",
			Template:    "object-html",
			SelfClosing: IsSelfClosing("object"),
		},
	}
}

func (component *Component) AddObject() *Component {
	newObject := NewObject()
	return component.AddChild(newObject)
}

func NewBig() *Component {
	return &Component{
		Node: Node{
			Name: "big",
		},
		Render: RenderSpec{
			Type:        "html",
			Template:    "big-html",
			SelfClosing: IsSelfClosing("big"),
		},
	}
}

func (component *Component) AddBig() *Component {
	newBig := NewBig()
	return component.AddChild(newBig)
}

func NewNav() *Component {
	return &Component{
		Node: Node{
			Name: "nav",
		},
		Render: RenderSpec{
			Type:        "html",
			Template:    "nav-html",
			SelfClosing: IsSelfClosing("nav"),
		},
	}
}

func (component *Component) AddNav() *Component {
	newNav := NewNav()
	return component.AddChild(newNav)
}

func NewAddress() *Component {
	return &Component{
		Node: Node{
			Name: "address",
		},
		Render: RenderSpec{
			Type:        "html",
			Template:    "address-html",
			SelfClosing: IsSelfClosing("address"),
		},
	}
}

func (component *Component) AddAddress() *Component {
	newAddress := NewAddress()
	return component.AddChild(newAddress)
}

func NewLi() *Component {
	return &Component{
		Node: Node{
			Name: "li",
		},
		Render: RenderSpec{
			Type:        "html",
			Template:    "li-html",
			SelfClosing: IsSelfClosing("li"),
		},
	}
}

func (component *Component) AddLi() *Component {
	newLi := NewLi()
	return component.AddChild(newLi)
}

func NewRp() *Component {
	return &Component{
		Node: Node{
			Name: "rp",
		},
		Render: RenderSpec{
			Type:        "html",
			Template:    "rp-html",
			SelfClosing: IsSelfClosing("rp"),
		},
	}
}

func (component *Component) AddRp() *Component {
	newRp := NewRp()
	return component.AddChild(newRp)
}

func NewImg() *Component {
	return &Component{
		Node: Node{
			Name: "img",
		},
		Render: RenderSpec{
			Type:        "html",
			Template:    "img-html",
			SelfClosing: IsSelfClosing("img"),
		},
	}
}

func (component *Component) AddImg() *Component {
	newImg := NewImg()
	return component.AddChild(newImg)
}

func NewForm() *Component {
	return &Component{
		Node: Node{
			Name: "form",
		},
		Render: RenderSpec{
			Type:        "html",
			Template:    "form-html",
			SelfClosing: IsSelfClosing("form"),
		},
	}
}

func (component *Component) AddForm() *Component {
	newForm := NewForm()
	return component.AddChild(newForm)
}

func NewSlot() *Component {
	return &Component{
		Node: Node{
			Name: "slot",
		},
		Render: RenderSpec{
			Type:        "html",
			Template:    "slot-html",
			SelfClosing: IsSelfClosing("slot"),
		},
	}
}

func (component *Component) AddSlot() *Component {
	newSlot := NewSlot()
	return component.AddChild(newSlot)
}

func NewRuby() *Component {
	return &Component{
		Node: Node{
			Name: "ruby",
		},
		Render: RenderSpec{
			Type:        "html",
			Template:    "ruby-html",
			SelfClosing: IsSelfClosing("ruby"),
		},
	}
}

func (component *Component) AddRuby() *Component {
	newRuby := NewRuby()
	return component.AddChild(newRuby)
}

func NewNobr() *Component {
	return &Component{
		Node: Node{
			Name: "nobr",
		},
		Render: RenderSpec{
			Type:        "html",
			Template:    "nobr-html",
			SelfClosing: IsSelfClosing("nobr"),
		},
	}
}

func (component *Component) AddNobr() *Component {
	newNobr := NewNobr()
	return component.AddChild(newNobr)
}

func NewOption() *Component {
	return &Component{
		Node: Node{
			Name: "option",
		},
		Render: RenderSpec{
			Type:        "html",
			Template:    "option-html",
			SelfClosing: IsSelfClosing("option"),
		},
	}
}

func (component *Component) AddOption() *Component {
	newOption := NewOption()
	return component.AddChild(newOption)
}

func NewRt() *Component {
	return &Component{
		Node: Node{
			Name: "rt",
		},
		Render: RenderSpec{
			Type:        "html",
			Template:    "rt-html",
			SelfClosing: IsSelfClosing("rt"),
		},
	}
}

func (component *Component) AddRt() *Component {
	newRt := NewRt()
	return component.AddChild(newRt)
}

func NewQ() *Component {
	return &Component{
		Node: Node{
			Name: "q",
		},
		Render: RenderSpec{
			Type:        "html",
			Template:    "q-html",
			SelfClosing: IsSelfClosing("q"),
		},
	}
}

func (component *Component) AddQ() *Component {
	newQ := NewQ()
	return component.AddChild(newQ)
}

func NewSub() *Component {
	return &Component{
		Node: Node{
			Name: "sub",
		},
		Render: RenderSpec{
			Type:        "html",
			Template:    "sub-html",
			SelfClosing: IsSelfClosing("sub"),
		},
	}
}

func (component *Component) AddSub() *Component {
	newSub := NewSub()
	return component.AddChild(newSub)
}

func NewSmall() *Component {
	return &Component{
		Node: Node{
			Name: "small",
		},
		Render: RenderSpec{
			Type:        "html",
			Template:    "small-html",
			SelfClosing: IsSelfClosing("small"),
		},
	}
}

func (component *Component) AddSmall() *Component {
	newSmall := NewSmall()
	return component.AddChild(newSmall)
}

func NewColgroup() *Component {
	return &Component{
		Node: Node{
			Name: "colgroup",
		},
		Render: RenderSpec{
			Type:        "html",
			Template:    "colgroup-html",
			SelfClosing: IsSelfClosing("colgroup"),
		},
	}
}

func (component *Component) AddColgroup() *Component {
	newColgroup := NewColgroup()
	return component.AddChild(newColgroup)
}

func NewMap() *Component {
	return &Component{
		Node: Node{
			Name: "map",
		},
		Render: RenderSpec{
			Type:        "html",
			Template:    "map-html",
			SelfClosing: IsSelfClosing("map"),
		},
	}
}

func (component *Component) AddMap() *Component {
	newMap := NewMap()
	return component.AddChild(newMap)
}

func NewFigure() *Component {
	return &Component{
		Node: Node{
			Name: "figure",
		},
		Render: RenderSpec{
			Type:        "html",
			Template:    "figure-html",
			SelfClosing: IsSelfClosing("figure"),
		},
	}
}

func (component *Component) AddFigure() *Component {
	newFigure := NewFigure()
	return component.AddChild(newFigure)
}

func NewCite() *Component {
	return &Component{
		Node: Node{
			Name: "cite",
		},
		Render: RenderSpec{
			Type:        "html",
			Template:    "cite-html",
			SelfClosing: IsSelfClosing("cite"),
		},
	}
}

func (component *Component) AddCite() *Component {
	newCite := NewCite()
	return component.AddChild(newCite)
}

func NewCaption() *Component {
	return &Component{
		Node: Node{
			Name: "caption",
		},
		Render: RenderSpec{
			Type:        "html",
			Template:    "caption-html",
			SelfClosing: IsSelfClosing("caption"),
		},
	}
}

func (component *Component) AddCaption() *Component {
	newCaption := NewCaption()
	return component.AddChild(newCaption)
}

func NewVideo() *Component {
	return &Component{
		Node: Node{
			Name: "video",
		},
		Render: RenderSpec{
			Type:        "html",
			Template:    "video-html",
			SelfClosing: IsSelfClosing("video"),
		},
	}
}

func (component *Component) AddVideo() *Component {
	newVideo := NewVideo()
	return component.AddChild(newVideo)
}

func NewVar() *Component {
	return &Component{
		Node: Node{
			Name: "var",
		},
		Render: RenderSpec{
			Type:        "html",
			Template:    "var-html",
			SelfClosing: IsSelfClosing("var"),
		},
	}
}

func (component *Component) AddVar() *Component {
	newVar := NewVar()
	return component.AddChild(newVar)
}

func NewMarquee() *Component {
	return &Component{
		Node: Node{
			Name: "marquee",
		},
		Render: RenderSpec{
			Type:        "html",
			Template:    "marquee-html",
			SelfClosing: IsSelfClosing("marquee"),
		},
	}
}

func (component *Component) AddMarquee() *Component {
	newMarquee := NewMarquee()
	return component.AddChild(newMarquee)
}

func NewNoscript() *Component {
	return &Component{
		Node: Node{
			Name: "noscript",
		},
		Render: RenderSpec{
			Type:        "html",
			Template:    "noscript-html",
			SelfClosing: IsSelfClosing("noscript"),
		},
	}
}

func (component *Component) AddNoscript() *Component {
	newNoscript := NewNoscript()
	return component.AddChild(newNoscript)
}

func NewTable() *Component {
	return &Component{
		Node: Node{
			Name: "table",
		},
		Render: RenderSpec{
			Type:        "html",
			Template:    "table-html",
			SelfClosing: IsSelfClosing("table"),
		},
	}
}

func (component *Component) AddTable() *Component {
	newTable := NewTable()
	return component.AddChild(newTable)
}

func NewBr() *Component {
	return &Component{
		Node: Node{
			Name: "br",
		},
		Render: RenderSpec{
			Type:        "html",
			Template:    "br-html",
			SelfClosing: IsSelfClosing("br"),
		},
	}
}

func (component *Component) AddBr() *Component {
	newBr := NewBr()
	return component.AddChild(newBr)
}

func NewP() *Component {
	return &Component{
		Node: Node{
			Name: "p",
		},
		Render: RenderSpec{
			Type:        "html",
			Template:    "p-html",
			SelfClosing: IsSelfClosing("p"),
		},
	}
}

func (component *Component) AddP() *Component {
	newP := NewP()
	return component.AddChild(newP)
}

func NewEm() *Component {
	return &Component{
		Node: Node{
			Name: "em",
		},
		Render: RenderSpec{
			Type:        "html",
			Template:    "em-html",
			SelfClosing: IsSelfClosing("em"),
		},
	}
}

func (component *Component) AddEm() *Component {
	newEm := NewEm()
	return component.AddChild(newEm)
}

func NewFigcaption() *Component {
	return &Component{
		Node: Node{
			Name: "figcaption",
		},
		Render: RenderSpec{
			Type:        "html",
			Template:    "figcaption-html",
			SelfClosing: IsSelfClosing("figcaption"),
		},
	}
}

func (component *Component) AddFigcaption() *Component {
	newFigcaption := NewFigcaption()
	return component.AddChild(newFigcaption)
}

func NewAudio() *Component {
	return &Component{
		Node: Node{
			Name: "audio",
		},
		Render: RenderSpec{
			Type:        "html",
			Template:    "audio-html",
			SelfClosing: IsSelfClosing("audio"),
		},
	}
}

func (component *Component) AddAudio() *Component {
	newAudio := NewAudio()
	return component.AddChild(newAudio)
}

func NewB() *Component {
	return &Component{
		Node: Node{
			Name: "b",
		},
		Render: RenderSpec{
			Type:        "html",
			Template:    "b-html",
			SelfClosing: IsSelfClosing("b"),
		},
	}
}

func (component *Component) AddB() *Component {
	newB := NewB()
	return component.AddChild(newB)
}

func NewMenu() *Component {
	return &Component{
		Node: Node{
			Name: "menu",
		},
		Render: RenderSpec{
			Type:        "html",
			Template:    "menu-html",
			SelfClosing: IsSelfClosing("menu"),
		},
	}
}

func (component *Component) AddMenu() *Component {
	newMenu := NewMenu()
	return component.AddChild(newMenu)
}
