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

type MasterPasswordType int

const (
	MaximalPassword MasterPasswordType = iota
	LongPassword
	MediumPassword
	ShortPassword
	BasicPassword
	PINPassword
	NamePassword
	PhrasePassword
)

func (t MasterPasswordType) GetTemplates() []string {
	switch t {
	case MaximalPassword:
		return []string{"anoxxxxxxxxxxxxxxxxx", "axxxxxxxxxxxxxxxxxno"}
	case LongPassword:
		return []string{
			"CvcvnoCvcvCvcv",
			"CvcvCvcvnoCvcv",
			"CvcvCvcvCvcvno",
			"CvccnoCvcvCvcv",
			"CvccCvcvnoCvcv",
			"CvccCvcvCvcvno",
			"CvcvnoCvccCvcv",
			"CvcvCvccnoCvcv",
			"CvcvCvccCvcvno",
			"CvcvnoCvcvCvcc",
			"CvcvCvcvnoCvcc",
			"CvcvCvcvCvccno",
			"CvccnoCvccCvcv",
			"CvccCvccnoCvcv",
			"CvccCvccCvcvno",
			"CvcvnoCvccCvcc",
			"CvcvCvccnoCvcc",
			"CvcvCvccCvccno",
			"CvccnoCvcvCvcc",
			"CvccCvcvnoCvcc",
			"CvccCvcvCvccno"}
	case MediumPassword:
		return []string{"CvcnoCvc", "CvcCvcno"}
	case ShortPassword:
		return []string{"Cvcn"}
	case BasicPassword:
		return []string{"aaanaaan", "aannaaan", "aaannaaa"}
	case PINPassword:
		return []string{"nnnn"}
	case NamePassword:
		return []string{"cvccvcvcv"}
	case PhrasePassword:
		return []string{"cvcc cvc cvccvcv cvc", "cvc cvccvcvcv cvcv", "cv cvccv cvc cvcvccv" }
	default:
		return nil
	}
}
