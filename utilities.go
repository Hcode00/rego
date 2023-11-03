package rego

import (
	"fmt"
	"io"
	"os"
	"strings"
)

func MakeHTMLFile(content, filename, dir string) string {
	if err := os.MkdirAll(HTML_DIR, os.ModePerm); err != nil {
		fmt.Printf("Failed to create HTML_DIR: %v\n", err)
		panic(err)
	}
	filePath := dir + "/" + filename
	if _, err := os.Stat(filePath); err == nil {
		if err := os.Remove(filePath); err != nil {
			fmt.Printf("Failed to remove file: %v\n", err)
			panic(err)
		}
	}
	file, err := os.OpenFile(filePath, os.O_RDWR|os.O_CREATE, os.ModePerm)
	if err != nil {
		fmt.Printf("Failed to create file: %v\n", err)
		panic(err)
	}
	defer file.Close()
	if _, err := io.WriteString(file, content); err != nil {
		fmt.Printf("Failed to write content to file: %v\n", err)
		panic(err)
	}

	fmt.Printf("File created: %s\n", filePath)
	return filePath
}

func ReplaceLang(Page Page, template string) string {
	language := Page.GetLang()
	if language == "" {
		language = LANG
	}
	template = strings.Replace(template, "((LANG))", language, 1)
	return template
}

func HeadToTemplate(Page Page, template string) string {
	str := ""
	if Page.Head.Title != "" {
		str += Page.GetTitleHTML()
	} else {
		Page.SetTitle("Rego")
		str += Page.GetTitleHTML()
	}
	str += Page.GetLinksHTML()
	str += Page.GetScriptsHTML()
	str += Page.GetMetasHTML()
	template = strings.Replace(template, "((HEAD))", str, 1)

	return template
}

func BodyAttrToTemplate(Page Page, template string) string {
	template = strings.Replace(template, "((BODY.ATTR))", Page.Body.GetAttr(), 1)
	return template
}

func BodyElementsToTemplate(Page Page, template string) string {
	if len(Page.Body.Elements) > 0 {
		template = strings.Replace(template, "((BODY.ELEMENTS))", Page.GetElementsHTML(), 1)
	} else {
		text := `<h1>Rego</h1><br><p>Rego is a Go library for generating HTML.</p>`
		template = strings.Replace(template, "((BODY.ELEMENTS))", text, 1)
	}
	return template
}
