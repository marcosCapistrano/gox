package handlers

import (
	"platform/pipeline"
	"strings"
	"text/template"
)

type NavbarContext struct {
	NavLinks     map[string]string
	SelectedLink string
}

type SubNavbarContext struct {
	NavLinks     map[string]string
	SelectedLink string
}

type StoreHandler struct {
}

type StoreHandlerContext struct {
	NavbarContext    NavbarContext
	SubNavbarContext SubNavbarContext
}

func (handler StoreHandler) Execute(ctx *pipeline.ComponentContext) {
	templ, err := template.ParseFiles("templates/layout.html", "templates/store/store.html", "templates/components/subnavbar.html", "templates/components/navbar.html")
	if err != nil {
		panic(err)
	}

	selectedLink := parseSelectedNavbarLink(ctx.URL.Path)
	selectedSubLink := parseSelectedSubnavbarLink(ctx.URL.Path)

	var context StoreHandlerContext = StoreHandlerContext{
		NavbarContext: NavbarContext{
			NavLinks: map[string]string{
				"loja":    "Loja",
				"oficina": "Oficina",
			},
			SelectedLink: selectedLink,
		},

		SubNavbarContext: SubNavbarContext{
			NavLinks: map[string]string{
				"pecas":        "Pecas",
				"equipamentos": "Equipamentos",
				"acessorios":   "Acessórios",
			},
			SelectedLink: selectedSubLink,
		},
	}

	templ.Execute(ctx.ResponseWriter, context)
}

func parseSelectedNavbarLink(path string) string {
	pathItems := strings.Split(path, "/")
	if len(pathItems) > 1 {
		return pathItems[1]
	}

	return ""
}

func parseSelectedSubnavbarLink(path string) string {
	pathItems := strings.Split(path, "/")
	if len(pathItems) > 2 {
		return pathItems[2]
	}

	return ""
}
