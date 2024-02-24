package renderer

import "maps"

type Options struct {
	Depth           *int           `json:"depth,omitempty"`
	MatchingDepth   *int           `json:"matching_depth,omitempty"`
	CurrentClass    string         `json:"current_class,omitempty"`
	AncestorClass   string         `json:"ancestor_class,omitempty"`
	FirstClass      string         `json:"first_class,omitempty"`
	LastClass       string         `json:"last_class,omitempty"`
	LeafClass       string         `json:"leaf_class,omitempty"`
	BranchClass     string         `json:"branch_class,omitempty"`
	CurrentAsLink   bool           `json:"current_as_link,omitempty"`
	AllowSafeLabels bool           `json:"allow_safe_labels,omitempty"`
	ClearMatcher    bool           `json:"clear_matcher,omitempty"`
	Extras          map[string]any `json:"extras,omitempty"`
}

// NewOptions creates a new instance of Options with default values and applies the provided options.
// The returned Options object can be further modified using the available setter methods.
// Example usage:
// options := NewOptions(
//
//	WithDepth(3),
//	WithCurrentClass("active"),
//	WithExtras(map[string]any{"key": "value"}),
//
// )
// myOptions := options.SetAncestorClass("ancestor")
// myOptions.SetFirstClass("first")
// myOptions.SetLastClass("last")
// ...
// Usage within a function:
// renderer := NewTemplateRenderer(
//
//	myTheme,
//	myMatcher,
//	NewOptions(WithDepth(2), WithCurrentClass("active")),
//
// )
// ...
// renderer := NewListRenderer(
//
//	myMatcher,
//	NewOptions(WithCurrentClass("active"), WithLeafClass("leaf")),
//
// )
func NewOptions(options ...Option) *Options {
	o := &Options{
		CurrentClass:  "current",
		AncestorClass: "current-ancestor",
		FirstClass:    "first",
		LastClass:     "last",
		CurrentAsLink: true,
		ClearMatcher:  true,
		Extras:        map[string]any{},
	}
	return o.Apply(options...)
}

// SubDepth decrements the value of the Depth field in the Options struct by 1 if it is not nil.
// It then returns a pointer to the modified Options struct.
// Example usage: options.SubDepth().SubMatchingDepth()
func (o *Options) SubDepth() *Options {
	if o.Depth != nil {
		*o.Depth--
	}
	return o
}

// SetDepth sets the `Depth` field of the `Options` struct to the specified value and returns a pointer to the `Options` struct.
// Parameters:
// - depth: An integer value representing the depth.
// Returns:
// - *Options: A pointer to the `Options` struct with the updated `Depth` field.
func (o *Options) SetDepth(depth int) *Options {
	o.Depth = &depth
	return o
}

// IsStop returns true if the depth of the Options is less than or equal to zero, indicating that
// the processing should stop.
func (o *Options) IsStop() bool {
	return o.Depth != nil && *o.Depth <= 0
}

// SubMatchingDepth decreases the value of o.MatchingDepth by 1 if it is not nil and greater than 0.
func (o *Options) SubMatchingDepth() *Options {
	if o.MatchingDepth != nil && *o.MatchingDepth > 0 {
		*o.MatchingDepth--
	}
	return o
}

// SetMatchingDepth sets the value of the MatchingDepth field in the Options struct.
// It takes an int as a parameter, matchingDepth, and assigns the address of that int to the MatchingDepth field.
// It then returns the pointer to the Options struct.
func (o *Options) SetMatchingDepth(matchingDepth int) *Options {
	o.MatchingDepth = &matchingDepth
	return o
}

// SetCurrentClass sets the value of the CurrentClass field in the Options struct and returns the modified Options struct.
func (o *Options) SetCurrentClass(currentClass string) *Options {
	o.CurrentClass = currentClass
	return o
}

// SetAncestorClass sets the value of the AncestorClass field in the Options struct
// and returns a pointer to the modified Options struct.
// The AncestorClass field represents the class of the ancestor node in a tree structure.
//
// Usage:
//
//	options := &Options{}
//	options.SetAncestorClass("ancestorClass")
//
//	// Usage in Apply method
//	options.Apply(WithAncestorClass("ancestorClass"))
//
// Parameters:
//   - ancestorClass: The class of the ancestor node.
//
// Returns:
//
//	A pointer to the modified Options struct.
func (o *Options) SetAncestorClass(ancestorClass string) *Options {
	o.AncestorClass = ancestorClass
	return o
}

// SetFirstClass sets the value of the FirstClass field in the Options struct and returns a pointer to the updated Options struct.
func (o *Options) SetFirstClass(firstClass string) *Options {
	o.FirstClass = firstClass
	return o
}

// SetLastClass sets the value of the LastClass field in the Options struct and returns a pointer to the modified Options struct.
func (o *Options) SetLastClass(lastClass string) *Options {
	o.LastClass = lastClass
	return o
}

