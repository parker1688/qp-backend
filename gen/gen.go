package main

import (
	"bytes"
	"fmt"
	"github.com/gobuffalo/packr/v2"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"text/template"
)

func createPkg(p *ModulesData, err error) {
	box := packr.New("all", p.Path)
	if err = os.MkdirAll(genDir, 0755); err != nil {
		return
	}
	for _, name := range box.List() {
		if name == "go.mod.tmpl" {
			continue
		}
		tmpl, _ := box.FindString(name)
		dir := strings.ReplaceAll(filepath.Join(genDir, filepath.Dir(name)), "{{packageName}}", p.PackageName)
		if err = os.MkdirAll(dir, 0755); err != nil {
			continue
		}
		var saveFileName string
		if strings.Index(name, ".go") > -1 {
			saveFileName = p.PackageName + ".go"
		}
		if strings.Index(name, ".jsx") > -1 {
			saveFileName = "index.jsx"
		}
		if err = write(filepath.Join(dir, saveFileName), tmpl, p); err != nil {
			continue
		}
	}
}

func write(path, tpl string, data interface{}) (err error) {
	outParse, err := parse(tpl, data)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	routParse := strings.Replace(string(outParse), "{@@@@@{", "{{", -1)

	err = ioutil.WriteFile(path, []byte(routParse), 0644)
	cmd := exec.Command("gofmt", "-w", path)
	cmd.Run()
	return err
}

func parse(s string, data interface{}) ([]byte, error) {
	t, err := template.New("").Parse(s)
	if err != nil {
		return nil, err
	}
	var buf bytes.Buffer
	if err = t.Execute(&buf, data); err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}
