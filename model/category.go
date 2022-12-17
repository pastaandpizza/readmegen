package model

import (
	"sort"

	"golang.org/x/text/collate"
	"golang.org/x/text/language"
)

type RawCategories map[string][]RawRecipe

func (rc RawCategories) GetSortedKeys() []string {
	keys := []string{}
	for k := range rc {
		keys = append(keys, k)
	}

	cl := collate.New(language.Polish)
	sort.SliceStable(keys, func(i, j int) bool {
		return cl.CompareString(keys[i], keys[j]) == -1
	})

	return keys
}

type TemplateCategory struct {
	Name  string
	Items []TemplateRecipe
}

func NewTemplateCategory(name string, items []RawRecipe) TemplateCategory {
	newItems := []TemplateRecipe{}
	for _, item := range items {
		newItems = append(newItems, item.ToTemplateRecipe())
	}

	cl := collate.New(language.Polish)
	sort.SliceStable(newItems, func(i, j int) bool {
		return cl.CompareString(newItems[i].Title, newItems[j].Title) == -1
	})

	return TemplateCategory{
		Name:  name,
		Items: newItems,
	}
}
