package main

import (
	"encoding/json"
	"fmt"
	"strconv"
	"time"
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
	strDate := i.Date
	date, err := time.Parse(time.RFC3339, strDate)
	// If the date is parsed then convert it
	// to human readable format
	if err == nil {
		duration := time.Since(date)
		hours := int(duration.Hours())
		days := hours / 24
		weeks := days / 7
		months := weeks / 4
		if months > 0 {
			strDate = fmt.Sprintf("%d months ago", months)
		} else if weeks > 0 {
			strDate = strconv.Itoa(weeks) + " weeks ago"
		} else if days > 1 {
			strDate = strconv.Itoa(days) + " days ago"
		} else if days == 1 {
			strDate = "Yesterday"
		} else if hours > 0 {
			strDate = strconv.Itoa(hours) + " hours ago"
		} else {
			strDate = "Just now"
		}
	}
	str := fmt.Sprintf("%-*d"+seperator+" %-*s"+seperator+" %-*s", iW, i.Id, tW, i.Title, dW, strDate)
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
