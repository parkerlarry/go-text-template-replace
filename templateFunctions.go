package main

import "text/template"

func defineAdd() template.FuncMap {
	//defineAdd returns custom function for addition

	funcMap := template.FuncMap{

		// add x y returns x + y
		"add": func(x int, y int) int {
			return x + y
		},
	}
	return funcMap
}
