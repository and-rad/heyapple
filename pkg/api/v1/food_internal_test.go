package api

import (
	"heyapple/pkg/core"
	"net/http"
	"net/http/httptest"
	"net/url"
	"reflect"
	"strings"
	"testing"
)

func Test_getFoodFilter(t *testing.T) {
	for idx, data := range []struct {
		query url.Values
		out   core.Filter
	}{
		{ //00// no data
			out: core.Filter{},
		},
		{ //01// ignore invalid data
			query: url.Values{
				"kcal":  {"a lot", "even more"}, // invalid types
				"fat":   {"20", "fifty"},        // invalid second type
				"power": {"9000"},               // property doesn't xist
			},
			out: core.Filter{},
		},
		{ //02// ignore default value
			query: url.Values{"kcal": {"0", "900"}},
			out:   core.Filter{},
		},
		{ //03// valid values
			query: url.Values{
				"kcal": {"50", "150"},
				"carb": {"20", "20"},
				"prot": {"15.5", "35"},
				"fib":  {"33"},
			},
			out: core.Filter{
				"kcal": core.FloatRange{50, 150},
				"carb": float32(20),
				"prot": core.FloatRange{15.5, 35},
				"fib":  float32(33),
			},
		},
	} {
		query := "/?" + data.query.Encode()
		req := httptest.NewRequest(http.MethodGet, query, strings.NewReader(""))

		if out := getFoodFilter(req); !reflect.DeepEqual(out, data.out) {
			t.Errorf("test case %d: filter mismatch \nhave: %v\nwant: %v", idx, out, data.out)
		}
	}
}
