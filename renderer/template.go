package renderer

import (
	"context"
	"html/template"

	"github.com/gowool/menu"
	"github.com/gowool/menu/internal"
)

var _ Renderer = TemplateRenderer{}

// MenuTemplate is the constant that holds the path to the menu template file.
const MenuTemplate = "@menu/menu.html"

// Theme is an interface that provides the method HTML for generating HTML code based on a provided template and data.
// HTML takes a context, a template string, and data and returns a string representation of the generated HTML code. It returns an error if there was an issue generating the HTML.
// Example usage:
// theme := // create an instance of a struct that implements the Theme interface
// template := "my-template"
// data := // provide data for the template rendering
// html, err := theme.HTML(ctx, template, data)
//
//	if err != nil {
//	    // handle error
//	}
//
// // use the generated HTML code
type Theme interface {
	HTML(ctx context.Context, template string, data any) (string, error)
}

// TemplateRenderer is a type that represents a renderer for templates.
// It is used to render HTML templates based on a given theme and matcher.
// The renderer provides options for customizing the rendering process.
type TemplateRenderer struct {
	theme   Theme
	matcher menu.Matcher
	options *Options
}

// NewTemplateRenderer creates a new TemplateRenderer with the given theme, matcher, and options.
func NewTemplateRenderer(theme Theme, matcher menu.Matcher, options ...Option) TemplateRenderer {
	return TemplateRenderer{
		theme:   theme,
		matcher: matcher,
		options: NewOptions(options...),
	}
}

// Render is a method of the TemplateRenderer struct that renders a menu item using the specified options and theme.
// It takes a context object, a pointer to a menu.Item object, and a variadic list of options as parameters.
// It returns a string (the rendered content) and an error (if any occurred during rendering).
//
// The function starts by creating a copy of the options and applying the passed options to it.
// It then calls the HTML method of the theme to render the menu item with the specified template and data.
// The data passed to the template includes the context object, the menu item, the options, the matcher, and helper functions for converting attributes and classes.
//
// If the "ClearMatcher" option is set to true, the matcher is cleared after rendering the content.
//
// The rendered content and any error that occurred during rendering are returned as the result of the function.
func (r TemplateRenderer) Render(ctx context.Context, item *menu.Item, options ...Option) (string, error) {
	opts := r.options.Copy().Apply(options...)

	content, err := r.theme.HTML(ctx, opts.Extra("template", MenuTemplate).(string), map[string]any{
		"Ctx":     ctx,
		"Item":    item,
		"Options": opts,
		"Matcher": r.matcher,
		"Classes": internal.HTMLClassesAny,
		"Attributes": func(attributes map[string]any) template.HTMLAttr {
			return template.HTMLAttr(internal.HTMLAttributes(attributes))
		},
	})

	if opts.ClearMatcher {
		r.matcher.Clear()
	}

	return content, err
}
