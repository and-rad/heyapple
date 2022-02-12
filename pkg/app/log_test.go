package app_test

import (
	"bytes"
	"errors"
	"fmt"
	"heyapple/pkg/app"
	"testing"
)

func TestLog_Log(t *testing.T) {
	for idx, data := range []struct {
		in  interface{}
		out string
	}{
		{ //00// nil data
			out: fmt.Sprintf("Log  : <nil>\n"),
		},
		{ //01// arbitrary objects
			in: struct {
				Name string
				Age  int
			}{Name: "Joe", Age: 23},
			out: fmt.Sprintf("Log  : {Joe 23}\n"),
		},
		{ //02// strings
			in:  "Problem with the moonwalk",
			out: fmt.Sprintf("Log  : Problem with the moonwalk\n"),
		},
		{ //03// errors
			in:  errors.New("Danger Will Robinson"),
			out: fmt.Sprintf("Log  : Danger Will Robinson\n"),
		},
	} {
		buf := &bytes.Buffer{}
		app.NewLog(buf).Log(data.in)

		if out := buf.String(); out != data.out {
			t.Errorf("test case %d: output mismatch \nhave: %v\nwant: %v", idx, out, data.out)
		}
	}
}

func TestLog_Warn(t *testing.T) {
	for idx, data := range []struct {
		in  interface{}
		out string
	}{
		{ //00// nil data
			out: fmt.Sprintf("Warn : <nil>\n"),
		},
		{ //01// arbitrary objects
			in: struct {
				Name string
				Age  int
			}{Name: "Joe", Age: 23},
			out: fmt.Sprintf("Warn : {Joe 23}\n"),
		},
		{ //02// strings
			in:  "Problem with the moonwalk",
			out: fmt.Sprintf("Warn : Problem with the moonwalk\n"),
		},
		{ //03// errors
			in:  errors.New("Danger Will Robinson"),
			out: fmt.Sprintf("Warn : Danger Will Robinson\n"),
		},
	} {
		buf := &bytes.Buffer{}
		app.NewLog(buf).Warn(data.in)

		if out := buf.String(); out != data.out {
			t.Errorf("test case %d: output mismatch \nhave: %v\nwant: %v", idx, out, data.out)
		}
	}
}

func TestLog_Error(t *testing.T) {
	for idx, data := range []struct {
		in  interface{}
		out string
	}{
		{ //00// nil data
			out: fmt.Sprintf("Error: <nil>\n"),
		},
		{ //01// arbitrary objects
			in: struct {
				Name string
				Age  int
			}{Name: "Joe", Age: 23},
			out: fmt.Sprintf("Error: {Joe 23}\n"),
		},
		{ //02// strings
			in:  "Problem with the moonwalk",
			out: fmt.Sprintf("Error: Problem with the moonwalk\n"),
		},
		{ //03// errors
			in:  errors.New("Danger Will Robinson"),
			out: fmt.Sprintf("Error: Danger Will Robinson\n"),
		},
	} {
		buf := &bytes.Buffer{}
		app.NewLog(buf).Error(data.in)

		if out := buf.String(); out != data.out {
			t.Errorf("test case %d: output mismatch \nhave: %v\nwant: %v", idx, out, data.out)
		}
	}
}
