package main

type Breviate struct {
	Id       int    `json:"id"`
	Name     string `json:"name"`
	Keywords string `json:"keywords"`
	Year     byte   `json:"year"`
	Volume   byte   `json:"volume"`
}

type Breviates []Breviate
