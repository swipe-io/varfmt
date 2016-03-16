# varfmt
Convert string into Go variable name

[![Build Status](https://travis-ci.org/achiku/varfmt.svg?branch=master)](https://travis-ci.org/achiku/varfmt)
[![GitHub license](https://img.shields.io/badge/license-MIT-blue.svg)](https://raw.githubusercontent.com/achiku/varfmt/master/LICENSE)

## Why created

Although there are a lot of needs to generate go struct from various sources (DDL, yml, toml, JSON, etc), I couldn't find  common library to format struct field name that follows go naming convention. Plus, I found really nice little piece of code that formats variable name in [ChimeraCoder/gojson](https://github.com/ChimeraCoder/gojson).


## How to use

```go
package varfmt_test

import (
	"fmt"

	"github.com/achiku/varfmt"
)

func ExamplePublicVarName() {
	malformattedVars := []string{
		"foo",
		"foo_bar",
		"fooBar",
		"foo_bar_buz",
		"foo_bar___buz",
		"foo_id",
	}

	for _, s := range malformattedVars {
		fmt.Printf("%s -> %s\n", s, varfmt.PublicVarName(s))
	}
	// Output:
	// foo -> Foo
	// foo_bar -> FooBar
	// fooBar -> FooBar
	// foo_bar_buz -> FooBarBuz
	// foo_bar___buz -> FooBarBuz
	// foo_id -> FooID

}
```

## Credits

Code is almost entirely borrowed from https://github.com/ChimeraCoder/gojson . Huge thanks for [ChimeraCoder](https://github.com/ChimeraCoder) !!
