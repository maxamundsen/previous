package main

import (
	"fmt"
	"os"
	"path/filepath"
	"previous/basic"
	"previous/security"
	"regexp"
	"strings"
)

var dedupMap = make(map[string]bool)

// For each input directory, walk the filesystem tree for *.go files, and expand calls to `InlineStyle(..)`
func expandInlineStyles() {
	fmt.Printf("Compiling Inline Styles")

	inlineStyleRegex := regexp.MustCompile("InlineStyle\\((((?:'[^']*')|(?:\"[^\"]*\")|(`(?:[^`]|[\r\n])*?`)))\\)")

	// these are the directories that get scanned
	var dirs = [2]string{"handlers", "components"}

	var matches []string

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

				// only pre-process page & component files
				inlineStyleRegex.ReplaceAllStringFunc(s, func(match string) string {
					submatch := inlineStyleRegex.FindStringSubmatch(match)

					if len(submatch) > 1 {
						submatch[1] = strings.TrimPrefix(submatch[1], "\"")
						submatch[1] = strings.TrimPrefix(submatch[1], "'")
						submatch[1] = strings.TrimPrefix(submatch[1], "`")
						submatch[1] = strings.TrimSuffix(submatch[1], "\"")
						submatch[1] = strings.TrimSuffix(submatch[1], "'")
						submatch[1] = strings.TrimSuffix(submatch[1], "`")
						submatch[1] = strings.ReplaceAll(submatch[1], "\n", " ")
						submatch[1] = strings.ReplaceAll(submatch[1], "\t", "")

						rawInput := submatch[1] // this is the input before we process it

						// Skip duplicates
						_, found := dedupMap[rawInput]
						if !found {
							cssHash, _ := security.HighwayHash58(rawInput)
							cssHash = basic.GetFirstNChars(cssHash, 8)

							//@TODO handle custom shorthand media queries + alternate syntax here
							submatch[1] = expandMe(submatch[1], cssHash)
							submatch[1] = expandMedia(submatch[1])
							submatch[1] = transformSpacing(submatch[1])

							matches = append(matches, submatch[1])
							dedupMap[rawInput] = true
						}
					}

					// We aren't actually replacing anything in the src file, we just needed to iterate over the regex matches
					return ""
				})
			}

			return nil
		})

		handleErr(err)
	}

	outputCSS := "/* " + METAGEN_AUTO_COMMENT + " */\n"
	outputCSS += strings.Join(matches, "")
	os.WriteFile("./wwwroot/css/style.metagen.css", []byte(outputCSS), 0664)

	printStatus(true)
}

// Expand spacing macro
// Ex:
//
//	padding: $(5);
//
// => padding: calc(var(--spacing) * 5);
func transformSpacing(input string) string {
	re := regexp.MustCompile(`\$\((\d+)\)`)

	transformed := re.ReplaceAllStringFunc(input, func(match string) string {
		number := re.FindStringSubmatch(match)[1]
		return fmt.Sprintf("calc(var(--spacing) * %s)", number)
	})

	return transformed
}

// Expand shorthand media queries.
// Ex:
//
//	media md { ... }
//
// => media screen and (min-width: 768px) { ... }
func expandMedia(input string) string {
	re := regexp.MustCompile(`(?:@media)\s(xs-|sm-|md-|lg-|xl-|sm|md|lg|xl|xx|dark|light)`)

	transformed := re.ReplaceAllStringFunc(input, func(match string) string {
		mq := ""

		if len(match) >= 2 {
			switch re.FindStringSubmatch(match)[1] {
			case "dark":
				mq = "(prefers-color-scheme: dark)"
			case "light":
				mq = "(prefers-color-scheme: light)"
			case "xs-":
				mq = "screen and (max-width: 639px)"
			case "sm-":
				mq = "screen and (max-width: 767px)"
			case "md-":
				mq = "screen and (max-width: 1023px)"
			case "lg-":
				mq = "screen and (max-width: 1279px)"
			case "xl-":
				mq = "screen and (max-width: 1535px)"
			case "sm":
				mq = "screen and (min-width: 640px)"
			case "md":
				mq = "screen and (min-width: 768px)"
			case "lg":
				mq = "screen and (min-width: 1024px)"
			case "xl":
				mq = "screen and (min-width: 1280px)"
			case "xx":
				mq = "screen and (min-width: 1536px)"
			}
		}

		return fmt.Sprintf("@media %s", mq)
	})

	return transformed
}

// Replace "me" selector with inline style attribute
// Ex:
//
//	me { ... }
//
// => [__inlinecss_{REPLACEMENT_ID}] { ... }
func expandMe(input string, replacementId string) string {
	re := regexp.MustCompile(`(?:^|\.|\s|[^a-zA-Z0-9\-\_{}()\[\]\<\>])me\b`)

	transformed := re.ReplaceAllStringFunc(input, func(match string) string {
		return fmt.Sprintf("[__inlinecss_%s]", replacementId)
	})

	return transformed
}
