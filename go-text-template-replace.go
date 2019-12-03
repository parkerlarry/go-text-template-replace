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

func writeCFG(
	fileName string,
	textTemplate *template.Template,
	tokenMap map[string]interface{}) error {

	cfgFile, err := os.Create(fileName)

	if err != nil {
		log.Fatal(err)
	}

	defer cfgFile.Close()

	err = textTemplate.Execute(cfgFile, tokenMap)

	if err != nil {
		log.Fatal(err)
	}

	return err
}

func templateReplace(f string, i string, c string) {
	//templateReplace:
	/*Synopsis:
	input:	-f: file name of a text template with .tmpl extensions
					-i: input file with tokens to replace in template with .toml extension
	output: -c: a configuration file with the output of the text replace with .cfg extension

		readTMPL opens the template file and parses it as a tree, and returns a text/template object.
		I call the return value textTemplate.

		readTOML opens the token file and parses it as a tree. After that the tree is converted into a map
		containing key and value pairs of the tokens and their values. I call the return value tokenMap.

		To replace the text template with values in tokenMap,
		textTemplaet is executed with tokenMap, and the output is stored in outputFile.

		To assert whether the text replace was successful, outputFile is opened with readTOML, which returns a map
		I call outputTokenMap. If the text replace is sucessful outputMap and tokenMap should have the same contents.
	*/

	textTemplate, _ := readTMPL(f)

	tokenMap, _ := readTOML(i)

	writeCFG(c, textTemplate, tokenMap)

	outputTokenMap, _ := readTOML(c)

	if !reflect.DeepEqual(tokenMap, outputTokenMap) {
		log.Fatalf("Replacement of %s with tokens in %s failed", f, i)
	}

}

func main() {

	var (
		templateFileName = flag.String("f", "testdata/testTemplate.tmpl", " -f <input template file>")
		tokenFileName    = flag.String("i", "testdata/testTokens.toml", "-i <input token file>")
		configFileName   = flag.String("c", "testdata/testOutput.cfg", "-c <output config file>")
	)

	flag.Parse()

	templateReplace(*templateFileName, *tokenFileName, *configFileName)

}
