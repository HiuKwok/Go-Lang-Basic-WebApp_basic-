package main

import (
	"text/template"
	"os"
	"log"
)

var tpl *template.Template

type sage struct {
	Name  string
	Motto string
}


type car struct {
	Manufacturer string
	Model        string
	Doors        int
}


func parseNormal (fileName string) {

	tpl = template.Must(template.ParseFiles(fileName))
	err := tpl.ExecuteTemplate(os.Stdout, "wisdom.gohtml", `Release self-focus; embrace other-focus.`)
	if err != nil {
		log.Fatalln(err)
	}
}



func parseRange (fileName string) {

	tpl = template.Must(template.ParseFiles(fileName))
	sages := []string{"Gandhi", "MLK", "Buddha", "Jesus", "Muhammad"}

	err := tpl.Execute(os.Stdout, sages)
	if err != nil {
		log.Fatalln(err)
	}
}

func parseStruct (fileName string) {

	tpl = template.Must(template.ParseFiles(fileName))
	buddha := sage{
		Name:  "Buddha",
		Motto: "The belief of no beliefs",
	}

	err := tpl.Execute(os.Stdout, buddha)
	if err != nil {
		log.Fatalln(err)
	}
}


func parseSliceStruct (fileName string) {

	tpl = template.Must(template.ParseFiles(fileName))

	buddha := sage{
		Name:  "Buddha",
		Motto: "The belief of no beliefs",
	}

	gandhi := sage{
		Name:  "Gandhi",
		Motto: "Be the change",
	}

	mlk := sage{
		Name:  "Martin Luther King",
		Motto: "Hatred never ceases with hatred but with love alone is healed.",
	}

	jesus := sage{
		Name:  "Jesus",
		Motto: "Love all",
	}

	muhammad := sage{
		Name:  "Muhammad",
		Motto: "To overcome evil with good is good, to resist evil by evil is evil.",
	}

	sages := []sage{buddha, gandhi, mlk, jesus, muhammad}

	err := tpl.Execute(os.Stdout, sages)
	if err != nil {
		log.Fatalln(err)
	}
}


func parseStructSliceStruct (fileName string) {

	tpl = template.Must(template.ParseFiles(fileName))

	b := sage{
		Name:  "Buddha",
		Motto: "The belief of no beliefs",
	}

	g := sage{
		Name:  "Gandhi",
		Motto: "Be the change",
	}

	m := sage{
		Name:  "Martin Luther King",
		Motto: "Hatred never ceases with hatred but with love alone is healed.",
	}

	f := car{
		Manufacturer: "Ford",
		Model:        "F150",
		Doors:        2,
	}

	c := car{
		Manufacturer: "Toyota",
		Model:        "Corolla",
		Doors:        4,
	}

	sages := []sage{b, g, m}
	cars := []car{f, c}

	data := struct {
		Wisdom    []sage
		Transport []car
	}{
		sages,
		cars,
	}

	err := tpl.Execute(os.Stdout, data)
	if err != nil {
		log.Fatalln(err)
	}
}


func main() {
	//parseRange("./res/parseRange.gohtml")
	//parseNormal("./res/wisdom.gohtml")
	//parseStruct("./res/parseStruct.gohtml")
	//parseSliceStruct("./res/parseSliceStruct.gohtml")
	parseStructSliceStruct("./res/structSliceStruct.gohtml")
}