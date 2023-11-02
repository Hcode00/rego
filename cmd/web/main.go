package main

import (
	"github.com/Hcode30/rego/cmd/web/rego"
)

func main() {
	HomePage := rego.Page{
		Lang: "ar",
		Head: rego.Head{
			Title: "Home Page",
		},
		Body: rego.Body{
			Attr:     "class='home'",
			Elements: []rego.Element{}},
	}
    println(rego.MakeTemplate(HomePage))

}
