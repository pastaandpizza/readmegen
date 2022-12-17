package main

import (
	"flag"
	"fmt"
	"os"
	"text/template"
	"time"

	"github.com/pastaandpizza/readmegen/constants"
	"github.com/pastaandpizza/readmegen/model"
	"github.com/pastaandpizza/readmegen/resolvers"
)

var templateFile, outFile string

func init() {
	flag.StringVar(&templateFile, "template", constants.DefaultTemplateFile, "template file")
	flag.StringVar(&outFile, "out", constants.DefaultOutputFile, "output file")
	flag.Parse()
}

func main() {
	// Load template
	tmpl, err := template.
		New("").
		Funcs(template.FuncMap{
			"now": func() string {
				return time.Now().Format("2006-01-02 15:04:05")
			},
		}).
		ParseFiles(templateFile)

	if err != nil {
		fmt.Printf("Error while parsing template '%s': %s", templateFile, err.Error())
		os.Exit(1)
	}

	// Load metadata from files
	categorizedRecipes := model.RawCategories{}
	recipesWithNoMetadata := []model.RawRecipe{}
	recipesWithNoCategory := []model.RawRecipe{}

	for _, filename := range flag.Args() {
		rawRecipe, err := resolvers.ResolveRecipe(filename)
		if err != nil {
			fmt.Printf("Error while parsing recipe '%s': %s", filename, err.Error())
			os.Exit(1)
		}

		if rawRecipe.IsMissingMeta {
			recipesWithNoMetadata = append(recipesWithNoMetadata, rawRecipe)
			continue
		}

		if len(rawRecipe.Categories) == 0 {
			recipesWithNoCategory = append(recipesWithNoCategory, rawRecipe)
			continue
		}

		for _, category := range rawRecipe.Categories {
			if _, ok := categorizedRecipes[category]; !ok {
				categorizedRecipes[category] = []model.RawRecipe{}
			}
			categorizedRecipes[category] = append(categorizedRecipes[category], rawRecipe)
		}
	}

	// Generate categories
	categories := []model.TemplateCategory{}
	for _, category := range categorizedRecipes.GetSortedKeys() {
		categories = append(categories, model.NewTemplateCategory(
			category,
			categorizedRecipes[category],
		))
	}
	if len(recipesWithNoCategory) != 0 {
		categories = append(categories, model.NewTemplateCategory(
			constants.CategoryNoCategory,
			recipesWithNoCategory,
		))
	}
	if len(recipesWithNoMetadata) != 0 {
		categories = append(categories, model.NewTemplateCategory(
			constants.CategoryMissingMetadata,
			recipesWithNoMetadata,
		))
	}

	// Generate output file
	out, err := os.Create(outFile)
	if err != nil {
		fmt.Printf("Error while creating output file '%s': %s", outFile, err.Error())
		os.Exit(1)
	}
	defer out.Close()

	tmpl.ExecuteTemplate(out, constants.TemplateHeader, nil)
	for _, category := range categories {
		tmpl.ExecuteTemplate(out, constants.TemplateCategory, category)
	}
	tmpl.ExecuteTemplate(out, constants.TemplateFooter, nil)
}
