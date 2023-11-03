package main

import (
	"time"

	r "github.com/Hcode30/rego/cmd/web/rego"
)

func main() {
	start := time.Now()
	htmlGenerationPage := r.Page{
		Lang: "en",
		Head: r.Head{
			Title: "HTML Generation Module",
			Link: []r.Link{
				{
					Rel:  "stylesheet",
					Type: "text/css",
					Href: "/static/styles.css",
					As:   "style",
				},
				{
					Rel:  "icon",
					Type: "image/png",
					Href: "favicon.png",
				},
			},
			Script: []r.Script{
				{
					Type: "text/javascript",
					Src:  "script.js",
				},
			},
			Meta: []r.Meta{
				{
					Name:    "description",
					Content: "An HTML generation module for web development.",
				},
			},
		},
		Body: r.Body{
			Attr: "html-generation-module",
			Elements: []r.Element{
				{
					Tag:     "header",
					Class:   "header",
					Content: "HTML Generation Module",
				},
				{
					Tag:   "nav",
					Class: "nav-menu",
					Content: []r.Element{
						{
							Tag: "ul",
							Content: []r.Element{
								{
									Tag:     "li",
									Content: "Home",
								},
								{
									Tag:     "li",
									Content: "Features",
								},
								{
									Tag:     "li",
									Content: "Documentation",
								},
								{
									Tag:     "li",
									Content: "Contact",
								},
							},
						},
					},
				},
				{
					Tag:   "main",
					Class: "content",
					Content: []r.Element{
						{
							Tag:     "h1",
							Id:      "title",
							Content: "Welcome to the HTML Generation Module",
						},
						{
							Tag:     "p",
							Content: "This module helps web developers generate HTML content easily.",
						},
						{
							Tag:     "p",
							Content: "You can customize it for your specific needs.",
						},
					},
				},
				{
					Tag:     "footer",
					Class:   "footer",
					Content: "© 2023 HTML Generation Module, Inc.",
				},
			},
		},
	}
	defer println("duration:", time.Since(start).String())

	server := r.NewHTMLServer()

	server.RegisterTemplate("/", htmlGenerationPage)
	server.RegisterHTML("/Rego", base.GetTemplate())

	server.StartRouter(8080)

}

var base r.Page = r.Page{
	Lang: "Rego",
	Body: r.Body{
		Attr: "id=\"rego\"",
		Elements: []r.Element{
			{Tag: "h1",
				Class:   "header",
				Content: "Rego"},
		},
	},
	Head: r.Head{
		Title: "Rego",
		Link: []r.Link{
			{
				Rel:  "stylesheet",
				Type: "text/css",
				Href: "styles.css",
			},
		},
	},
}
