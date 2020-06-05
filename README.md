# JSON Packages

[![Go Report Card](https://goreportcard.com/badge/github.com/itsmurugappan/json)](https://goreportcard.com/report/github.com/itsmurugappan/json)

Useful JSON go packages

```
go get github.com/itsmurugappan/json
```

### 1. JSON Diff 

Compares to JSON strings and gives the top level fields that has changed. 

Example

```
import (
  "fmt"
  json "github.com/itsmurugappan/json/pkg"
)

.....

a := `{"a": "1", "b": [{"a": "2"},{"a": "3"}], "c": {"a": "1"}}`
b := `{"c": {"a": "1", "b": "1"}, "a": "2", "b": [{"a": "2"},{"a": "3"}]}`

fmt.Println(json.Diff(a,b))


#OUT

[a,c]
```

### 2. JSON Parse and Validate

Parse and Validate a HTTP Request Body into the given json Struct

[Example](./pkg/transform_test.go)


#### Credits

The main packages I have used

1. https://github.com/nsf/jsondiff
2. https://github.com/go-playground/validator