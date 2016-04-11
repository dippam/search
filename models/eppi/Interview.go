package main

import "time"
import "../place"

type Interview struct {
	AgeGroup     byte   `json:"ageGroup"`
	Birthplace   Place  `json:"birthplace"`
	Childhood    Place  `json:"childhood"`
	Code         string `json:"code"`
	Date         Date   `json:"date"`
	Denomination byte   `json:"denomination"`
	Duration     uint16 `json:"duration"`
	Gender       byte   `json:"gender"`
	Id           byte   `json:"id"`
	Residence    Place  `json:"residence"`
	Summary      string `json:"summary"`
}

type Interviews []Interview

type Denomination struct {
	Id   byte   `json:"id"`
	Name string `json:"name"`
}
