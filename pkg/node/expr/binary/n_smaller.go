package binary

import (
	"github.com/z7zmey/php-parser/pkg/freefloating"
	"github.com/z7zmey/php-parser/pkg/node"
	"github.com/z7zmey/php-parser/pkg/position"
	"github.com/z7zmey/php-parser/pkg/walker"
)

// Smaller node
type Smaller struct {
	FreeFloating freefloating.Collection
	Position     *position.Position
	Left         node.Node
	Right        node.Node
}

// NewSmaller node constructor
func NewSmaller(Variable node.Node, Expression node.Node) *Smaller {
	return &Smaller{
		FreeFloating: nil,
		Left:         Variable,
		Right:        Expression,
	}
}

// SetPosition sets node position
func (n *Smaller) SetPosition(p *position.Position) {
	n.Position = p
}

// GetPosition returns node positions
func (n *Smaller) GetPosition() *position.Position {
	return n.Position
}

func (n *Smaller) GetFreeFloating() *freefloating.Collection {
	return &n.FreeFloating
}

// Attributes returns node attributes as map
func (n *Smaller) Attributes() map[string]interface{} {
	return nil
}

// Walk traverses nodes
// Walk is invoked recursively until v.EnterNode returns true
func (n *Smaller) Walk(v walker.Visitor) {
	if v.EnterNode(n) == false {
		return
	}

	if n.Left != nil {
		v.EnterChildNode("Left", n)
		n.Left.Walk(v)
		v.LeaveChildNode("Left", n)
	}

	if n.Right != nil {
		v.EnterChildNode("Right", n)
		n.Right.Walk(v)
		v.LeaveChildNode("Right", n)
	}

	v.LeaveNode(n)
}