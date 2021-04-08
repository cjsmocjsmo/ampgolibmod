// LICENSE = GNU General Public License, version 2 (GPLv2)
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
	// "os"
	"fmt"
	// "strconv"
	"gopkg.in/mgo.v2/bson"
)

//Set Constants
// const (
// 	OffSet = 10
// )

// GDistAlbum exported
func GDistAlbum() (DAlbum []string) {
	sess := DBcon()
	defer sess.Close()
	MAINc := sess.DB("tempdb1").C("meta1")
	MAINc.Find(nil).Distinct("album", &DAlbum)
	return
}

// GDistAlbum3 exported
func GDistAlbum3() (DAlbum []map[string]string) {
	sess := DBcon()
	defer sess.Close()
	MAINc := sess.DB("maindb").C("maindb")
	var dlist []string
	MAINc.Find(nil).Distinct("album", &dlist)
	for _, d := range dlist {
		DMainc := sess.DB("maindb").C("maindb")
		b1 := bson.M{"album": d}
		b2 := bson.M{}
		var Boo map[string]string = make(map[string]string)
		DMainc.Find(b1).Select(b2).One(&Boo)
		DAlbum = append(DAlbum, Boo)
	}
	return
}

// InsAlbumID exported
func InsAlbumID(alb string) {
	uuid, _ := UUID()
	sess := DBcon()
	defer sess.Close()
	TAlbIc := sess.DB("tempdb2").C("albumid")
	DALBI := map[string]string{"album": alb, "albumid": uuid}
	TAlbIc.Insert(&DALBI)
}

// GtempAlbInfo exported
func GtempAlbInfo(DAlb string) (AlbInfo2 map[string]string) {
	sess := DBcon()
	defer sess.Close()
	AMPc := sess.DB("maindb").C("maindb")
	AMPc.Find(bson.M{"album": DAlb}).One(&AlbInfo2)
	return
}

//GAlbInfo exportec
func GAlbInfo(DAlb map[string]string) (AlbInfo2 map[string]string) {
	sess := DBcon()
	defer sess.Close()
	AMPc := sess.DB("maindb").C("maindb")
	AMPc.Find(bson.M{"album": DAlb["album"]}).One(&AlbInfo2)
	return
}

//  exported
type p2 struct {
	Titlez []string
}

// AlbPipeline exported
func AlbPipeline(DAlb map[string]string) []string {
	var P2 []p2
	sess := DBcon()
	defer sess.Close()
	AMPc := sess.DB("maindb").C("maindb")
	pipeline2 := AMPc.Pipe([]bson.M{
		// bson.M{"$match" = bson.M{"Album" = DAlb}},
		{"$match": bson.M{"album": DAlb["album"]}},
		{"$group": bson.M{"_id": "title", "titlez": bson.M{"$addToSet": "$title"}}},
		{"$project": bson.M{"titlez": 1}},
	}).Iter()
	err := pipeline2.All(&P2)
	CheckError(err, "\n AlbPipeline: Agg Album pipeline2 fucked up")
	fmt.Printf("this is P2 %s", P2)
	return P2[0].Titlez
}

// AddTitleID exported
func AddTitleID(titlez []string) []map[string]string {
	// a := <- AlbPipchan
	var TAAID []map[string]string
	sess := DBcon()
	defer sess.Close()
	AMP2c := sess.DB("maindb").C("maindb")
	for _, boo := range titlez {
		var TAid map[string]string = make(map[string]string)
		AMP2c.Find(bson.M{"title": boo}).Select(bson.M{"title": 1, "fileID": 1, "_id": 0}).One(&TAid)
		TAAID = append(TAAID, TAid)
	}
	return TAAID
}

// GetSImage exported
func GetSImage(alb map[string]string) string {
	sess := DBcon()
	defer sess.Close()
	MAINc := sess.DB("maindb").C("maindb")
	var albinfo map[string]string = make(map[string]string)
	err := MAINc.Find(bson.M{"album": alb["album"]}).Select(bson.M{"_id": 0, "picID": 1}).One(&albinfo)
	CheckError(err, "\n GetSImage:  simage fucked up goaggalbum \n")
	JPGc := sess.DB("thumbnails").C("meta")
	var BsIMAGE map[string]string = make(map[string]string)

	b1 := bson.M{"imgID": albinfo["picID"]}
	b2 := bson.M{"_id": 0, "b64Image": 1}
	err = JPGc.Find(b1).Select(b2).One(&BsIMAGE)
	CheckError(err, "\n GetSImage: NapIMAGE fucked up \n")

	return BsIMAGE["b64Image"]
}

