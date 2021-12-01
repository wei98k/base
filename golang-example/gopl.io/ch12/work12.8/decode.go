package sexpr

import (
	"bytes"
	"fmt"
	"io"
	"reflect"
	"strconv"
	"text/scanner"
)

// 练习 12.8： sexpr.Unmarshal函数和json.Unmarshal一样，
// 都要求在解码前输入完整的字节slice。
// 定义一个和json.Decoder类似的sexpr.Decoder类型，
// 支持从一个io.Reader流解码。
// 修改sexpr.Unmarshal函数，使用这个新的类型实现。

type Decoder struct {
	*lexer
}

func NewDecoder(r io.Reader) *Decoder {
	dec := &Decoder{&lexer{scan: scanner.Scanner{Mode: scanner.GoTokens}}}
	dec.scan.Init(r)
	return dec
}

func (dec *Decoder) Decode(out interface{}) (err error) {
	defer func() {
		if x := recover(); x != nil {
			err = fmt.Errorf("error at %s: %v", dec.scan.Position, x)
		}
	}()
	if dec.token == 0 {
		dec.next()
	}
	read(dec.lexer, reflect.ValueOf(out).Elem())
	return nil
}

func Unmarshal(data []byte, out interface{}) error {
	return NewDecoder(bytes.NewReader(data)).Decode(out)
}

type lexer struct {
	scan  scanner.Scanner
	token rune
}

func (lex *lexer) next()        { lex.token = lex.scan.Scan() }
func (lex *lexer) text() string { return lex.scan.TokenText() }

func (lex *lexer) consume(want rune) {
	if lex.token != want {
		panic(fmt.Sprintf("got %q, want %q", lex.text(), want))
	}
	lex.next()
}

func read(lex *lexer, v reflect.Value) {
	switch lex.token {
	case scanner.Ident:
		if lex.text() == "nil" {
			v.Set(reflect.Zero(v.Type()))
			lex.next()
			return
		}
	case scanner.String:
		s, _ := strconv.Unquote(lex.text())
		v.SetString(s)
		lex.next()
		return
	case scanner.Int:
		i, _ := strconv.Atoi(lex.text())
		v.SetInt(int64(i))
		lex.next()
		return
	case '(':
		lex.next()
		readList(lex, v)
		lex.next()
		return
	}
	panic(fmt.Sprintf("unexpected token %q", lex.text()))
}

func readList(lex *lexer, v reflect.Value) {
	switch v.Kind() {
	case reflect.Array:
		for i := 0; !endList(lex); i++ {
			read(lex, v.Index(i))
		}
	case reflect.Slice:
		for !endList(lex) {
			item := reflect.New(v.Type().Elem()).Elem()
			read(lex, item)
			v.Set(reflect.Append(v, item))
		}
	case reflect.Struct:
		for !endList(lex) {
			lex.consume('(')
			if lex.token != scanner.Ident {
				panic(fmt.Sprintf("got token %q, want field name", lex.text()))
			}
			name := lex.text()
			lex.next()
			read(lex, v.FieldByName(name))
			lex.consume(')')
		}
	case reflect.Map:
		v.Set(reflect.MakeMap(v.Type()))
		for !endList(lex) {
			lex.consume('(')
			key := reflect.New(v.Type().Key()).Elem()
			read(lex, key)
			value := reflect.New(v.Type().Elem()).Elem()
			read(lex, value)
			v.SetMapIndex(key, value)
			lex.consume(')')
		}
	default:
		panic(fmt.Sprintf("cannot decode list into %v", v.Type()))
	}
}

func endList(lex *lexer) bool {
	switch lex.token {
	case scanner.EOF:
		panic("end of file")
	case ')':
		return true
	}
	return false
}
