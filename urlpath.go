// Copyright (c) CattleCloud LLC
// SPDX-License-Identifier: BSD-3-Clause

package urlpath

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// Typical usage:
//
//    // in the mux definition, e.g.
//    router.Handle("/v1/{foo}/{bar}, newHandler())
//
//    // in the handler implementation, e.g.
//    var foo string
//    var bar int
//    err := urlpath.Parse(request, urlpath.Schema{
//        "foo": urlpath.String(&foo),
//        "bar": urlpath.Int(&bar),
//    })

// A Parameter is a named element of a URL route,
// encoded such that a gorilla router interprets it
// as a path parameter.
type Parameter string

func (p Parameter) String() string {
	return "{" + string(p) + "}"
}

// Name returns the name of the parameter.
func (p Parameter) Name() string {
	return string(p)
}

// A Schema describes how path variables should be parsed.
// Currently only int and string types are supported.
type Schema map[Parameter]Parser

// Parse will parse the URL path vars from r given the
// element names and parsers defined in schema.
//
// This method only works with requests being processed by
// handlers of a gorilla/mux.
func Parse(r *http.Request, schema Schema) error {
	return ParseValues(mux.Vars(r), schema)
}

// ParseValues will parse the parameters in vars given the
// element names and parsers defined in schema.
//
// Most use cases will be parsing values coming from an *http.Request,
// which can be done conveniently with Parse.
func ParseValues(values map[string]string, schema Schema) error {
	for name, parser := range schema {
		value, exists := values[name.Name()]
		if !exists {
			return fmt.Errorf("url path element not present: %q", name)
		}

		if err := parser.Parse(value); err != nil {
			return fmt.Errorf("could not parse url path variable: %w", err)
		}
	}
	return nil
}

// A Parser parses raw input into a destination variable.
type Parser interface {
	Parse(string) error
}

type StringType interface {
	~string
}

type stringParser[T StringType] struct {
	destination *T
}

func (p *stringParser[T]) Parse(s string) error {
	*p.destination = T(s)
	return nil
}

// String creates a parser that will parse a path element into s.
func String[T StringType](s *T) Parser {
	return &stringParser[T]{destination: s}
}

// IntType represents any type compatible with the Go integer built-in types,
// to be used as a destination for writing the value of an url path element.
type IntType interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 |
		~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64
}

type intParser[T IntType] struct {
	destination *T
}

func (p *intParser[T]) Parse(s string) error {
	i, err := strconv.Atoi(s)
	if err != nil {
		return err
	}
	*p.destination = T(i)
	return nil
}

// Int creates a Parser that will parse a path element into i.
func Int[T IntType](i *T) Parser {
	return &intParser[T]{destination: i}
}
