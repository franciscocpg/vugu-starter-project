// +build ignore

package main

import (
	"flag"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
	"time"

	"github.com/gobuffalo/packr/v2/jam"
	"github.com/vugu/vugu/distutil"
	"github.com/vugu/vugu/simplehttp"
)

func main() {

	cleanPackr := flag.Bool("clean-packr", false, "Remove packr files")
	clean := flag.Bool("clean", false, "Remove dist dir before starting")
	dist := flag.String("dist", "dist", "Directory to put distribution files in")
	flag.Parse()

	if *cleanPackr {
		jam.Clean()
		return
	}

	start := time.Now()

	if *clean {
		os.RemoveAll(*dist)
	}

	os.MkdirAll(*dist, 0755) // create dist dir if not there

	// copy static files
	distutil.MustCopyDirFiltered(".", *dist, nil)

	// find and copy wasm_exec.js
	distutil.MustCopyFile(distutil.MustWasmExecJsPath(), filepath.Join(*dist, "wasm_exec.js"))

	// check for vugugen and go get if not there
	if _, err := exec.LookPath("vugugen"); err != nil {
		fmt.Sprint(distutil.MustExec("go", "get", "github.com/vugu/vugu/cmd/vugugen"))
	}

	// run go generate
	fmt.Sprint(distutil.MustExec("go", "generate", "."))

	jam.Clean()
	// run go build for wasm binary
	fmt.Sprint(distutil.MustEnvExec([]string{"GOOS=js", "GOARCH=wasm"}, "go", "build", "-o", filepath.Join(*dist, "main.wasm"), "."))

	// STATIC INDEX FILE:
	req, _ := http.NewRequest("GET", "/index.html", nil)
	outf, err := os.OpenFile(filepath.Join(*dist, "index.html"), os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0644)
	distutil.Must(err)
	defer outf.Close()
	template.Must(template.New("_page_").Parse(simplehttp.DefaultPageTemplateSource)).Execute(outf, map[string]interface{}{"Request": req})

	jam.Pack(jam.PackOptions{})

	// BUILD GO SERVER:
	fmt.Sprint(distutil.MustExec("go", "build", "-o", "bin/server", "."))

	log.Printf("dist.go complete in %v", time.Since(start))
}
