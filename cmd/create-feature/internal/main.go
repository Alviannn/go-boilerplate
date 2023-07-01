package main

import (
	"cmd/create-feature/internal/constants"
	"cmd/create-feature/internal/helpers"
	"flag"
	"fmt"
	"log"
	"os"
	"path"
	"strings"
)

type CreateTemplateParams struct {
	DomainsDirectory string

	DomainName  string
	FeatureName string
}

func main() {
	flagSet := flag.NewFlagSet("create-feature", flag.ExitOnError)
	domainName := flagSet.String("domain", "", "The domain or parent of the feature, for example users, snake case.")
	featureName := flagSet.String("name", "", "The feature name, must be in snake case.")

	cliArgs := os.Args[1:]
	if err := flagSet.Parse(cliArgs); err != nil {
		log.Fatal(err)
	}

	*domainName = strings.Trim(*domainName, " ")
	*featureName = strings.Trim(*featureName, " ")

	if *featureName == "" || *domainName == "" {
		flagSet.Usage()
		return
	}

	params := CreateTemplateParams{
		DomainsDirectory: "./internal/domains",
		DomainName:       *domainName,
		FeatureName:      *featureName,
	}

	if err := WriteDomainFiles(params); err != nil {
		log.Fatal(err)
	}
	if err := WriteFeatureFiles(params); err != nil {
		log.Fatal(err)
	}
}

func ReplaceTemplateContent(template string, domainName string, featureName string) string {
	domainPackageName := helpers.SnakeToPackageName(domainName)
	domainPascalName := helpers.SnakeToPascalCase(domainName)

	packageName := helpers.SnakeToPackageName(featureName)
	pascalName := helpers.SnakeToPascalCase(featureName)

	template = strings.ReplaceAll(template, "{module_name}", helpers.GetModuleName())

	template = strings.ReplaceAll(template, "{package_name}", packageName)
	template = strings.ReplaceAll(template, "{pascal_name}", pascalName)

	template = strings.ReplaceAll(template, "{domain_package_name}", domainPackageName)
	template = strings.ReplaceAll(template, "{domain_pascal_name}", domainPascalName)
	template = strings.ReplaceAll(template, "{domain_snake_name}", domainName)

	return template
}

func WriteFeatureFiles(params CreateTemplateParams) (err error) {
	domainPackageName := helpers.SnakeToPackageName(params.DomainName)
	packageName := helpers.SnakeToPackageName(params.FeatureName)

	featureDir := path.Join(params.DomainsDirectory, domainPackageName, packageName)
	if err = os.Mkdir(featureDir, os.ModePerm); err != nil {
		return
	}

	templateList, err := os.ReadDir(constants.TemplatesFeatureDir)
	if err != nil {
		return
	}

	for _, template := range templateList {
		fileName := template.Name()
		goFileName := strings.Replace(fileName, constants.TemplateFileExt, constants.GoFileExt, 1)

		newFilePath := path.Join(featureDir, goFileName)
		templateFilePath := path.Join(constants.TemplatesFeatureDir, fileName)

		templateContent, err := helpers.ReadFileAsString(templateFilePath)
		if err != nil {
			return err
		}

		output := ReplaceTemplateContent(templateContent, params.DomainName, params.FeatureName)
		if err := helpers.WriteFileFromString(newFilePath, output); err != nil {
			return err
		}
	}

	return
}

func WriteDomainFiles(params CreateTemplateParams) (err error) {
	domainPackageName := helpers.SnakeToPackageName(params.DomainName)

	domainDir := path.Join(params.DomainsDirectory, domainPackageName)
	if err = os.Mkdir(domainDir, os.ModePerm); err != nil {
		return nil
	}

	templateList, err := os.ReadDir(constants.TemplatesDomainDir)
	if err != nil {
		return
	}

	for _, template := range templateList {
		fileName := template.Name()
		goFileName := strings.Replace(fileName, constants.TemplateFileExt, constants.GoFileExt, 1)
		domainFileName := fmt.Sprintf("%s_%s", params.DomainName, goFileName)

		newFilePath := path.Join(domainDir, domainFileName)
		templateFilePath := path.Join(constants.TemplatesDomainDir, fileName)

		templateContent, err := helpers.ReadFileAsString(templateFilePath)
		if err != nil {
			return err
		}

		output := ReplaceTemplateContent(templateContent, params.DomainName, params.FeatureName)
		if err := helpers.WriteFileFromString(newFilePath, output); err != nil {
			return err
		}
	}

	return
}
