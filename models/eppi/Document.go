package main

import (
	"../institution"
	"../person"
	"../subject"
	"./breviate"
	"time"
)

type Document struct {
	Id           int      `json:"id"`
	Abstract     string   `json:"abstract"`
	Additional   string   `json:"additional"`
	Appointed    Date     `json:"appointed"`
	AuthRef      int      `json:"authRef"`
	Breviate     Breviate `json:"breviate"`
	BreviatePage int      `json:"breviatePage"`
	Chairman     Person   `json:"chairman"`
	CorpAuthors  string   `json:"corpAuthors"`

	// Only four types of document.
	Doctype   string      `json:"doctype"`
	HeldBy    Institution `json:"heldBy"`
	LCSubject Subject     `json:"lcSubject"`
	Sectioned int         `json:"sectioned"`

	// Paper number is arabic, but some documents contain
	// multiple paper numbers.
	PaperNo        string      `json:"paperNo"`
	PersonalAuthor Person      `json:"personalAuthor"`
	Published      Publication `json:"published"`
	Series         string      `json:"series"`

	// Session might be better included in breviates instead.
	Session   Session `json:"session"`
	Source    string  `json:"source"`
	StartPage int     `json:"startPage"`
	Subvol    string  `json:"Subvol"`
	Tables    string  `json:"tables"`
	Title     string  `json:"title"`
}

type Documents []Document

type Publication struct {
	Date      Date        `json:"date"`
	Publisher Institution `json:"publisher"`
}
