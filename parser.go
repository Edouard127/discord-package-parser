package main

import (
	"encoding/json"
	"fmt"
	"github.com/gocarina/gocsv"
	"os"
	"path/filepath"
	"strings"
)

var extensions = []string{".json", ".csv"}
var export map[string][]string

func DoParse(path string) map[string][]string {
	parseChannels(path)

	for channel, _ := range export {
		parseMessages(channel)
	}

	return export
}

func parseChannels(path string) {
	filepath.Walk(path, func(path string, info os.FileInfo, err error) error {
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

			export[string(data.Id)] = nil // Persist the map between function calls
		}

		return nil
	})
}

func parseMessages(channel string) {
	var data []Message
	var cast []string

	file, rt := findMessageFile(
		fmt.Sprintf("./messages/c%s/messages", channel), 0)

	switch rt {
	case -1:
		fmt.Printf("No messages found at %s, this is probably an empty channel\n", channel)
		return
	case 0:
		json.NewDecoder(file).Decode(&data)
	case 1:
		gocsv.UnmarshalFile(file, &data)
	}

	for _, v := range data {
		cast = append(cast, string(v.Id))
	}

	export[channel] = cast
	file.Close()

	return
}

func findMessageFile(path string, index int) (*os.File, int) {
	if index >= len(extensions) {
		return nil, -1
	}

	file, err := os.Open(path + extensions[index])
	if err != nil {
		return findMessageFile(path, index+1)
	}

	return file, index
}
