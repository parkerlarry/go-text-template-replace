package main

import (
	"flag"
	"log"
	"os"
	"text/template"

	"github.com/pelletier/go-toml"
)

var (
	templateFileName = flag.String("f", "myInput.tmpl", " -f <input file>")
	tokenFileName    = flag.String("i", "myValues.toml", "-i <input token file>")
	configFileName   = flag.String("c", "myApp.cfg", "-c <output file>")
)

func readTMPL(fileName string) *template.Template {
	//readTMPL parses the the TMPL file, and returns the text/template object

	//add custom functions defined in "./templateFunctions.go"
	textTemplate, err := template.New(fileName).Funcs(
		defineAdd(),
	).ParseFiles(fileName)

	if err != nil {
		log.Fatal(err)
	}

	return textTemplate
}

func readTOML(fileName string) map[string]interface{} {
	//readTOML loads the TOML file, and returns a map of the syntax tree

	tomlTree, err := toml.LoadFile(fileName)

	if err != nil {
		log.Fatal(err)
	}

	return tomlTree.ToMap()
}

func main() {

	flag.Parse()

	textTemplate := readTMPL(*templateFileName)

	tokenMap := readTOML(*tokenFileName)

	//TODO: store output to configFileName, with replaced tokens in tokenMap
	textTemplate.Execute(os.Stdout, tokenMap)

}
