package renderer

import (
	"context"
	"fmt"
	"html"
	"maps"
	"strings"

	"github.com/gowool/menu"
	"github.com/gowool/menu/internal"
)

var _ Renderer = ListRenderer{}

// ListRenderer is a type that implements the Renderer interface and is responsible for rendering menus in list format.
// Render method of the ListRenderer type is used to render the menu and return the generated HTML string.
// It takes the root menu item along with optional rendering options and returns the rendered menu as a string.
// The options can be used to customize the rendering behavior such as depth, matching depth, CSS classes, etc.
// It uses the renderList method to recursively render the menu items and their children.
// renderList method recursively renders the menu items and their children in list format.
// It takes the current menu item, its children attributes, and the rendering options.
// It checks if the rendering should stop based on the depth option or if the current item has no children or if the display of children is disabled.
// If any of these conditions are met, it returns an empty string.
// Otherwise, it iterates over the children of the current item and recursively calls the renderItem method to render each child.
// It then formats the rendered children as an unordered list and returns it as a string.
// renderItem method renders a single menu item and its children.
// It takes the current menu item and the rendering options.
// It checks if the item should be displayed based on its display flag.
// It then determines the CSS classes to be applied to the item based on its current state, such as current, ancestor, first, last, etc.
// It creates a clone of the item's attributes, adds the CSS classes, and formats them as HTML attributes.
// It then determines the current item's level and formats the HTML tag for the list item accordingly.
// It calls the renderLink method to render the link or span element based on the item's URI and the rendering options.
// It then formats the rendered item along with its children as a list item and returns it as a string.
// renderLink method renders either a link or a span element for a menu item.
// It takes the current menu item and the rendering options.
// If the item has a non-empty URI and it's not currently considered as the current item or if the currentAsLink option is enabled, it renders a link element.
// Otherwise, it renders a span element.
// It uses the renderLabel method to render the label content for the link or span element.
// It then formats the rendered
type ListRenderer struct {
	matcher menu.Matcher
	options *Options
}

// NewListRenderer creates a new instance of ListRenderer with the given matcher and options. The matcher is used to determine the current item and its ancestor. The options are used
func NewListRenderer(matcher menu.Matcher, options ...Option) ListRenderer {
	return ListRenderer{
		matcher: matcher,
		options: NewOptions(options...),
	}
}

// Render renders the menu item and its children into a HTML list.
// It accepts a context, the menu item to render, and optional rendering options.
// It returns the rendered content as a string and an error if any.
func (r ListRenderer) Render(ctx context.Context, item *menu.Item, options ...Option) (string, error) {
	opts := r.options.Copy().Apply(options...)

	content := r.renderList(ctx, item, item.ChildrenAttributes, opts)

	if opts.ClearMatcher {
		r.matcher.Clear()
	}

	return content, nil
}

// renderList renders a list of items and their children in HTML format.
//
// If the options indicate that the rendering should stop or if the item
// has no children or is not set to display its children, an empty string
// is returned.
//
// The method constructs an HTML string by appending the formatted list
// opening tag, the rendered children, and the formatted list closing tag.
//
// The rendered children are obtained by calling the renderChildren method
// and passing it the parent item, a context, and the options.
//
// The method then constructs the opening and closing tags by calling the
// format method, passing in the appropriate arguments.
//
// Finally, the method returns the resulting HTML string.
func (r ListRenderer) renderList(ctx context.Context, item *menu.Item, attributes map[string]any, options *Options) string {
	if options.IsStop() || !item.HasChildren() || !item.DisplayChildren {
		return ""
	}

	level := item.Level()

	var b strings.Builder
	b.WriteString(r.format(fmt.Sprintf("<ul%s>", internal.HTMLAttributes(attributes)), "ul", level, options))
	b.WriteString(r.renderChildren(ctx, item, options))
	b.WriteString(r.format("</ul>", "ul", level, options))

	return b.String()
}

// renderChildren renders the children of a menu item with the given context and options.
func (r ListRenderer) renderChildren(ctx context.Context, item *menu.Item, options *Options) string {
	options = options.SubDepth().SubMatchingDepth()

	var b strings.Builder
	for _, child := range item.Children {
		b.WriteString(r.renderItem(ctx, child, options.Copy()))
	}
	return b.String()
}

