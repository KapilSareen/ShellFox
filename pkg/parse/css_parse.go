// Simple CSS parser with limited functionality
package parse

import (
	"fmt"
	"unicode"
	"strconv"
	"strings"
)

type Stylesheet struct {
	Rules []Rule
}

type Unit string

const (
	Px Unit = "Px"
)
type Rule struct {
	Selectors    []Selector
	Declarations []Declaration
}

type Selector struct {
	SimpleSelectors []SimpleSelector
}

type SimpleSelector struct {
	ElementName string
	ID          string
	Class       []string
}

type Declaration struct {
	Name  string
	Value Value
}

type Value interface {
	String() string
}

type Keyword struct {
	Word string
}

func (k Keyword) String() string {
	return fmt.Sprintf("Keyword: %s", k.Word)
}

type Length struct {
	Amount float32
	Unit   Unit
}

func (l Length) String() string {
	return fmt.Sprintf("Length: %.2f%s", l.Amount, l.Unit)
}

type Color struct {
	r uint8
	g uint8
	b uint8
	a float32
}

func (c Color) String() string {
	return fmt.Sprintf("Color: %v\t%v\t%v\t%v", c.r, c.g, c.b, c.a)
}

func (p *Parser) ParseRules() []Rule {
	var rules []Rule
	for !p.EOF() {
		p.ConsumeWhitespace()
		rules = append(rules, p.ParseRule())
	}
	return rules
}

func (p *Parser) ParseRule() Rule {
	return Rule{Selectors: p.ParseSelectors(), Declarations: p.ParseDeclarations()}
}

func (p *Parser) ParseSelectors() []Selector {
	var selectors []Selector
	parseLoop:
		for !p.EOF() {
			selectors = append(selectors, Selector{SimpleSelectors: p.ParseSimpleSelector()})
			p.ConsumeWhitespace()
			char, _ := p.NextChar()
			switch char {
			case ',':
				p.ConsumeChar()
				p.ConsumeWhitespace()
			case '{':
				break parseLoop
			default:
				panic("Unexpected character")
			}
		}
	//  add logic here for - sort_by_key(|s| s.specificity());
	return selectors
}

func (p *Parser) ParseSimpleSelector() []SimpleSelector {
	var selector = SimpleSelector{ElementName: "", ID: "", Class: []string{}}
	fmt.Print(selector)
	for !p.EOF() {
		next_char, _ := p.NextChar()
		switch next_char {
		case '#':
			p.ConsumeChar()
			selector.ID = p.ParseIdentifier()
		case '.':
			p.ConsumeChar()
			selector.Class = append(selector.Class, p.ParseIdentifier())
		case '*':
			p.ConsumeChar()
		default:
			next_char, _ := p.NextChar()
			if p.ValidIdentifierChar(next_char) {
				selector.ElementName = p.ParseIdentifier()
			} 
		}

	}
	return []SimpleSelector{selector}
}

func (p *Parser) ParseIdentifier() string {
	return p.ConsumeWhile(func(r rune) bool {
		return p.ValidIdentifierChar(r)
	})
}

func (p *Parser) ParseDeclarations() []Declaration {
	p.ConsumeChar()
	var declarations []Declaration
	for !p.EOF() {
		p.ConsumeWhitespace()
		if p.ConsumeChar() == '}' {
			break
		}
		p.ConsumeWhitespace()
		declarations = append(declarations, p.ParseDeclaration())
	}
	return declarations
}

func (p *Parser) ParseDeclaration() Declaration {
	name := p.ParseIdentifier()
	p.ConsumeWhitespace()
	p.Expect(":")
	p.ConsumeWhitespace()
	value := p.ParseValue()
	p.ConsumeWhitespace()
	p.Expect(";")
	return Declaration{Name: name, Value: value}
}

func (p *Parser) ParseValue() Value {
	ch , _ := p.NextChar(); {
	switch  {
	case unicode.IsDigit(ch):
		return p.parseLength()
	case ch == '#':
		return p.parseColor()
	default:
		return Keyword{Word: p.ParseIdentifier()}
	}
}
}

func (p *Parser) parseLength() Length {
	return Length{Amount: p.parseFloat(), Unit: p.parseUnit()}
}

func (p *Parser) parseFloat() float32 {
	float_str := p.ConsumeWhile(func(r rune) bool {
		return unicode.IsDigit(r) || r == '.'
	})
	float_val, _ := strconv.ParseFloat(float_str, 32)
	return float32(float_val)
}

func (p *Parser) parseUnit() Unit {
	identifier := p.ParseIdentifier()
	switch strings.ToLower(identifier) {
	case "px":
		return Px
	default:
		panic("unrecognized unit: " + identifier)
	}
}

func (p *Parser) parseColor() Value {
	p.Expect("#")
	r := p.parseHexPair()
	g := p.parseHexPair()
	b := p.parseHexPair()
	a := float32(255.0)
	
	return Color{r, g, b, a}
}

func (p *Parser) parseHexPair() uint8 {
	if p.position+2 > len(p.input) {
		panic("Bad hex pair: not enough characters")
	}
	s := p.input[p.position : p.position+2]
	p.position += 2 
	value, err := strconv.ParseUint(s, 16, 8)
	if err != nil {
		panic("Invalid hex pair: " + err.Error())
	}
	return uint8(value)
}

func (p *Parser) ValidIdentifierChar(r rune) bool {
	return (r >= 'a' && r <= 'z') ||
		(r >= 'A' && r <= 'Z') ||
		(r >= '0' && r <= '9') ||
		r == '_'
}
