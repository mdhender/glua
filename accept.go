// GLUA - a Lua-like interpreter
//
// MIT License
//
// Copyright (c) 2021 Michael D Henderson
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in all
// copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
// SOFTWARE.

package main

import (
	"bytes"
	"fmt"
)

func (p parser) accept() (parser, *node) {
	if eof(p) {
		return p, nil
	}
	return p, nil
}

type CHUNK struct {
	block *BLOCK
}

// chunk ::= block
func (p parser) accept_chunk() (parser, *CHUNK) {
	if eof(p) {
		return p, nil
	}
	pSaved, chunk := p, &CHUNK{}

	p, chunk.block = p.accept_block()
	if chunk.block == nil {
		return pSaved, nil
	}
	return p, chunk
}

type BLOCK struct {
	stat    []*STAT
	retstat *RETSTAT
}

// block ::= {stat} [retstat]
func (p parser) accept_block() (parser, *BLOCK) {
	block := &BLOCK{}

	var stat *STAT
	for p, stat = p.accept_stat(); stat != nil; p, stat = p.accept_stat() {
		block.stat = append(block.stat, stat)
	}

	p, block.retstat = p.accept_retstat()

	return p, block
}

type STAT struct{}

// stat ::= ';'
//      | varlist '=' explist
//      | functioncall
//      | label
//      | 'break'
//      | 'goto' Name
//      | 'do' block 'end'
//      | 'while' exp 'do' block 'end'
//      | 'repeat' block 'until' exp
//      | 'if' exp 'then' block {'elseif' exp 'then' block} ['else' block] 'end'
//      | 'for' Name '=' exp ',' exp [',' exp] 'do' block 'end'
//      | 'for' namelist 'in' explist 'do' block 'end'
//      | 'function' funcname funcbody
//      | 'local' 'function' Name funcbody
//      | 'local' attnamelist ['=' explist]
func (p parser) accept_stat() (parser, *STAT) {
	if eof(p) {
		return p, nil
	}
	panic("!implemented")
}

type RETSTAT struct {
	explist *EXPLIST
}

// retstat ::= 'return' [explist] [';']
func (p parser) accept_retstat() (parser, *RETSTAT) {
	if eof(p) {
		return p, nil
	}
	pSaved, retstat := p, &RETSTAT{}

	var kwReturn *KEYWORD
	if p, kwReturn = p.accept_Keyword('r', 'e', 't', 'u', 'r', 'n'); kwReturn == nil {
		return pSaved, nil
	}

	// [explist]
	p, retstat.explist = p.accept_explist()

	// [';']
	p, _ = p.accept_Literal(',')

	return p, retstat
}

type EXPLIST struct {
	exp []*EXP
}

// explist ::= exp {',' exp}
func (p parser) accept_explist() (parser, *EXPLIST) {
	if eof(p) {
		return p, nil
	}
	pSaved, explist := p, &EXPLIST{}

	var exp *EXP
	if p, exp = p.accept_exp(); exp == nil {
		return pSaved, nil
	}
	explist.exp = append(explist.exp, exp)

	// {',' exp}
	for pp, comma := p.accept_Literal(','); comma != nil; pp, comma = pp.accept_Literal(',') {
		if pp, exp = pp.accept_exp(); exp == nil {
			break
		}
		p, explist.exp = pp, append(explist.exp, exp)
	}

	return p, explist
}

type EXP struct {
	literal       []byte
	literalString []byte
	numeral       []byte
	functiondef   *FUNCTIONDEF
}

// exp ::= 'nil'
//     | 'false'
//     | 'true'
//     | Numeral
//     | LiteralString
//     | '...'
//     | functiondef
//     | prefixexp
//     | tableconstructor
//     | exp binop exp
//     | unop exp
func (p parser) accept_exp() (parser, *EXP) {
	if eof(p) {
		return p, nil
	}
	pSaved, exp := p, &EXP{}

	if p, exp.literal = p.accept_Literal('n', 'i', 'l'); exp.literal != nil {
		return p, exp
	}
	if p, exp.literal = p.accept_Literal('f', 'a', 'l', 's', 'e'); exp.literal != nil {
		return p, exp
	}
	if p, exp.literal = p.accept_Literal('t', 'r', 'u', 'e'); exp.literal != nil {
		return p, exp
	}
	if p, exp.numeral = p.accept_Numeral(); exp.numeral != nil {
		return p, exp
	}
	if p, exp.literalString = p.accept_LiteralString(); exp.literalString != nil {
		return p, exp
	}
	if p, exp.literal = p.accept_Literal('.', '.', '.'); exp.literal != nil {
		return p, exp
	}

	fmt.Println("todo:", pSaved)
	panic("!implemented")
}

type FUNCTIONDEF struct {
	funcbody *FUNCBODY
}

// functiondef ::= 'function' funcbody
func (p parser) accept_functiondef() (parser, *FUNCTIONDEF) {
	if eof(p) {
		return p, nil
	}
	pSaved, functiondef := p, &FUNCTIONDEF{}
	var function []byte
	if p, function = p.accept_Literal('f', 'u', 'n', 'c', 't', 'i', 'o', 'n'); function == nil {
		return pSaved, nil
	}
	if p, functiondef.funcbody = p.accept_funcbody(); functiondef.funcbody == nil {
		return pSaved, nil
	}
	return p, functiondef
}

type FUNCBODY struct {
	parlist *PARLIST
	block   *BLOCK
}
type PARLIST struct {
	namelist *NAMELIST
}
type NAMELIST struct {
	name [][]byte
}

