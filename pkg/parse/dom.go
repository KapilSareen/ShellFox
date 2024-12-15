package parse
import (
	"fmt"
)

type Node struct {
    Children []*Node  
    NodeType NodeType 
}
type NodeType interface{}

type TextNode struct {
	Text string
}

type ElementNode struct {
	Data ElementData
}

type ElementData struct {
	TagName string
	Attr map[string]string
}

func Text(data string) *Node {
    return &Node{
        Children: []*Node{}, 
        NodeType: TextNode{Text: data}, 
    }
}

func Element(tagName string, attr map[string]string, children ...*Node) *Node {
	return &Node{
		Children: children, 
		NodeType: ElementNode{Data: ElementData{TagName: tagName, Attr: attr}}, 
	}
}


func (n *Node) PrintDOM(indent string) {
    switch nodeType := n.NodeType.(type) {
    case TextNode:
        fmt.Printf("%s%s\n", indent, nodeType.Text)
    case ElementNode:
        fmt.Printf("%s<%s", indent, nodeType.Data.TagName)
        if len(nodeType.Data.Attr) > 0 {
            for key, value := range nodeType.Data.Attr {
                fmt.Printf(" %s=\"%s\"", key, value)
            }
        }
        if len(n.Children) > 0 {
            fmt.Println(">")
            for _, child := range n.Children {
                child.PrintDOM(indent + "│  ")
            }
            fmt.Printf("%s└</%s>\n", indent, nodeType.Data.TagName)
        } else {
            fmt.Println(" />")
        }
    }
}

