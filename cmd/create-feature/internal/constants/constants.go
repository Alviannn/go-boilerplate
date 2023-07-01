package constants

import "path"

var (
	TemplatesDir        = "./cmd/create-feature/templates"
	TemplatesDomainDir  = path.Join(TemplatesDir, "domains")
	TemplatesFeatureDir = path.Join(TemplatesDir, "features")

	TemplateFileExt = ".template"
	GoFileExt       = ".go"
)
