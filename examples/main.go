package main

import (
	"github.com/ngyewch/go-chartjs"
	"github.com/ngyewch/go-chartjs/examples/line"
	"github.com/ngyewch/go-chartjs/examples/resources"
	"github.com/urfave/cli/v2"
	"html/template"
	"log"
	"os"
)

var (
	app = &cli.App{
		Name:   "go-chartjs-examples",
		Usage:  "go-chartjs examples",
		Action: doMain,
	}
)

func main() {
	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}

type ExamplesData struct {
	Charts []chartjs.IChartConfiguration
}

func doMain(context *cli.Context) error {
	var examplesData ExamplesData

	line1, err := line.Example1()
	if err != nil {
		return err
	}
	examplesData.Charts = append(examplesData.Charts, line1)

	line2, err := line.Example2()
	if err != nil {
		return err
	}
	examplesData.Charts = append(examplesData.Charts, line2)

	templates, err := template.New("templates").
		ParseFS(resources.TemplateFS, "templates/*.gohtml")
	if err != nil {
		return err
	}

	f, err := os.Create("examples.html")
	if err != nil {
		return err
	}
	defer func(f *os.File) {
		_ = f.Close()
	}(f)
	err = templates.ExecuteTemplate(f, "examples.gohtml", examplesData)
	if err != nil {
		return err
	}

	return nil
}
