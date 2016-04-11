package main

import "time"

type Document struct {
	Id             uint16 `json:"id"`
	Abstract       string `json:"abstract"`
	Additional     string `json:"additional"`
	Appointed      string `json:"appointed"`
	AuthRef        string `json:"authRef"`
	BreviatePage   string `json:"breviatePage"`
	Chairman       string `json:"chairman"`
	CorpAuthors    string `json:"corpAuthors"`
	Doctype        string `json:"doctype"`
	HeldBy         string `json:"heldBy"`
	LCSubject      string `json:"lcSubject"`
	PagesSectioned string `json:"pagesSectioned"`
	PaperNo        string `json:"paperNo"`
	PersonalAuthor string `json:"personalAuthor"`
	Published      string `json:"published"`
	Publisher      string `json:"publisher"`
	Series         string `json:"series"`
	Session        string `json:"session"`
	Source         string `json:"source"`
	StartPage      string `json:"startPage"`
	Subvol         string `json:"Subvol"`
	Tables         string `json:"tables"`
	Title          string `json:"titlr"`
}

type Documents []Document
