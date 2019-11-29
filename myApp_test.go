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
	{*templateFileName, nil},
	{*tokenFileName, nil},
}

func TestInputFileExists(t *testing.T) {
	//TestFileExists tests whether all files exists

	for _, entry := range inputFileNameTable {
		fileInfo, err := os.Stat(entry.in)
		if err != nil {
			if os.IsNotExist(err) {
				t.Log("File does not exist.")
			}
		}
		t.Logf("File info: %v\n", fileInfo)
	}
}

func TestReadInputFile(t *testing.T) {
	//TestReadFile tests whether input files can be read

	for _, entry := range inputFileNameTable {
		data := readFile(entry.in)

		t.Logf("File contents: %v\n", data)
	}
}
