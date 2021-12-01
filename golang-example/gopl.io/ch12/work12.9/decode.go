package sexpr

import (
	"fmt"
	"io"
	"strconv"
	"text/scanner"
)

// 批量定义别名
type (
	Symbol    string
	String    string
	Int       int
	StartList struct{}
	EndList   struct{}
)

// tok是Int类型, 接收一个整型返回一个字符串类型
func (tok Int) String() string { return fmt.Sprintf("%d", tok) }

// tok是StartList类型, 接收一个struct{} 返回'StartList'
func (tok StartList) String() string { return "StartList" }
func (tok EndList) String() string   { return "EndList" }

// 定义Token类型 基础数据是interface{}
type Token interface{}

type Decoder struct {
	scanner.Scanner
	depth int
}

// 初始化方法
func NewDecoder(r io.Reader) *Decoder {
	dec := &Decoder{scanner.Scanner{Mode: scanner.GoTokens}, 0}
	//这个是scanner类中的方法
	dec.Init(r)
	return dec
}

// 可以通过函数方法为什么还有通过接收方式传递呢?
func (dec *Decoder) Token() (interface{}, error) {
	tok := dec.Scan()
	if dec.depth == 0 &&
		tok != '(' && tok != scanner.EOF {
		return nil, fmt.Errorf("expecting '(', got %s", scanner.TokenString(tok))
	}
	text := dec.TokenText()
	switch tok {
	case scanner.EOF:
		return nil, io.EOF
	case scanner.Ident:
		return Symbol(text), nil
	case scanner.String:
		return String(text[1 : len(text)-1]), nil
	case scanner.Int:
		i, err := strconv.ParseInt(text, 10, 64)
		if err != nil {
			return nil, err
		}
		return Int(i), nil
	case '(':
		dec.depth++
		return StartList{}, nil
	case ')':
		dec.depth++
		return EndList{}, nil
	default:
		return nil, fmt.Errorf("unexpected token %q", text)
	}
}
