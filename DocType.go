package main

type DocType struct {
	Id   uint16 `json:"id"`
	Name string `json:"name"`
}

type DocTypes []DocType
