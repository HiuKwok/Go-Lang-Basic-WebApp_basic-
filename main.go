package main

import (
	"text/template"
	"log"
	"os"
)

var x int



func main() {

	tpl, err := template.ParseFiles("./res/tpl.gohtml")

	//Error handling
	if err!= nil {
		log.Fatalln(err)
	}
	nf, err := os.Create("./temp/index.html")

	if err != nil {
		log.Println("Error creating file", err)
	}

	defer nf.Close()



	//Write to html file instead
	err = tpl.Execute(nf, nil)
	if err != nil {
		log.Fatalln(err)
	}


}
