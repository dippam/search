package main

type Session struct {
	Id   int  `json:"id"`
	From byte `json:"from"`
	To   byte `json:"to"`
}

type Sessions []Session
