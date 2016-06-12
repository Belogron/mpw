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

import "encoding/binary"

func unit32ToByteSlice(num uint32) []byte {
	bs := make([]byte, 4)
	binary.BigEndian.PutUint32(bs, num)
	return bs
}
