package rego

type Page struct {
	Lang string
	Head Head
	Body Body
	HTML string
}
type Head struct {
	Title  string
	Link   []Link
	Script []Script
	Meta   []Meta
	HTML   string
}
type Body struct {
	Attr     string
	Elements []Element
	HTML     string
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

type HTMLServer struct {
	templates map[string]Page
}


