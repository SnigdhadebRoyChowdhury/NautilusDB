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
	"fmt"
	"os"
)

func createDir(dir_name string) error {
	err := os.Mkdir(dir_name, 0750)
	if err != nil {
		return err
	}

	return nil
}

func changeDir(dir_name string) error {
	err := os.Chdir(dir_name)

	if err != nil {
		return err
	}

	return nil
}

// The below function checks whether an object exists in the
func CheckDir(dir_name string) error {
	info, err := os.Stat(dir_name)
	if err != nil {
		if os.IsNotExist(err) {
			err = createDir(dir_name)
			if err != nil {
				return err
			}
			return nil
		} else {
			return err
		}
	}
	fmt.Println(info)
	return nil
}
