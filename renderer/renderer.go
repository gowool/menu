package renderer

import (
	"context"

	"github.com/gowool/menu"
)

// Renderer is an interface for rendering menu items.
//
// Usage Example:
//
//	_ Renderer = ListRenderer{}
type Renderer interface {
	Render(ctx context.Context, item *menu.Item, options ...Option) (string, error)
}
