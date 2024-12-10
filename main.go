package main

import (
	"encoding/csv"
	"log"
	"os"
)

func main() {
	data := DoParse("./messages")

	if len(data) == 0 {
		log.Println("No channels found")
		return
	}

	log.Printf("Found %d channels\n", len(data))

	file, err := os.Create("messages.csv")
	if err != nil {
		log.Fatal(err)
	}

	writer := csv.NewWriter(file)
	writer.Write([]string{"ChannelId", "MessageIds"})

	for channel, messages := range data {
		// Write the map as a flat map
		writer.Write(append([]string{channel}, messages...))
	}

	// Ensure full data integrity
	writer.Flush()
	if err := writer.Error(); err != nil {
		log.Fatal("Error flushing writer:", err)
	}

	// Count messages for log output
	var i int
	for range data {
		i++
	}

	log.Printf("Successfully written %d channels and %d messages\n", len(data), i)
}
