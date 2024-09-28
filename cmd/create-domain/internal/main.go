package main

import (
	"flag"
	"log"
	"os"
	"path"
	"strings"

	"go-boilerplate/cmd/create-domain/internal/constants"
	"go-boilerplate/cmd/create-domain/internal/helpers"
)

type CreateTemplateParams struct {
	DomainsDirectory string
	DomainName       string
}

func main() {
	flagSet := flag.NewFlagSet("create-domain", flag.ExitOnError)
	domainName := flagSet.String("domain", "", "The domain or parent of the feature, for example 'super_hero', in snake case.")

	cliArgs := os.Args[1:]
	if err := flagSet.Parse(cliArgs); err != nil {
		log.Fatal(err)
	}

	*domainName = strings.Trim(*domainName, " ")
	*domainName = strings.ToLower(*domainName)

	if *domainName == "" {
		flagSet.Usage()
		return
	}

	params := CreateTemplateParams{
		DomainsDirectory: constants.DomainsDirPath,
		DomainName:       *domainName,
	}

	if err := WriteTemplateFiles(params); err != nil {
		log.Fatal(err)
	}

	log.Printf("Created new domain called %s!\n", *domainName)
	log.Print("Don't forget to add the domains module to './internal/domains/modules.go' file.\n")
}

func ReplaceTemplateContent(template string, domainName string) string {
	moduleName := helpers.GetModuleName()
	domainPascalName := helpers.SnakeToPascalCase(domainName)
	domainCamelName := helpers.SnakeToCamelCase(domainName)

	template = strings.ReplaceAll(template, "{module_name}", moduleName)
	template = strings.ReplaceAll(template, "{domain_pascal_name}", domainPascalName)
	template = strings.ReplaceAll(template, "{domain_snake_name}", domainName)
	template = strings.ReplaceAll(template, "{domain_camel_name}", domainCamelName)

	return template
}

func WriteTemplateFiles(params CreateTemplateParams) (err error) {
	domainsDir := path.Join(params.DomainsDirectory, params.DomainName)
	if err = os.Mkdir(domainsDir, os.ModePerm); err != nil {
		return
	}

	filePathList, err := helpers.ReadCompleteFilesInDir(constants.TemplatesDir)
	if err != nil {
		return
	}

	for _, filePath := range filePathList {
		outFilePath := strings.ReplaceAll(filePath, path.Clean(constants.TemplatesDir), path.Clean(domainsDir))
		outFilePath = ReplaceTemplateContent(outFilePath, params.DomainName)
		outFilePath = strings.Replace(outFilePath, constants.TemplateFileExt, constants.GoFileExt, 1)

		parentOutFilePath := path.Dir(outFilePath)
		if err = os.MkdirAll(parentOutFilePath, os.ModePerm); err != nil {
			return err
		}

		templateContent, err := helpers.ReadFileAsString(filePath)
		if err != nil {
			return err
		}

		templateContent = ReplaceTemplateContent(templateContent, params.DomainName)
		if err = helpers.WriteFileFromString(outFilePath, templateContent); err != nil {
			return err
		}
	}

	return
}
