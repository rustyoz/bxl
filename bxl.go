package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

type Decoder struct {
	Source_buffer []byte
	Is_filled     bool
	Max_level     int
	Node_count    int
	Leaf_count    int
	Level         int
	Root          *Node
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

func (d *Decoder) NewNode(parent *Node, symbol int) *Node {
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
			n.Right = n.decoder.NewNode(n, -1)
			return n.Right
		}
		if n.Left == nil {
			n.Left = n.decoder.NewNode(n, -1)
			return n.Left
		}
		return ret
	} else {
		if n.Right == nil {
			n.Right = n.decoder.NewNode(n, symbol)
			return n.Right
		} else if n.Left == nil {
			n.Left = n.decoder.NewNode(n, symbol)
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

func SwapNodes(n1 *Node, n2 *Node, n3 *Node) {
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

func (d *Decoder) CreateTree() {
	var n = d.NewNode(nil, 0)
	d.Root = n
	for n != nil {
		n = d.Root.addChild(d.Leaf_count)
		if n != nil && n.isLeaf() {
			d.Leaf_count = d.Leaf_count + 1
		}
	}
}

func (d *Decoder) nextbit() byte {
	var result byte
	if d.source_bit < 0 {
		d.source_bit = 7
		d.source_char = d.Source_buffer[d.source_index]
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
		SwapNodes(grand_parent, parent, current)
		SwapNodes(grand_parent, parent_sibling, parent)
		SwapNodes(parent, current, parent_sibling)
		parent.weight = parent.Right.weight + parent.Left.weight
		grand_parent.weight = current.weight + parent.weight
		updateTree(parent)
		updateTree(grand_parent)
		updateTree(current)
	}
}

func (d *Decoder) decode() string {
	var output bytes.Buffer

	d.source_index = 4
	fmt.Println("input bytes: ", len(d.Source_buffer))
	for d.source_index < len(d.Source_buffer) {
		var bits int
		var node = d.Root
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
	}
	return output.String()
}

func main() {

	file, err := os.Open(os.Args[1])

	if err != nil {
		log.Fatal(err)

	}

	var d Decoder
	d.CreateTree()
	d.Source_buffer, err = ioutil.ReadAll(file)
	if err != nil {
		log.Fatalln("Error reading File: ", os.Args[1])
	}

	os.Create(os.Args[1] + ".txt")
	outfile, err := os.OpenFile(os.Args[1]+".txt", os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		fmt.Println(err)
		log.Fatalln("Error opening file: ", os.Args[1]+".txt")
	}
	output := d.decode()

	var characters int
	characters, err = outfile.WriteString(output)
	fmt.Println("Characters: ", characters)
	if err != nil {
		fmt.Println(err)
	}
	outfile.Close()

}
