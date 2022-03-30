package api

import (
	"heyapple/internal/core"
	"net/http"
	"net/http/httptest"
	"net/url"
	"reflect"
	"strings"
	"testing"
)

func Test_getRecipeFilter(t *testing.T) {
	for idx, data := range []struct {
		query url.Values
		out   core.Filter
	}{
		{ //00// no data
			out: core.Filter{},
		},
		{ //01// filter by name
			query: url.Values{
				"name": {"My Pie"},
			},
			out: core.Filter{
				"name": "My Pie",
			},
		},
		{ //02// ignore invalid values
			query: url.Values{
				"size":     {"My Pie"},
				"cooktime": {"30", "longer"},
				"preptime": {"12.5"},
				"kcal":     {"lots"},
				"carb":     {"12", "more"},
			},
			out: core.Filter{},
		},
		{ //03// valid values
			query: url.Values{
				"size":     {"3"},
				"preptime": {"15", "30"},
				"cooktime": {"120"},
				"misctime": {"45"},
				"kcal":     {"100", "350"},
				"prot":     {"25.5"},
			},
			out: core.Filter{
				"size":     3,
				"preptime": core.IntRange{15, 30},
				"cooktime": 120,
				"misctime": 45,
				"kcal":     core.FloatRange{100, 350},
				"prot":     float32(25.5),
			},
		},
	} {
		query := "/?" + data.query.Encode()
		req := httptest.NewRequest(http.MethodGet, query, strings.NewReader(""))

		if out := getRecipeFilter(req); !reflect.DeepEqual(out, data.out) {
			t.Errorf("test case %d: filter mismatch \nhave: %v\nwant: %v", idx, out, data.out)
		}
	}
}