// renderItem takes a context, an item, and options, and renders the item as an HTML list item.
// If the item should not be displayed, it returns an empty string.
// It retrieves the item's classes and appends additional classes based on its properties and context.
// The method then constructs the attributes, including the classes, for the <li> element.
// It constructs a string builder and appends the opening <li> tag, followed by the rendered link for the item.
// If the item has children and should be displayed, it appends the appropriate classes for a branch element.
// Otherwise, it appends the appropriate classes for a leaf element.
// It then constructs the attributes for the children list, and appends the rendered list to the string builder.
// Finally, it appends the closing </li> tag and returns the constructed string.
func (r ListRenderer) renderItem(ctx context.Context, item *menu.Item, options *Options) string {
	if !item.Display {
		return ""
	}

	classes := make([]string, 0, 5)
	classes = append(classes, item.Attribute("class", "").(string))

	if r.matcher.IsCurrent(ctx, item) {
		classes = append(classes, options.CurrentClass)
	} else if r.matcher.IsAncestor(ctx, item, options.MatchingDepth) {
		classes = append(classes, options.AncestorClass)
	}

	if item.ActsLikeFirst() {
		classes = append(classes, options.FirstClass)
	}
	if item.ActsLikeLast() {
		classes = append(classes, options.LastClass)
	}

	if !options.IsStop() && item.HasChildren() {
		if item.DisplayChildren {
			classes = append(classes, options.BranchClass)
		}
	} else {
		classes = append(classes, options.LeafClass)
	}

	attributes := maps.Clone(item.Attributes)
	attributes["class"] = internal.HTMLClasses(classes)

	level := item.Level()

	var b strings.Builder
	b.WriteString(r.format(fmt.Sprintf("<li%s>", internal.HTMLAttributes(attributes)), "li", level, options))
	b.WriteString(r.renderLink(ctx, item, options))

	classes = []string{
		item.ChildrenAttribute("class", "").(string),
		fmt.Sprintf("menu-level-%d", item.Level()),
	}
	attributes = maps.Clone(item.ChildrenAttributes)
	attributes["class"] = internal.HTMLClasses(classes)

	b.WriteString(r.renderList(ctx, item, attributes, options))
	b.WriteString(r.format("</li>", "li", level, options))

	return b.String()
}

// renderLink renders a link element or a span element based on the item and options.
// It returns the formatted link or span element.
func (r ListRenderer) renderLink(ctx context.Context, item *menu.Item, options *Options) string {
	var text string
	if item.URI != "" && (!r.matcher.IsCurrent(ctx, item) || options.CurrentAsLink) {
		text = r.renderLinkElement(item, options)
	} else {
		text = r.renderSpanElement(item, options)
	}
	return r.format(text, "link", item.Level(), options)
}

// renderLinkElement formats a link element for a menu item.
// It escapes the URI, applies link attributes and renders the label.
func (r ListRenderer) renderLinkElement(item *menu.Item, options *Options) string {
	return fmt.Sprintf(`<a href="%s"%s>%s</a>`, html.EscapeString(item.URI), internal.HTMLAttributes(item.LinkAttributes), r.renderLabel(item, options))
}

// renderSpanElement renders a span element with the label of the menu item.
// It formats the element using the internal.HTMLAttributes function to handle HTML attributes,
// and calls the renderLabel method to render the label itself. The resulting HTML element is returned as a string.
// The function accepts the menu item and the options as parameters.
func (r ListRenderer) renderSpanElement(item *menu.Item, options *Options) string {
	return fmt.Sprintf("<span%s>%s</span>", internal.HTMLAttributes(item.LabelAttributes), r.renderLabel(item, options))
}

// renderLabel renders the label of a menu item.
//
// This method takes an item and options as input and returns the rendered label
// as a string. The rendered label is the menu item's label with HTML special
// characters escaped, unless the "AllowSafeLabels" option is set to true and the
// item has the "safe_label" extra attribute set to true.
//
// Parameters:
//   - item: The menu item whose label should be rendered.
//   - options: The options to be used during rendering.
//
// Returns:
//
//	The rendered label as a string.
//
// Example usage:
//
//	renderer := ListRenderer{}
//	options := &Options{AllowSafeLabels: true}
//	label := renderer.renderLabel(item, options)
func (r ListRenderer) renderLabel(item *menu.Item, options *Options) string {
	if options.AllowSafeLabels && item.Extra("safe_label", false).(bool) {
		return item.Label
	}
	return html.EscapeString(item.Label)
}

// format formats the given content based on the type and level parameters, as well as the options provided.
// If the "compressed" extra option is set to true, the content is returned as is. Otherwise, the content is indented
// according to the level parameter and returned with a newline character appended at the end.
// The type parameter determines the indentation spacing as follows:
// - "ul" or "link": level * 4 spaces
// - "li": level * 4 - 2 spaces
// Parameters:
//   - content: the content to be formatted
//   - typ: the type of content
//   - level: the level of the content
//   - options: additional options to customize the format
//
// Returns:
//   - the formatted content
func (r ListRenderer) format(content, typ string, level int, options *Options) string {
	if options.Extra("compressed", false).(bool) {
		return content
	}

	spacing := 0
	switch typ {
	case "ul", "link":
		spacing = level * 4
	case "li":
		spacing = level*4 - 2
	}

	return strings.Repeat(" ", spacing) + content + "\n"
}