// AlbvieW exported
type AlbvieW struct {
	Artist   string              `bson:"artist"`
	ArtistID string              `bson:"artistID"`
	Album    string              `bson:"album"`
	AlbumID  string              `bson:"albumID"`
	Songs    []map[string]string `bson:"songs"`
	Page     int                 `bson:"page"`
	NumSongs string              `bson:"numsongs"`
	HSImage  string              `bson:"hsimage"`
	Idx      int                 `bson:"idx"`
}

//InsAlbViewID exported
func InsAlbViewID(av AlbvieW) {
	sess := DBcon()
	defer sess.Close()
	AVc := sess.DB("albview").C("albview")
	AVc.Insert(av)
}

// GAlbVCount exported
func GAlbVCount() (AlbV []AlbvieW) {
	// var AlbV []AlbvieW
	sess := DBcon()
	defer sess.Close()
	ALBc := sess.DB("albview").C("albview")
	err := ALBc.Find(nil).All(&AlbV)
	CheckError(err, "GALBVCount: albumcount has fucked up")
	return
}

//AlbumOffset exported
func AlbumOffset() {
	sess := DBcon()
	defer sess.Close()
	ALBcc := sess.DB("albview").C("albview")
	ALBview := GAlbVCount()
	fmt.Printf("THIS IS ALBview FOR ALBUMVIEW %v", ALBview[0].Page)
	fmt.Printf("THIS IS ALBview FOR ALBUMVIEW %v", ALBview[0].Idx)

	var Albcount int = 0
	var page1 int = 1
	for _, alb := range ALBview {
		Albcount++
		switch {
		case Albcount < OffSet:
			// spage := strconv.Itoa(Page)
			var BOO AlbvieW
			BOO.Artist = alb.Artist
			BOO.ArtistID = alb.ArtistID
			BOO.Album = alb.Album
			BOO.AlbumID = alb.AlbumID
			BOO.Songs = alb.Songs
			BOO.Page = page1
			BOO.NumSongs = alb.NumSongs
			BOO.HSImage = alb.HSImage
			BOO.Idx = alb.Idx
			fmt.Printf("\n\n this is boo album page %s \n\n", BOO.Album)
			fmt.Printf("\n\n this is boo page %v \n\n", BOO.Page)
			ALBcc.Update(bson.M{"ArtistID": alb.ArtistID}, BOO)

		case Albcount == OffSet:
			Albcount = 0
			page1++
			// Sspage := strconv.Itoa(Page)
			var MOO AlbvieW
			MOO.Artist = alb.Artist
			MOO.ArtistID = alb.ArtistID
			MOO.Album = alb.Album
			MOO.AlbumID = alb.AlbumID
			MOO.Songs = alb.Songs
			MOO.Page = page1
			MOO.NumSongs = alb.NumSongs
			MOO.HSImage = alb.HSImage
			MOO.Idx = alb.Idx
			ALBcc.Update(bson.M{"AlbumID": alb.AlbumID}, MOO)

			// else if Albcount == OffSet+1 {
			// 		Albcount = 0
			// 		page ++
			// 		st1 := strconv.Itoa(page)
			// 		var GOO AlbvieW
			// 		GOO.Artist = alb.Artist
			// 		GOO.ArtistID = alb.ArtistID
			// 		GOO.Album = alb.Album
			// 		GOO.AlbumID = alb.AlbumID
			// 		GOO.Songs = alb.Songs
			// 		GOO.Page = st1
			// 		GOO.NumSongs = alb.NumSongs
			// 		GOO.HSImage = alb.HSImage
			// 		GOO.Idx = alb.Idx
			// 		ALBcc.Update(bson.M{"albumID":alb.AlbumID}, GOO)
			// 		} else {
			// 			fmt.Println("WTF AlbvieW")
			// 		}
			// 	}

		}
	}
}
