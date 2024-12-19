package main

import (
	"encoding/csv"
	"github.com/alexflint/go-arg"
	"log"
	"os"
)

var options Options

func main() {
	arg.MustParse(&options)

	data := DoParse("./messages")

	if len(data) == 0 {
		log.Println("No channels found")
		return
	}

	log.Printf("Found %d channels (including empty channels)\n", len(data))

	file, err := os.Create("messages.csv")
	if err != nil {
		log.Fatal(err)
	}

	writer := csv.NewWriter(file)
	writer.Write([]string{"ChannelID", "MessageIDs"})

	var messageCount int
	var ignoreCount int

	for channel, messages := range data {
		if len(messages) > 0 {
			// As of 2024-12-18, Discord want the message format as follow
			// https://docs.google.com/spreadsheets/d/1XvVHgET0LYrUiDvRy2cPfBMQTIr3AulYkpLbUVdQjGk/
			for _, message := range messages {
				writer.Write([]string{channel, message})
			}

			messageCount += len(messages)
		} else {
			ignoreCount++
		}
	}

	// Ensure full data integrity
	writer.Flush()
	if err := writer.Error(); err != nil {
		log.Fatal("Error flushing writer:", err)
	}

	log.Printf("Successfully written %d channels and %d messages\n", len(data)-ignoreCount, messageCount)
}
