package menu

type Node interface {
	Name() string
	Options() []Option
	Children() []Node
}

type SimpleNode struct {
	name     string
	options  []Option
	children []Node
}

func NewSimpleNode(name string, options []Option, children []Node) SimpleNode {
	return SimpleNode{
		name:     name,
		options:  options,
		children: children,
	}
}

func (n SimpleNode) Name() string {
	return n.name
}

func (n SimpleNode) Options() []Option {
	return n.options
}

func (n SimpleNode) Children() []Node {
	return n.children
}
