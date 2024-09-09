package main

import (
	"encoding/json"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
)

var export []string

func main() {
	if err := WalkFiles("./messages", parseChannelId); err != nil {
		panic(err)
	}

	if len(export) == 0 {
		fmt.Println("No channels found")
		return
	}

	fmt.Println("Exporting", len(export), "channels to channels.txt")

	file, _ := os.Create("channels.txt")

	for _, id := range export {
		file.WriteString(id + "\n")
	}

	file.Close()
}

type channel struct {
	Id string `json:"id"`
}

func parseChannelId(path string, info fs.FileInfo, _ error) error {
	if info.Name() == "channel.json" {
		var data channel

		// Read the file
		file, err := os.Open(path)
		if err != nil {
			return err
		}

		err = json.NewDecoder(file).Decode(&data)
		if err != nil {
			return err
		}

		export = append(export, data.Id)
	}

	return nil
}

// WalkFiles walks the file tree rooted at root, calling walkFn for each file or
// directory in the tree, including root.
func WalkFiles(root string, walkFn filepath.WalkFunc) error {
	return filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if info == nil {
			return err
		}

		if info.IsDir() {
			return nil
		}

		return walkFn(path, info, err)
	})
}
