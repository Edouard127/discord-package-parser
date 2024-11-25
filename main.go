package main

import (
	"encoding/json"
	"fmt"
	"github.com/gocarina/gocsv"
	"os"
	"path/filepath"
	"strings"
)

var export = make(map[Channel][]Message)

func main() {
	if err := ParseChannel("./messages"); err != nil {
		panic(err)
	}

	if len(export) == 0 {
		fmt.Println("No channels found")
		return
	}

	fmt.Printf("Found %d channels\n", len(export))

	for channel := range export {
		ParseMessages(channel)
	}

	out, _ := os.Create("messages.txt")

	for channel, messages := range export {
		out.WriteString(channel.Id.String() + ":\n")

		for _, message := range messages {
			out.WriteString(fmt.Sprintf("%s,", message.Id))
		}

		out.Seek(-1, 2)
		out.WriteString("\n\n")
	}

	out.Close()

	fmt.Println("Done")
}

func ParseChannel(path string) error {
	return filepath.Walk(path, func(path string, info os.FileInfo, err error) error {
		if info == nil {
			return err
		}

		fileName := strings.TrimSuffix(info.Name(), filepath.Ext(info.Name()))
		if fileName == "channel" {
			var data Channel

			file, err := os.Open(path)
			if err != nil {
				return err
			}

			err = json.NewDecoder(file).Decode(&data)
			if err != nil {
				return err
			}

			export[data] = nil
		}

		return nil
	})
}

func ParseMessages(channel Channel) {
	var data []Message
	var err error

	file, rt := FindMessageFile(
		fmt.Sprintf("./messages/c%s/messages", channel.Id), 0)
	switch rt {
	case -1:
		fmt.Printf("No messages found at %s, this is probably an empty channel\n", channel.Id)
		return
	case 0:
		err = json.NewDecoder(file).Decode(&data)
	case 1:
		err = gocsv.UnmarshalFile(file, &data)
	}

	if err != nil {
		fmt.Println(err)
		return
	}

	file.Close()
	export[channel] = data
	return
}

func FindMessageFile(path string, index int) (*os.File, int) {
	if index >= len(extensions) {
		return nil, -1
	}

	file, err := os.Open(path + extensions[index])
	if err != nil {
		return FindMessageFile(path, index+1)
	}

	return file, index
}

type Channel struct {
	Id    json.Number `json:"id"`
	Name  string      `json:"name"`
	Guild Guild       `json:"guild"`
}

type Guild struct {
	Id   json.Number `json:"id"`
	Name string      `json:"name"`
}

type Message struct {
	Id        json.Number `csv:"ID" json:"ID"`
	Timestamp string      `csv:"Timestamp" json:"Timestamp"`
	Contents  string      `csv:"Contents" json:"Contents"`
}

var extensions = []string{".json", ".csv"}
