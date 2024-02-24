package menu

import (
	"errors"
	"fmt"
	"slices"
)

var ErrItemBelongsToAnotherMenu = errors.New("cannot add menu item as child, it already belongs to another menu (e.g. has a parent)")

// Item represents an item in a menu.
type Item struct {
	Name               string         `json:"name,omitempty"`
	URI                string         `json:"uri,omitempty"`
	Label              string         `json:"label,omitempty"`
	Position           int            `json:"position,omitempty"`
	DisplayChildren    bool           `json:"display_children,omitempty"`
	Display            bool           `json:"display,omitempty"`
	Current            *bool          `json:"current,omitempty"`
	Attributes         map[string]any `json:"attributes,omitempty"`
	LinkAttributes     map[string]any `json:"link_attributes,omitempty"`
	ChildrenAttributes map[string]any `json:"children_attributes,omitempty"`
	LabelAttributes    map[string]any `json:"label_attributes,omitempty"`
	Extras             map[string]any `json:"extras,omitempty"`
	Parent             *Item          `json:"parent,omitempty"`
	Children           []*Item        `json:"children,omitempty"`
}

func Must(item *Item, err error) *Item {
	if err != nil {
		panic(err)
	}
	return item
}

// NewItem creates a new Item with the specified name and options. It initializes the Item with default attribute maps
// and sets the Display and DisplayChildren fields to true. The function applies each option to the Item sequentially,
// returning an error if any of the options fail. If successful, it returns the created Item and a nil error.
func NewItem(name string, options ...Option) (*Item, error) {
	item := &Item{
		Name:               name,
		Attributes:         map[string]any{},
		LinkAttributes:     map[string]any{},
		ChildrenAttributes: map[string]any{},
		LabelAttributes:    map[string]any{},
		Extras:             map[string]any{},
		Display:            true,
		DisplayChildren:    true,
	}

	for _, option := range options {
		if err := option(item); err != nil {
			return nil, err
		}
	}

	return item, nil
}

// String returns the name of an Item. If the name is empty, it returns "n/a".
func (i *Item) String() string {
	if i.Name == "" {
		return "n/a"
	}
	return i.Name
}

// SetIsCurrent sets the IsCurrent property of an Item to true by assigning a pointer to a boolean value to its Current field.
func (i *Item) SetIsCurrent() {
	current := true
	i.Current = &current
}

// SetNotCurrent sets the Current field of an Item to false.
func (i *Item) SetNotCurrent() {
	current := false
	i.Current = &current
}

// IsCurrent returns true if the item is marked as current.
func (i *Item) IsCurrent() bool {
	return i.Current != nil && *i.Current
}

// Attribute returns the value of the specified attribute from the Attributes map for the given item.
// If the attribute is not found, it returns the default value.
func (i *Item) Attribute(name string, def any) any {
	if attribute, ok := i.Attributes[name]; ok {
		return attribute
	}
	return def
}

// LinkAttribute returns the value of the specified link attribute from the LinkAttributes map for the given item.
// If the attribute is not found, it returns the default value.
func (i *Item) LinkAttribute(name string, def any) any {
	if attribute, ok := i.LinkAttributes[name]; ok {
		return attribute
	}
	return def
}

// ChildrenAttribute returns the value of the provided attribute name from the ChildrenAttributes map of an Item.
// If the attribute is not found, it returns the default value.
func (i *Item) ChildrenAttribute(name string, def any) any {
	if attribute, ok := i.ChildrenAttributes[name]; ok {
		return attribute
	}
	return def
}

// LabelAttribute returns the attribute value associated with the given name from the LabelAttributes map of the Item.
// If the attribute is not found, it returns the default value.
func (i *Item) LabelAttribute(name string, def any) any {
	if attribute, ok := i.LabelAttributes[name]; ok {
		return attribute
	}
	return def
}

// Extra returns the value of the specified extra info for an Item.
// If the info is not found, it returns the default value provided or nil.
func (i *Item) Extra(name string, def ...any) any {
	if extra, ok := i.Extras[name]; ok {
		return extra
	}
	if len(def) > 0 {
		return def[0]
	}
	return nil
}

// IsRoot returns true if the Item has no parent, indicating that it is the root item in the tree structure. Otherwise, it returns false.
func (i *Item) IsRoot() bool {
	return i.Parent == nil
}

// Root returns the root item of the item hierarchy.
// If the item has no parent, it is considered the root itself.
func (i *Item) Root() *Item {
	if i.Parent == nil {
		return i
	}
	return i.Parent.Root()
}

