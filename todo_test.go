package main

import (
	"os"
	"testing"
)

func TestStruct(t *testing.T) {
	t.Log("TestStruct")
	item := item{0, "Title", "date", false}
	if item.id != 0 {
		t.Error("item.id != 0")
	}
	if item.title != "Title" {
		t.Error("item.title != Title")
	}
	if item.date != "date" {
		t.Error("item.date != date")
	}
	if item.status != false {
		t.Error("item.status != false")
	}
	os.Exit(0)
}