//funcbody ::= '(' [parlist] ')' block 'end'
func (p parser) accept_funcbody() (parser, *FUNCBODY) {
	if eof(p) {
		return p, nil
	}
	panic("!implemented")
}

// binop ::=  '+' | '-' | '*' | '/' | '//' | '^' | '%' |
//            '&' | '~' | '|' | '>>' | '<<' | '..' |
//            '<' | '<=' | '>' | '>=' | '==' | '~=' |
//            and | or
func (p parser) accept_binop() (parser, *node) {
	if eof(p) {
		return p, nil
	}
	switch p.buf[0] {
	case '+':
		p.buf = p.buf[1:]
		return p, &node{}
	case '-':
		p.buf = p.buf[1:]
		return p, &node{}
	case '*':
		p.buf = p.buf[1:]
		return p, &node{}
	case '^':
		p.buf = p.buf[1:]
		return p, &node{}
	case '%':
		p.buf = p.buf[1:]
		return p, &node{}
	case '&':
		p.buf = p.buf[1:]
		return p, &node{}
	case '|':
		p.buf = p.buf[1:]
		return p, &node{}
	case '.': // '..'
		if match := []byte{'.', '.'}; bytes.HasPrefix(p.buf, match) {
			p.buf = p.buf[len(match):]
			return p, &node{}
		}
		p.buf = p.buf[1:]
		return p, &node{}
	case '=': // '=='
		if match := []byte{'=', '='}; bytes.HasPrefix(p.buf, match) {
			p.buf = p.buf[len(match):]
			return p, &node{}
		}
		p.buf = p.buf[1:]
		return p, &node{}
	case '/': // '/' or '//'
		if match := []byte{'/', '/'}; bytes.HasPrefix(p.buf, match) {
			p.buf = p.buf[len(match):]
			return p, &node{}
		}
		p.buf = p.buf[1:]
		return p, &node{}
	case '<': // '<' | '<<' | '<='
		if match := []byte{'<', '<'}; bytes.HasPrefix(p.buf, match) {
			p.buf = p.buf[len(match):]
			return p, &node{}
		}
		if match := []byte{'<', '='}; bytes.HasPrefix(p.buf, match) {
			p.buf = p.buf[len(match):]
			return p, &node{}
		}
		p.buf = p.buf[1:]
		return p, &node{}
	case '>': // '>' | '>>' | '>='
		if match := []byte{'>', '>'}; bytes.HasPrefix(p.buf, match) {
			p.buf = p.buf[len(match):]
			return p, &node{}
		}
		if match := []byte{'>', '='}; bytes.HasPrefix(p.buf, match) {
			p.buf = p.buf[len(match):]
			return p, &node{}
		}
		p.buf = p.buf[1:]
		return p, &node{}
	case '~': // '~' | '~='
		if match := []byte{'~', '='}; bytes.HasPrefix(p.buf, match) {
			p.buf = p.buf[len(match):]
			return p, &node{}
		}
		p.buf = p.buf[1:]
		return p, &node{}
	case 'a': // 'and'
		if match := []byte{'a', 'n', 'd'}; bytes.HasPrefix(p.buf, match) {
			p.buf = p.buf[len(match):]
			return p, &node{}
		}
	case 'o': // 'or'
		if match := []byte{'o', 'r'}; bytes.HasPrefix(p.buf, match) {
			p.buf = p.buf[len(match):]
			return p, &node{}
		}
	}
	return p, nil
}

// fieldsep ::= ',' | ';'
func (p parser) accept_fieldsep() (parser, *node) {
	if eof(p) {
		return p, nil
	}
	switch p.buf[0] {
	case ',':
		p.buf = p.buf[1:]
		return p, &node{}
	case ';':
		p.buf = p.buf[1:]
		return p, &node{}
	}
	return p, nil
}

type PREFIXEXP struct {
	variable     *VARIABLE
	functioncall *FUNCTIONCALL
	exp          *EXP
}

// prefixexp ::= var | functioncall | '(' exp ')'
func (p parser) accept_prefixexp() (parser, *PREFIXEXP) {
	if eof(p) {
		return p, nil
	}
	pSaved, prefixexp := p, &PREFIXEXP{}
	if p, prefixexp.variable = p.accept_var(); prefixexp.variable != nil {
		return p, prefixexp
	}
	if p, prefixexp.functioncall = p.accept_functioncall(); prefixexp.functioncall != nil {
		return p, prefixexp
	}
	var oparen, cparen []byte
	if p, oparen = p.accept_Literal('('); oparen != nil {
		if p, prefixexp.exp = p.accept_exp(); prefixexp.exp != nil {
			if p, cparen = p.accept_Literal(')'); cparen != nil {
				return p, prefixexp
			}
		}
	}
	return pSaved, nil
}

type VARIABLE struct{}

func (p parser) accept_var() (parser, *VARIABLE) {
	panic("!implemented")
}

type FUNCTIONCALL struct{}

func (p parser) accept_functioncall() (parser, *FUNCTIONCALL) {
	panic("!implemented")
}

// unop ::= '-' | not | '#' | '~'
func (p parser) accept_unop() (parser, *node) {
	if eof(p) {
		return p, nil
	}
	switch p.buf[0] {
	case '-':
		p.buf = p.buf[1:]
		return p, &node{}
	case '#':
		p.buf = p.buf[1:]
		return p, &node{}
	case 'n': // 'not'
		if match := []byte{'n', 'o', 't'}; bytes.HasPrefix(p.buf, match) {
			p.buf = p.buf[len(match):]
			return p, &node{}
		}
	case '~':
		p.buf = p.buf[1:]
		return p, &node{}
	}
	return p, nil
}
