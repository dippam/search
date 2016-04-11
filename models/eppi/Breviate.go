package main

type Breviate struct {
	Id       uint16 `json:"id"`
	Name     string `json:"name"`
	Keywords string `json:"name"`
	Year     byte   `json:"name"`
	Volume   byte   `json:"name"`
}

type Breviates []Breviate
