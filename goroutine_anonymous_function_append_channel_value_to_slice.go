package main

import "time"

func processMessages(messages []string) []string {

	// make channel and initialize slice to hold processed words later
	c := make(chan string, len(messages))
	var processed []string

	// loop over every msg in messages
	for _, msg := range messages {
		// go routine with anonymous function to put the result of process() into a channel
		go func(m string) {
			processedMessage := process(m)
			c <- processedMessage
		}(msg) // pass msg into anonymous function
	}

	// loop the amount we expect to receive messages, append slice with values received from channel
	for i := 0; i < len(messages); i++ {
		processed = append(processed, <-c)
	}

	// close channel
	close(c)

	return processed
}

func process(message string) string {
	time.Sleep(1 * time.Second)
	return message + "-processed"
}
