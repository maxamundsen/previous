package main

import (
	"errors"
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"os"
	"os/exec"
	"path"
	"path/filepath"
	. "previous/basic"
	"runtime/debug"
	"strings"
)

// Generates routes from files named `*.page.go` found recursively inside the `/src/pages` directory.
// When parsing files, we search for the first function suffixed with `Page`, if one is not found, return an error and fail compilation.
type RouteInfo struct {
	FileDef  string
	URL      string
	PageName string
	Package  string
	Import   string

	Identity      bool `note:"true"`
	Protected     bool `note:"true"`
	CookieSession bool `note:"true"`
	EnableCors    bool `note:"true"`

	// Used to determine if page should be pre-rendered at compile time.
	Static    bool `note:"true"`
	StaticAPI bool `note:"true"`

	// http verbs
	HttpPost   bool `note:"true"`
	HttpGet    bool `note:"true"`
	HttpPut    bool `note:"true"`
	HttpPatch  bool `note:"true"`
	HttpDelete bool `note:"true"`
}

func generatePageData() {
	const root = "pages" // Use the /pages directory for autogenerating routes

	bi, _ := debug.ReadBuildInfo()
	parts := strings.Split(bi.Path, "/")
	module_name := parts[0]

	fmt.Printf("Generating HTTP routes")

	var routeList []RouteInfo

	// Walk the filesystem recursively
	err := filepath.Walk(root, func(pathStr string, info os.FileInfo, err error) error {
		pathStr = filepath.ToSlash(pathStr)

		handleErr(err)

		// remove old metagen files first
		if strings.HasSuffix(info.Name(), ".page.metagen.go") {
			err := os.Remove(pathStr)
			if err != nil {
				return fmt.Errorf("error deleting file %s: %v", pathStr, err)
			}
		}

		// Only process .go files
		if strings.HasSuffix(info.Name(), ".page.go") {
			file, err := os.Open(pathStr)
			handleErr(err)

			defer file.Close()

			fs := token.NewFileSet()
			node, err := parser.ParseFile(fs, pathStr, file, parser.ParseComments)
			handleErr(err)

			found := false

			for _, decl := range node.Decls {
				if found {
					break
				}

				// Look for function declarations
				if funcDecl, ok := decl.(*ast.FuncDecl); ok {
					ri := RouteInfo{}

					if funcDecl, ok := decl.(*ast.FuncDecl); ok {
						identifier := funcDecl.Name.Name
						// page functions MUST end with the word "Page" (case sensitive)
						if strings.HasSuffix(identifier, "Page") {
							found = true
						} else {
							continue
						}
					}

					handleErr(parseNotesFromDocComment(decl, file, &ri))

					relativePath := strings.TrimPrefix(pathStr, root)
					relativePath = strings.TrimSuffix(relativePath, ".page.go")

					dirPath := filepath.Dir(relativePath)
					lastDir := filepath.Base(dirPath)

					// Combine last directory and function name
					// Also replace underscore characters "_" with hyphen characters "-"
					ri.URL = strings.ReplaceAll(relativePath, "_", "-")
					ri.Package = lastDir
					ri.PageName = funcDecl.Name.Name

					ri.FileDef = file.Name()

					// strip names from route and use the base package name if file is index
					if strings.HasSuffix(pathStr, "/index.page.go") {
						ri.URL = path.Dir(ri.URL)
					}

					// if the route is in the "root" folder, make sure it imports the correct package.
					ri.Package = filepath.ToSlash(ri.Package)

					if ri.Package == "/" {
						ri.Package = "pages"
						ri.Import = module_name + "/" + ri.Package
						ri.PageName = strings.ReplaceAll(ri.PageName, "/", "pages")
					} else {
						ri.Import = module_name + "/" + root + "/" + removeLastPart(strings.TrimPrefix(relativePath, "/"))
					}

					routeList = append(routeList, ri)
				}
			}

			if !found {
				return fmt.Errorf("\n`%s`: Attempted to generate route, but no function suffixed with `Page` found.", file.Name())
			}
		}

		return nil
	})

	handleErr(err)

	// -- GENERATE PAGEINFO STRUCTS --
	// generate recursive structs representing pages
	// this is used in order to reference a page without needing to actually write out the link as a string literal
	// it also lets you jump to page code whenever that page is referenced in view links or something like that.
	// essentially making dead page links a compile time error if you use this structure.

	newpath := filepath.Join(".", ".metagen/pageinfo")
	dirErr := os.MkdirAll(newpath, os.ModePerm)
	handleErr(dirErr)

	structCode := METAGEN_AUTO_COMMENT
	structCode += "\npackage pageinfo\n\n"

	structCode += "import \"net/http\"\n"

	structCode += "type middleware struct {\n"
	structCode += "\tIdentity      bool\n"
	structCode += "\tProtected     bool\n"
	structCode += "\tCookieSession bool\n"
	structCode += "\tEnableCors    bool\n"
	structCode += "}\n\n"

	structCode += "type PageInfo struct {\n"
	structCode += "\tstatic         bool\n"
	structCode += "\turl            string\n"
	structCode += "\tfileDef        string\n"
	structCode += "\tmiddleware middleware\n"
	structCode += "}\n\n"

	structCode += "func (info PageInfo) IsStatic() bool {\n"
	structCode += "\treturn info.static\n"
	structCode += "}\n\n"

	structCode += "func (info PageInfo) Url() string {\n"
	structCode += "\treturn info.url\n"
	structCode += "}\n\n"

	structCode += "func (info PageInfo) FileDef() string {\n"
	structCode += "\treturn info.fileDef\n"
	structCode += "}\n\n"

	structCode += "func (info PageInfo) Middleware() middleware {\n"
	structCode += "\treturn info.middleware\n"
	structCode += "}\n\n"

	structCode += "func GetPageInfoMap() map[string]PageInfo {\n"
	structCode += "\tnewMap := make(map[string]PageInfo)\n\n"
	structCode += "\tfor k, v := range pageInfoMap {\n"
	structCode += "\t\tnewMap[k] = v\n"
	structCode += "\t}\n\n"
	structCode += "\treturn newMap\n"
	structCode += "}\n\n"

	structCode += "func Reflect(r *http.Request) PageInfo {\n"
	structCode += "\treturn pageInfoMap[r.URL.Path]\n"
	structCode += "}\n\n"

	structCode += "var (\n"
	structCode += "\tpageInfoMap map[string]PageInfo //maps URLs to PageInfo\n"
	structCode += ")\n\n"

	var pageTree Tree

	pageTree.Name = "Root"

	for i, route := range routeList {
		// If you are staticapi, you are also static
		if route.StaticAPI {
			routeList[i].Static = true
		}

		if route.PageName == "IndexPage" && route.Import != module_name+"/pages" {
			parts := GetPathParts(route.URL)
			parts = append(parts, "index")

			AddStringPartsToTree(&pageTree, parts)
		} else {
			AddStringPartsToTree(&pageTree, GetPathParts(route.URL))
		}
	}

	generateRecursivePageInfoStructs(&structCode, &pageTree, 0)

	structCode += "\n\n"
	structCode += "func init() {\n"
	structCode += "\tpageInfoMap = make(map[string]PageInfo)\n\n"

	for _, v := range routeList {
		structExpansion := strings.ReplaceAll(strings.ReplaceAll(strings.TrimPrefix(v.URL, "/"), "/", "."), "-", "_")
		parts := strings.Split(structExpansion, ".")

		for i, _ := range parts {
			parts[i] = CapitalizeFirstLetter(parts[i])
		}

		structExpansion = strings.Join(parts, ".")

		if structExpansion == "" {
			structExpansion = "Index"
		} else if v.PageName == "IndexPage" {
			structExpansion += ".Index"
		}

		// initialize each pageinfo struct
		structCode += fmt.Sprintf("\tRoot.%s.url = \"%s\"\n", structExpansion, v.URL)
		structCode += fmt.Sprintf("\tRoot.%s.static = %t\n", structExpansion, v.Static)
		structCode += fmt.Sprintf("\tRoot.%s.fileDef = \"/%s\"\n", structExpansion, v.FileDef)
		structCode += fmt.Sprintf(
			"\tRoot.%s.middleware = middleware{\n\t\tIdentity: %t,\n\t\tProtected: %t,\n\t\tCookieSession: %t,\n\t\tEnableCors: %t,\n\t}\n",
			structExpansion,
			v.Identity,
			v.Protected,
			v.CookieSession,
			v.EnableCors,
		)

		// add to pageinfomap *after* providing the value (duh)
		structCode += fmt.Sprintf("\tpageInfoMap[\"%s\"] = Root.%s\n\n", v.URL, structExpansion)
	}

	structCode += "}\n"

	structCode_b := []byte(structCode)

	structFileErr := os.WriteFile("./.metagen/pageinfo/pageinfo.metagen.go", structCode_b, 0644)
	handleErr(structFileErr)

	// -- GENERATE ROUTE FILE --
	routeCode := METAGEN_AUTO_COMMENT + "\npackage main\n"

	routeCode += "\nimport (\n\t\"net/http\"\n"

	for _, v := range routeList {
		if v.Identity || v.CookieSession || v.EnableCors {
			routeCode += "	. \"" + module_name + "/middleware\"\n"
			break
		}
	}

	// the go import system is horrible, so the following is a package auto-importer/ de-duplicator for code generation.
	// search through all routes, figure out their package based on the import, and possibly rename it if there already exists an import in
	// the same namespace.

	packageMap := make(map[string][]RouteInfo)
	seenImport := make(map[string]bool)

	// group by package
	for _, route := range routeList {
		// entry de-duplication
		if !seenImport[route.Import] {
			packageMap[route.Package] = append(packageMap[route.Package], route)
		}

		seenImport[route.Import] = true
	}

	// iterate over each sub-group
	for _, routes := range packageMap {
		// don't give the first instance a named import, only subsequent entries
		if len(routes) > 0 {
			routes[0].Import = fmt.Sprintf("\"%s\"", routes[0].Import)
		}

		if len(routes) > 1 {
			// modify the package name in the routeList to include the index of the sublist
			// (this is the whole point -- automatically giving duplicate packages named imports)
			for i := 1; i < len(routes); i += 1 {
				for j, v := range routeList {
					if v.Import == routes[i].Import {
						routeList[j].Package = fmt.Sprintf("%s%d", routes[i].Package, i)
					}
				}

				routes[i].Import = fmt.Sprintf("%s%d \"%s\"", routes[i].Package, i, routes[i].Import)
			}
		}
	}

	var result []RouteInfo
	for _, routes := range packageMap {
		result = append(result, routes...)
	}

	for i := range result {
		routeCode += fmt.Sprintf("\t%s\n", result[i].Import)
	}

	routeCode += ")\n"
	routeCode += "\nfunc mapAutoRoutes(mux *http.ServeMux) {\n"

	for _, routeInfo := range routeList {
		printablePage := routeInfo.Package + "." + routeInfo.PageName

		if routeInfo.Static || routeInfo.StaticAPI {
			staticErr := generateStaticPage(module_name, routeInfo)
			if staticErr != nil {
				fmt.Println(staticErr.Error())
			}

			printablePage += "_STATIC"
		}

		if routeInfo.CookieSession {
			printablePage = fmt.Sprintf("LoadSessionFromCookie(%s)", printablePage)
		}

		if routeInfo.Identity {
			printablePage = fmt.Sprintf("LoadIdentity(%s, %t)", printablePage, routeInfo.Protected)
		}

		if routeInfo.EnableCors {
			printablePage = fmt.Sprintf("EnableCors(%s)", printablePage)
		}

		httpVerb := ""
		if routeInfo.HttpGet {
			httpVerb = "GET "
		} else if routeInfo.HttpPost {
			httpVerb = "POST "
		} else if routeInfo.HttpPut {
			httpVerb = "PUT "
		} else if routeInfo.HttpPatch {
			httpVerb = "PATCH "
		} else if routeInfo.HttpDelete {
			httpVerb = "DELETE "
		}

		routeCode += fmt.Sprintf("\tmux.HandleFunc(\"%s%s\", %s)\n", httpVerb, routeInfo.URL, printablePage)
	}

	routeCode += "}"

	in := []byte(routeCode)

	fileErr := os.WriteFile("./cmd/server/generated_routes.metagen.go", in, 0644)
	handleErr(fileErr)

	os.RemoveAll("./cmd/metagen/.staticgen")

	printStatus(true)
}

