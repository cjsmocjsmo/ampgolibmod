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
	"fmt"
	"github.com/bogem/id3v2"
	"log"
	"os"
	"path/filepath"
	"strconv"
)

func getDirPath(apath string) string {
	return filepath.Dir(apath)
}

func getFileInfo(apath string) (filename string, size string) {
	ltn, err := os.Open(apath)
	CheckError(err, "getFileInfo: file open has fucked up")
	defer ltn.Close()
	ltnInfo, _ := ltn.Stat()
	filename = ltnInfo.Name()
	size = strconv.FormatInt(ltnInfo.Size(), 10)
	return
}

func getExtention(apath string) string {
	return filepath.Ext(apath)
}

func getMetaData(apath string) (Artist string, Album string, Title string, Genre string) {
	tag, err := id3v2.Open(apath, id3v2.Options{Parse: true})
	CheckError(err, "Error while opening mp3 file")
	defer tag.Close()
	Artist = tag.Artist()
	Album = tag.Album()
	Title = tag.Title()
	Genre = tag.Genre()
	return
}

//DumpArtToFile is exported
func DumpArtToFile(apath string) string {
	tag, err := id3v2.Open(apath, id3v2.Options{Parse: true})
	CheckError(err, "Error while opening mp3 file")
	defer tag.Close()
	pictures := tag.GetFrames(tag.CommonID("Attached picture"))
	outFile2 := ""
	for _, f := range pictures {
		pic, ok := f.(id3v2.PictureFrame)
		if !ok {
			log.Fatal("Couldn't assert picture frame")
		}

		// outFile2 := "./" + tag.Artist() + "_-_" + tag.Album() + ".jpg"
		// print("this is outfile")
		// print(outFile)
		outFile2 = "/root/thumb/" + tag.Artist() + "_-_" + tag.Album() + ".jpg"
		print("this is outfile2")
		print(outFile2)
		// f, err := os.Create(outFile)
		// defer f.Close()
		// CheckError(err, "outfile1 creation has fucked up")
		// n2, err := f.Write(pic.Picture)
		// CheckError(err, "outfile1 Write has fucked up")

	// 	g, err := os.Create(outFile2)
	// 	defer g.Close()
	// 	if err != nil {
	// 		fmt.Println(f)
	// 		fmt.Println(outFile2)
	// 		fmt.Println(err)
	// 	}
	// 	// CheckError(err, "outfile creation has fucked up")
	// 	n3, err := g.Write(pic.Picture)
	// 	CheckError(err, "outfile2 Write has fucked up")

	// 	// fmt.Println(n2, "bytes written successfully")
	// 	fmt.Println(n3, "bytes written successfully")
	// }
	return outFile2
}

//Tagmap exported
type Tagmap struct {
	// ID bson.ObjectId `bson:"_id,omitempty"`
	Dirpath   string `bson:"dirpath"`
	Filename  string `bson:"filename"`
	Extension string `bson:"extension"`
	FileID    string `bson:"fileID"`
	Filesize  string `bson:"filesize"`
	Artist    string `bson:"artist"`
	ArtistID  string `bson:"artistID"`
	Album     string `bson:"album"`
	AlbumID   string `bson:"albumID"`
	Title     string `bson:"title"`
	Genre     string `bson:"genre"`
	Page      string `bson:"page"`
	PicID     string `bson:"picID"`
	PicDB     string `bson:"picDB"`
	PicCol    string `bson:"picCol"`
	Idx       string `bson:"idx"`
}

//TagMap exported
func TagMap(apath string) (TagMap Tagmap) {
	picpath := DumpArtToFile(apath)
	zoo := Thumbnails(picpath)
	fname, size := getFileInfo(apath)
	artist, album, title, genre := getMetaData(apath)
	TagMap.Dirpath = getDirPath(apath)
	TagMap.Filename = fname
	TagMap.Extension = getExtention(apath)
	TagMap.FileID, _ = UUID()
	TagMap.Filesize = size
	TagMap.Artist = artist
	TagMap.ArtistID = "None"
	TagMap.Album = album
	TagMap.AlbumID = "None"
	TagMap.Title = title
	TagMap.Genre = genre
	TagMap.PicID = zoo.ImgID
	TagMap.PicDB = "None"
	TagMap.PicCol = "None"
	TagMap.Idx = "None"

	ses := DBcon()
	defer ses.Close()
	tagz := ses.DB("tempdb1").C("meta1")
	tagz.Insert(TagMap)
	return
}
