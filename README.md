# urlpath

[![Go Reference](https://pkg.go.dev/badge/cattlecloud.net/go/urlpath.svg)](https://pkg.go.dev/cattlecloud.net/go/urlpath)
[![License](https://img.shields.io/github/license/cattlecloud/urlpath?color=7C00D8&style=flat-square&label=License)](https://github.com/cattlecloud/urlpath/blob/main/LICENSE)
[![Build](https://img.shields.io/github/actions/workflow/status/cattlecloud/urlpath/ci.yaml?style=flat-square&color=0FAA07&label=Tests)](https://github.com/cattlecloud/urlpath/actions/workflows/ci.yaml)

`urlpath` provides a way to parse URL path elements using a schema

For users of [gorilla/mux](https://github.com/gorilla/mux).

### Getting Started

The `urlpath` package can be added to a project by running:

```shell
go get cattlecloud.net/go/urlpath@latest
```

```go
import "cattlecloud.net/go/urlpath"
```

### Examples

##### mux definition

Make use of gorilla's path variables.

```go
router.Handle("/v1/{category}/{name}, newHandler())
```

##### parsing schema

Create a `Schema` and call `Parse` to extract the path variables.

```go
var (
  category int
  name     string
)

err := urlpath.Parse(request, urlpath.Schema {
  "category": urlpath.Int(&category),
  "name":     urlpath.String(&name),
})
```

### License

The `cattlecloud.net/go/urlpath` module is open source under the [BSD](LICENSE) license.
