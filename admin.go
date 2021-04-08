///////////////////////////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////////////////////////
// LICENSE: GNU General Public License, version 2 (GPLv2)
// Copyright 2016, Charlie J. Smotherman
//
// This program is free software; you can redistribute it and/or
// modify it under the terms of the GNU General Public License v2
// as published by the Free Software Foundation.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU General Public License for more details.
//
// You should have received a copy of the GNU General Public License
// along with this program; if not, write to the Free Software
// Foundation, Inc., 59 Temple Place - Suite 330, Boston, MA  02111-1307, USA.
///////////////////////////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////////////////////////

package ampgolib

import (
	"crypto/rand"
	"encoding/hex"
)

//GK exported
// func GK() string {
// 	token := make([]byte, 32)
// 	rand.Read(token)
// 	return base64.StdEncoding.EncodeToString(token)
// }

// Gk exported
// var Gk = GK()
// var Gk string = "GK"

// // Utoken exported
// func Utoken(args map[string]string) string {
// 	h := args["username"] + ", " + args["password"] + ", " + Gk
// 	h1 := sha512.New()
// 	h1.Write([]byte(h))
// 	return hex.EncodeToString(h1.Sum(nil))
// }

// UUID exported
func UUID() (string, error) {
	uuid := make([]byte, 16)
	n, err := rand.Read(uuid)
	if n != len(uuid) || err != nil {
		return "", err
	}
	uuid[8] = 0x80
	uuid[4] = 0x40
	boo := hex.EncodeToString(uuid)
	return boo, nil
}

// DropTempDBs exported
func DropTempDBs() {
	go func() {
		sesCopy := DBcon()
		defer sesCopy.Close()
		// tmp1 := sesCopy.DB("tempdb1").C("meta1")
		// tmp1.DropCollection()
		sesCopy.DB("tempdb1").DropDatabase()
		sesCopy.DB("tempdb2").DropDatabase()
	}()
}
