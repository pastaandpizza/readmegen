package resolvers

import (
	"bytes"
	"os"
	"regexp"

	"github.com/adrg/frontmatter"
	"github.com/pastaandpizza/readmegen/model"
)

var titlePattern *regexp.Regexp = regexp.MustCompile("(?m)^# (.*?)$")

func ResolveRecipe(filename string) (model.RawRecipe, error) {
	rawData, err := os.ReadFile(filename)
	if err != nil {
		return model.RawRecipe{}, err
	}

	recipe := model.RawRecipe{
		Filename: filename,
	}

	_, err = frontmatter.MustParse(bytes.NewBuffer(rawData), &recipe)
	if err != nil || recipe.Title == "" {
		// No frontmatter, search for title
		haystack := string(rawData)
		titleElements := titlePattern.FindStringSubmatch(haystack)
		if len(titleElements) < 2 {
			// No title found
			recipe.Title = filename
			recipe.IsMissingMeta = true
		} else {
			// Title found
			recipe.Title = titleElements[1]
		}
	}

	return recipe, nil
}
