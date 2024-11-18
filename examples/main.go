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

	addExamples := func(suppliers ...func() (chartjs.IChartConfiguration, error)) error {
		for _, supplier := range suppliers {
			chartConfiguration, err := supplier()
			if err != nil {
				return err
			}
			examplesData.Charts = append(examplesData.Charts, chartConfiguration)
		}
		return nil
	}

	err := addExamples(
		line.Example1,
		line.Example2,
		line.Example3,
		line.Example4,
		line.Example5,
		line.Example6,
	)
	if err != nil {
		return err
	}

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
