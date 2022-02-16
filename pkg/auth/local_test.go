package auth_test

import (
	"heyapple/internal/mock"
	"heyapple/pkg/app"
	"heyapple/pkg/auth"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"

	"github.com/and-rad/scs/v2"
)

func TestLocalLogin(t *testing.T) {
	for idx, data := range []struct {
		db *mock.DB
		in url.Values

		status int
	}{
		{ //00// no data
			db:     mock.NewDB(),
			status: http.StatusUnauthorized,
		},
		{ //01// connection failure
			db:     mock.NewDB().WithError(mock.ErrDOS),
			in:     url.Values{"email": {"a@a.a"}, "pass": {"password123"}},
			status: http.StatusInternalServerError,
		},
		{ //02// empty DB
			db:     mock.NewDB(),
			in:     url.Values{"email": {"a@a.a"}, "pass": {"password123"}},
			status: http.StatusUnauthorized,
		},
		{ //03// user doesn't exist
			db:     mock.NewDB().WithUser(app.User{ID: 1, Email: "b@b.b", Pass: "ahdjhehlk"}),
			in:     url.Values{"email": {"a@a.a"}, "pass": {"password123"}},
			status: http.StatusUnauthorized,
		},
		{ //04// wrong password
			db:     mock.NewDB().WithUser(app.User{ID: 1, Email: "a@a.a", Pass: "ahdjhehlk"}),
			in:     url.Values{"email": {"a@a.a"}, "pass": {"password123"}},
			status: http.StatusUnauthorized,
		},
		{ //05// success
			db:     mock.NewDB().WithUser(app.User{ID: 1, Email: "a@a.a", Pass: "$2a$10$ADm2JBRbt8UvB0uI7NNFBupOdTq7XKae6Dvc7NfVCnw89rPZr3.zK"}),
			in:     url.Values{"email": {"a@a.a"}, "pass": {"password123"}},
			status: http.StatusOK,
		},
	} {
		req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(data.in.Encode()))
		req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
		res := httptest.NewRecorder()

		auth.LocalLogin(scs.New(), data.db)(res, req, nil)

		if status := res.Result().StatusCode; status != data.status {
			t.Errorf("test case %d: status mismatch \nhave: %v\nwant: %v", idx, status, data.status)
		}
	}
}
