package main

import (
	"os"
	"reflect"
	"testing"
)

// IDEA: Arithmetic in text/templates is a bad idea.
// No control over integer types, but for now test simple stuff

func TestAddFunction(t *testing.T) {

	//define test files for this test

	var addTemplateFileName = "testdata/testAdd.tmpl"
	var addTokenFileName = "testdata/testAdd.toml"
	var addConfigFileName = "testdata/testAdd.cfg"

	testAddTemplate, err := readTMPL(addTemplateFileName)

	if err != nil {
		t.Errorf("Could not create template file %s\n", addTemplateFileName)
	}

	testAddTokenMap, err := readTOML(addTokenFileName)

	if err != nil {
		t.Errorf("Could not create token file %s\n ", addTokenFileName)
	}

	outputFile, err := os.Create(addConfigFileName)

	if err != nil {
		t.Errorf("Could not create output file %s\n: ", addConfigFileName)
	}

	defer outputFile.Close()

	err = testAddTemplate.Execute(outputFile, testAddTokenMap)

	if err != nil {
		t.Errorf("Could not execute template %v: ", testAddTemplate.Name())
	}

	outputTokenMap, err := readTOML(addConfigFileName)

	if err != nil {
		t.Errorf("Could not read output file %s: ", addConfigFileName)
	}

	if !reflect.DeepEqual(testAddTokenMap, outputTokenMap) {
		t.Errorf("Text/template replace of template in %s with tokens in %s failed",
			addTemplateFileName,
			addTokenFileName,
		)
	}
}
