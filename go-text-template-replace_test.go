package main

import (
	"io/ioutil"
	"log"
	"os"
	"testing"
)

//used for testing
func openFile(fileName string) *os.File {
	// open file opens a file, returns the file Object

	file, err := os.Open(fileName)

	if err != nil {
		log.Fatal(err)
	}

	return file
}

//used for testing
func readFile(fileName string) []byte {
	//readFile reads in a file and returns the raw bytes

	file := openFile(fileName)

	data, err := ioutil.ReadAll(file)
	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	return data
}

var inputFileNameTable = []struct {
	in  string
	out error
}{
	{"testdata/testTemplate.tmpl", nil},
	{"testdata/testTokens.toml", nil},
}

func TestInputFileExists(t *testing.T) {
	t.Parallel()
	//TestFileExists tests whether all files exists

	for _, entry := range inputFileNameTable {
		_, err := os.Stat(entry.in)

		if err != nil {
			if os.IsNotExist(err) {
				t.Errorf("File %s does not exist.\n", entry.in)
			}
		}
	}
}

func TestReadInputFile(t *testing.T) {
	//TestReadFile tests whether input files can be read
	t.Parallel()

	for _, entry := range inputFileNameTable {
		data := readFile(entry.in)

		if data == nil {
			t.Errorf("Could not read file %s\n", entry.in)

		}
	}
}

func TestTemplateReplace(t *testing.T) {

	//define test files for this test

	var TemplateFileName = "testdata/testTemplate.tmpl"
	var TokenFileName = "testdata/testTokens.toml"
	var ConfigFileName = "testdata/testOutput.cfg"

	templateReplace(TemplateFileName, TokenFileName, ConfigFileName)
}
