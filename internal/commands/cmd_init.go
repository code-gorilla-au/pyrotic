package commands

import (
	"fmt"
	"log"
	"os"
	"path"
	"path/filepath"

	"github.com/code-gorilla-au/pyrotic/internal/chalk"
	"github.com/spf13/cobra"
)

func initCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "init",
		Short: "init pyrotic",
		Long:  "creates required setup for pyrotic",
		Run:   initFunc,
	}
}

func initFunc(cmd *cobra.Command, args []string) {
	log.Println(chalk.Green("creating initial setup"), templatePath)
	dirPath := path.Join(templatePath, "new", fmt.Sprintf("new%s", templateSuffix))
	if err := os.MkdirAll(filepath.Dir(dirPath), 0750); err != nil {
		log.Println("error creating", err)
		return
	}

	file, err := os.Create(dirPath)
	if err != nil {
		log.Println("error creating base template ", err)
		return
	}
	defer file.Close()
	if _, err := file.Write([]byte(initTemplate)); err != nil {
		log.Println("error creating file")
		return
	}
}

var initTemplate = `---
to:
---
package main

func main() {}
`