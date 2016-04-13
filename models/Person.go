package main

type Person struct {
	Id       int    `json:"id"`
	Forename string `json:"forename"`
	Surname  string `json:"surname"`
	Title    string `json:"title"`
}

type People []Person
