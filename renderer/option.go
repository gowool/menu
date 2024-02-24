package renderer

// Option represents a function that modifies an *Options object.
//
// Usage example:
//
//	func WithDepth(depth *int) Option {
//	    return func(options *Options) {
//	        if depth == nil {
//	            options.Depth = nil
//	        } else {
//	            *options.Depth = *depth
//	        }
//	    }
//	}
//
//	opt := &Options{}
//	WithDepth(5)(opt)
type Option func(*Options)

// WithDepth is a function that sets the value of the Depth field in the Options struct. It takes a pointer to an int as a parameter and returns an Option function.
// The returned Option function updates the Options struct by assigning the value of the depth parameter to the Depth field, or setting the Depth field to nil if the depth parameter
func WithDepth(depth *int) Option {
	return func(options *Options) {
		if depth == nil {
			options.Depth = nil
		} else {
			*options.Depth = *depth
		}
	}
}

// WithMatchingDepth is a function that returns an Option for setting the matching depth of an Options object.
// The matching depth determines the maximum number of matching elements in the search hierarchy.
// If the given matchingDepth is nil, the Options' matching depth will be set to nil, meaning there is no maximum depth.
// Otherwise, the Options' matching depth will be set to the value of the matchingDepth pointer.
// If matchingDepth is greater than zero, the Options' matching depth will be decreased by one.
// The Options is passed by reference and modified directly.
func WithMatchingDepth(matchingDepth *int) Option {
	return func(options *Options) {
		if matchingDepth == nil {
			options.MatchingDepth = nil
		} else {
			*options.MatchingDepth = *matchingDepth
		}
	}
}

// WithCurrentClass is a function that returns an Option function. The returned Option function sets the CurrentClass field of an Options struct.
// Usage example:
// options := &Options{}
// WithCurrentClass("className")(options)
func WithCurrentClass(currentClass string) Option {
	return func(options *Options) {
		options.SetCurrentClass(currentClass)
	}
}

// WithAncestorClass is a function that creates an Option to set the ancestor class in the Options struct.
// It takes a string parameter, ancestorClass, and returns an Option.
// The returned Option sets the ancestorClass field in the Options struct.
//
// Example usage:
//
//	options := &Options{}
//	option := WithAncestorClass("AncestorClass")
//	option(options)
func WithAncestorClass(ancestorClass string) Option {
	return func(options *Options) {
		options.SetAncestorClass(ancestorClass)
	}
}

// WithFirstClass returns an Option function that sets the FirstClass field of the Options struct.
func WithFirstClass(firstClass string) Option {
	return func(options *Options) {
		options.SetFirstClass(firstClass)
	}
}

// WithLastClass is a function that creates an Option for setting the LastClass field in the Options struct.
// It takes a string parameter representing the last class and returns an Option function.
// The returned Option function sets the LastClass field in the Options struct when called.
// Example usage:
//
//	WithLastClass("lastClass") // returns an Option function to set LastClass field
//	WithLastClass("lastClass")(options) // sets LastClass field in options
func WithLastClass(lastClass string) Option {
	return func(options *Options) {
		options.SetLastClass(lastClass)
	}
}

// WithLeafClass is a function that returns an Option to set the leafClass field of Options.
// It takes a string parameter representing the leaf class and returns an Option function that sets the leafClass field to the provided value.
// The Options type represents a set of configuration options.
// The leafClass field is used to specify the leaf class value.
// Usage:
//
//	leafClassOption := WithLeafClass("exampleLeafClass")
//	options := &Options{}
//	leafClassOption(options)
//
//	// Alternative usage
//	options := &Options{}
//	options.Apply(WithLeafClass("exampleLeafClass"))
//
// The Options type has other fields and methods that can be used to configure additional options and apply a set of options to an Options object.
// For more information and examples, refer to the documentation for Options and other Option functions.
func WithLeafClass(leafClass string) Option {
	return func(options *Options) {
		options.SetLeafClass(leafClass)
	}
}

// WithBranchClass is a function that creates an Option to set the BranchClass field of the Options struct.
// It takes in a string parameter branchClass, and returns a function that sets the BranchClass field of the Options struct to the provided value.
func WithBranchClass(branchClass string) Option {
	return func(options *Options) {
		options.SetBranchClass(branchClass)
	}
}

// WithCurrentAsLink is a function that returns an Option, which sets the value of the CurrentAsLink field in the Options struct.
// The CurrentAsLink field determines whether the current node in a tree structure should be treated as a link.
// If currentAsLink is true, the current node will be treated as a link, otherwise, it will not be treated as a link.
//
// Example usage:
// options := &Options{}
// opt := WithCurrentAsLink(true)
// opt(options)
//
// This will set CurrentAsLink to true in the options object.
func WithCurrentAsLink(currentAsLink bool) Option {
	return func(options *Options) {
		options.SetCurrentAsLink(currentAsLink)
	}
}

// WithAllowSafeLabels is a function that returns an Option for setting the AllowSafeLabels field in the Options struct.
func WithAllowSafeLabels(allowSafeLabels bool) Option {
	return func(options *Options) {
		options.SetAllowSafeLabels(allowSafeLabels)
	}
}

// WithClearMatcher is a function that returns an Option function. The Option function sets the ClearMatcher field of the Options struct to the provided value.
// Usage example:
// options := &Options{}
// clearMatcherOption := WithClearMatcher(true)
// clearMatcherOption(options)
func WithClearMatcher(clearMatcher bool) Option {
	return func(options *Options) {
		options.SetClearMatcher(clearMatcher)
	}
}

// WithExtras is a function that returns an Option for setting the Extras field in the Options struct.
// It takes a map[string]any as input and sets the Extras field in the Options struct to the provided map.
// Usage example:
//
//	extras := map[string]any{"key1": value1, "key2": value2}
//	withExtras := WithExtras(extras)
//	options := &Options{}
//	withExtras(options)
func WithExtras(extras map[string]any) Option {
	return func(options *Options) {
		options.SetExtras(extras)
	}
}

// WithExtra is a function that creates an Option which adds an extra value to the Options struct.
// The extra value is stored in the Extras map with the specified name.
//
// Parameters:
//   - name: the name of the extra value.
//   - value: the value of the extra.
//
// Returns:
//   - Option: an Option function that adds the extra value to the Options struct.
//
// Example:
//
//	options := &Options{}
//	extraOption := WithExtra("key", "value")
//	extraOption(options)
//	// Now options.Extras["key"] contains "value"
//
// Note:
//
//	The Options struct is modified in-place by calling the Option function.
//	To apply multiple options at once, use the Apply method of Options.
func WithExtra(name string, value any) Option {
	return func(options *Options) {
		options.AddExtra(name, value)
	}
}
