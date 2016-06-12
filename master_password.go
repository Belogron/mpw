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
	"crypto/hmac"
	"crypto/sha256"
	"bytes"
)

const hmacBeginning string = "com.lyndir.masterpassword"

type MasterPassword struct {
	seed []byte
	masterPasswordType MasterPasswordType
	counter uint32
	siteName string
}

// Creates a MasterPassword with the given MasterKey, site and counter. The type has to be set later or default to maximum security
func GenerateMasterPassword(masterKey *MasterKey, siteName string, counter uint32) *MasterPassword {
	masterPassword := new(MasterPassword)
	masterPassword.counter = counter
	masterPassword.siteName = siteName
	masterPassword.seed = generateSeed(masterKey, siteName, counter)
	return masterPassword
}

// Creates a MasterPassword with the given MasterKey, site, counter and type.
func GenerateMasterPasswordWithType(masterKey *MasterKey, siteName string, counter uint32, mpType MasterPasswordType) *MasterPassword {
	masterPassword := GenerateMasterPassword(masterKey, siteName, counter)
	masterPassword.masterPasswordType = mpType
	return masterPassword
}

func generateSeed(masterKey *MasterKey, siteName string, counter uint32) []byte {
	hmac := hmac.New(sha256.New, *masterKey)

	hmac.Write([]byte(hmacBeginning))
	hmac.Write(unit32ToByteSlice(uint32(len(siteName))))
	hmac.Write([]byte(siteName))
	hmac.Write(unit32ToByteSlice(counter))

	return hmac.Sum(nil)
}

func (mp MasterPassword) SetType(mpType MasterPasswordType) {
	mp.masterPasswordType = mpType
}

func (mp MasterPassword) GetType() MasterPasswordType {
	return mp.masterPasswordType
}

// Generates an returns the password for the specified site
func (mp MasterPassword) GetPassword() string {
	templates := mp.masterPasswordType.GetTemplates()
	template := templates[int(mp.seed[0]) % len(templates)]
	var buffer bytes.Buffer

	for idx, runeValue := range template {
		buffer.WriteRune(fillTemplate(runeValue, mp.seed[idx + 1]))
	}
	return buffer.String()
}

// Returns the corresponding password letter to a template letter and a seed number
// If the template letter is illegal, the template letter is returned (used for spaces in phrase password)
func fillTemplate(template rune, pos byte) rune {
	var letters string
	switch template {
	case 'V':
		letters = "AEIOU"
	case 'C':
		letters = "BCDFGHJKLMNPQRSTVWXYZ"
	case 'v':
		letters = "aeiou"
	case 'c':
		letters = "bcdfghjklmnpqrstvwxyz"
	case 'A':
		letters = "AEIOUBCDFGHJKLMNPQRSTVWXYZ"
	case 'a':
		letters = "AEIOUaeiouBCDFGHJKLMNPQRSTVWXYZbcdfghjklmnpqrstvwxyz"
	case 'n':
		letters = "0123456789"
	case 'o':
		letters = "@&%?,=[]_:-+*$#!'^~;()/."
	case 'x':
		letters = "AEIOUaeiouBCDFGHJKLMNPQRSTVWXYZbcdfghjklmnpqrstvwxyz0123456789!@#$%^&*()"
	default:
		return template
	}
	return []rune(letters)[int(pos) % len(letters)]
}