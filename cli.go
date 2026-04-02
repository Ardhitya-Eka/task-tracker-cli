package main

import (
	"fmt"
	"strconv"
	"strings"
)

func RunCLI(data *Data, args []string) {
	if len(args) < 2 {
		fmt.Println("Command not provided")
		return
	}

	command := args[1]
	argsJoin := args[2:]
	result := strings.Join(argsJoin, " ")
	switch command {

	case "add":
		if len(args) < 3 {
			fmt.Println("Usage: add <description>")
			return
		}

		err := add(result)
		if err != nil {
			fmt.Println("Failed to Add Description, Erorr :", err)
			return
		}

	case "delete":
		id, err := strconv.Atoi(args[2])
		if err != nil {
			fmt.Println("Invalid ID")
			return
		}

		delete(id)

	case "update":

		if len(args) < 4 {
			fmt.Println("Usage: update <id> <description>")
			return
		}

		id, err := strconv.Atoi(args[2])
		if err != nil {
			fmt.Println("Invalid ID")
			return
		}

		joinArgs := args[3:]

		description := strings.Join(joinArgs, " ")

		err = updateDescription(id, description)
		if err != nil {
			fmt.Println("Error:", err)
			return
		}

		fmt.Println("Updating Task, ID :", id)

	case "in-progres":
		id, _ := strconv.Atoi(args[2])
		err := changeStatus(id, StatusInProgress)
		if err != nil {
			fmt.Println("Failed to change status, Error :", err)
			return
		}

		fmt.Println("Data Succesfully Update, ID : ", id)
	case "done":
		id, _ := strconv.Atoi(args[2])
		err := changeStatus(id, StatusDone)
		if err != nil {
			fmt.Println("Failed to change status, Error :", err)
			return
		}

		fmt.Println("Data Succesfully Update, ID : ", id)

	case "--help":
		fmt.Println("Usage add : add <descrition>")
		fmt.Println("Usage update : update <id> <description>")
		fmt.Println("Usage in-progres : in-progres <id>")
		fmt.Println("Usage done : done <id>")
		fmt.Println("Usage delete : delete <id>")

	default:
		fmt.Println("run : <--help>")
	}
}
