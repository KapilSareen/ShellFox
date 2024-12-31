package parse

import (
	"fmt"
	"strings"
	"unicode"
	"unicode/utf8"
)
type Parser struct {
	position int
	input string
}

func (p *Parser) NextChar() (rune, int) {
    if p.position >= len(p.input) {
		return 0, 0 
    }
    r, size := utf8.DecodeRuneInString(p.input[p.position:])
    return r, size
}	

func (p *Parser) StartsWith(a string) bool {
	return strings.HasPrefix(p.input[p.position:], a)
}

func (p *Parser) EOF() bool {
	return p.position >= len(p.input)
}

func (p *Parser) Expect(s string) {
	if !p.StartsWith(s) {
		panic(fmt.Sprintf("Expected %q at position %d", s, p.position))
	}
	p.position += len(s)
}


// return current char and advance the position
func (p *Parser) ConsumeChar() rune {
	char, size := p.NextChar()
	p.position += size
	return char
}	

func (p *Parser) ConsumeWhile(test func(rune) bool) string {
	result := ""
	char, _:=p.NextChar()

	for !p.EOF() && test(char) { 
		result += string(p.ConsumeChar())
		char, _=p.NextChar()
	}
	return result
}

// consume all whitespace characters
func (p *Parser) ConsumeWhitespace(){
	p.ConsumeWhile(unicode.IsSpace)
}

func (p *Parser) ParseName() string{
    isAlphaNumeric := func(r rune) bool {
        return (r >= 'a' && r <= 'z') || (r >= 'A' && r <= 'Z') || (r >= '0' && r <= '9')
    }
    return p.ConsumeWhile(isAlphaNumeric)
}

// Parse a quoted attribute value 
func (p *Parser) ParseAttrValue() string {
	openQuote := p.ConsumeChar()
	if openQuote != '"' && openQuote != '\'' {
		panic(fmt.Sprintf("Expected quote at position %d", p.position))
	}
	value := p.ConsumeWhile(func(r rune) bool {
		return r != openQuote
	})
	p.Expect(string(openQuote))
	return value
}

func (p *Parser) ParseAttr() (string, string) {
	key := p.ParseName()
	p.ConsumeWhitespace()
	p.Expect("=")
	p.ConsumeWhitespace()
	value := p.ParseAttrValue()
	return key, value
}

func (p *Parser) ParseAttributes() map[string]string {
	attributes := make(map[string]string)
	for {
		p.ConsumeWhitespace()
		if p.StartsWith(">") || p.EOF() {
			break
		}
		key, value := p.ParseAttr()
		attributes[key] = value
	}
	return attributes
}

func (p *Parser) ParseElement() *Node {
	p.Expect("<")
	tagName := p.ParseName()
	attributes := p.ParseAttributes()
	p.Expect(">")
	children := p.ParseNodes()
	p.Expect("</")
	p.ParseName()
	p.Expect(">")
	return Element(tagName, attributes, children...)
}

func (p *Parser) ParseNode() *Node {
	char,_:=p.NextChar()
	if char == '<' {
		return p.ParseElement()
	}
	return Text(p.ConsumeWhile(func(r rune) bool {
		return r != '<'
	}))
}

func (p *Parser) ParseNodes() []*Node {
	nodes := []*Node{}
	for !p.EOF() {
		p.ConsumeWhitespace()
		if p.StartsWith("</") {
			break
		}
		nodes = append(nodes, p.ParseNode())
	}
	return nodes
}

func Parse(html string) *Node { // need to be improved
	// Extract text between <body></body>
	html = extractBody(html)
	// Remove script tags
	html = RemoveScriptTags(html)
	// Parse the html
	parser := &Parser{input: html, position: 0}
	nodes := parser.ParseNodes()
	if len(nodes) == 1 {
		return nodes[0]
	}
	return Element("html", nil, nodes...)
}

func NewParser(html string) *Parser {
	return &Parser{input: html, position: 0}
}

// Extracts the body content from the html - not fully functional
func extractBody(html string) string {
	start := "<body"
	end := "</body>"
	bodyContent := ""
	bodyStart := strings.Index(html, start)
	bodyEnd := strings.Index(html, end)
	if bodyStart != -1 && bodyEnd != -1 {
		bodyContent = html[bodyStart+len(start) : bodyEnd]
		return bodyContent
	} else {
		return "Body tags not found"
	}
}

// faulthy implementation
func RemoveScriptTags(html string) string {
	for {
		start := strings.Index(html, "<script")
		if start == -1 {
			break
		}
		end := strings.Index(html[start:], "</script>")
		if end == -1 {
			break
		}
		end += start + len("</script>")
		html = html[:start] + html[end:]
	}
	return html
}