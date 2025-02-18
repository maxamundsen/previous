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
var dedupMap = make(map[string]int)

// For each input directory, walk the filesystem tree for *.go files, and expand calls to `InlineStyle(..)`
func expandInlineStyles() {
	module_name := getCurrentModuleName()

	inlineStyleRegex := regexp.MustCompile("InlineStyle\\(((?:\"|'|`).*?(?:\"|'|`))\\)")

	var dirs = [2]string{"pages", "components"}
	var matches []string

	count := 0

	for _, dir := range dirs {
		err := filepath.Walk(dir, func(pathStr string, info os.FileInfo, err error) error {
			if strings.HasSuffix(info.Name(), ".go") {
				pathStr = filepath.ToSlash(pathStr)
				handleErr(err)

				file, err := os.Open(pathStr)
				handleErr(err)

				defer file.Close()

				content, contentErr := os.ReadFile(pathStr)
				if contentErr != nil {
					return contentErr
				}

				s := string(content)
				var output string

				if strings.HasSuffix(info.Name(), ".page.go") || strings.HasSuffix(info.Name(), ".component.go") {
					// only pre-process page & component files
					output = inlineStyleRegex.ReplaceAllStringFunc(s, func(match string) string {
						submatch := inlineStyleRegex.FindStringSubmatch(match)
						count += 1

						if len(submatch) > 1 {
							submatch[1] = strings.TrimPrefix(submatch[1], "\"")
							submatch[1] = strings.TrimPrefix(submatch[1], "'")
							submatch[1] = strings.TrimPrefix(submatch[1], "`")
							submatch[1] = strings.TrimSuffix(submatch[1], "\"")
							submatch[1] = strings.TrimSuffix(submatch[1], "'")
							submatch[1] = strings.TrimSuffix(submatch[1], "`")

							//@TODO handle custom shorthand media queries + alternate syntax here
							submatch[1] = expandThis(submatch[1], count)
							submatch[1] = expandMedia(submatch[1])
							submatch[1] = transformSpacing(submatch[1])

							cssOutput := submatch[1]

							matches = append(matches, cssOutput)
						}

						// minify
						// input = strings.ReplaceAll(input, "\n", "")
						// input = strings.ReplaceAll(input, "\t", "")
						return fmt.Sprintf("Attr(\"__INLINECSS_%d\")", count)
					})

					output = strings.Replace(output, fmt.Sprintf(`"%s/pages`, module_name), fmt.Sprintf(`"%s/.metagen/pages`, module_name), -1)
					output = strings.Replace(output, fmt.Sprintf(`"%s/components`, module_name), fmt.Sprintf(`"%s/.metagen/components`, module_name), -1)

				} else {
					// we still want to copy other go files, but don't preprocess them.
					output = s
				}

				output = METAGEN_AUTO_COMMENT + "\n" + output

				mkdirErr := os.MkdirAll("./.metagen/"+removeLastPart(pathStr)+"/", 0755)
				if mkdirErr != nil {
					return mkdirErr
				}

				writeErr := os.WriteFile(".metagen/"+pathStr, []byte(output), 0664)
				if writeErr != nil {
					return writeErr
				}
			}

			return nil
		})

		handleErr(err)
	}

	outputCSS := "/* " + METAGEN_AUTO_COMMENT + " */\n"
	outputCSS += strings.Join(matches, "")
	os.WriteFile("./wwwroot/css/style.metagen.css", []byte(outputCSS), 0664)
}

func transformSpacing(input string) string {
	re := regexp.MustCompile(`\$\((\d+)\)`)

	transformed := re.ReplaceAllStringFunc(input, func(match string) string {
		number := re.FindStringSubmatch(match)[1]
		return fmt.Sprintf("calc(var(--spacing) * %s)", number)
	})

	return transformed
}

func expandMedia(input string) string {
	input = strings.ReplaceAll(input, "$dark", "(prefers-color-scheme: dark)")
	input = strings.ReplaceAll(input, "$light", "(prefers-color-scheme: light)")
	input = strings.ReplaceAll(input, "$xs-", "screen and (max-width: 639px)")
	input = strings.ReplaceAll(input, "$sm-", "screen and (max-width: 767px)")
	input = strings.ReplaceAll(input, "$md-", "screen and (max-width: 1023px)")
	input = strings.ReplaceAll(input, "$lg-", "screen and (max-width: 1279px)")
	input = strings.ReplaceAll(input, "$xl-", "screen and (max-width: 1535px)")
	input = strings.ReplaceAll(input, "$sm", "screen and (min-width: 640px)")
	input = strings.ReplaceAll(input, "$md", "screen and (min-width: 768px)")
	input = strings.ReplaceAll(input, "$lg", "screen and (min-width: 1024px)")
	input = strings.ReplaceAll(input, "$xl", "screen and (min-width: 1280px)")
	input = strings.ReplaceAll(input, "$xx", "screen and (min-width: 1536px)")
	return input
}

func expandThis(input string, id int) string {
	input = strings.ReplaceAll(input, "$this", fmt.Sprintf("[__INLINECSS_%d]", id))

	return input
}