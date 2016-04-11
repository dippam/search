package main

type Volume struct {
	Id   uint16 `json:"id"`
	Name byte   `json:"name"`
}

type Volumes []Volume
