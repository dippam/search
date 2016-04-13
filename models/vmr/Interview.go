package main

import "time"
import "../place"

type Interview struct {
	Id           byte         `json:"id"`
	AgeGroup     byte         `json:"ageGroup"`
	Birthplace   Place        `json:"birthplace"`
	Childhood    Place        `json:"childhood"`
	Code         string       `json:"code"`
	Date         Date         `json:"date"`
	Denomination Denomination `json:"denomination"`
	Duration     int          `json:"duration"`
	Gender       Gender       `json:"gender"`
	Residence    Place        `json:"residence"`
	Summary      string       `json:"summary"`
}

type Interviews []Interview

type Denomination struct {
	Id   byte   `json:"id"`
	Name string `json:"name"`
}

type Gender struct {
	Name string `json:"name"`
}
