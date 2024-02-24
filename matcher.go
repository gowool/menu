package menu

import (
	"context"
	"sync"
)

var _ Matcher = (*CoreMatcher)(nil)

// Matcher represents an interface for matching items.
// It provides methods for checking whether an item is current or an ancestor.
// It also provides a method for clearing the state of the matcher.
type Matcher interface {
	// IsCurrent checks whether an item is current
	IsCurrent(ctx context.Context, item *Item) bool

	// IsAncestor checks whether an item is the ancestor of a current item
	IsAncestor(ctx context.Context, item *Item, depth *int) bool

	// Clear clears the state of the matcher
	Clear()
}

// CoreMatcher represents a matcher that determines the current state of an item.
type CoreMatcher struct {
	voters []Voter
	cache  map[*Item]bool
	mu     sync.RWMutex
}

// NewCoreMatcher creates a new instance of the CoreMatcher with the given voters.
// It initializes the cache with an empty map.
// The voters are used to determine whether an item is current.
// The CoreMatcher has the following methods:
// - IsCurrent: checks if an item is current based on the registered voters.
// - IsAncestor: checks if an item is an ancestor of any current items within a certain depth.
// - Clear: clears the cache.
//
// Example usage:
//
//	v := NewCoreMatcher(voter1, voter2)
//	isCurrent := v.IsCurrent(ctx, item)
//
// Parameters:
//   - voters: a list of Voter implementations.
//
// Returns:
//   - Pointer to the initialized CoreMatcher.
func NewCoreMatcher(voters ...Voter) *CoreMatcher {
	return &CoreMatcher{
		voters: voters,
		cache:  map[*Item]bool{},
	}
}

// IsCurrent checks whether an item is considered current.
//
// If the "Current" field of the item is not nil, it returns the value of the field.
// If the item is found in the cache, it returns the cached value.
// Otherwise, it iterates over the registered voters and calls the "MatchItem" method on each voter.
// If a voter returns a non-nil value, it considers the item as current and breaks the loop.
// It then caches the value and returns it.
func (m *CoreMatcher) IsCurrent(ctx context.Context, item *Item) bool {
	if item.Current != nil {
		return *item.Current
	}

	m.mu.RLock()
	if current, ok := m.cache[item]; ok {
		m.mu.RUnlock()
		return current
	}

	var current bool
	for _, voter := range m.voters {
		if v := voter.MatchItem(ctx, item); v != nil {
			current = *v
			break
		}
	}

	m.mu.RUnlock()
	m.mu.Lock()
	defer m.mu.Unlock()

	m.cache[item] = current
	return current
}

// IsAncestor checks whether the given item is an ancestor of any current item in the hierarchy, up to the specified depth.
// If the depth is not nil, it first checks if the depth is zero. If it is, it returns false.
// Then, it iterates over each child of the given item. If the child is a current item or an ancestor (recursive call to IsAncestor), it returns true.
// If none of the children match the condition, it returns false.
func (m *CoreMatcher) IsAncestor(ctx context.Context, item *Item, depth *int) bool {
	if depth != nil {
		if *depth == 0 {
			return false
		}
		*depth = *depth - 1
	}

	for _, child := range item.Children {
		if m.IsCurrent(ctx, child) || m.IsAncestor(ctx, child, depth) {
			return true
		}
	}
	return false
}

// Clear eliminates all the items from the cache map,
// synchronizing the access with a read-write lock.
func (m *CoreMatcher) Clear() {
	m.mu.Lock()
	defer m.mu.Unlock()

	m.cache = map[*Item]bool{}
}
