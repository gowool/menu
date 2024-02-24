package main

import (
	"context"
	"fmt"
	"html/template"
	"io/fs"
	"net/url"
	"strings"

	sprig "github.com/go-task/slim-sprig"

	"github.com/gowool/menu"
	"github.com/gowool/menu/renderer"
	"github.com/gowool/menu/views"
)

func main() {
	u, _ := url.Parse("http://localhost/blog/article-test-1")
	ctx := context.WithValue(context.Background(), "url", u)

	item := menu.Must(menu.NewItem("root",
		menu.WithChild(menu.Must(menu.NewItem("home",
			menu.WithLabel("Home"),
			menu.WithURI("/"),
		))),
		menu.WithChild(menu.Must(menu.NewItem("about",
			menu.WithLabel("About"),
			menu.WithURI("/about"),
			menu.WithPosition(2),
		))),
		menu.WithChild(menu.Must(menu.NewItem("blog",
			menu.WithLabel("Blog"),
			menu.WithURI("/blog"),
			menu.WithPosition(1),
			menu.WithChild(menu.Must(menu.NewItem("article1",
				menu.WithLabel("Article 1"),
				menu.WithURI("/blog/article-test-1"),
			))),
		))),
	))

	item.ReorderChildren()

	matcher := menu.NewCoreMatcher(menu.URLVoter{})

	printMenu(ctx, renderer.NewTemplateRenderer(newTheme(), matcher), item)
	printMenu(ctx, renderer.NewListRenderer(matcher), item)
}

func printMenu(ctx context.Context, render renderer.Renderer, item *menu.Item) {
	str, err := render.Render(ctx, item)
	if err != nil {
		panic(err)
	}
	fmt.Println(str)
}

type Theme struct {
	t *template.Template
}

func newTheme() Theme {
	return Theme{t: buildTemplate()}
}

func (t Theme) HTML(_ context.Context, name string, data any) (string, error) {
	var b strings.Builder
	err := t.t.ExecuteTemplate(&b, name, data)
	return b.String(), err
}

func buildTemplate() *template.Template {
	funcMap := sprig.FuncMap()
	funcMap["raw"] = func(s string) template.HTML {
		return template.HTML(s)
	}

	files, err := fs.Glob(views.FS, "menu/*.html")
	if err != nil {
		panic(err)
	}

	t := template.Must(template.
		New(renderer.MenuTemplate).
		Funcs(funcMap).
		Parse(readFile(renderer.MenuTemplate[1:])))

	for _, f := range files {
		if f == renderer.MenuTemplate[1:] {
			continue
		}
		_ = template.Must(t.New("@" + f).Parse(readFile(f)))
	}

	return t
}

func readFile(path string) string {
	data, err := fs.ReadFile(views.FS, path)
	if err != nil {
		panic(err)
	}
	return string(data)
}
