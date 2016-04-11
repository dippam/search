package main

import (
	"../institution"
	"../person"
	"../subject"
	"session"
	"time"
)

type Document struct {
	Id             uint16      `json:"id"`
	Abstract       string      `json:"abstract"`
	Additional     string      `json:"additional"`
	Appointed      Date        `json:"appointed"`
	AuthRef        string      `json:"authRef"`
	BreviatePage   string      `json:"breviatePage"`
	Chairman       Person      `json:"chairman"`
	CorpAuthors    string      `json:"corpAuthors"`
	Doctype        string      `json:"doctype"`
	HeldBy         Institution `json:"heldBy"`
	LCSubject      Subject     `json:"lcSubject"`
	PagesSectioned string      `json:"pagesSectioned"`
	PaperNo        string      `json:"paperNo"`
	PersonalAuthor string      `json:"personalAuthor"`
	Publication    Publication `json:"publisher"`
	Series         string      `json:"series"`
	Session        Session     `json:"session"`
	Source         string      `json:"source"`
	StartPage      string      `json:"startPage"`
	Subvol         string      `json:"Subvol"`
	Tables         string      `json:"tables"`
	Title          string      `json:"titlr"`
}

type Documents []Document

type Publication struct {
	Publisher Institution `json:"publisher"`
	Date      Date        `json:"date"`
}
