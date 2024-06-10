# pointerhelpers

`pointerhelpers` is a Go package designed to reduce boilerplate code when working with pointers. This package provides helper functions for various data types, including `bool`, `int` (and its variations), `float` (and its variations), `complex` (and its variations), and `string`. By using this package, you can streamline your code and enhance readability.

## Installation

To install the `pointerhelpers` package, use the following command:

```bash
go get github.com/xorima/pointerhelpers
```

## Usage

The `pointerhelpers` package provides functions to obtain pointers to values and to safely retrieve values from pointers. Below are the examples for each supported type:

### Boolean

```go
package main

import (
	"fmt"
	"github.com/xorima/pointerhelpers"
)

func main() {
	b := true
	bPtr := pointerhelpers.Bool(b)
	fmt.Println(*bPtr) // Output: true

	nilBoolPtr := (*bool)(nil)
	fmt.Println(pointerhelpers.BoolValue(nilBoolPtr)) // Output: false
}
```

### Integers

#### int

```go
package main

import (
	"fmt"
	"github.com/xorima/pointerhelpers"
)

func main() {
	i := 42
	iPtr := pointerhelpers.Int(i)
	fmt.Println(*iPtr) // Output: 42

	nilIntPtr := (*int)(nil)
	fmt.Println(pointerhelpers.IntValue(nilIntPtr)) // Output: 0
}
```

#### Other integer types

The package supports other integer types similarly: `int8`, `int16`, `int32`, `int64`, `uint`, `uint8`, `uint16`, `uint32`, `uint64`.

### Floating Point Numbers

#### float32

```go
package main

import (
	"fmt"
	"github.com/xorima/pointerhelpers"
)

func main() {
	f := float32(3.14)
	fPtr := pointerhelpers.Float32(f)
	fmt.Println(*fPtr) // Output: 3.14

	nilFloat32Ptr := (*float32)(nil)
	fmt.Println(pointerhelpers.Float32Value(nilFloat32Ptr)) // Output: 0
}
```

#### float64

```go
package main

import (
	"fmt"
	"github.com/xorima/pointerhelpers"
)

func main() {
	f := 3.14
	fPtr := pointerhelpers.Float64(f)
	fmt.Println(*fPtr) // Output: 3.14

	nilFloat64Ptr := (*float64)(nil)
	fmt.Println(pointerhelpers.Float64Value(nilFloat64Ptr)) // Output: 0
}
```

### Complex Numbers

#### complex64

```go
package main

import (
	"fmt"
	"github.com/xorima/pointerhelpers"
)

func main() {
	c := complex64(1 + 2i)
	cPtr := pointerhelpers.Complex64(c)
	fmt.Println(*cPtr) // Output: (1+2i)

	nilComplex64Ptr := (*complex64)(nil)
	fmt.Println(pointerhelpers.Complex64Value(nilComplex64Ptr)) // Output: (0+0i)
}
```

#### complex128

```go
package main

import (
	"fmt"
	"github.com/xorima/pointerhelpers"
)

func main() {
	c := complex(1 + 2i)
	cPtr := pointerhelpers.Complex128(c)
	fmt.Println(*cPtr) // Output: (1+2i)

	nilComplex128Ptr := (*complex128)(nil)
	fmt.Println(pointerhelpers.Complex128Value(nilComplex128Ptr)) // Output: (0+0i)
}
```

#### Embedding Complex64Helper

For users who want to embed `Complex64Helper` into their custom types to gain the pointer helper functionalities directly, the package provides the `Complex64Helper` struct.

```go
package main

import (
	"fmt"
	"github.com/xorima/pointerhelpers"
)

type MyStruct struct {
	pointerhelpers.Complex64Helper
}

func main() {
	s := MyStruct{}
	c := complex64(1 + 2i)
	cPtr := s.Complex64(c)
	fmt.Println(*cPtr) // Output: (1+2i)

	nilComplex64Ptr := (*complex64)(nil)
	fmt.Println(s.Complex64Value(nilComplex64Ptr)) // Output: (0+0i)
}
```

### String

```go
package main

import (
	"fmt"
	"github.com/xorima/pointerhelpers"
)

func main() {
	s := "hello"
	sPtr := pointerhelpers.String(s)
	fmt.Println(*sPtr) // Output: hello

	nilStringPtr := (*string)(nil)
	fmt.Println(pointerhelpers.StringValue(nilStringPtr)) // Output: ""
}
```

## Full API Reference

### Functions

- `Bool(v bool) *bool`
- `BoolValue(v *bool) bool`
- `Int(v int) *int`
- `IntValue(v *int) int`
- `Int8(v int8) *int8`
- `Int8Value(v *int8) int8`
- `Int16(v int16) *int16`
- `Int16Value(v *int16) int16`
- `Int32(v int32) *int32`
- `Int32Value(v *int32) int32`
- `Int64(v int64) *int64`
- `Int64Value(v *int64) int64`
- `Uint(v uint) *uint`
- `UintValue(v *uint) uint`
- `Uint8(v uint8) *uint8`
- `Uint8Value(v *uint8) uint8`
- `Uint16(v uint16) *uint16`
- `Uint16Value(v *uint16) uint16`
- `Uint32(v uint32) *uint32`
- `Uint32Value(v *uint32) uint32`
- `Uint64(v uint64) *uint64`
- `Uint64Value(v *uint64) uint64`
- `Float32(v float32) *float32`
- `Float32Value(v *float32) float32`
- `Float64(v float64) *float64`
- `Float64Value(v *float64) float64`
- `Complex64(v complex64) *complex64`
- `Complex64Value(v *complex64) complex64`
- `Complex128(v complex128) *complex128`
- `Complex128Value(v *complex128) complex128`
- `String(v string) *string`
- `StringValue(v *string) string`

### Structs

- `Complex64Helper`

The `Complex64Helper` struct can be embedded into your custom types to easily incorporate complex64 pointer helper methods.

## Contributing

Contributions are welcome! Please open an issue or submit a pull request with any improvements or bug fixes.

## License

This project is licensed under the MIT License - see the LICENSE file for details.

---

This README now includes information about the `Complex64Helper` struct and its usage. Feel free to customize it further based on your specific requirements.
