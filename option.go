package menu

import "maps"

// Option represents a function that can be used to modify an Item.
// It takes a pointer to an Item and returns an error if any.
type Option func(item *Item) error

// WithURI is a function that returns an Option for setting the URI of an Item.
func WithURI(uri string) Option {
	return func(item *Item) error {
		item.URI = uri
		return nil
	}
}

// WithLabel is a function that returns an Option for setting the label of an Item.
// The Option function updates the Item's Label field with the provided label parameter.
// It returns nil if the operation is successful, otherwise an error.
//
// Example usage:
//
//	item := &Item{}
//	err := WithLabel("Example Label")(item)
//	if err != nil {
//	    fmt.Println("Error:", err)
//	    return
//	}
//	fmt.Println("Item label:", item.Label)
//
// Parameters:
//   - label: the label to be set for the Item
//
// Returns:
//   - Option: the Option function that sets the label of an Item
//   - error: nil if the operation is successful, otherwise an error
func WithLabel(label string) Option {
	return func(item *Item) error {
		item.Label = label
		return nil
	}
}

// WithPosition is a function that creates an Option for setting the Position field of an Item.
// The Position field represents the order in which the Item should be displayed.
// The option created by WithPosition takes an integer parameter representing the desired position.
// The option function sets the Position field of the provided Item to the specified position.
// Example usage:
// item := &Item{}
// option := WithPosition(1)
// option(item)
// After applying the option, the Item's Position field will be set to 1.
func WithPosition(position int) Option {
	return func(item *Item) error {
		item.Position = position
		return nil
	}
}

// WithDisplayChildren is a function that returns an Option, which is used to set the DisplayChildren field of an Item struct.
// Option is a function type that takes a pointer to an Item and returns an error.
// It is used to configure the properties of an Item.
func WithDisplayChildren(displayChildren bool) Option {
	return func(item *Item) error {
		item.DisplayChildren = displayChildren
		return nil
	}
}

// WithDisplay is a function that returns an Option for setting the display property of an Item. The display property determines whether or not the Item is displayed.
// Parameters:
//   - display: a boolean indicating whether the Item should be displayed (true) or hidden (false).
//
// Returns:
//   - an Option function that modifies the display property of an Item.
//
// Example usage:
//
//	item := &Item{}
//	option := WithDisplay(true)
//	option(item)
//
// Note: This function is part of the Option pattern, where an Option is a function that modifies the properties of an Item.
func WithDisplay(display bool) Option {
	return func(item *Item) error {
		item.Display = display
		return nil
	}
}

// WithCurrent takes a pointer to a bool as its argument and returns an Option.
// The returned Option function sets the Current field of the provided Item to the value of the provided bool pointer.
// It returns nil error.
// Example usage: opt := WithCurrent(&current)
func WithCurrent(current *bool) Option {
	return func(item *Item) error {
		if current == nil {
			item.Current = nil
		} else {
			*item.Current = *current
		}
		return nil
	}
}

// WithAttributes is a function that returns an Option for setting the attributes of an Item.
// It takes a map of attribute names to values and updates the Attributes field of the Item with those values.
// The Option is a function that takes a pointer to an Item and returns an error.
// It sets the Attributes field of the Item and returns nil.
//
// Example:
// item := &Item{Name: "example"}
// attributes := map[string]interface{}{"color": "red", "size": "large"}
// option := WithAttributes(attributes)
// err := option(item)
// // item.Attributes is now set to the provided attributes
func WithAttributes(attributes map[string]any) Option {
	return func(item *Item) error {
		item.Attributes = maps.Clone(attributes)
		return nil
	}
}

// WithAttribute sets the specified attribute name and value for an item.
// The attribute is stored in the item's Attributes map.
// The function returns an Option function, which can be used to apply the attribute to an item.
// An error is returned if the attribute cannot be set.
//
// Parameters:
// - name: the name of the attribute
// - value: the value of the attribute
//
// Returns:
// - Option: the option function to apply the attribute to an item
//
// Example Usage:
// item := &Item{Name: "Example"}
// opt := WithAttribute("color", "red")
// err := opt(item)
//
//	if err != nil {
//	    log.Fatal(err)
//	}
//
// fmt.Println(item) // Output: {Name: "Example", Attributes: {"color": "red"}}
func WithAttribute(name string, value any) Option {
	return func(item *Item) error {
		item.Attributes[name] = value
		return nil
	}
}

// WithLinkAttributes is a function that returns an Option function to set the link attributes of an Item.
// The link attributes are specified as a map[string]any, where the keys are attribute names and the values are attribute values.
//
// Usage example:
// attributes := map[string]any{"class": "link", "target": "_blank"}
// option := WithLinkAttributes(attributes)
// item := &Item{}
// option(item)
// fmt.Println(item.LinkAttributes) // Output: map[class:link target:_blank]
func WithLinkAttributes(attributes map[string]any) Option {
	return func(item *Item) error {
		item.LinkAttributes = maps.Clone(attributes)
		return nil
	}
}

