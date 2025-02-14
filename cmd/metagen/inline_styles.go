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
var dedupNum = 1

// For each input directory, walk the filesystem tree for *.go files, and expand calls to `InlineStyle(..)`
func expandInlineStyles(dirs []string) {
	// module_name := getCurrentModuleName()

	inlineStyleRegex := regexp.MustCompile(`InlineStyle\((.*?)\)`)

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

				var b []byte

				file.Read(b)

				s := string(b)

				output := inlineStyleRegex.ReplaceAllStringFunc(s, func(match string) string {
					submatch := inlineStyleRegex.FindStringSubmatch(match)
					if len(submatch) > 1 {
						matches = append(matches, submatch[1])
					}

					// minify
					// input = strings.ReplaceAll(input, "\n", "")
					// input = strings.ReplaceAll(input, "\t", "")

					return fmt.Sprintf("Attr(\"__INLINECSS_%d\")", dedupNum)
				})

				fmt.Printf("%s", output)

			}

			return nil
		})

		handleErr(err)
	}
}
