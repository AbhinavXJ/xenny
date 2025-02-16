package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	cmd := exec.Command("whoami")
	output, err := cmd.Output()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	cmd1 := exec.Command("hostname")
	output1, err1 := cmd1.Output()
	if err1 != nil {
		fmt.Println("Error:", err1)
		return
	}

	hostname := strings.TrimSuffix(string(output), "\n") + "@" + strings.TrimSuffix(string(output1), "\n")

	for {
		fmt.Print(">")

		fmt.Print(hostname + "# ")

		cmd := exec.Command("pwd")

		output, err := cmd.Output()
		if err != nil {
			fmt.Println("Error:", err)
			return
		}
		pwd_string := strings.TrimSuffix(string(output), "\n")

		fmt.Print(pwd_string)
		fmt.Print("$ ")

		input, err := reader.ReadString('\n')

		if err != nil {
			fmt.Fprintln(os.Stderr, "Error occurred")
		}

		if err := executeInput(input); err != nil {
			fmt.Fprintln(os.Stderr, "Error occurred")

		}

	}

}

func executeInput(input string) error {

	input = strings.TrimSuffix(input, "\n")

	args := strings.Split(input, " ")

	switch args[0] {

	case "cd":
		if len(args) < 2 {
			return errors.New("Path not provided")
		}
		return os.Chdir(args[1])

	case "exit":
		os.Exit(0)

	}

	cmd := exec.Command(args[0], args[1:]...)

	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout

	return cmd.Run()

}
