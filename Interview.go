package main

import "time"

type Interview struct {
	Code          string    `json:"code"`
	Summary       bool      `json:"summary"`
	Duration      uint16    `json:"duration"`
	Date          Date      `json:"date"`
	AgeGroup      byte      `json:"ageGroup"`
	Residence     byte      `json:"residence"`
	Childhood     uint16    `json:"childhood"`
	Birthplace    uint16    `json:"birthplace"`
	Gender        byte      `json:"gender"`
	Denomination  byte      `json:"denomination"`
}

type Interviews []Interview
