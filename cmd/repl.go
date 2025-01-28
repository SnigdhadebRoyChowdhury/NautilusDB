/*
	TyniDB is database that has been created to understand how databases work under the hood
    Copyright (C) 2025  Snigdhadeb Roy Chowdhury

    This program is free software: you can redistribute it and/or modify
    it under the terms of the GNU General Public License as published by
    the Free Software Foundation, either version 3 of the License, or
    (at your option) any later version.

    This program is distributed in the hope that it will be useful,
    but WITHOUT ANY WARRANTY; without even the implied warranty of
    MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
    GNU General Public License for more details.

    You should have received a copy of the GNU General Public License
    along with this program.  If not, see <https://www.gnu.org/licenses/>.

	For communications or inquiries, contact:
	Email: snigdhadeb_roychowdhury@outlook.com
*/

package cmd

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func repl() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Welcome to NautilusDB")
	fmt.Println("Type 'exit' to quit")
	for {
		// Display a prompt
		fmt.Print("> ")

		// Read input from the user
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)

		// Exit condition
		if input == "exit" {
			fmt.Println("Goodbye!")
			break
		}

		// Evaluate and respond
		fmt.Printf("You typed: %s\n", input)
	}
}
