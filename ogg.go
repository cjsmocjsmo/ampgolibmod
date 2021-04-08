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
	"bufio"
	"crypto/rand"
	"crypto/sha512"
	"encoding/hex"
	"fmt"
	"os"
	"path"
	"strconv"
	"strings"
	// "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

//tags from filenames

func nameSplit(f string) map[string]string {
	Tracks := ""
	Artists := ""
	Albums := ""
	Songs := ""

	fs1 := strings.Replace(f, ".ogg", "", 1)
	fs := strings.Split(fs1, "_-_")
	fsl := len(fs)
	switch {
	case fsl == 4:
		fs2 := strings.Split(fs[0], "/")
		fsl2 := len(fs2)
		fsl2--
		Tracks = fs2[fsl2]
		Artists = strings.Replace(fs[1], "_", " ", -1)
		Albums = strings.Replace(fs[2], "_", " ", -1)
		Songs = strings.Replace(fs[3], "_", " ", -1)
	case fsl == 3:
		fs2 := strings.Split(fs[0], "/")
		fsl2 := len(fs2)
		fsl2--
		Tracks = fs2[fsl2]
		Artists = strings.Replace(fs[1], "_", " ", -1)
		Albums = strings.Replace(fs[1], "_", " ", -1)
		Songs = strings.Replace(fs[2], "_", " ", -1)
	}
	ASS := map[string]string{
		"Track":  Tracks,
		"Artist": Artists,
		"Album":  Albums,
		"Song":   Songs,
	}
	return ASS
}

//TagsFromFilenames exported
func TagsFromFilenames(OGGTagMetachan chan<- map[string]string, file string) {
	TASS := nameSplit(file)
	u := make([]byte, 16)
	rand.Read(u)
	u[8] = 0x80
	u[4] = 0x40
	uuid := hex.EncodeToString(u)
	ltn, err := os.Open(file)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer ltn.Close()
	ltnInfo, _ := ltn.Stat()
	var Lsize int64 = ltnInfo.Size()
	fsize := strconv.FormatInt(Lsize, 10)
	lbuf := make([]byte, Lsize)
	ltnReader := bufio.NewReader(ltn)
	ltnReader.Read(lbuf)
	h := sha512.New()
	h.Write([]byte(lbuf))
	mp3hash := hex.EncodeToString(h.Sum(nil))
	tagmap := map[string]string{
		"Filename": file,
		"FileId":   uuid,
		"Filesize": fsize,
		"MP3Hash":  mp3hash,
		"Artist":   TASS["Artist"],
		"Album":    TASS["Album"],
		"Song":     TASS["Song"],
		"Genre":    "ogg",
		"Track":    TASS["Track"],
	}
	OGGTagMetachan <- tagmap
}

//OGGInsertMainDB exported
func OGGInsertMainDB(OGGTagMetachan <-chan map[string]string, Tartistchan chan<- map[string]string,
	Talbumchan chan<- map[string]string, fileN string) {
	sesCopy := DBcon()
	defer sesCopy.Close()
	JPGc := sesCopy.DB("goimggo").C("imgdb")
	NAPc := sesCopy.DB("goimggo").C("goimgdb")
	MAINc := sesCopy.DB("goampgo").C("maindb")
	TM := <-OGGTagMetachan
	EXT := path.Ext(fileN)
	dp, _ := path.Split(fileN)
	var ArtTS map[string]string = make(map[string]string)
	err := JPGc.Find(bson.M{"Dirpath": dp}).One(&ArtTS)
	if err != nil {
		var NoArtS map[string]string
		err := NAPc.Find(nil).One(&NoArtS)
		if err != nil {
			fmt.Printf("this is NAPc err \n %v %v \n", err, dp)
		}
		mainI := map[string]string{
			"Dirpath":   dp,
			"Filename":  TM["Filename"],
			"FileId":    TM["FileId"],
			"Extension": EXT,
			"Filesize":  TM["Filesize"],
			"MP3Hash":   TM["MP3Hash"],
			"Artist":    TM["Artist"],
			"ArtistId":  "none",
			"Album":     TM["Album"],
			"AlbumId":   "none",
			"Title":     TM["Song"],
			"Genre":     TM["Genre"],
			"ImgInfo":   NoArtS["ImgId"],
			"LFilesize": NoArtS["LFilesize"],
			"SFilesize": NoArtS["SFilesize"],
			"Page":      "0",
			"Track":     TM["Track"],
		}
		//fmt.Printf("\n this is maini %v \n", mainI)
		MAINc.Insert(mainI)
	} else {
		mainI := map[string]string{
			"Dirpath":   dp,
			"Filename":  TM["Filename"],
			"FileId":    TM["FileId"],
			"Extension": EXT,
			"Filesize":  TM["Filesize"],
			"MP3Hash":   TM["MP3Hash"],
			"Artist":    TM["Artist"],
			"ArtistId":  "none",
			"Album":     TM["Album"],
			"AlbumId":   "none",
			"Title":     TM["Song"],
			"Genre":     TM["Genre"],
			"ImgInfo":   ArtTS["ImgId"],
			"LFilesize": ArtTS["LFilesize"],
			"SFilesize": ArtTS["SFilesize"],
			"Page":      "0",
			"Track":     TM["Track"],
		}
		//fmt.Printf("\n this is maini %v \n", mainI)
		MAINc.Insert(mainI)
	}
	Tmpartistmap := map[string]string{"DirPath": dp, "Artist": TM["Artist"]}
	Tartistchan <- Tmpartistmap
	Tmpalbummap := map[string]string{"DirPath": dp, "Album": TM["Album"]}
	Talbumchan <- Tmpalbummap
}
