package main

import(
	"flag"
	"os"
	"golang.org/x/net/html"
	// "fmt"
	"github.com/jntullo/link"
	"fmt"
)

func main(){
	fileName := flag.String("file", "ex1.html", "the file to parse")
	flag.Parse()

	file, err := os.Open(*fileName)
	if err != nil {
		panic(err)
	}

	// Documentation: https://godoc.org/golang.org/x/net/html
	doc, err := html.Parse(file)
	links := link.Traverse(doc)
	fmt.Println(links)
}


