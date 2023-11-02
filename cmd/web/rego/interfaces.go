package rego

import "fmt"

type Page struct {
	Lang string
	Head Head
	Body Body
}
type Head struct {
	Title  string
	Link   []Link
	Script []Script
	Meta   []Meta
}
type Body struct {
	Attr     string
	Elements []Element
}
type Link struct {
	Rel  string
	Type string
	Href string
	As   string
	HTML string
}
type Meta struct {
	Name     string
	Content  string
	Charset  string
	Property string
	HTML     string
}
type Script struct {
	Type string
	Attr string
	Src  string
	HTML string
}
type Element struct {
	Tag     string
	Attr    string
	Class   string
	Id      string
	Content interface{}
	HTML    string
}

// -----   Page   -----
// Language
func (p *Page) SetLang(lang string) {
	LANGS := []string{"af", "ar", "az", "be", "bg", "bs-Latn", "ca", "cs", "cy-GB", "da", "de", "dv", "el", "en", "es", "et", "eu", "fa", "fi", "fo", "fr", "gl", "gu", "hi", "hr", "hu", "hy", "id", "is", "it", "ja", "ka", "kk", "kn", "ko", "kok-IN", "ky", "lt", "lv", "mi-NZ", "mk", "mn", "mr", "ms", "mt-MT", "nb-NO", "nl", "nn-NO", "no", "ns-ZA", "pa", "pl", "pt", "quz-BO", "quz-EC", "quz-PE", "ro", "ru", "sa", "se-FI", "se-NO", "se-SE", "sk", "sl", "sma-NO", "sma-SE", "smj-NO", "smj-SE", "smn-FI", "sms-FI", "sq", "sr", "sv", "sw", "syr", "ta", "te", "th", "tn-ZA", "tr", "tt", "uk", "ur", "uz", "vi", "xh-ZA", "zh-CN", "zh-HK", "zh-CHS", "zh-CHT", "zh-MO", "zh-SG", "zh-TW", "zu-ZA"}
	// make sure lang is valid
	for _, l := range LANGS {
		if l == lang {
			p.Lang = lang
		}
	}
}

func (p *Page) GetLang() string {
	return p.Lang
}

// Title
func (p *Page) SetTitle(title string) {
	p.Head.Title = title
}
func (p *Page) GetTitle() string {
	return p.Head.Title
}
func (p *Page) GetTitleHTML() string {
	return "<title>" + p.Head.Title + "</title>\n"
}

// Link
func (p *Page) AddLink(Rel, Href, Type, As string) {
	link := Link{
		Rel:  Rel,
		Href: Href,
		As:   As,
		Type: Type,
		HTML: fmt.Sprintf(`<link> rel="%s" href="%s" type="%s" as="%s"`, Rel, Href, Type, As) + "\n",
	}
	println(link.HTML)
	p.Head.Link = append(p.Head.Link, link)
}
func (p *Page) GetLinks() []Link {
	return p.Head.Link
}
func (p *Page) GetLinksHTML() string {
	str := ""
	for _, link := range p.Head.Link {
		if link.HTML == "" {
			link.HTML = fmt.Sprintf(`<link> rel="%s" href="%s" type="%s" as="%s"`, link.Rel, link.Href, link.Type, link.As) + "\n"
			str += link.HTML
		}
	}
	return str
}

// Script
func (p *Page) AddScript(Type, Src, Attr string) {
	script := Script{
		Type: Type,
		Src:  Src,
		Attr: Attr,
		HTML: fmt.Sprintf(`<script type="%s" src="%s" %s></script>`, Type, Src, Attr) + "\n",
	}
	p.Head.Script = append(p.Head.Script, script)
}
func (p *Page) GetScripts() []Script {
	return p.Head.Script
}
func (p *Page) GetScriptsHTML() string {
	str := ""
	for _, script := range p.Head.Script {
		if script.HTML == "" {
			script.HTML = fmt.Sprintf(`<script type="%s" src="%s" %s></script>`, script.Type, script.Src, script.Attr) + "\n"
		}
		str += script.HTML
	}
	return str
}

