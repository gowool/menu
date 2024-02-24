package menu

import (
	"context"
	"net/url"
)

// Voter represents an interface for determining whether an item is current.
//
// MatchItem checks whether an item is current.
//
// If the voter is not able to determine a result,
// it should return nil to let other voters do the job.
type Voter interface {
	// MatchItem checks whether an item is current.
	//
	// If the voter is not able to determine a result,
	// it should return nil to let other voters do the job.
	MatchItem(ctx context.Context, item *Item) *bool
}

// URLVoter represents a type that implements the Voter interface for determining whether an item's URI matches a given URI.
// MatchItem checks whether an item's URI matches the URI provided in the context.
//
// If the URLVoter is not able to determine a result,
// it should return nil to let other voters do the job.
// Usage example:
//
// ```go
//
//	func (v URLVoter) MatchItem(ctx context.Context, item *Item) *bool {
//		if _url, ok := ctx.Value("url").(*url.URI); ok && _url.Path == item.URI {
//			return &ok
//		}
//		return nil
//	}
//
// ```
type URLVoter struct{}

// MatchItem is a method of the URLVoter type that checks if the URI of an Item matches with the URI stored in the context.
// If the URLs match, it returns a pointer to a boolean value set to true. Otherwise, it returns nil.
// It takes in a context.Context and a pointer to an Item as parameters.
// The context should contain a value with the key "url" that is of type *url.URL.
// The item's URI is compared with the URI from the context's value.
//
// Example usage:
//
//	item := &Item{URI: "/example"}
//	url, _ := url.Parse("/example")
//	ctx := context.WithValue(context.Background(), "url", url)
//	result := urlVoter.MatchItem(ctx, item)
//	if result != nil && *result {
//	    fmt.Println("URLs match!")
//	}
func (v URLVoter) MatchItem(ctx context.Context, item *Item) *bool {
	if _url, ok := ctx.Value("url").(*url.URL); ok && _url.Path == item.URI {
		return &ok
	}
	return nil
}
