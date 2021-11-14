package main

import (
	"encoding/json"
	"fmt"
	"strconv"
)

type item struct {
	Id     int    `json:"id"`     // Id of the item
	Title  string `json:"title"`  // the item's name
	Date   string `json:"date"`   // Date added
	Status bool   `json:"status"` // Status of the item
}

func (i *item) toString() string {
	str := strconv.Itoa(i.Id) + " " + i.Title + " " + i.Date + " " + strconv.FormatBool(i.Status)
	return str
}

func (i *item) toFormattedString(seperator string, iW, tW, dW int) string {
	str := fmt.Sprintf("%-*d"+seperator+" %-*s"+seperator+" %-*s", iW, i.Id, tW, i.Title, dW, i.Date)
	return str
}

func (i *item) toJSON() string {
	js, err := json.Marshal(i)
	if err != nil {
		//fmt.Println(err)
		return ""
	}
	return string(js)
}

func parseJSON(js string) *item {
	i := item{}
	err := json.Unmarshal([]byte(js), &i)
	if err != nil {
		//fmt.Println("error: ", err)
		return nil
	}
	return &i
}
