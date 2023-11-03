package rego

import (
	"fmt"
)

// Lang

func (p *Page) SetLang(lang string) {
	LANGS := []string{"af", "ar", "az", "be", "bg", "bs-Latn", "ca", "cs", "cy-GB", "da", "de", "dv", "el", "en", "es", "et", "eu", "fa", "fi", "fo", "fr", "gl", "gu", "hi", "hr", "hu", "hy", "id", "is", "it", "ja", "ka", "kk", "kn", "ko", "kok-IN", "ky", "lt", "lv", "mi-NZ", "mk", "mn", "mr", "ms", "mt-MT", "nb-NO", "nl", "nn-NO", "no", "ns-ZA", "pa", "pl", "pt", "quz-BO", "quz-EC", "quz-PE", "ro", "ru", "sa", "se-FI", "se-NO", "se-SE", "sk", "sl", "sma-NO", "sma-SE", "smj-NO", "smj-SE", "smn-FI", "sms-FI", "sq", "sr", "sv", "sw", "syr", "ta", "te", "th", "tn-ZA", "tr", "tt", "uk", "ur", "uz", "vi", "xh-ZA", "zh-CN", "zh-HK", "zh-CHS", "zh-CHT", "zh-MO", "zh-SG", "zh-TW", "zu-ZA"}
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
	p.Head.SetTitle(title)
}
func (p *Page) GetTitle() string {
	return p.Head.GetTitle()
}
func (p *Page) GetTitleHTML() string {
	return p.Head.GetTitleHTML()
}
func (h *Head) SetTitle(title string) {
	h.Title = title
}
func (h *Head) GetTitle() string {
	return h.Title
}
func (h *Head) GetTitleHTML() string {
	return "<title>" + h.Title + "</title>\n"
}
// Style
func (h *Head) AddStyle(block string) {
    st := Style{
        block,
    }
    h.Style = append(h.Style, st)
}

// Link
func (h *Head) AddLink(Rel, Href, Type, As string) {
	link := Link{
		Rel:  Rel,
		Href: Href,
		As:   As,
		Type: Type,
		HTML: GenerateLinkHTML(Rel, Href, Type, As),
	}
	h.Link = append(h.Link, link)
}
func (h *Head) GetLinks() []Link {
	return h.Link
}
func (h *Head) GetLinksHTML() string {
	str := ""
	for _, link := range h.Link {
		if link.HTML == "" {
			link.HTML = GenerateLinkHTML(link.Rel, link.Href, link.Type, link.As)
		}
		str += link.HTML
	}
	return str
}
func (p *Page) AddLink(Rel, Href, Type, As string) {
	p.Head.AddLink(Rel, Href, Type, As)
}
func (p *Page) GetLinks() []Link {
	return p.Head.GetLinks()
}
func (p *Page) GetLinksHTML() string {
	return p.Head.GetLinksHTML()
}

func (h *Head) AddCSSLink(Href string) {
	h.AddLink("stylesheet", Href, "text/css", "style")
}
func (p *Page) AddCSSLink(Href string) {
	p.Head.AddCSSLink(Href)
}

// Script
func (h *Head) AddScript(Type, Src, Attr string) {
	script := Script{
		Type: Type,
		Src:  Src,
		Attr: Attr,
		HTML: GenerateScriptHTML(Type, Src, Attr),
	}
	h.Script = append(h.Script, script)
}
func (h *Head) GetScripts() []Script {
	return h.Script
}
func (h *Head) GetScriptsHTML() string {
	str := ""
	for _, script := range h.Script {
		if script.HTML == "" {
			script.HTML = GenerateScriptHTML(script.Type, script.Src, script.Attr)
		}
		str += script.HTML
	}
	return str
}
func (p *Page) AddScript(Type, Src, Attr string) {
	p.Head.AddScript(Type, Src, Attr)
}
func (p *Page) GetScripts() []Script {
	return p.Head.GetScripts()
}
func (p *Page) GetScriptsHTML() string {
	return p.Head.GetScriptsHTML()
}

// Meta
func (h *Head) AddMeta(Name, Content, Property, Charset string) {
	var meta Meta
	meta = Meta{
		Name:     Name,
		Content:  Content,
		Charset:  Charset,
		Property: Property,
		HTML:     GenerateMetaHTML(Name, Content, Property, Charset),
	}

	h.Meta = append(h.Meta, meta)
}
func (h *Head) GetMetas() []Meta {
	return h.Meta
}
func (h *Head) GetMetasHTML() string {
	str := ""
	for _, meta := range h.Meta {
		if meta.HTML == "" {
			meta.HTML = GenerateMetaHTML(meta.Name, meta.Content, meta.Property, meta.Charset)
		}
		str += meta.HTML
	}
	return str
}

func (p *Page) AddMeta(Name, Content, Property, Charset string) {
	p.Head.AddMeta(Name, Content, Property, Charset)
}
func (p *Page) GetMetas() []Meta {
	return p.Head.GetMetas()
}
func (p *Page) GetMetasHTML() string {
	return p.Head.GetMetasHTML()
}

// Body Attr
func (b *Body) SetAttr(attr string) {
	b.Attr = attr
}
func (b *Body) GetAttr() string {
	return b.Attr
}
func (p *Page) SetBodyAttr(attr string) {
	p.Body.SetAttr(attr)
}
func (p *Page) GetBodyAttr() string {
	return p.Body.GetAttr()
}

