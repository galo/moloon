package cmd

import (
	"fmt"
	"io/ioutil"
	"log"

	"github.com/galo/moloon/internal/rest"
	"github.com/go-chi/docgen"
	"github.com/spf13/cobra"
)

var (
	routes bool
)

// gendocCmd represents the gendoc command
var gendocCmd = &cobra.Command{
	Use:   "gendoc",
	Short: "Generate project documentation",
	Long:  `Generate project documentation.`,
	Run: func(cmd *cobra.Command, args []string) {
		if routes {
			genRoutesDoc()
		}
	},
}

func init() {
	RootCmd.AddCommand(gendocCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// gendocCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// gendocCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	gendocCmd.Flags().BoolVarP(&routes, "routes", "r", false, "create api routes markdown file")
}

func genRoutesDoc() {
	api, _ := rest.New(false)
	fmt.Print("generating routes markdown file: ")
	md := docgen.MarkdownRoutesDoc(api, docgen.MarkdownOpts{
		ProjectPath: "github.azc.ext.hp.com/galo/pym",
		Intro:       "Pym REST API.",
	})
	if err := ioutil.WriteFile("routes.md", []byte(md), 0644); err != nil {
		log.Println(err)
		return
	}
	fmt.Println("OK")
}
