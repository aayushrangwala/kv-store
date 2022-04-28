package main

import (
	"bufio"
	"cryptowatch/backend-go/internal/factory"
	"fmt"
	"os"
	"strings"

	"cryptowatch/backend-go/internal/cmd"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print("> ")
		text, _ := reader.ReadString('\n')

		// split the text into command strings
		operation := strings.Fields(text)

		datastore := factory.NewStoreFactory().GetInmemoryStore()
		command, err := cmd.GetCommandExecutor(datastore, operation...)
		if err != nil {
			fmt.Println(err.Error())

			continue
		}

		if err := command.Execute(); err != nil {
			fmt.Println(err.Error())

			continue
		}
	}
}
