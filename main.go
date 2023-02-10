package main

import "fmt"

func main() {
	client := NewClient()
	defer client.Close()

	// Create some rooms
	client.CreateRoom(map[string]interface{}{
		"id": "room-1",
		"categories": map[string]interface{}{
			"cosiness": 10,
			"light":    2,
		},
	})

	// Add some messages to the rooms
	client.SendMessage("room-1", map[string]interface{}{
		"sender":    "human-1",
		"recipient": "human-2",
		"category":  "welcome",
		"message":   "Oh hai there",
	})

	client.SendMessage("room-1", map[string]interface{}{
		"sender":    "human-2",
		"recipient": "human-1",
		"category":  "welcome",
		"message":   "Right back at you",
	})

	// Let's render out some data
	messages := client.GetAllMessagesForRoom("room-1")
	for _, message := range messages {
		fmt.Printf("Message: %v\n", message)
	}
}
