package main

type Institution struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

type Institutions []Institution
