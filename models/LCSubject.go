package main

type LCSubject struct {
	Id   uint16 `json:"id"`
	Name string `json:"name"`
}

type LCSubjects []LCSubject
