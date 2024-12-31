package main

import (
	"fmt"
	"go/ast"
	"os"
	"reflect"
	"regexp"
	"strings"
	. "saral/basic"
)

// given a doc string and destination struct, look at the struct for any
// boolean fields tagged `Note`. If valid notes are found in the given string
// set the tagged booleans to true on the input struct.
func parseNotesFromDocComment(decl ast.Decl, file *os.File, dest any) error {
	re := regexp.MustCompile(`@(\w+)`)

	var identifier string
	var docstring string

	var docNotes []string
	var validNotes []string

	// check if decl is a function
	if funcDecl, ok := decl.(*ast.FuncDecl); ok {
		identifier = funcDecl.Name.Name
		docstring = funcDecl.Doc.Text()
	}

	matches := re.FindAllStringSubmatch(docstring, -1)

	for _, match := range matches {
		// match[1] contains the first capture group (the word after '@')
		docNotes = append(docNotes, match[1])
	}

	v := reflect.ValueOf(dest).Elem()
	t := v.Type()

	if v.Kind() != reflect.Struct {
		return fmt.Errorf("\ndestination struct must be a pointer to a struct")
	}

	for i := 0; i < v.NumField(); i++ {
		field := t.Field(i)

		// Look for the `Note` tag in the struct field
		if field.Tag == "note" {
			if Contains[string](docNotes, field.Name) {
				v.Field(i).SetBool(true)
			}

			validNotes = append(validNotes, field.Name)
		}
	}

	for _, v := range docNotes {
		if !Contains[string](validNotes, v) {
			return fmt.Errorf("\nUnknown note `@%s`, Identifier: `%s`, Location: `%s`\n\tValid values are: %v", v, identifier, file.Name(), validNotes)
		}
	}

	return nil
}

// helpers
func removeLastPart(s string) string {
	// Find the last occurrence of "/"
	lastSlashIndex := strings.LastIndex(s, "/")

	// If there is no "/" in the string, return the string itself
	if lastSlashIndex == -1 {
		return s
	}

	// Return everything before the last "/"
	return s[:lastSlashIndex]
}

var reset = "\033[0m"
var red = "\033[31m"
var green = "\033[32m"
var yellow = "\033[33m"
var blue = "\033[34m"
var magenta = "\033[35m"
var cyan = "\033[36m"
var gray = "\033[37m"
var white = "\033[97m"

func printStatus(b bool) {
	var status string

	if b {
		status = green + "SUCCESS" + reset
	} else {
		status = red + "FAILED" + reset
	}

	fmt.Printf("... %s", status)
}

func handleCmdOutput(out []byte, err error) {
	if err != nil {
		printStatus(false)
		println("")
		fmt.Printf("%s", out)
		os.Exit(1)
	} else {
		printStatus(true)
	}
}

func handleErr(err error) {
	if err != nil {
		printStatus(false)
		println(err.Error())
		os.Exit(1)
	}
}