// SetLeafClass sets the value of the `LeafClass` field in the `Options` struct and returns a pointer to the modified `Options` object.
func (o *Options) SetLeafClass(leafClass string) *Options {
	o.LeafClass = leafClass
	return o
}

// SetBranchClass sets the branch class in the Options struct and returns a pointer to Options.
// It takes a string parameter called branchClass.
// Example usage:
// opt := &Options{}
// opt.SetBranchClass("Branch")
// Returns:
// &Options{..., BranchClass: "Branch", ...}
func (o *Options) SetBranchClass(branchClass string) *Options {
	o.BranchClass = branchClass
	return o
}

// SetCurrentAsLink sets the `currentAsLink` field of the `Options` struct and returns the modified `Options` object.
// Setting `currentAsLink` to `true` indicates that the current node should be rendered as a link.
func (o *Options) SetCurrentAsLink(currentAsLink bool) *Options {
	o.CurrentAsLink = currentAsLink
	return o
}

// SetAllowSafeLabels sets the value of AllowSafeLabels in the Options struct and returns a pointer to the modified Options struct.
func (o *Options) SetAllowSafeLabels(allowSafeLabels bool) *Options {
	o.AllowSafeLabels = allowSafeLabels
	return o
}

// SetClearMatcher sets the `ClearMatcher` field in the `Options` struct and returns a pointer to the modified struct.
//
// Example:
// opts := &Options{}
// opts.SetClearMatcher(true)
//
// The `SetClearMatcher` method can be used in conjunction with the `WithClearMatcher` function to set the `ClearMatcher` field in the `Options` struct using functional options.
// Example:
// opts := &Options{}
// opts.Apply(WithClearMatcher(true))
//
// Parameters:
//
//	clearMatcher: A boolean value indicating whether to clear the matcher. If true, the matcher will be cleared; otherwise, it will not be cleared.
//
// Returns:
//
//	*Options: A pointer to the modified `Options` struct.
func (o *Options) SetClearMatcher(clearMatcher bool) *Options {
	o.ClearMatcher = clearMatcher
	return o
}

// SetExtras sets the extras map for the Options object.
// If the provided extras map is nil, it sets an empty map for extras.
// Otherwise, it clones the provided extras map and sets it as extras.
// Returns a pointer to the Options object.
func (o *Options) SetExtras(extras map[string]any) *Options {
	if extras == nil {
		o.Extras = map[string]any{}
	} else {
		o.Extras = maps.Clone(extras)
	}
	return o
}

// AddExtra adds an extra value to the Options.Extras map.
//
// Parameters:
// - name: the name of the extra value.
// - value: the value to be added.
//
// Returns:
// - *Options: the Options object with the extra value added.
func (o *Options) AddExtra(name string, value any) *Options {
	o.Extras[name] = value
	return o
}

// Extra returns the value of the specified extra property from the Options struct. If the property is not found, it returns the default value.
func (o *Options) Extra(name string, def ...any) any {
	if value, ok := o.Extras[name]; ok {
		return value
	}
	if len(def) > 0 {
		return def[0]
	}
	return nil
}

// Copy creates a copy of the Options object.
// It creates a new Options object and copies the values from the original object.
// If Depth is not nil, it creates a new int variable and assigns the value of Depth to it.
// It assigns a pointer to the new int variable to the new Options object's Depth field.
// If MatchingDepth is not nil, it creates a new int variable and assigns the value of MatchingDepth to it.
// It assigns a pointer to the new int variable to the new Options object's MatchingDepth field.
// It clones the Extras field using the maps.Clone function and assigns the cloned map to the new Options object's Extras field.
// It returns a pointer to the new Options object.
func (o *Options) Copy() *Options {
	newOptions := *o

	if o.Depth != nil {
		depth := *o.Depth
		newOptions.Depth = &depth
	}
	if o.MatchingDepth != nil {
		depth := *o.MatchingDepth
		newOptions.MatchingDepth = &depth
	}
	newOptions.Extras = maps.Clone(o.Extras)

	return &newOptions
}

// Apply applies the given list of options to the Options object.
// It iterates over the list of options and calls each option passing the Options object as an argument.
// Returns the modified Options object.
func (o *Options) Apply(options ...Option) *Options {
	for _, option := range options {
		option(o)
	}
	return o
}

// Slice returns a slice of Option functions that correspond to the current state of the Options object.
func (o *Options) Slice() []Option {
	return []Option{
		WithDepth(o.Depth),
		WithMatchingDepth(o.MatchingDepth),
		WithCurrentClass(o.CurrentClass),
		WithAncestorClass(o.AncestorClass),
		WithFirstClass(o.FirstClass),
		WithLastClass(o.LastClass),
		WithLeafClass(o.LeafClass),
		WithBranchClass(o.BranchClass),
		WithAllowSafeLabels(o.AllowSafeLabels),
		WithClearMatcher(o.ClearMatcher),
		WithExtras(o.Extras),
	}
}