// Level returns the level of the item in the hierarchy.
// If the item has no parent, it is considered to be at level 0.
// Each level is determined by the level of its parent item plus 1.
func (i *Item) Level() int {
	if i.Parent == nil {
		return 0
	}
	return i.Parent.Level() + 1
}

// Copy creates a deep copy of the Item and its children.
func (i *Item) Copy() (*Item, error) {
	item := *i
	item.Parent = nil
	item.Children = make([]*Item, 0, len(i.Children))

	for _, child := range i.Children {
		c, err := child.Copy()
		if err != nil {
			return nil, err
		}
		if _, err = item.AddChild(c); err != nil {
			return nil, err
		}
	}

	return &item, nil
}

// AddChild adds a child item to the current item. It accepts a `child` parameter of type `any`,
// which can be either an `*Item` or any other value. If `child` is an `*Item`, it checks if the child already
// belongs to another menu (i.e., it has a non-nil parent). If so, it returns an error `ErrItemBelongsToAnotherMenu`.
// Otherwise, it sets the parent of the child to the current item and appends the child to the list of children.
// If `child` is not an `*Item`, it creates a new item with a name obtained by formatting `child` as a string
// and using the options passed as variadic arguments. It sets the parent of the newly created child to the current item
// and appends it to the list of children. The method returns the child item added and a possible error.
func (i *Item) AddChild(child any, options ...Option) (childItem *Item, err error) {
	switch child := child.(type) {
	case *Item:
		if child.Parent != nil {
			return nil, ErrItemBelongsToAnotherMenu
		}
		childItem = child
	default:
		name := fmt.Sprintf("%v", child)
		if childItem, err = NewItem(name, options...); err != nil {
			return nil, err
		}
	}

	childItem.Parent = i
	i.Children = append(i.Children, childItem)

	return childItem, nil
}

// Child returns the child item with the specified name, if it exists. If no child with the given name is found, nil is returned.
func (i *Item) Child(name string) *Item {
	for _, child := range i.Children {
		if child.Name == name {
			return child
		}
	}
	return nil
}

// ReorderChildren sorts the child items of an Item based on their Position field.
// The sorting is done in ascending order.
func (i *Item) ReorderChildren() {
	slices.SortFunc(i.Children, func(a, b *Item) int {
		return a.Position - b.Position
	})
}

// HasChildren checks if the item has any children that are set to be displayed.
func (i *Item) HasChildren() bool {
	for _, child := range i.Children {
		if child.Display {
			return true
		}
	}
	return false
}

// FirstChild returns the first child of an Item instance.
func (i *Item) FirstChild() *Item {
	return i.Children[0]
}

// LastChild returns the last child of the current item.
//
// Example:
//
//	func (i *Item) IsLast() bool {
//	    if i.Parent == nil {
//	        return false
//	    }
//	    return i.Parent.LastChild() == i
//	}
func (i *Item) LastChild() *Item {
	return i.Children[len(i.Children)-1]
}

// IsLast returns true if the item is the last child of its parent, false otherwise.
func (i *Item) IsLast() bool {
	if i.Parent == nil {
		return false
	}
	return i.Parent.LastChild() == i
}

// IsFirst returns true if the item is the first child of its parent,
// otherwise it returns false.
func (i *Item) IsFirst() bool {
	if i.Parent == nil {
		return false
	}
	return i.Parent.FirstChild() == i
}

// ActsLikeFirst checks if an Item acts like the first item in the menu hierarchy.
func (i *Item) ActsLikeFirst() bool {
	// root items are never "marked" as first
	if i.Parent == nil {
		return false
	}

	// a menu acts like first only if it is displayed
	if !i.Display {
		return false
	}

	// if we're first and visible, we're first, period.
	if i.IsFirst() {
		return true
	}

	for _, child := range i.Parent.Children {
		// loop until we find a visible menu. If its this menu, we're first
		if child.Display {
			return child.Name == i.Name
		}
	}

	return false
}

func (i *Item) ActsLikeLast() bool {
	// root items are never "marked" as first
	if i.Parent == nil {
		return false
	}

	// a menu acts like last only if it is displayed
	if !i.Display {
		return false
	}

	// if we're last and visible, we're last, period.
	if i.IsLast() {
		return true
	}

	for j := len(i.Parent.Children) - 1; j >= 0; j-- {
		// loop until we find a visible menu. If its this menu, we're first
		if i.Parent.Children[j].Display {
			return i.Parent.Children[j].Name == i.Name
		}
	}

	return false
}
