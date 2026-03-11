package main

import (
	"fmt"
	"os"

	"github.com/pschlump/englishtocron"
)

func main() {
	texts := []string{
		"every 15 seconds",
		"every minute",
		"every day at 4:00 pm",
		"at 10:00 am",
		"Run at midnight on the 1st and 15th of the month",
		"on Sunday at 12:00",
	}

	for _, text := range texts {
		cron, err := englishtocron.New(text)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error parsing '%s': %v\n", text, err)
			continue
		}
		fmt.Printf("%s: %s\n", text, cron)
	}
}
