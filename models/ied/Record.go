package main

import "time"

type Record struct {
	Id          uint16 `json:"id"`
	UUID        string `json:"uuid"`
	Serial      string `json:"serial"`
	Name        string `json:"name"`
	Institution uint16 `json:"institution"`
	Source      Source `json:"source"`
	Category    byte   `json:"category"`
	Date        Date   `json:"date"`
	Log         LogMsg `json:"date"`
}

type Records []Record

type Source struct {
	Name string `json:"name"`
	Date Date   `json:"date"`
}

type LogMsg struct {
	Author string `json:"author"`
	Date   Date   `json:"date"`
	Action string `json:"action"`
}
