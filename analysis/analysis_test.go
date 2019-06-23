package analysis

import (
	"testing"
)

func TestIsRedirectTarget(t *testing.T) {
	var result bool
	result = IsRedirectTarget("/foo", "/foo")
	if result != true {
		t.Fatal("failed")
	}

	result = IsRedirectTarget("/foo", "/bar")
	if result != false {
		t.Fatal("failed")
	}

	result = IsRedirectTarget("/foo/bar", "/foo")
	if result != true {
		t.Fatal("failed")
	}
}

func TestGetConfigIdx(t *testing.T) {
	var result int
	config := []string{
		"/bar",
		"/foo",
	}
	result, _ = GetConfigIdx("/foo", config)
	if result != 1 {
		t.Fatal("failed")
	}

	result, _ = GetConfigIdx("/piyo", config)
	if result != -1 {
		t.Fatal("failed")
	}
}

func TestGetRedirectPath(t *testing.T) {
	var result string

	result, _ = GetRedirectPath("/foo", "/foo", true)
	if result != "" {
		t.Fatal(result)
	}

	result, _ = GetRedirectPath("/foo/bar", "/foo", true)
	if result != "/bar" {
		t.Fatal(result)
	}

	result, _ = GetRedirectPath("/foo/bar", "/foo", false)
	if result != "/foo/bar" {
		t.Fatal(result)
	}
}
