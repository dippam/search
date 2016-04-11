package main

type Institution struct {
	Id   uint16 `json:"id"`
	Name string `json:"name"`
}

type Institutions []Institution
