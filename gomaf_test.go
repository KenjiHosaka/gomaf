package gomaf

import "testing"

func TestDiff_SuccessDifferentValue(t *testing.T) {
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

	if err != nil {
		t.Fatalf("failed test %#v", err)
	}

	if len(result) != 1 {
		t.Fatalf("failed test %#v", result)
	}

	if result["name"] != "" {
		t.Fatalf("failed test %#v", result)
	}

	if result["age"] != nil {
		t.Fatalf("failed test %#v", result)
	}

	t.Logf("success test %#v", result)
}

func TestDiff_SuccessNoMetaTag(t *testing.T) {
	oldVal := struct {
		Name   string
		MaxAge int
	}{
		Name:   "gomaf",
		MaxAge: 20,
	}

	newVal := struct {
		Name   string
		MaxAge int
	}{
		Name:   "gomaf",
		MaxAge: 0,
	}

	result, err := Diff(oldVal, newVal)

	if err != nil {
		t.Fatalf("failed test %#v", err)
	}

	if len(result) != 1 {
		t.Fatalf("failed test %#v", result)
	}

	if result["name"] != nil {
		t.Fatalf("failed test %#v", result)
	}

	if result["max_age"] == nil {
		t.Fatalf("failed test %#v", result)
	}

	if result["max_age"] != 0 {
		t.Fatalf("failed test %#v", result)
	}

	t.Logf("success test %#v", result)
}

func TestDiff_SuccessEmpty(t *testing.T) {
	oldVal := struct {
		Name string
		Age  int
	}{}

	newVal := struct {
		Name string
		Age  int
	}{}

	result, err := Diff(oldVal, newVal)

	if err != nil {
		t.Fatalf("failed test %#v", err)
	}

	if len(result) != 0 {
		t.Fatalf("failed test %#v", result)
	}

	t.Logf("success test %#v", result)
}

func TestDiff_SuccessSameValue(t *testing.T) {
	oldVal := struct {
		Name string
		Age  int
	}{
		Name: "gomaf",
		Age:  20,
	}

	newVal := struct {
		Name string
		Age  int
	}{
		Name: "gomaf",
		Age:  20,
	}

	result, err := Diff(oldVal, newVal)
	if err != nil {
		t.Fatalf("failed test %#v", err)
	}

	if len(result) != 0 {
		t.Fatalf("failed test %#v", result)
	}

	t.Logf("success test %#v", result)
}

func TestDiff_FailedDifferentType(t *testing.T) {
	oldVal := struct {
		Name string
		Age  int
	}{
		Name: "gomaf",
		Age:  20,
	}

	newVal := struct {
		Name string
		Age  string
	}{
		Name: "gomaf",
		Age:  "20",
	}

	result, err := Diff(oldVal, newVal)

	if err == nil {
		t.Fatalf("failed test %#v", err)
	}

	if result != nil {
		t.Fatalf("failed test %#v", result)
	}

	t.Logf("success test %#v", result)
}
