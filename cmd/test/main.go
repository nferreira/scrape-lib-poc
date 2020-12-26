package main

import (
	"github.com/PuerkitoBio/goquery"
	"github.com/nferreira/godom/pkg/model"
	"github.com/robertkrimen/otto"
	"golang.org/x/net/html"
	"log"
	//"os"
	"strings"
)
func main() {
	var rootNode *html.Node
	var err error
	docString := strings.NewReader("<html class=\"cool super\">" +
		"<head>" +
		"<title>This is a test</title>" +
		"<script>" +
		"var para = document.createElement(\"p\");" +
		"var node = document.createTextNode(\"This is new.\");" +
		"para.appendChild(node);" +
		"var element = document.getElementById(\"div1\");" +
		"element.appendChild(para);" +
		"</script>" +
		"</head>" +
		"<body>Test</body>" +
		"<div id=\"div1\">" +
		"<p id=\"p1\">This is a paragraph.</p>" +
		"<p id=\"p2\">This is another paragraph.</p>" +
		"</div>" +
		"</html")
	rootNode, err = html.Parse(docString)

	doc := goquery.NewDocumentFromNode(rootNode)

	if err != nil {
		log.Panicf("Error: %s", err.Error())
	}

	vm := otto.New()
	document := model.NewDocument(vm, doc)
	vm.Set("document", document)
	vm.Set(
		"window",
		model.NewWindow(vm, document),
	)

	vm.Set("name", "Nadilson")
	_, err = vm.Run(`
    	abc = 2 + 2;
    	console.log("The value of abc is " + abc); // 4
		console.log("My name is " + name);

		console.log(window);
		console.log(window.document.getElementById("div1"));
	`)
	if err != nil {
		panic(err)
	}

	//html.Render(os.Stdout, rootNode)
}
