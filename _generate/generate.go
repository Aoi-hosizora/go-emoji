package main

import (
	"bytes"
	"flag"
	"go/format"
	"log"
	"os"
	"text/template"
)

var (
	fUsage    = flag.Bool("h", false, "show help")
	fPackage  = flag.String("pkg", "emoji", "package name")
	fFilename = flag.String("o", "../emoji.go", "output filename")
)

type templateData struct {
	Package string
	List    []*emojiData
}

const codeTemplate = `
package {{ .Package }}

//go:generate ./_generate/generate.out -pkg "emoji" -o "emoji.go"

// THIS FILE WAS GENERATED BY github.com/Aoi-hosizora/go-emoji/_generate.
// FOR DETAILS, PLEASE VISIT https://unicode.org/emoji/charts/emoji-list.html.
// DO NOT EDIT THIS FILE!

const (
	{{ range $idx, $val := .List }}
	// {{ $val.Var }} represents "{{ $val.Name }}" ({{ $val.Unicode }}) in "{{ $val.Subgroup }}" from "{{ $val.Group }}".
	// Other keywords: {{ $val.Keyword }}.
	{{ $val.Var }} = "{{ $val.Unicode }}"
	{{ end }}
)
`

func main() {
	flag.Parse()
	if *fUsage {
		flag.Usage()
		return
	}

	// http get
	list, err := getEmojiList()
	if err != nil {
		log.Fatalln("Failed to get emoji list:", err)
	}

	// template
	t, err := template.New("template").Parse(codeTemplate)
	if err != nil {
		log.Fatalln("Failed to parse template:", err)
	}

	buf := &bytes.Buffer{}
	err = t.Execute(buf, &templateData{
		Package: *fPackage,
		List:    list,
	})
	if err != nil {
		log.Fatalln("Failed to execute template:", err)
	}

	// fmt
	bs, err := format.Source(buf.Bytes())
	if err != nil {
		log.Fatalln("Failed to format generated code:", err)
	}

	// file
	_ = os.Remove(*fFilename)
	file, err := os.Create(*fFilename)
	if err != nil {
		log.Fatalln("Failed to create code file:", err)
	}
	defer file.Close()

	_, err = file.Write(bs)
	if err != nil {
		log.Fatalln("Failed to write to code file:", err)
	}
}
