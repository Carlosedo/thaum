package main

import (
	"bufio"
	"os"
	"strings"

	constants "github.com/Flaque/thaum/thaum/constants"
	files "github.com/Flaque/thaum/thaum/files"
	output "github.com/Flaque/thaum/thaum/output"
	"github.com/urfave/cli"
)

var overwrite bool

func askForVariables(template files.Template) files.Template {
	reader := bufio.NewReader(os.Stdin)

	output.Space()
	for v := range template.Variables {
		output.VariableLabel(v)
		value, _ := reader.ReadString('\n')
		template.Variables[v] = strings.TrimSpace(value)
	}
	output.Space()

	return template.Update()
}

// Lists all available templates
func listTemplates() {
	templateNames, err := files.ThaumTemplates()
	if err != nil {
		output.ErrorAsObject(err)
		os.Exit(1)
	}
	output.ListTemplates(templateNames)
}

// Called when thaum is actually run.
func onRun(c *cli.Context) error {
	template := c.Args().Get(0)

	if len(c.Args()) == 0 {
		listTemplates()
		return nil
	}

	templateStruct, err := files.GetTemplate(template)
	if err != nil {
		output.ErrorAsObject(err)
		os.Exit(1)
	}

	if (overwrite) { output.Warning("Thaum will overwrite files. Be careful!") }
	templateStruct = askForVariables(templateStruct)
	files.Compile(templateStruct, overwrite)

	return nil
}

func buildApp() *cli.App {
	cli.AppHelpTemplate = constants.HelpTemplate

	app := cli.NewApp()
	app.Name = "thaum"
	app.Usage = "Generate micro-boilerplates"
	app.Action = onRun
	app.Version = "0.6.0"
	app.Flags = []cli.Flag {
    cli.BoolFlag{
      Name: "force, f",
      Usage: "forces overwrite of existing files. Use with caution.",
			Destination: &overwrite,
    },
  }

	return app
}

func main() {
	app := buildApp()
	app.Run(os.Args)
}
