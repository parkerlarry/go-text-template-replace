package main

import "testing"

// IDEA: Arithmetic in text/templates is a bad idea.
// No control over integer types, but for now test simple stuff

func TestAddFunction(t *testing.T) {

	//define test files for this test

	var addTemplateFileName = "testdata/testAdd.tmpl"
	var addTokenFileName = "testdata/testAdd.toml"
	var addConfigFileName = "testdata/testAdd.cfg"

	templateReplace(addTemplateFileName, addTokenFileName, addConfigFileName)
}
