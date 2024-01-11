package main

import (
	"fmt"

	"github.com/alecthomas/kong"

	"bytes"
	"io/ioutil"

	"github.com/yuin/goldmark"
	"github.com/yuin/goldmark/extension"
	"github.com/yuin/goldmark/parser"
	"github.com/yuin/goldmark/renderer/html"
	"github.com/yuin/goldmark/text"
)

type Context struct {
	Debug bool
}

type ConvertCmd struct {
	Paths []string `arg name:"path" help:"Paths to file or files to be converted." type:"path"`
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func (r *ConvertCmd) Run(ctx *Context) error {
	fmt.Println("convert", r.Paths)

	md := goldmark.New(
		goldmark.WithExtensions(extension.GFM),
		goldmark.WithParserOptions(
			parser.WithAutoHeadingID(),
		),
		goldmark.WithRendererOptions(
			html.WithHardWraps(),
			html.WithXHTML(),
		),
	)

	for _, path := range r.Paths {
		source, err := ioutil.ReadFile(path)
		check(err)
		var b bytes.Buffer

		doc := md.Parser().Parse(text.NewReader(source))
		// TODO Turn doc into json
		errRender := md.Renderer().Render(&b, source, doc)
		check(errRender)

		fmt.Println(b.String())

	}

	return nil
}

var cli struct {
	Debug bool `help:"Enable debug mode."`

	Convert ConvertCmd `cmd help:"Converts Markdown files to JSON."`
}

func main() {
	ctx := kong.Parse(&cli)
	// Call the Run() method of the selected parsed command.
	err := ctx.Run(&Context{Debug: cli.Debug})
	ctx.FatalIfErrorf(err)
}
