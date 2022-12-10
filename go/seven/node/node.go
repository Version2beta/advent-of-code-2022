package node

import (
	"fmt"
)

type NodeType string

const (
	Directory NodeType = "dir"
	File      NodeType = "file"
)

type Node struct {
	Type     NodeType
	name     string
	Size     int
	ancestor *Node
	children []*Node
}

func (n *Node) Root() *Node {
	for {
		if n.IsRoot() {
			break
		}
		n = n.ancestor
	}
	return n
}

func (n *Node) Parent() *Node {
	if n.IsRoot() {
		return nil
	}
	return n.ancestor
}

func (n *Node) Ancestors() []*Node {
	ancestors := []*Node{}
	child := n
	for {
		if child.IsRoot() {
			break
		}
		ancestors = append(ancestors, child.ancestor)
		child = child.ancestor
	}
	return reverse(ancestors)
}

func (n *Node) AddChildren(children ...*Node) {
	for _, child := range children {
		n.children = append(n.children, child)
		child.ancestor = n
	}
}

func (n *Node) DirSizes() int {
	if n.Type != Directory {
		return n.Size
	}

	n.Size = 0
	for _, child := range n.children {
		n.Size = n.Size + child.DirSizes()
	}

	return n.Size
}

func (n *Node) Children() []*Node {
	if n.IsLeaf() {
		return nil
	}
	return n.children
}

func (n *Node) FindOrCreateChild(nodeType NodeType, name string, size int) *Node {
	for _, child := range n.children {
		if child.Type == nodeType && child.name == name && child.Size == size {
			return child
		}
	}
	n.AddChildren(&Node{Type: nodeType, name: name, Size: size})
	return n.children[len(n.children)-1]
}

func (n *Node) IsRoot() bool {
	return n.ancestor == nil
}

func (n *Node) IsLeaf() bool {
	return len(n.children) == 0
}

func (n *Node) Name() string {
	if n.Type == Directory {
		return n.name + "/"
	}
	return n.name
}

func (n *Node) Path() string {
	var path string
	for _, ancestor := range append(n.Ancestors(), n) {
		path = path + ancestor.Name()
	}
	return path + fmt.Sprintf(" (%d bytes)", n.Size)
}

func (n *Node) List() string {
	return n.list("")
}

func (n *Node) list(acc string) string {
	acc = acc + n.Path() + "\n"
	for _, child := range n.children {
		acc = child.list(acc)
	}
	return acc
}

func (n *Node) SizeFilter(filter int, acc int) int {
	for _, child := range n.children {
		if child.Type == Directory && filter >= child.Size {
			acc = acc + child.Size
		}
		acc = child.SizeFilter(filter, acc)
	}

	return acc
}

func (n *Node) FindSpace(totalDiskSpace, requiredDiskSpace int) *Node {
	haveDiskSpace := totalDiskSpace - n.Size
	needDiskSpace := requiredDiskSpace - haveDiskSpace
	return n.findSpace(haveDiskSpace, needDiskSpace, n)
}

func (n *Node) findSpace(have, need int, best *Node) *Node {
	for _, child := range n.children {
		if child.Type == Directory {
			fmt.Printf(
				"%d > %d && %d < %d == %t\n",
				child.Size, need, child.Size, best.Size,
				child.Size > need && child.Size < best.Size,
			)
			if child.Size > need && child.Size < best.Size {
				best = child
			}
			best = child.findSpace(have, need, best)
		}
	}
	return best
}

func reverse(s []*Node) []*Node {
	var rev []*Node
	for i := 1; i <= len(s); i++ {
		rev = append(rev, s[len(s)-i])
	}
	return rev
}
