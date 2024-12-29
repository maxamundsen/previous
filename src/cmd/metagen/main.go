package main

import (
	"log"
	"fmt"
	"os"
	"os/exec"
	"runtime"
)

var DEBUG bool
var OS string

// metagen - code generator application
//
// Generates code for other applications, such as the server application
func main() {
	// check operating system before running build command
	OS = runtime.GOOS

	println("Metagen - Code Generator and Build System")
	println("")


	// handle build arguments
	args := os.Args[1:]

	DEBUG = true
	for _, arg := range args {
		switch arg {
		case "release":
			DEBUG = false
		case "debug":
		default:
		}
	}

	compileTailwindCSS()

	// codegen
	setDebugConstant()
	generateHTTPRoutes()

	// actual compilation
	compileServer()
	compileMigrator()
}

func compileTailwindCSS() {
	tailwindcmd := ""

	fmt.Printf("Compiling TailwindCSS")

	if OS == "windows" {
		log.Fatal("OS unsupported. Compilation failed.")
	} else if OS == "darwin" {
		tailwindcmd = "tailwindcss-macos-arm64"
	}

	out, err := exec.Command("./" + tailwindcmd, "-i", "styles/global.css", "-o", "wwwroot/css/style.css", "--minify").CombinedOutput()
	if err != nil {
		printStatus(false)
		fmt.Printf("%s", out)
		os.Exit(1)
	} else {
		printStatus(true)
	}

	println("")
}

func compileServer() {
	var out []byte
	var err error

	if DEBUG {
		fmt.Printf("Compiling Server Binary (DEBUG MODE)")
		// include extra flags for the GC
		out, err = exec.Command("go", "build", `-gcflags=all="-N -l"`, "./cmd/server").CombinedOutput()
	} else {
		fmt.Printf("Compiling Server Binary (RELEASE MODE)")
		out, err = exec.Command("go", "build", "./cmd/server").CombinedOutput()
	}

	if err != nil {
		printStatus(false)
		println("")
		fmt.Printf("%s", out)
		os.Exit(1)
	} else {
		printStatus(true)
	}

	println("")
}

func compileMigrator() {

}

// based
func generateHTTPRoutes() {

}

// set debug constant inside the "config" package
func setDebugConstant() {
	fmt.Printf("Generating DEBUG Config")

	code := `package config

const (
`

	if DEBUG {
		code += `	DEBUG = true
`
	} else {
		code += `	DEBUG = false
`
	}

	code += `)
`

	// open file and write code to it

	in := []byte(code)

	err := os.WriteFile("./config/debug.go", in, 0644)
	if err != nil {
		printStatus(false)
		fmt.Println(err)
		os.Exit(1)
	} else {
		printStatus(true)
	}

	println("")
}

func printStatus(b bool) {
	var status string

	if b {
		status = "... SUCCESS"
	} else {
		status = "... FAILED"
	}

	fmt.Printf("%s", status)
}