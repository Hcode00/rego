package rego

var NUM_OF_PAGES int= 0
var LANG string= "en"
var STATIC_DIR string= "./ui/static"
var HTML_DIR string= "./ui/html"
var BASE_FILE string= "index.html"
var INITIAL_TEMPLATE string= 
`<!DOCTYPE html>
<html lang="((LANG))">
<head>
((HEAD))
</head>
<body ((BODY.ATTR))>
((BODY.ELEMENTS))
</body>
</html>`




