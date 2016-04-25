package main

import (
	"fmt"
	"math"
	"os"
	"strings"
	"text/template"

	"github.com/davewalk/adventures-in-go/templates/transformers"
)

type Product struct {
	Brand         string
	OriginalPrice float64
	Price         float64
	Url           string
	ImageUrl      string
	Title         string
	Categories    []string
	Args          []map[string]interface{}
}

func main() {
	funcMap := template.FuncMap{
		"uppercase":     strings.ToUpper,
		"url":           transformers.Url,
		"branded":       transformers.Branded,
		"lowest":        math.Min,
		"sliceToString": transformers.SliceToString,
		"append":        transformers.Append,
	}

	tmpl, err := template.New("cse.tmpl").Funcs(funcMap).ParseFiles("cse.tmpl")
	if err != nil {
		panic(err)
	}

	var args []map[string]interface{}
	arg1 := make(map[string]interface{})
	arg1["value"] = "val1"
	args = append(args, arg1)
	arg2 := make(map[string]interface{})
	arg2["value"] = "val2"
	args = append(args, arg2)
	fmt.Println(args)

	p := Product{
		Brand:         "Nike",
		Price:         89.99,
		OriginalPrice: 84.99,
		Url:           "nike.com",
		ImageUrl:      "http://nike.com",
		Title:         "Nike Air Max Tavas - Men's",
		Categories:    []string{"shoes, classic, something else"},
		Args:          args,
	}

	outfile, err := os.Create("output.csv")
	if err != nil {
		panic(err)
	}

	err = tmpl.Execute(outfile, p)
	if err != nil {
		panic(err)
	}

	fmt.Println("done")
}
