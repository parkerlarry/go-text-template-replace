package main

import (
	"flag"
	"log"
	"os"
	"path/filepath"
	"reflect"
	"text/template"

	"github.com/pelletier/go-toml"
)

func readTMPL(fileName string) (*template.Template, error) {
	//readTMPL parses the the TMPL file, and returns the text/template object

	//add custom functions defined in "./templateFunctions.go"

	//FIXED: New () needs base name for files in abs path
	textTemplate, err := template.New(filepath.Base(fileName)).Funcs(
		defineAdd(),
	).ParseFiles(fileName)
	//
	if err != nil {
		log.Fatal(err)
	}

	return textTemplate, err
}

func readTOML(fileName string) (map[string]interface{}, error) {
	//readTOML loads the TOML file, and returns a map of the syntax tree

	tomlTree, err := toml.LoadFile(fileName)

	if err != nil {
		log.Fatal(err)
	}

	return tomlTree.ToMap(), err
}

func templateReplace(f string, i string, c string) {

	textTemplate, _ := readTMPL(f)

	tokenMap, _ := readTOML(i)

	outputFile, err := os.Create(c)

	if err != nil {
		log.Fatal(err)
	}

	defer outputFile.Close()

	err = textTemplate.Execute(outputFile, tokenMap)

	if err != nil {
		log.Fatal(err)
	}

	outputTokenMap, err := readTOML(c)

	if err != nil {
		log.Fatal(err)
	}

	if !reflect.DeepEqual(tokenMap, outputTokenMap) {
		log.Fatalf("Text/template replace of template in %s with tokens in %s failed",
			f,
			i,
		)
	}

}

func main() {

	var (
		templateFileName = flag.String("f", "testdata/testTemplate.tmpl", " -f <input file>")
		tokenFileName    = flag.String("i", "testdata/testTokens.toml", "-i <input token file>")
		configFileName   = flag.String("c", "testdata/testOutput.cfg", "-c <output file>")
	)

	flag.Parse()

	templateReplace(*templateFileName, *tokenFileName, *configFileName)

}
