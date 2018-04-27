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
package fs

import (
	"os"
	"io/ioutil"
)

type Filesystem struct {

}

func (fs *Filesystem) Mkdir(path string) {
	os.Mkdir(path, 0777)
}

func (fs* Filesystem) Rm(path string) {
	os.Remove(path)
}

func (fs* Filesystem) GetDirContents(path string) ([]os.FileInfo, error){
	files, err := ioutil.ReadDir("./")
	return files, err
}

func (fs* Filesystem) CreateEmptyFile(path string, sizeKb int64) {
	fp, err := os.Create(path)
	if err != nil {
		zeroByte := make([]byte, 1)
		fp.Write(zeroByte)
		fp.Close()
	}
}
