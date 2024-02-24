package menu

import (
	"context"
	"errors"
	"fmt"
)

// ErrUnsupported represents an error indicating unsupported data.
var ErrUnsupported = errors.New("unsupported data")

// Loader is an interface that represents a data loader.
type Loader interface {
	// Load loads the data and returns an *Item and an error. The ctx parameter is a context.Context object that can be used for cancellation or propagation of deadlines.
	Load(ctx context.Context, data any) (*Item, error)

	// Supports checks if the loader supports the given data. It returns `true` if the loader supports the data, otherwise `false`.
	Supports(data any) bool
}

// NodeLoader represents a data loader for nodes.
type NodeLoader struct{}

// NewNodeLoader returns a new instance of NodeLoader.
func NewNodeLoader() NodeLoader {
	return NodeLoader{}
}

// Load processes the given data and returns a new Item representing the loaded data and its children, if any. If the data is not of type Node, an error is returned. The context.Context
func (l NodeLoader) Load(ctx context.Context, data any) (*Item, error) {
	node, ok := data.(Node)
	if !ok {
		return nil, fmt.Errorf("%w: expected Node, got %T", ErrUnsupported, data)
	}

	item, err := NewItem(node.Name(), node.Options()...)
	if err != nil {
		return nil, err
	}

	for _, childNode := range node.Children() {
		child, err := l.Load(ctx, childNode)
		if err != nil {
			return nil, err
		}

		if _, err = item.AddChild(child); err != nil {
			return nil, err
		}
	}

	return item, nil
}

// Supports checks if the given data is of type Node. Returns true if it is, false otherwise.
func (l NodeLoader) Supports(data any) bool {
	_, ok := data.(Node)
	return ok
}
