package rego

var NUM_OF_PAGES int= 0
var LANG string= "en"
var STATIC_DIR string= "./ui/static"
var HTML_DIR string= "./ui/html"
var INITIAL_TEMPLATE string= `
{{define "page"}}
<!DOCTYPE html>
<html lang="{{.Lang}}">
<head>
{{template "head" .Head.getHead()}}
</head>
<body {{template "body.attr"}}>
{{template "body" .Body.getBody()}}
</body>
`




