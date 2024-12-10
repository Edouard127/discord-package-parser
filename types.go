package main

import "encoding/json"

type Channel struct {
	Id json.Number `json:"id"`
}

type Message struct {
	Id json.Number `csv:"ID" json:"ID"`
}
