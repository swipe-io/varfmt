# varfmt

[![Build Status](https://travis-ci.org/swipe-io/varfmt.svg?branch=master)](https://travis-ci.org/swipe-io/varfmt)
[![GitHub license](https://img.shields.io/badge/license-MIT-blue.svg)](https://raw.githubusercontent.com/swipe-io/varfmt/master/LICENSE)
[![Doc](https://img.shields.io/badge/godoc-reference-blue.svg)](https://godoc.org/github.com/swipe-io/varfmt)
[![Go Report Card](https://goreportcard.com/badge/github.com/swipe-io/varfmt)](https://goreportcard.com/report/github.com/swipe-io/varfmt)

Convert string into Go variable name following Go naming convention

## Why created

Although there are a lot of needs to generate go struct from various sources (DDL, yml, toml, JSON, etc), I couldn't find any common library to format struct field names that follow go naming convention. Plus, I found really nice little piece of code that formats variable name in [ChimeraCoder/gojson](https://github.com/ChimeraCoder/gojson).


## How to use

### ExportVarName

```go
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
```

### UnExportVarName

```go
package varfmt_test

import (
	"fmt"

	"github.com/swipe-io/varfmt"
)

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
```

## Credits

Code is being almost entirely borrowed from https://github.com/ChimeraCoder/gojson. 
Huge thanks to [ChimeraCoder](https://github.com/ChimeraCoder) !!
