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
	"os"
)

func CreateDir(dir_name string) error {
	err := os.Mkdir(dir_name, 0750)
	if err != nil {
		return err
	}

	return nil
}

func ChangeDir(dir_name string) error {
	err := os.Chdir(dir_name)
	if err != nil {
		return err
	}

	return nil
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

func CheckRootLocation(dir_name string) error {
	if CheckExists(dir_name) {
		err := ChangeDir(dir_name)
		if err != nil {
			return err
		}

		return nil
	} else {
		err := CreateDir(dir_name)
		if err != nil {
			return err
		}

		err = ChangeDir(dir_name)
		if err != nil {
			return err
		}

		return nil
	}
}
