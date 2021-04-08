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

// import (
// 	// "fmt"
// 	"gopkg.in/mgo.v2"
// )

//TextSearchIndexes exported
// func TextSearchIndexes() {
// 	sesC := DBcon()
// 	defer sesC.Close()
// 	ARTc := sesC.DB("artistview").C("artistviews")
// 	ALBc := sesC.DB("albview").C("albview")
// 	MAINc := sesC.DB("maindb").C("maindb")
// 	ARindex := mgo.Index{
// 		Key: []string{"$text:artist"},
// 	}

// 	// err := ARTc.EnsureIndex(ARindex)
// 	// CheckError(err, "Artisttext index fucked up")
// 	// fmt.Println("ArtistText index created")

// 	// ALindex := mgo.Index{
// 	// 	Key: []string{"$text:album"},
// 	// }
// 	// err = ALBc.EnsureIndex(ALindex)
// 	// CheckError(err, "Albumtext index fucked up")
// 	// fmt.Println("AlbumText index created")

// 	// Mindex := mgo.Index{
// 	// 	Key: []string{"$text:title"},
// 	// }
// 	// err = MAINc.EnsureIndex(Mindex)
// 	// CheckError(err, "Songtext index fucked up")
// 	// fmt.Println("Songtext index created")
// }
