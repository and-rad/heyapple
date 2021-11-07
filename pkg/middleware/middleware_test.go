package middleware_test

import (
	"heyapple/internal/mock"
	"heyapple/pkg/middleware"
	"net/http"
	"net/http/httptest"
	"reflect"
	"strings"
	"testing"
)

func TestHeaders(t *testing.T) {
	for idx, data := range []struct {
		header http.Header
		next   mock.Handler
		status int
	}{
		{ //00//
			next: mock.Handler{Body: "okay"},
			header: http.Header{
				"Access-Control-Allow-Origin": {"*"},
				"Content-Type":                {"application/json; charset=utf-8"},
			},
			status: http.StatusOK,
		},
	} {
		req := httptest.NewRequest(http.MethodGet, "/", strings.NewReader(""))
		res := httptest.NewRecorder()
		middleware.Headers(data.next).ServeHTTP(res, req)

		if header := res.Result().Header; !reflect.DeepEqual(header, data.header) {
			t.Errorf("test case %d: header mismatch \nhave: %v \nwant: %v", idx, header, data.header)
		}

		if status := res.Result().StatusCode; status != data.status {
			t.Errorf("test case %d: status mismatch \nhave: %v \nwant: %v", idx, status, data.status)
		}

		if body := res.Body.String(); body != data.next.Body {
			t.Errorf("test case %d: body mismatch \nhave: %v \nwant: %v", idx, body, data.next.Body)
		}
	}
}

func TestOptions(t *testing.T) {
	for idx, data := range []struct {
		inhead  http.Header
		outhead http.Header
		status  int
	}{
		{ //00// no headers sent
			outhead: http.Header{},
			status:  http.StatusNoContent,
		},
		{ //01// no relevant headers sent
			inhead:  http.Header{"X-whazzup": {"whazzzzuuuuup"}},
			outhead: http.Header{},
			status:  http.StatusNoContent,
		},
		{ //02// CORS preflight header
			inhead: http.Header{
				"Access-Control-Request-Method": {"POST"},
			},
			outhead: http.Header{
				"Access-Control-Allow-Headers": {"Content-Type"},
				"Access-Control-Allow-Methods": {"DELETE, GET, OPTIONS, POST"},
				"Access-Control-Allow-Origin":  {"*"},
				"Access-Control-Max-Age":       {"86400"},
			},
			status: http.StatusNoContent,
		},
	} {
		req := httptest.NewRequest(http.MethodGet, "/", strings.NewReader(""))
		req.Header = data.inhead
		res := httptest.NewRecorder()
		middleware.Options(res, req)

		if header := res.Result().Header; !reflect.DeepEqual(header, data.outhead) {
			t.Errorf("test case %d: header mismatch \nhave: %v \nwant: %v", idx, header, data.outhead)
		}

		if status := res.Result().StatusCode; status != data.status {
			t.Errorf("test case %d: status mismatch \nhave: %v \nwant: %v", idx, status, data.status)
		}
	}
}
