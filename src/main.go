/*
    This program is free software: you can redistribute it and/or modify
    it under the terms of the GNU Affero General Public License as
    published by the Free Software Foundation, either version 3 of the
    License, or (at your option) any later version.

    This program is distributed in the hope that it will be useful,
    but WITHOUT ANY WARRANTY; without even the implied warranty of
    MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
    GNU Affero General Public License for more details.

    You should have received a copy of the GNU Affero General Public License
    along with this program.  If not, see <https://www.gnu.org/licenses/>.
 */

package main

import (
    "os"
    "fmt"
    "log"

    "clamor/commands"
)

//Lets enter into the cobra cli stuff.
func main() {
    // open a file
    f, err := os.OpenFile("/logs/service.log", os.O_APPEND | os.O_CREATE | os.O_RDWR, 0666)
    if err != nil {
        fmt.Printf("error opening file: %v", err)
    }
    defer f.Close()
    // assign it to the standard logger
    log.SetOutput(f)

    log.Output(1, "starting service.")
    
    commands.Execute()
}
