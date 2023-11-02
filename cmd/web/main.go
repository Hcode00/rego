package main

import (
	r "github.com/Hcode30/rego/cmd/web/rego"
)

func main() {
	htmlGenerationPage := r.Page{
		Lang: "en",
		Head: r.Head{
			Title: "HTML Generation Module",
			Link: []r.Link{
				{
					Rel:  "stylesheet",
					Type: "text/css",
					Href: "styles.css",
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
                            Id:   "title",
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

    pageText := r.MakeTemplate(htmlGenerationPage)
    // r.MakeTemplate(htmlGenerationPage)
	println(pageText)

}
