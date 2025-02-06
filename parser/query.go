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
)

var (
	DDL_IDENTIFIER = []string{"use", "create"}
)

func checkValueExists(v string, list []string) bool {
	for _, val := range list {
		if val == v {
			return true
		}
	}
	return false
}

func CommandType(str string) {
	lower_str := strings.ToLower(str)
	switch {
	case checkValueExists(strings.Split(lower_str, " ")[0], DDL_IDENTIFIER) && strings.Split(lower_str, " ")[0] == "use":
		useDatabase(strings.Split(lower_str, " ")[1])
	default:
		fmt.Println("Not able to parse incorrect command. Please check the available commands for NautilusDB.")
	}

}
