package bxlbin

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
)

func NewDecoder() *Decoder {
	var d Decoder
	d.NewTree()
	return &d
}

type Decoder struct {
	source_buffer []byte
	is_filled     bool
	max_level     int
	node_count    int
	leaf_count    int
	level         int
	root          *Node
	source_index  int
	source_char   byte
	source_bit    int
}

type Node struct {
	decoder *Decoder
	level   int
	Parent  *Node
	Left    *Node
	Right   *Node
	symbol  int
	weight  int
}

func (d *Decoder) newNode(parent *Node, symbol int) *Node {
	var n Node
	n.symbol = -1
	n.decoder = d
	if parent != nil {
		n.Parent = parent
		n.level = parent.level + 1
	} else {
		n.level = 0
	}
	if n.level > 7 {
		n.symbol = symbol
	}
	return &n
}

func (n *Node) addChild(symbol int) (ret *Node) {
	if n.level < 7 {
		if n.Right != nil {
			ret = n.Right.addChild(symbol)
			if ret != nil {
				return ret
			}
		}
		if n.Left != nil {
			ret = n.Left.addChild(symbol)
			if ret != nil {
				return ret
			}
		}
		if n.Right == nil {
			n.Right = n.decoder.newNode(n, -1)
			return n.Right
		}
		if n.Left == nil {
			n.Left = n.decoder.newNode(n, -1)
			return n.Left
		}
		return ret
	} else {
		if n.Right == nil {
			n.Right = n.decoder.newNode(n, symbol)
			return n.Right
		} else if n.Left == nil {
			n.Left = n.decoder.newNode(n, symbol)
			return n.Left
		} else {
			return nil
		}

	}
}

func (n *Node) isLeaf() bool {
	return n.level > 7
}

func (n *Node) sibling(s *Node) *Node {
	if s == n.Left {
		return n.Right
	} else {
		return n.Left
	}
}

func (n *Node) needSwapping() bool {
	if n.Parent != nil && n.Parent.Parent != nil && n.weight > n.Parent.weight {
		return true
	}
	return false
}

func swapNodes(n1 *Node, n2 *Node, n3 *Node) {
	if n3 != nil {
		n3.Parent = n1
	}
	if n1.Right == n2 {
		n1.Right = n3
		return
	}
	if n1.Left == n2 {
		n1.Left = n3
		return
	}

}

func (d *Decoder) NewTree() {
	var n = d.newNode(nil, 0)
	d.root = n
	for n != nil {
		n = d.root.addChild(d.leaf_count)
		if n != nil && n.isLeaf() {
			d.leaf_count = d.leaf_count + 1
		}
	}
}

func (d *Decoder) nextbit() byte {
	var result byte
	if d.source_bit < 0 {
		d.source_bit = 7
		d.source_char = d.source_buffer[d.source_index]
		result = d.source_char & (1 << uint(d.source_bit))
		d.source_index = d.source_index + 1
	} else {
		result = d.source_char & (1 << uint(d.source_bit))
	}
	d.source_bit = d.source_bit - 1

	return result
}

func updateTree(current *Node) {
	if current.Parent != nil && current.needSwapping() {
		var parent = current.Parent
		var grand_parent = parent.Parent
		var parent_sibling = grand_parent.sibling(parent)
		swapNodes(grand_parent, parent, current)
		swapNodes(grand_parent, parent_sibling, parent)
		swapNodes(parent, current, parent_sibling)
		parent.weight = parent.Right.weight + parent.Left.weight
		grand_parent.weight = current.weight + parent.weight
		updateTree(parent)
		updateTree(grand_parent)
		updateTree(current)
	}
}

func (d *Decoder) Decode(input io.Reader) (string, error) {
	var err error
	d.source_buffer, err = ioutil.ReadAll(input)
	if err != nil {
		return "", fmt.Errorf("Decode: error reading input", err)
	}
	var output bytes.Buffer

	d.source_index = 4

	var outchar int
	outchar = uncompressed_size(d.source_buffer[0:4])
	for d.source_index < (len(d.source_buffer) - 1) {
		var bits int
		var node = d.root
		for !node.isLeaf() {
			if d.nextbit() != 0 {
				node = node.Left
				//	fmt.Print("Left ")

			} else {
				node = node.Right
				//	fmt.Print("Right ")
			}
			bits = bits + 1
		}
		output.WriteByte(byte(node.symbol))
		node.weight = node.weight + 1
		updateTree(node)
		outchar = outchar - 1
	}
	return output.String(), nil
}

func uncompressed_size(b []byte) int {
	/* Uncompressed size =
	B0b7 * 1<<0 + B0b6 * 1<<1 + ... + B0b0 * 1<<7 +
	B1b7 * 1<<0 + B1b6 * 1<<1 + ... + B2b0 * 1<<7 +
	B2b7 * 1<<0 + B2b6 * 1<<1 + ... + B3b0 * 1<<7 +
	B3b7 * 1<<0 + B3b6 * 1<<1 + ... + B4b0 * 1<<7
	*/
	var size int
	var mask uint
	var i int
	i = 7
	for i >= 0 {
		if b[0]&(1<<uint(i)) != 0 {
			size |= (1 << uint(mask))
		}
		mask = mask + 1
		i = i - 1
	}
	i = 7
	for i >= 0 {
		if b[1]&(1<<uint(i)) != 0 {
			size |= (1 << uint(mask))
		}
		mask = mask + 1
		i = i - 1
	}
	i = 7
	for i >= 0 {
		if b[2]&(1<<uint(i)) != 0 {
			size |= (1 << uint(mask))
		}
		mask = mask + 1
		i = i - 1
	}
	i = 7
	for i >= 0 {
		if b[3]&(1<<uint(i)) != 0 {
			size |= (1 << uint(mask))
		}
		mask = mask + 1
		i = i - 1
	}
	return size
}
