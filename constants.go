package rego

var Lang string = "en"
var NewBodyLine = "\n"
var NewHeadLine = "\n"
var HTMLDir string = "./ui/html/"
var InitialTemplate string = `<!DOCTYPE html>
<html lang="((LANG))">
<head>
((HEAD))
</head>
<body ((BODY.ATTR))>
((BODY.ELEMENTS))
</body>
</html>`

var NotFoundTemplate string = `<!DOCTYPE html>
<html>
<head>
    <title>404 - Not Found</title>
    <style>
        body {
            font-family: Arial, sans-serif;
            background-color: #f4f4f4;
            margin: 0;
            padding: 0;
            display: flex;
            justify-content: center;
            align-items: center;
            height: 100vh;
        }

        .container {
            text-align: center;
        }

        h1 {
            color: #333;
            font-size: 100px;
            font-weight: bold;
            margin: 0;
        }

        p {
            color: #666;
            font-size: 20px;
        }
    </style>
</head>
<body>
    <div class="container">
        <h1>404</h1>
        <p>Not Found</p>
    </div>
</body>
</html>
`