// WithLinkAttribute is a function that defines an option for modifying the link attributes of an Item. It adds or updates the specified attribute and its value in the LinkAttributes
func WithLinkAttribute(name string, value any) Option {
	return func(item *Item) error {
		item.LinkAttributes[name] = value
		return nil
	}
}

// WithChildrenAttributes is a function that returns an option to set the children attributes of an Item.
//
// The attributes are provided as a map[string]any, where the keys are the attribute names and the values are
// the attribute values. These attributes will be set for the children of the Item.
//
// The returned option function modifies the provided Item by setting its ChildrenAttributes field to the
// provided attributes.
//
// Example usage:
//
//	attributes := map[string]any {
//	    "color": "blue",
//	    "size": 12,
//	}
//	option := WithChildrenAttributes(attributes)
//	item := &Item{}
//	option(item)
func WithChildrenAttributes(attributes map[string]any) Option {
	return func(item *Item) error {
		item.ChildrenAttributes = maps.Clone(attributes)
		return nil
	}
}

// WithChildrenAttribute is a function that creates an Option to add a children attribute to an Item.
// It takes a name string and a value of type any as parameters.
// The Option function updates the item's ChildrenAttributes field with the given name and value, and returns nil.
//
// Example usage:
// item := &Item{}
// option := WithChildrenAttribute("color", "red")
// err := option(item)
//
//	if err != nil {
//	 	fmt.Println("Error:", err)
//	}
func WithChildrenAttribute(name string, value any) Option {
	return func(item *Item) error {
		item.ChildrenAttributes[name] = value
		return nil
	}
}

// WithLabelAttributes sets the label attributes of an item with the given attributes. It returns an Option type function that can be used to modify an item.
// The attributes are provided as a map[string]any, where the keys are the attribute names and the values are the attribute values.
//
// Example usage:
// item := &Item{}
// attributes := map[string]any{"class": "label"}
// option := WithLabelAttributes(attributes)
// option(item)
//
// Parameters:
// - attributes: The label attributes to be set for the item.
//
// Returns:
// - An Option function that sets the label attributes of the item.
func WithLabelAttributes(attributes map[string]any) Option {
	return func(item *Item) error {
		item.LabelAttributes = maps.Clone(attributes)
		return nil
	}
}

// WithLabelAttribute is a function that creates an Option to add a label attribute to an Item.
// It takes a name string and a value any as parameters.
// The returned Option function adds the label attribute to the given Item by setting the value under the specified name in the LabelAttributes map.
// It returns an error if there is an issue adding the label attribute.
// Example usage:
//
//	option := WithLabelAttribute("color", "red")
//	item := &Item{}
//	err := option(item)
func WithLabelAttribute(name string, value any) Option {
	return func(item *Item) error {
		item.LabelAttributes[name] = value
		return nil
	}
}

// WithExtras is a function that returns an Option which sets the Extras field of an Item. The Extras field is a map[string]any that contains any additional data associated with the
func WithExtras(extras map[string]any) Option {
	return func(item *Item) error {
		item.Extras = maps.Clone(extras)
		return nil
	}
}

// WithExtra adds extra information to an Item.
// The name parameter specifies the key of the extra information.
// The value parameter specifies the value of the extra information.
// It returns an Option function that can be used to apply the extra information to an Item.
func WithExtra(name string, value any) Option {
	return func(item *Item) error {
		item.Extras[name] = value
		return nil
	}
}

// WithSafeLabel is a function that returns an Option for setting the "safe_label" extra attribute of an Item.
func WithSafeLabel(safeLabel bool) Option {
	return WithExtra("safe_label", safeLabel)
}

// WithParent is an option function that sets the parent of an Item.
// It takes a pointer to an Item as a parameter and assigns the given parent to it.
// It returns an error if any error occurs during the assignment.
func WithParent(parent *Item) Option {
	return func(item *Item) error {
		item.Parent = parent
		return nil
	}
}

// WithChildren is a function that returns an Option for setting the children of an Item object.
// It takes a slice of *Item as the children parameter and an optional variadic parameter options of type Option.
// The function iterates over the children slice and adds each child to the Item object using the AddChild method.
// The children are added in the order provided in the slice.
// If any error occurs during the addition of a child, the function stops adding any more children and returns the error.
// If all children are successfully added, the function returns nil indicating no error.
// The Children field of the Item object is set to a new empty slice with the capacity equal to the length of the children slice before adding any children.
// The function signature is:
//
//	func WithChildren(children []*Item, options ...Option) Option {
//	    ...
//	}
func WithChildren(children []*Item, options ...Option) Option {
	return func(item *Item) error {
		item.Children = make([]*Item, 0, len(children))
		for _, child := range children {
			if _, err := item.AddChild(child, options...); err != nil {
				return err
			}
		}
		return nil
	}
}

// WithChild is a function that returns an Option for adding a child Item to another Item.
// The child Item is passed as the first argument to WithChild, and additional options can be provided
// as variadic arguments. The child Item is added to the parent Item's Children slice.
// If the child Item already belongs to another parent, an error of type ErrItemBelongsToAnotherMenu
// is returned.
func WithChild(child *Item, options ...Option) Option {
	return func(item *Item) error {
		_, err := item.AddChild(child, options...)
		return err
	}
}
