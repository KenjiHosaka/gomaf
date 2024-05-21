# gomaf
[![codecov](https://codecov.io/github/KenjiHosaka/gomaf/graph/badge.svg?token=GTJQU8G30F)](https://codecov.io/github/KenjiHosaka/gomaf)

Output the difference of struct to map

## Installation
```
go get github.com/KenjiHosaka/gomaf
```

## How to use
```golang

oldVal := struct {
    Name string `gomaf:"name"`
    Age  int    `gomaf:"age"`
}{
    Name: "gomaf",
    Age:  20,
}

newVal := struct {
    Name string `gomaf:"name"`
    Age  int    `gomaf:"age"`
}{
    Name: "",
    Age:  20,
}

result, err := Diff(oldVal, newVal)
// result: {"name":""}

```