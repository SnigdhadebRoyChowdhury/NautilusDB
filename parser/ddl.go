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

package parser

import (
	"fmt"
	"strings"

	"NautilusDB/location"
)

func whichDDLCommand(cmd_str string) {
	switch {
	case strings.Split(cmd_str, " ")[0] == "USE":
		useDatabase(strings.Split(cmd_str, " ")[1])
	case strings.Split(cmd_str, " ")[0] == "CREATE" && strings.Split(cmd_str, " ")[1] == "DATABASE":
		createDatabase(strings.Split(cmd_str, " ")[2])
	default:
		fmt.Println("Incorrect command. Please check the available commands for NautilusDB.")
	}
}

func useDatabase(db_name string) {
	if location.CheckExists(db_name) {
		location.ChangeDir(db_name)
	} else {
		fmt.Println("ERROR: Database does not exist. Please use CREATE DATABASE command to create the database.")
	}
}

func createDatabase(db_name string) {
	if location.CheckExists(db_name) {
		fmt.Println("ERROR: The database " + db_name + " already exists. ")
	} else {
		err := location.CreateDir(db_name)

		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println("Database " + db_name + " created...")
		}
	}

}
