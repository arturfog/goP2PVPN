// Copyright (C) 2018  Artur Fogiel
// This file is part of goP2PVPN.
//
// goP2PVPN is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// goP2PVPN is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU General Public License for more details.
//
// You should have received a copy of the GNU General Public License
// along with goP2PVPN.  If not, see <http://www.gnu.org/licenses/>.
package cli

import (
	"os"
	"../fs"
	"fmt"
)

type CLI struct {
	fs fs.Filesystem
}

// Exec executes shell command given as parameter
func (cli* CLI) LS(path string) {
	fmt.Println("LS !!!")
	cli.fs.GetDirContents(path)
}

func (cli* CLI) CP(src string, dst string) error {
	err := os.Rename(src, dst)
	return err
}

func (cli* CLI) RM(path string) error {
	err := os.Remove(path)
	return err
}

func (cli* CLI) MKDIR(path string, perm os.FileMode) error {
	fmt.Println("MKDIR !!!")
	err := os.Mkdir(path, perm)
	return err
}

func (cli* CLI) WGET(url string) {

}

func (cli* CLI) IFCONFIG(iface string) {

}
