package main

import "encoding/json"

type Options struct {
	Ignore []string `arg:"-i, --ignore" description:"Ignore certain servers or channels"`
}

type Channel struct {
	Id    json.Number `json:"id"`
	Guild struct {
		Id json.Number `json:"id"`
	} `json:"guild"`
}

type Message struct {
	Id json.Number `csv:"ID" json:"ID"`
}
