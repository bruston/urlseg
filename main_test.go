package main

import "testing"

func TestPaths(t *testing.T) {
	url := "https://www.example.com/foo/bar/baz.js"
	expected := []string{
		"https://www.example.com/foo/bar/baz.js",
		"https://www.example.com/",
		"https://www.example.com/foo",
		"https://www.example.com/foo/bar",
	}
	segs := paths(url)
	if len(segs) != 4 {
		t.Errorf("expecting 4 results, got %d", len(segs))
	}
	for i := range segs {
		if segs[i] != expected[i] {
			t.Errorf("expecting %s but got %s", expected[i], segs[i])
		}
	}
}
