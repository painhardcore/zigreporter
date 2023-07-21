package main

import (
	"reflect"
	"testing"

	gofeed "github.com/mmcdole/gofeed"
)

func TestEscapeMarkdown(t *testing.T) {
	// Test case 1: No special characters
	input := "Hello, world!"
	expected := "Hello, world!"
	output := escapeMarkdown(input)
	if output != expected {
		t.Errorf("Expected %q, but got %q", expected, output)
	}

	// Test case 2: One special character
	input = "Hello, _world_!"
	expected = "Hello, \\_world\\_!"
	output = escapeMarkdown(input)
	if output != expected {
		t.Errorf("Expected %q, but got %q", expected, output)
	}

	// Test case 3: Multiple special characters
	input = "Hello, *world*! [Link](https://example.com)"
	expected = "Hello, \\*world\\*! \\[Link\\]\\(https://example.com\\)"
	output = escapeMarkdown(input)
	if output != expected {
		t.Errorf("Expected %q, but got %q", expected, output)
	}

	// Test case 4: Empty string
	input = ""
	expected = ""
	output = escapeMarkdown(input)
	if output != expected {
		t.Errorf("Expected %q, but got %q", expected, output)
	}
}
func TestGetFieldValues(t *testing.T) {
	// Test case 1: Single field
	item := &gofeed.Item{Title: "Test post"}
	fields := []string{"Title"}
	expected := []interface{}{"Test post"}
	output := getFieldValues(item, fields)
	if !reflect.DeepEqual(output, expected) {
		t.Errorf("Expected %v, but got %v", expected, output)
	}

	// Test case 2: Multiple fields
	item = &gofeed.Item{Title: "Test post", Link: "https://example.com/post"}
	fields = []string{"Title", "Link"}
	expected = []interface{}{"Test post", "https://example.com/post"}
	output = getFieldValues(item, fields)
	if !reflect.DeepEqual(output, expected) {
		t.Errorf("Expected %v, but got %v", expected, output)
	}

	// Test case 3: Empty fields
	item = &gofeed.Item{Title: "Test post", Link: "https://example.com/post"}
	fields = []string{}
	expected = []interface{}{}
	output = getFieldValues(item, fields)
	if !reflect.DeepEqual(output, expected) {
		t.Errorf("Expected %v, but got %v", expected, output)
	}
}
