/*
	NautilusDB is database that has been created to understand how databases work under the hood
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
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var name string
var age int
var isActive bool

func showLicense() {
	year := "2025"
	author := "Snigdhadeb Roy Chowdhury"

	// Display the copyright and license message
	fmt.Printf("Copyright (C) %s %s\n", year, author)
	fmt.Println("This program comes with ABSOLUTELY NO WARRANTY.")
	fmt.Println("This is free software, and you are welcome to redistribute it")
	fmt.Println("under certain conditions.")
	fmt.Println("Please check the license attached in the repository for further clarifications")
	fmt.Println()
}

var rootCmd = &cobra.Command{
	Use:   "NautilusDB",
	Short: "NautilusDB is database that has been created to understand how databases work under the hood",
	Long:  "NautilusDB is database that has been created to understand how databases work under the hood",
	Run: func(cmd *cobra.Command, args []string) {
		// fmt.Printf("Name: %s\n", name)
		// fmt.Printf("Age: %d\n", age)
		// fmt.Printf("Is Active: %v\n", isActive)
		showLicense()
		repl()
	},
}

func init() {
	// Define flags and bind them to variables
	rootCmd.Flags().StringVarP(&name, "name", "n", "", "Your name")
	rootCmd.Flags().IntVarP(&age, "age", "a", 0, "Your age")
	rootCmd.Flags().BoolVarP(&isActive, "active", "v", false, "Is the user active?")
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "Oops. An error while executing Zero '%s'\n", err)
		os.Exit(1)
	}
}
