package main

type DocType struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

type DocTypes []DocType
