package commands

import (
	"log"
	"os"
	"path/filepath"
	"time"

	"github.com/code-gorilla-au/pyrotic/internal/chalk"
	"github.com/code-gorilla-au/pyrotic/internal/engine"
	"github.com/spf13/cobra"
)

func generateCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "generate <generator-name> --name <file-name>",
		Short: "Generate template",
		Long:  `Generate template by generator name. `,
		Run:   generateFunc(),
	}
}

func generateFunc() cmdFunc {
	return func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			log.Println("at least 1 generator must be provided")
			return
		}

		generatorName := args[0]
		dirPath := filepath.Join(flagTemplatePath, generatorName)
		_, err := os.ReadDir(dirPath)
		if err != nil {
			log.Println("generator not found:", generatorName)
			return
		}

		sharedPath := filepath.Join(flagTemplatePath, flagSharedFolder)

		log.Println(chalk.Green("running generator:"), generatorName)
		e, err := engine.New(flagDryrun, dirPath, sharedPath, flagTemplateSuffix)
		if err != nil {
			log.Println("error creating engine ", err)
			return
		}

		startTime := time.Now()
		err = e.Generate(engine.Data{Name: flagGeneratorName, MetaArgs: flagMetaArgs})
		if err != nil {
			return
		}
		log.Println(chalk.Green("generated in "), time.Since(startTime))
	}
}
