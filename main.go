package main

import (
	"html/template"
	"os"
	"log"
)

func parseBasic (fname string ){
	var tpl *template.Template = template.Must(template.ParseFiles(fname))

	xs :=[] string {"zero", "one", "two", "three", "four", "five",}

	err := tpl.Execute(os.Stdout, xs)
	if err != nil {
		log.Fatalln(err)
	}
}

func parseIf (fname string ){
	var tpl *template.Template = template.Must(template.ParseFiles(fname))

	u1 := user{
		Name:  "Buddha",
		Motto: "The belief of no beliefs",
		Admin: false,
	}

	u2 := user{
		Name:  "Gandhi",
		Motto: "Be the change",
		Admin: true,
	}

	u3 := user{
		Name:  "",
		Motto: "Nobody",
		Admin: true,
	}

	users := []user{u1, u2, u3}

	err := tpl.Execute(os.Stdout, users)
	if err != nil {
		log.Fatalln(err)
	}
}

func parseCompare (fname string ) {
	var tpl *template.Template = template.Must(template.ParseFiles(fname))

	g1 := struct {
		Score1 int
		Score2 int
	}{
		7,
		9,
	}

	err := tpl.Execute(os.Stdout, g1)
	if err != nil {
		log.Fatalln(err)
	}
}

type user struct {
	Name  string
	Motto string
	Admin bool
}

func main() {

	//parseBasic("./res/global_functions/basic.gohtml")
	//parseIf("./res/global_functions/if.gohtml")
	//parseCompare("./res/global_functions/compare.gohtml")

}
