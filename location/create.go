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

package location

import (
	"errors"
	"os"
)

// The below function has been created to re-use the error checking in each function in this package
func checkError(err error) error {
	if err != nil {
		return err
	}
	return nil
}

func CreateDir(dir_name string) error {
	err := os.Mkdir(dir_name, 0750)
	return checkError(err)
}

func ChangeDir(dir_name string) error {
	err := os.Chdir(dir_name)
	return checkError(err)
}

func CheckRoot(dir_name string) string {
	if dir_name == "nautilus" {
		return dir_name
	} else {
		path := "nautilus/" + dir_name
		return path
	}
}

func CheckExists(dir_name string) bool {
	_, err := os.Stat(dir_name)
	if err != nil {
		if os.IsNotExist(err) {
			return false
		} else {
			panic(err)
		}
	}
	return true
}

func CheckLocation(dir_name string) error {
	path := CheckRoot(dir_name)
	if CheckExists(path) {
		err := ChangeDir(path)
		return checkError(err)
	} else if path == "nautilus" {
		err := CreateDir(path)
		return checkError(err)
	} else {
		return errors.New("the path does not exist")
	}
}
