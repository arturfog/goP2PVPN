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
	"fmt"
	"log"
	"os/exec"
)

type Shell struct {
}

func (sh *Shell) Exec(cmd string, arg string) {
	out, err := exec.Command(cmd, arg).Output()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Command finished with out: " + string(out))
}
