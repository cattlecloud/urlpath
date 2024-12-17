// Copyright (c) CattleCloud LLC
// SPDX-License-Identifier: BSD-3-Clause

package urlpath

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
	"github.com/shoenig/test/must"
)

func Test_Parse(t *testing.T) {
	t.Parallel()

	router := mux.NewRouter()
	executed := false

	router.HandleFunc("/v1/{foo}/{bar}", func(_ http.ResponseWriter, r *http.Request) {
		var foo string
		var bar int

		err := Parse(r, Schema{
			"foo": String(&foo),
			"bar": Int(&bar),
		})

		must.NoError(t, err)
		must.EqOp(t, "blah", foo)
		must.EqOp(t, 31, bar)
		executed = true
	})

	w := httptest.NewRecorder()
	ctx := context.Background()
	request, err := http.NewRequestWithContext(ctx, http.MethodGet, "/v1/blah/31", nil)
	must.NoError(t, err)

	router.ServeHTTP(w, request)
	must.True(t, executed)
}

func Test_ParseValues(t *testing.T) {
	t.Parallel()

	var foo string
	var bar int

	values := map[string]string{
		"foo": "blah",
		"bar": "21",
	}

	err := ParseValues(values, Schema{
		"foo": String(&foo),
		"bar": Int(&bar),
	})

	must.NoError(t, err)
	must.EqOp(t, "blah", foo)
	must.EqOp(t, 21, bar)
}

func Test_ParseValues_incompatible(t *testing.T) {
	t.Parallel()

	var foo string
	var bar int

	values := map[string]string{
		"foo": "blah",
		"bar": "not an int",
	}

	err := ParseValues(values, Schema{
		"foo": String(&foo),
		"bar": Int(&bar),
	})

	must.Error(t, err)
}

func Test_ParseValues_missing(t *testing.T) {
	t.Parallel()

	var foo string
	var bar int

	values := map[string]string{
		"foo": "blah",
	}

	err := ParseValues(values, Schema{
		"foo": String(&foo),
		"bar": Int(&bar),
	})

	must.Error(t, err)
}

func Test_Parameter_String(t *testing.T) {
	t.Parallel()

	p := Parameter("foo")
	s := p.String()
	must.EqOp(t, "{foo}", s)
}

func Test_ParseValues_StringType_Parse(t *testing.T) {
	t.Parallel()

	values := map[string]string{
		"KEY": "value",
	}

	type ident string

	var id ident

	err := ParseValues(values, Schema{
		"KEY": String(&id),
	})
	must.NoError(t, err)
	must.Eq(t, "value", id)
}

func Test_ParseValues_IntType_Parse(t *testing.T) {
	t.Parallel()

	type years int
	var age years

	values := map[string]string{
		"age": "51",
	}

	err := ParseValues(values, Schema{
		"age": Int(&age),
	})

	must.NoError(t, err)
	must.Eq(t, 51, age)
}
