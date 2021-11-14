package main

import (
	"os"
	"testing"
)

func TestStruct(t *testing.T) {
	t.Log("TestStruct")
	item := item{0, "Title", "Date", false}
	if item.Id != 0 {
		t.Error(item.Id != 0)
	}
	if item.Title != "Title" {
		t.Error(item.Title != "Title")
	}
	if item.Date != "Date" {
		t.Error(item.Date != "Date")
	}
	if item.Status != false {
		t.Error(item.Status != false)
	}
}

func TestAddDeleteItem(t *testing.T) {
	t.Log("TestAddItem")
	tdmap := TodoList{}
	Title := "Complete assignment"
	i1 := addItem(&tdmap, Title)
	if len(tdmap) != 1 {
		t.Error("len(tdmap) != 1")
	}
	if tdmap[i1].Title != Title {
		t.Error(tdmap[i1].Title, " != ", Title)
	}
	Title = "Complete assignment 2"
	i2 := addItem(&tdmap, Title)
	if tdmap[i2].Title != Title {
		t.Error(tdmap[i2].Title, " != ", Title)
	}

	deleteItem(&tdmap, i1)
	if len(tdmap) != 1 {
		t.Error("len(tdmap) != 1")
	}
	if tdmap[i1] != nil {
		t.Error("tdmap[0] != nil\n", tdmap[0].toString())
	}
}

func TestMarkItem(t *testing.T) {
	t.Log("TestMarkItem")
	tdmap := TodoList{}
	Title := "Complete assignment"
	index := addItem(&tdmap, Title)
	if tdmap[index].Status != false {
		t.Error("tdmap[0].Status != false")
	}
	markItem(&tdmap, index)
	if tdmap[index].Status != true {
		t.Error("tdmap[0].Status != true")
	}
	unMarkItem(&tdmap, index)
	if tdmap[index].Status != false {
		t.Error("tdmap[0].Status != false")
	}
}

func TestJSONtoParse(t *testing.T) {
	t.Log("TestJSON")
	tdmap := TodoList{}
	Title := "Complete assignment"
	index := addItem(&tdmap, Title)
	json := tdmap[index].toJSON()
	i := parseJSON(json)

	if i.Id != tdmap[index].Id {
		t.Error("i.Id != tdmap[index].Id")
	}
	if i.Title != tdmap[index].Title {
		t.Error("i.Title != tdmap[index].Title")
	}
	if i.Date != tdmap[index].Date {
		t.Error("i.Date != tdmap[index].Date")
	}
	if i.Status != tdmap[index].Status {
		t.Error("i.Status != tdmap[index].Status")
	}
}

func TestFileWriteRead(t *testing.T) {
	t.Log("TestFileWriteRead")
	filename := "test.txt"
	tdmap := TodoList{}
	index := addItem(&tdmap, "Complete assignment")
	addItem(&tdmap, "Complete assignment 2")
	addItem(&tdmap, "Complete assignment 3")
	tdmap.writeToFile(filename)
	tdmap2 := TodoList{}
	readFromFile(tdmap2, filename)
	if len(tdmap) != len(tdmap2) {
		t.Error("len(tdmap) != len(tdmap2)")
	}
	if tdmap[index].Title != tdmap2[index].Title {
		t.Error(tdmap[index].Title, " != ", tdmap2[index].Title)
	}
	os.Remove(filename)
}
