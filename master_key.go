// mpw go libray - generates passwords using the masterpassword algorithm
// by Maarten Billemont (http://masterpasswordapp.com/)
// Copyright (C) 2016  Ole Wehrmeyer <me@belogron.de>

// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.

// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU General Public License for more details.

// You should have received a copy of the GNU General Public License
// along with this program.  If not, see <http://www.gnu.org/licenses/>.

package mpw

import (
	"golang.org/x/crypto/scrypt"
	"errors"
	"fmt"
	"crypto/sha256"
	"crypto/hmac"
	"bytes"
)

const scryptParamN int = 32768
const scryptParamr int = 8
const scryptParamp int = 2
const scryptParamLen = 64

const scryptSaltBeginning string = "com.lyndir.masterpassword"

var identiconLeftArm = [...]string{"╔", "╚", "╰", "═" }
var identiconRightArm = [...]string{"╗", "╝", "╯", "═" }
var identiconBody = [...]string{"█", "░", "▒", "▓", "☺", "☻" }
var identiconAccessory = [...]string{
	"◈", "◎", "◐", "◑", "◒", "◓", "☀", "☁", "☂", "☃", "☄", "★", "☆", "☎", "☏", "⎈", "⌂", "☘", "☢", "☣",
	"☕", "⌚", "⌛", "⏰", "⚡", "⛄", "⛅", "☔", "♔", "♕", "♖", "♗", "♘", "♙", "♚", "♛", "♜", "♝", "♞", "♟",
	"♨", "♩", "♪", "♫", "⚐", "⚑", "⚔", "⚖", "⚙", "⚠", "⌘", "⏎", "✄", "✆", "✈", "✉", "✌" }

type MasterKey []byte

// Generates the master key used for password derivation from a given user name and password
func GenerateMasterKey(userName, masterPassword string) (*MasterKey, error) {
	var masterKey = new(MasterKey)
	var err error

	var salt = assembleSalt(userName);

	*masterKey, err = scrypt.Key(
		[]byte(masterPassword),
		*salt,
		scryptParamN,
		scryptParamr,
		scryptParamp,
		scryptParamLen)
	if err != nil {
		fmt.Println(err)
		return nil, errors.New(err.Error())
	}

	return masterKey, nil
}

func GenerateIdenticon(userName, password string) (string, int) {

	hmac := hmac.New(sha256.New, []byte(password))

	hmac.Write([]byte(userName))

	identiconSeed := hmac.Sum(nil)
	var buffer bytes.Buffer

	buffer.Write([]byte(identiconLeftArm[ int(identiconSeed[0]) % len(identiconLeftArm)]))
	buffer.Write([]byte(identiconBody[ int(identiconSeed[1]) % len(identiconBody)]))
	buffer.Write([]byte(identiconRightArm[ int(identiconSeed[2]) % len(identiconRightArm)]))
	buffer.Write([]byte(identiconAccessory[ int(identiconSeed[3]) % len(identiconAccessory)]))


	return buffer.String(), int(identiconSeed[4] % 7 +1)
}

func assembleSalt(userName string) *[]byte {
	// constant salt part + 4 byte for integer + user name
	lenSaltBeg := len(scryptSaltBeginning)
	var saltLength = lenSaltBeg + len(userName) + 4
	var salt = make([]byte, saltLength)

	// Copy the first part over
	copy(salt[:lenSaltBeg], []byte(scryptSaltBeginning))
	//Copy the length in
	var userNameLen uint32 = uint32(len(userName))
	bs := unit32ToByteSlice(userNameLen)
	copy(salt[lenSaltBeg: lenSaltBeg + 4], bs)
	// Copy the user name
	copy(salt[lenSaltBeg + 4:], []byte(userName))
	fmt.Printf("salt %v\n", salt)
	return &salt
}
