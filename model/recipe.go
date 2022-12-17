package model

type RawRecipe struct {
	Title         string   `json:"title" yaml:"title" toml:"title"`
	Categories    []string `json:"categories" yaml:"categories" toml:"categories"`
	Filename      string   `json:"-" yaml:"-" toml:"-"`
	IsMissingMeta bool     `json:"-" yaml:"-" toml:"-"`
}

func (rr RawRecipe) ToTemplateRecipe() TemplateRecipe {
	tr := TemplateRecipe{
		Title:    rr.Title,
		Filename: rr.Filename,
	}

	if rr.Title == "" {
		tr.Title = rr.Filename
	}

	return tr
}

type TemplateRecipe struct {
	Title    string
	Filename string
}