// Elements
func (b *Body) AddElement(Tag, Attr, Class, Id string, Content interface{}) {
	var element Element
	element = Element{
		Tag:     Tag,
		Attr:    Attr,
		Class:   Class,
		Id:      Id,
		Content: Content,
		HTML:    generateElementHTML(element) + NEW_BODY_LINE,
	}
	b.Elements = append(b.Elements, element)
}
func (b *Body) GetElements() []Element {
	return b.Elements
}
func (b *Body) GetElementsHTML() string {
	str := ""
	for _, element := range b.Elements {
		html := generateElementHTML(element)
		str += html
	}
	return str
}
func (p *Page) AddElement(Tag, Attr, Class, Id string, Content interface{}) {
	p.Body.AddElement(Tag, Attr, Class, Id, Content)
}
func (p *Page) GetElements() []Element {
	return p.Body.GetElements()
}
func (p *Page) GetElementsHTML() string {
	return p.Body.GetElementsHTML()
}

// Making The Page
func (p *Page) MakeTemplate() {
	template := INITIAL_TEMPLATE
	template = ReplaceLang(*p, template)
	template = HeadToTemplate(*p, template)
	template = BodyAttrToTemplate(*p, template)
	template = BodyElementsToTemplate(*p, template)
    p.Head.UpdateHeadHTML()
	p.Body.UpdateBodyHTML()
	p.SetTemplate(template)
}
func (p *Page) SetBodyHTML(html string) {
	p.Body.HTML = html
}
func (p *Page) SetHeadHTML(html string) {
	p.Head.HTML = html
}
func (p *Page) SetTemplate(html string) {
	p.HTML = html
}
func (b *Body) UpdateBodyHTML() {
	str := ""
	str += "<body " + b.Attr + ">" + NEW_BODY_LINE
	str += b.GetElementsHTML()
	str += "</body>"
	b.HTML = str
}
func (h *Head) UpdateHeadHTML() {
	str := ""
	str += h.GetTitleHTML()
	str += h.GetLinksHTML()
	str += h.GetScriptsHTML()
	str += h.GetMetasHTML()
	h.HTML = str
}

func (p *Page) GetBodyHTML() string {
	if p.Body.HTML == "" {
		p.Body.UpdateBodyHTML()
	}
	return p.Body.HTML
}

func (p *Page) GetHeadHTML() string {
	if p.Head.HTML == "" {
		p.Head.UpdateHeadHTML()
	}
	return p.Head.HTML
}

func (p *Page) GetTemplate() string {
	if p.HTML == "" {
		p.MakeTemplate()
	}
	return p.HTML
}

// Resetting
// Reset Head
func (p *Page) ResetHead() {
    p.Head.Title = ""
    p.Head.Link = []Link{}
    p.Head.Script = []Script{}
    p.Head.Meta = []Meta{}
    p.MakeTemplate()
}

// Reset Body
func (p *Page) ResetBody() {
    p.Body.Attr = ""
    p.Body.Elements = []Element{}
    p.MakeTemplate()
}

// reset Page
func (p *Page) ResetPage() {
	p.ResetHead()
	p.ResetBody()
}

// Filtering functions
func generateNestedElementsHTML(elements []Element) string {
	str := ""
	for _, element := range elements {
		html := generateElementHTML(element)
		str += html
	}
	return str
}
func generateElementHTML(element Element) string {
	if element.Tag == "" {
		element.Tag = "div"
	}
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
		html += fmt.Sprintf(">%s</%s>", v, element.Tag)
	} else if nestedElements, isSlice := element.Content.([]Element); isSlice {
		// Handle nested elements
		html += ">" + NEW_BODY_LINE
		for _, nestedElement := range nestedElements {
			html += generateElementHTML(nestedElement)
		}
		html += fmt.Sprintf("</%s>", element.Tag)
	}

	return html + NEW_BODY_LINE
}
func GenerateLinkHTML(Rel, Href, Type, As string) string {
	linkHTML := fmt.Sprintf(`<link rel="%s" href="%s" type="%s"`, Rel, Href, Type)

	if As != "" {
		linkHTML += fmt.Sprintf(` as="%s"`, As)
	}


	linkHTML += ">" + NEW_HEAD_LINE
	return linkHTML
}

func GenerateScriptHTML(Type, Src, Attr string) string {
	scriptHTML := fmt.Sprintf(`<script type="%s" src="%s"`, Type, Src)

	if Attr != "" {
		scriptHTML += fmt.Sprintf(` %s`, Attr)
	}

	scriptHTML += "></script>" + NEW_HEAD_LINE
	return scriptHTML
}

func GenerateMetaHTML(Name, Content, Property, Charset string) string {
	metaHTML := "<meta"

	if Name != "" {
		metaHTML += fmt.Sprintf(` name="%s"`, Name)
	}

	if Content != "" {
		metaHTML += fmt.Sprintf(` content="%s"`, Content)
	}

	if Property != "" {
		metaHTML += fmt.Sprintf(` property="%s"`, Property)
	}

	if Charset != "" {
		metaHTML += fmt.Sprintf(` charset="%s"`, Charset)
	}

	metaHTML += "/>" + NEW_HEAD_LINE
	return metaHTML
}

// func (Page *Page) InjectToBody(html string) error {
// }
// func (Page *Page) InjectToHead(html string) error {
// }
