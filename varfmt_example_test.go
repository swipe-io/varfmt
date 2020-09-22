package varfmt_test

import (
	"fmt"

	"github.com/swipe-io/varfmt"
)

func ExampleExportVarName() {
	malformattedVars := []string{
		"foo",
		"foo_bar",
		"fooBar",
		"foo_bar_buz",
		"foo_bar___buz",
		"foo_id",
		"foo_tls",
		"foo_json",
	}

	for _, s := range malformattedVars {
		fmt.Printf("%s -> %s\n", s, varfmt.ExportVarName(s))
	}
	// Output:
	// foo -> Foo
	// foo_bar -> FooBar
	// fooBar -> FooBar
	// foo_bar_buz -> FooBarBuz
	// foo_bar___buz -> FooBarBuz
	// foo_id -> FooID
	// foo_tls -> FooTLS
	// foo_json -> FooJSON
}

func ExampleUnExportVarName() {
	malformattedVars := []string{
		"foo",
		"foo_bar",
		"fooBar",
		"foo_bar_buz",
		"foo_bar___buz",
		"foo_id",
		"foo_tls",
		"foo_json",
		"id_test",
		"https_test",
		"test_https",
		"test124",
	}

	for _, s := range malformattedVars {
		fmt.Printf("%s -> %s\n", s, varfmt.UnExportVarName(s))
	}
	// Output:
	// foo -> foo
	// foo_bar -> fooBar
	// fooBar -> fooBar
	// foo_bar_buz -> fooBarBuz
	// foo_bar___buz -> fooBarBuz
	// foo_id -> fooID
	// foo_tls -> fooTLS
	// foo_json -> fooJSON
	// id_test -> idTest
	// https_test -> httpsTest
	// test_https -> testHTTPS
	// test124 -> test124
}
