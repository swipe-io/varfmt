package varfmt_test

import (
	"testing"

	"github.com/swipe-io/varfmt"
)

func Test_ExportVarName(t *testing.T) {
	testData := map[string]string{
		"foo":           "Foo",
		"foo_bar":       "FooBar",
		"fooBar":        "FooBar",
		"foo_bar_buz":   "FooBarBuz",
		"foo_bar___buz": "FooBarBuz",
		"foo_id":        "FooID",
	}

	for k, v := range testData {
		f := varfmt.ExportVarName(k)
		if f != v {
			t.Errorf("expect %s got %s", v, f)
		}
	}
}

func Test_UnExportVarName(t *testing.T) {
	testData := map[string]string{
		"foo":           "foo",
		"foo_bar":       "fooBar",
		"fooBar":        "fooBar",
		"foo_bar_buz":   "fooBarBuz",
		"foo_bar___buz": "fooBarBuz",
		"foo_id":        "fooID",
		"id_foo":        "idFoo",
		"idHTTPS":       "idHTTPS",
	}
	for k, v := range testData {
		f := varfmt.UnExportVarName(k)
		if f != v {
			t.Errorf("expect %s got %s", v, f)
		}
	}
}
