package main

type Session struct {
	Id   uint16 `json:"id"`
	From byte   `json:"name"`
	To   byte   `json:"name"`
}

type Sessions []Session
