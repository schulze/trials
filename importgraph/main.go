// The program importgraph builds a graph of the packages imported by packages in sub-directories.
package main

import (
	"flag"
	"os"
	"os/exec"
	"log"
	"path/filepath"
	"go/build"
	"fmt"
	"io"
)

// TODO: 
// - different colors (e.g. for GOPATH, GOROOT)
// - flag to ignore standard packages (like everything in GOROOT)

var (
	baseDir =flag.String("base", "", "base directory")
	justDot = flag.Bool("dot", false, "output graph in Graphviz DOT format")
)

// packages is the global list of packages
var packages []*build.Package

func main() {
	flag.Parse()
	if *baseDir == "" {
		wd, err := os.Getwd()
		if err != nil {
			log.Fatal(err)
		}
		*baseDir = wd
	}

	filepath.Walk(*baseDir, packageFilter)

	// we could do some processing on the database now

	if *justDot {
		printGraph(os.Stdout, packages)
		return
	}
	
	cmd := exec.Command("dot", "-Tsvg")
	cmdErr, err := cmd.StderrPipe()
	if err != nil {
		log.Fatal(err)
	}
	cmdIn, err := cmd.StdinPipe()
	if err != nil {
		log.Fatal(err)
	}
	cmdOut, err := cmd.StdoutPipe()
	if err != nil {
		log.Fatal(err)
	}
	err = cmd.Start()
	if err != nil {
		log.Fatal(err)
	}

	printGraph(cmdIn, packages)
	cmdIn.Close()
	io.Copy(os.Stdout, cmdOut)
	io.Copy(os.Stderr, cmdErr)
}

func packageFilter(path string, info os.FileInfo, err error) error {
	fi, err := os.Stat(path)
	if err != nil {
		log.Fatal(err)
	}
	if fi.IsDir() {
		readPackage(path)
	}
	return nil
}

func readPackage(dir string) {
	p, err := build.ImportDir(dir, 0)
	if err != nil {  // we go on and just ignore this directory
		log.Print(err)
		return
	}
	packages = append(packages, p)
}

func printGraph(dst io.Writer, packages []*build.Package) {
	fmt.Fprint(dst, dotHeader)
	for _, p := range packages {
		packagePath := p.ImportPath
		for _, importPath  := range p.Imports {
			fmt.Fprintf(dst, "   \"%s\" -> \"%s\"\n", packagePath, importPath)
		}
	}
	fmt.Fprint(dst, dotFooter)
}

const dotHeader = `
digraph Imports {
   size="12.8"
   overlap=false
   ratio=fill
   rankdir=LR
`
const dotFooter = `
}
`
