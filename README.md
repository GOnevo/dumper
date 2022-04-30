# Introduction

Dumper is a library for printing pretty dump to stdout.

## Installation

```shell
go get github.com/gonevo/dumper
```

## Example

```go
package main

import (
	"github.com/gonevo/dumper"
)

func main() {
	variable := "I'm value of any type"
	dumper.DD(variable)
}
```

## Methods

##### `dumper.D()` prints pretty dump to stdout only.
##### `dumper.DD()` prints pretty dump to stdout and die.

## License

The dumper library is open-sourced software licensed under the [MIT license](https://opensource.org/licenses/MIT).
