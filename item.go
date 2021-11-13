package main

import "strconv"

type item struct {
	id     int    // id of the item
	title  string // the item's name
	date   string // date added
	status bool   // status of the item
}

func (i *item) toString() string {
	str := strconv.Itoa(i.id) + " " + i.title + " " + i.date + " " + strconv.FormatBool(i.status)
	return str
}

func (i *item) toJSON() string {
	str := "{\"id\":" + strconv.Itoa(i.id) + ",\"title\":\"" + i.title + "\",\"date\":\"" + i.date + "\",\"status\":" + strconv.FormatBool(i.status) + "}"
	return str
}

func parseJSON(str string) *item {
	var i item
	i.id = 0
	i.title = "unset"
	i.date = "unset"
	i.status = false
	return &i
}