// If @Static tag is set on page handler, compile the page and generate a new handler
// that serves the compiled output. This is similar to `#run` in Jai, where you can
// run arbitrary code at compile time.
//
// Unfortunately, Go does not have the ability to execute itself inside the compiler,
// so we must resort to generating and running a Go program for each page we want to prerender.
// After this generated program runs, the output is captured, and used to generate a static page handler,
// which is injected back into your program, and mapped correctly in the generated routes file.
//
// Honestly, this is pretty far out, since `metagen` is now acting as a second order metaprogram
// (a program that generates a program that generates a program),
// which is quite difficult to reason about. This is probably as far as we would like to push the metaprogramming stuff.
//
// This type of system increases compile time complexity, for the sake of runtime performance gains.
// Pages that take advantage of @Static are now just dumb HTML pages that contain NO server-side logic.
//
// -mta
// @todo: fix this - this waaaay slower than you thought :D
func generateStaticPage(module_name string, ri RouteInfo) error {
	if !ri.Static && !ri.StaticAPI {
		return errors.New("attmept to generate static page from non-static RouteInfo")
	}

	pageController := ri.Package + "." + ri.PageName

	metacode := METAGEN_AUTO_COMMENT
	metacode += "\npackage main\n\n"

	metacode += "import (\n"
	metacode += "\t\"fmt\"\n"
	metacode += "\t\"net/http\"\n"
	metacode += "\t\"net/http/httptest\"\n"
	metacode += fmt.Sprintf("\t. \"%s/middleware\"\n", module_name)
	metacode += fmt.Sprintf("\t. \"%s/preload\"\n", module_name)
	metacode += fmt.Sprintf("\t\"%s\"\n", ri.Import)
	metacode += ")\n\n"

	metacode += "func main() {\n"
	metacode += "\tPreload(PreloadOptions{\n"
	metacode += "\t\tShouldInitDatabase: true,\n"
	metacode += "\t})\n"
	metacode += "\tmux := http.NewServeMux()\n"
	metacode += fmt.Sprintf("\tmux.HandleFunc(\"%s\", LoadIdentity(LoadSessionFromCookie(%s), false))\n", ri.URL, pageController)
	metacode += fmt.Sprintf("\treq, _ := http.NewRequest(\"GET\", \"%s\", nil)\n", ri.URL)
	metacode += "\trr := httptest.NewRecorder()\n"
	metacode += "\tmux.ServeHTTP(rr, req)\n"
	metacode += "\tfmt.Printf(\"%s\", rr.Body.String())\n"
	metacode += "}\n"

	os.Mkdir("./cmd/metagen/.staticgen", 0755)

	err := os.WriteFile("./cmd/metagen/.staticgen/staticgen.go", []byte(metacode), 0644)
	if err != nil {
		return err
	}

	out, cmdErr := exec.Command("go", "run", "./cmd/metagen/.staticgen/staticgen.go").CombinedOutput()
	if cmdErr != nil {
		return err
	}

	newFilePath := strings.TrimSuffix(ri.FileDef, ".page.go")
	newFilePath += ".page.metagen.go"

	staticPageCode := METAGEN_AUTO_COMMENT + "\n"
	staticPageCode += fmt.Sprintf("package %s\n\n", ri.Package)
	staticPageCode += "import \"net/http\"\n\n"
	staticPageCode += fmt.Sprintf("func %s_STATIC(w http.ResponseWriter, r *http.Request) {\n", ri.PageName)

	if ri.StaticAPI {
		staticPageCode += "\tw.Header().Set(\"Content-Type\", \"application/json\")\n"
	} else {
		staticPageCode += "\tw.Header().Set(\"Content-Type\", \"text/html\")\n"
	}

	staticPageCode += "\tw.WriteHeader(http.StatusOK)\n"
	staticPageCode += fmt.Sprintf("\tw.Write([]byte(`%s`))\n", out)
	staticPageCode += "}\n"

	newFileErr := os.WriteFile(newFilePath, []byte(staticPageCode), 0644)
	if newFileErr != nil {
		return newFileErr
	}

	return nil
}

func generateRecursivePageInfoStructs(code *string, tree *Tree, level int) {
	if code == nil {
		return
	}

	if tree == nil {
		return
	}

	// Print the current node with indentation.
	sanitizedName := strings.ReplaceAll(CapitalizeFirstLetter(tree.Name), "-", "_")

	if tree.Children != nil {
		printVar := ""
		if level == 0 {
			printVar = "var "
		}

		*code += fmt.Sprintf("%s%s%s struct {\n", strings.Repeat("\t", level), printVar, sanitizedName)
	} else {
		*code += fmt.Sprintf("%s%s PageInfo\n", strings.Repeat("\t", level), sanitizedName)
	}

	// Recursively print each child.
	if tree.Children != nil {
		for _, child := range *tree.Children {
			generateRecursivePageInfoStructs(code, &child, level+1)
		}
		*code += strings.Repeat("\t", level) + "}\n"
	}

}
