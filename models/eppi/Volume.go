package main

type Volume struct {
	Id   int  `json:"id"`
	Name byte `json:"name"`
}

type Volumes []Volume
