package main

import (
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

// used to deduplicate css output
// string key is NOT the css, but an encoded []string by converting to []byte, which can be converted to a string
// this is because dynamic arrays (slices) cannot be used as map keys, but strings can
var dedupMap map[string]bool
var dedupNum = 0

// For each input directory, walk the filesystem tree for *.go files, and expand calls to `InlineStyle(..)`
func expandInlineStyles(dirs ...string) {
	// module_name := getCurrentModuleName()

	inlineStyleRegex := regexp.MustCompile("InlineStyle\\(((?:\"|'|`).*?(?:\"|'|`))\\)")

	var matches []string

	for _, dir := range dirs {
		err := filepath.Walk(dir, func(pathStr string, info os.FileInfo, err error) error {
			pathStr = filepath.ToSlash(pathStr)
			handleErr(err)

			// only process page & component files
			if strings.HasSuffix(info.Name(), ".page.go") || strings.HasSuffix(info.Name(), ".component.go") {
				file, err := os.Open(pathStr)
				handleErr(err)

				defer file.Close()

				content, contentErr := os.ReadFile(pathStr)
				if contentErr != nil {
					return contentErr
				}

				s := string(content)

				output := inlineStyleRegex.ReplaceAllStringFunc(s, func(match string) string {
					submatch := inlineStyleRegex.FindStringSubmatch(match)
					if len(submatch) > 1 {
						matches = append(matches, submatch[1])
					}

					// minify
					// input = strings.ReplaceAll(input, "\n", "")
					// input = strings.ReplaceAll(input, "\t", "")

					dedupNum += 1

					return fmt.Sprintf("Attr(\"__INLINECSS_%d\")", dedupNum)
				})

				mkdirErr := os.MkdirAll("./.metagen/"+ removeLastPart(pathStr) +"/", 0755)
				if mkdirErr != nil {
					return mkdirErr
				}

				output = METAGEN_AUTO_COMMENT + "\n" + output

				writeErr := os.WriteFile("./.metagen/" + pathStr, []byte(output), 0664)
				if writeErr != nil {
					return writeErr
				}
			}

			return nil
		})

		handleErr(err)
	}

	for i, _ := range matches {
		matches[i] = strings.TrimPrefix(matches[i], "\"")
		matches[i] = strings.TrimPrefix(matches[i], "'")
		matches[i] = strings.TrimPrefix(matches[i], "`")
		matches[i] = strings.TrimSuffix(matches[i], "\"")
		matches[i] = strings.TrimSuffix(matches[i], "'")
		matches[i] = strings.TrimSuffix(matches[i], "`")

		//@TODO handle custom shorthand media queries + alternate syntax here

		matches[i] = fmt.Sprintf("[__INLINECSS_%d] { %s }\n", i + 1, matches[i])
	}

	outputCSS := "/* " +  METAGEN_AUTO_COMMENT + " */\n"
	outputCSS += strings.Join(matches, "")
	os.WriteFile("./wwwroot/css/output.css", []byte(outputCSS), 0664)
}
