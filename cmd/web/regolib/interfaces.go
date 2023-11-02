package rego

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
	Attr    []Attr
	Class   []Class
	Id      Id
	Style   []Style
	Element []Element
}
type Link struct {
	Rel  string
	Type string
	Href string
	As   string
}
type Meta struct {
	Name     string
	Content  string
	CharSet  string
	Property string
}
type Script struct {
	Type string
	Attr []Attr
	Src  string
}
type Element struct {
	Tag     string
	Style   []Style
	Attr    []Attr
	Content interface{}
}

type Attr string
type Class string
type Id string
type Style string

// -----   Page   -----
// Page
func (p Page) setPage(page Page) {
    p = page
}
func (p Page) getPage() Page {
    return p
}
// Lang
func (b Page) setLang(lang string) {
	b.Lang = lang
}
func (b Page) getLang() string {
	return b.Lang
}

// -----   Head  -----
// Title
func (h Head) setTitle(title string) {
	h.Title = title
}
func (h Head) getTitle() string {
	return h.Title
}

// Link
func (h Head) setLink(link []Link) {
	h.Link = link
}
func (h Head) getLink() []Link {
	return h.Link
}

// Script
func (h Head) setScript(script []Script) {
	h.Script = script
}
func (h Head) getScript() []Script {
	return h.Script
}

// Meta
func (h Head) setMeta(meta []Meta) {
	h.Meta = meta
}
func (h Head) getMeta() []Meta {
	return h.Meta
}

// -----   Body  -----
// Attr
func (b Body) setAttr(attr []Attr) {
	b.Attr = attr
}
func (b Body) getAttr() []Attr {
	return b.Attr
}

// Class
func (b Body) setClass(class []Class) {
	b.Class = class
}
func (b Body) getClass() []Class {
	return b.Class
}

// Id
func (b Body) setId(id Id) {
	b.Id = id
}
func (b Body) getId() Id {
	return b.Id
}

// Style
func (b Body) setStyle(style []Style) {
	b.Style = style
}
func (b Body) getStyle() []Style {
	return b.Style
}

// Element
func (b Body) setElement(element []Element) {
	b.Element = element
}
func (b Body) getElement() []Element {
	return b.Element
}

// -----   Link  -----
// Rel
func (l Link) setRel(rel string) {
	l.Rel = rel
}
func (l Link) getRel() string {
	return l.Rel
}

// Type
func (l Link) setType(t string) {
	l.Type = t
}
func (l Link) getType() string {
	return l.Type
}

// Href
func (l Link) setHref(href string) {
	l.Href = href
}
func (l Link) getHref() string {
	return l.Href
}

// As
func (l Link) setAs(as string) {
	l.As = as
}
func (l Link) getAs() string {
	return l.As
}

// -----   Meta  -----
// Name
func (m Meta) setName(name string) {
	m.Name = name
}
func (m Meta) getName() string {
	return m.Name
}

// Content
func (m Meta) setContent(content string) {
	m.Content = content
}
func (m Meta) getContent() string {
	return m.Content
}

// CharSet
func (m Meta) setCharSet(charSet string) {
	m.CharSet = charSet
}
func (m Meta) getCharSet() string {
	return m.CharSet
}

// Property
func (m Meta) setProperty(property string) {
	m.Property = property
}
func (m Meta) getProperty() string {
	return m.Property
}

// -----   Script  -----
// Type
func (s Script) setType(t string) {
	s.Type = t
}
func (s Script) getType() string {
	return s.Type
}

// Attr
func (s Script) setAttr(attr []Attr) {
	s.Attr = attr
}
func (s Script) getAttr() []Attr {
	return s.Attr
}

// Src
func (s Script) setSrc(src string) {
	s.Src = src
}
func (s Script) getSrc() string {
	return s.Src
}

// -----   Element  -----
// Content
func (e *Element) SetContent(content interface{}) {
	e.Content = content
}

func (e *Element) GetContent() interface{} {
	switch v := e.Content.(type) {
	case string:
		return v
	case []string:
		return v
	case Element:
		return v
	case []Element:
		return v
	default:
		panic("unknown content type")
	}
}

// Tag
func (e *Element) SetTag(tag string) {
	e.Tag = tag
}
func (e *Element) GetTag() string {
	return e.Tag
}

// Style
func (e *Element) SetStyle(style []Style) {
	e.Style = style
}
func (e *Element) GetStyle() []Style {
	return e.Style
}

// Attr
func (e *Element) SetAttr(attr []Attr) {
	e.Attr = attr
}
func (e *Element) GetAttr() []Attr {
	return e.Attr
}
