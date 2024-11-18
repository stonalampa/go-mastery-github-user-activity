package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"main/api"
	"main/output-mapper"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("Enter a username or 'q!' to quit: ")
		scanner.Scan()
		input := strings.TrimSpace(scanner.Text())

		if (input == "q!") {
			fmt.Println("Exiting...")
			return;
		}

		fmt.Println("Fetching user activity for", input)
		response, err := api.FetchUserActivity(input)
		if err != nil {
			fmt.Println("Error fetching user activity:", err)
			continue
		}

		fmt.Println("--------------------------------")
		mappedEvents, err := outputMapper.MapUserActivity(response)
		if err != nil {
			fmt.Println("Error mapping user activity:", err)
			continue
		}

		for _, event := range mappedEvents {
			fmt.Println(event)
		}
		fmt.Println("--------------------------------")
	}
}
