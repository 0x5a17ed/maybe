# maybe ✨

[![Go Reference](https://pkg.go.dev/badge/github.com/0x5a17ed/maybe.svg)](https://pkg.go.dev/github.com/0x5a17ed/maybe)
[![License: 0BSD](https://img.shields.io/badge/License-0BSD-blue.svg)](https://opensource.org/licenses/0BSD)

A tiny Go package that unifies functions returning `(T, error)` and `(T, bool)` into a single interface with `Must()` semantics. Perfect for tests, helpers, optional logic, and expressive control flow.


## ✨ Features

- Wraps `func() (T, error)` and `func() (T, bool)` into a common interface
- Includes a `Must()` helper that panics on failure
- Lightweight, dependency-free, idiomatic Go, no reflection, no nonsense


## 📦 Installation

```bash
go get github.com/0x5a17ed/maybe
````


## 🚀 Quick Start

```go
package main

import (
	"fmt"
	"github.com/0x5a17ed/maybe"
)

func main() {
	fromErr := maybe.Wrap[int](func() (int, error) {
		return 42, nil
	})

	val, err := fromErr.Try()
	fmt.Println(val, err) // 42 nil

	// Use Must for testing or safe startup code
	message := maybe.Must[string](func() (string, bool) {
		return "Hello World!", true
	})
	fmt.Println(message) // outputs: "Hello World!"
}
```


## 💡 Use Cases

* Unifying different forms of "optional" logic
* Treating `(T, error)` and `(T, bool)` uniformly
* Expressive control flow
* Deduplicating error handling in setup code


## 📜 License

This project is licensed under the 0BSD Licence — see the [LICENCE](LICENSE) file for details.

---

<p align="center">Made with ❤️ for expressive control flow 🌟</p>
