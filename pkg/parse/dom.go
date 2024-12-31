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


func (n *Node) DOMString(indent string) string {
    var result string
    switch nodeType := n.NodeType.(type) {
    case TextNode:
        result += fmt.Sprintf("%s%s\n", indent, nodeType.Text)
    case ElementNode:
        result += fmt.Sprintf("%s<%s", indent, nodeType.Data.TagName)
        if len(nodeType.Data.Attr) > 0 {
            for key, value := range nodeType.Data.Attr {
                result += fmt.Sprintf(" %s=\"%s\"", key, value)
            }
        }
        if len(n.Children) > 0 {
            result += ">\n"
            for _, child := range n.Children {
                result += child.DOMString(indent + "│  ")
            }
            result += fmt.Sprintf("%s└</%s>\n", indent, nodeType.Data.TagName)
        } else {
            result += " />\n"
        }
    }
    return result
}