// Meta
func (p *Page) AddMeta(Name, Content, Property, Charset string) {
	var meta Meta
	if Charset == "" {
		meta = Meta{
			Name:     Name,
			Content:  Content,
			Property: Property,
			HTML:     fmt.Sprintf(`<meta name="%s" content="%s" property="%s">`, Name, Content, Property) + "\n",
		}
	} else {
		meta = Meta{
			Name:     Name,
			Content:  Content,
			Charset:  Charset,
			Property: Property,
			HTML:     fmt.Sprintf(`<meta name="%s" content="%s" charset="%s" property="%s">`, Name, Content, Charset, Property) + "\n",
		}
	}
	p.Head.Meta = append(p.Head.Meta, meta)
}
func (p *Page) GetMetas() []Meta {
	return p.Head.Meta
}
func (p *Page) GetMetasHTML() string {
	str := ""
	for _, meta := range p.Head.Meta {
		if meta.HTML == "" {
			if meta.Charset == "" {
				meta.HTML = fmt.Sprintf(`<meta name="%s" content="%s" property="%s">`, meta.Name, meta.Content, meta.Property) + "\n"
			} else {
				meta.HTML = fmt.Sprintf(`<meta name="%s" content="%s" charset="%s" property="%s">`, meta.Name, meta.Content, meta.Charset, meta.Property) + "\n"
			}
		}
		str += meta.HTML
	}
	return str
}

// -----   Body   -----
// Elements

func (p *Page) AddElement(Tag, Attr, Class, Id string, Content interface{}) {

	var element Element
	if Id == "" {
		element = Element{
			Tag:     Tag,
			Attr:    Attr,
			Class:   Class,
			Content: Content,
			HTML:    fmt.Sprintf(`<%s %s class="%s">%s</%s>`, Tag, Attr, Class, Content, Tag) + "\n",
		}
	} else {
		element = Element{
			Tag:     Tag,
			Attr:    Attr,
			Class:   Class,
			Id:      Id,
			Content: Content,
			HTML:    fmt.Sprintf(`<%s %s class="%s" id="%s">%s</%s>`, Tag, Attr, Class, Id, Content, Tag) + "\n",
		}
	}
	p.Body.Elements = append(p.Body.Elements, element)
}

func (p *Page) GetElements() []Element {
	return p.Body.Elements
}


func (p *Page) GetElementsHTML() string {
    str := ""
    for _, element := range p.Body.Elements {
        switch v := element.Content.(type) {
        case string:
            element.HTML = generateElementHTML(element)
        case []Element:
            element.HTML = generateNestedElementsHTML(v)
        case Element:
            element.HTML = generateElementHTML(element)
        default:
            // Handle other cases if necessary
        }
        str += element.HTML
    }
    return str
}

func generateNestedElementsHTML(elements []Element) string {
    str := ""
    for _, element := range elements {
        switch v := element.Content.(type) {
        case string:
            element.HTML = generateElementHTML(element)
        case []Element:
            element.HTML = generateNestedElementsHTML(v) // Recursively generate HTML for nested elements
        case Element:
            element.HTML = generateElementHTML(element)
        default:
            // Handle other cases if necessary
        }
        str += element.HTML
    }
    return str
}


func generateElementHTML(element Element) string {
    if element.Tag == "" {
        element.Tag = "div" // Default to 'div' if tag is not provided
    }

    // Initialize an empty string to hold the HTML
    html := fmt.Sprintf("<%s", element.Tag)

    if element.Attr != "" {
        html += " " + element.Attr
    }

    if element.Class != "" {
        html += fmt.Sprintf(" class='%s'", element.Class)
    }

    if element.Id != "" {
        html += fmt.Sprintf(" id='%s'", element.Id)
    }

    if v, isString := element.Content.(string); isString {
        html += fmt.Sprintf(">%s</%s>\n", v, element.Tag)
    } else {
        html += "></" + element.Tag + ">\n"
    }

    return html
}

