package rego

import (
	"os"
)

func initiateHTMLFile() {
	// make a new html file
	File_Path := HTML_DIR + "/index.html"
	html, err := os.OpenFile(File_Path, os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		panic(err)
	}
	defer html.Close()
	// write to the file
	_, err = html.WriteString(INITIAL_TEMPLATE)
	if err != nil {
		panic(err)
	}
	err = html.Sync()
	if err != nil {
		panic(err)
	}

}
