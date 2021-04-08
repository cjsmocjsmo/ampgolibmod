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
	// "sync"
	"strconv"
	// "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

func getFS() string {
	sesCopy := DBcon()
	defer sesCopy.Close()
	AMPc := sesCopy.DB("maindb").C("maindb")
	var F1S []map[string]string
	err := AMPc.Find(nil).Select(bson.M{"filesize": 1}).All(&F1S)
	if err != nil {
		fmt.Println("fisi fucked up")
		fmt.Println(err)
	}
	fstotal := 0
	for _, v := range F1S {
		s := v["filesize"]
		if s, err := strconv.Atoi(s); err == nil {
			fstotal += s
		}
	}
	fsTotal := calcBytes(int64(fstotal))
	return fsTotal
}

// func getLFilesize() string {
// 	sesCopy := DBcon()
// 	defer sesCopy.Close()
// 	AMPc := sesCopy.DB("maindb").C("maindb")
// 	var F2S []map[string]string
// 	err := AMPc.Find(nil).Select(bson.M{"Filesize":1}).All(&F2S)
// 	if err != nil {
// 		fmt.Println("lfs fucked up")
// 	}
// 	limgtotal := 0
// 	var wg sync.WaitGroup
// 	for _, x := range F2S{
// 		wg.Add(1)
// 		go func () {
// 			m := x["Filesize"]
// 			if s, err := strconv.Atoi(m); err == nil{
// 				limgtotal += s
// 			}
// 			wg.Done()
// 		}()
// 		wg.Wait()
// 	}
// 	limgTotal := calcBytes(int64(limgtotal))
// 	return limgTotal
// }

//ArtistCount exported
func ArtistCount() (ArtCount string) {
	sesCopy := DBcon()
	defer sesCopy.Close()
	AMPc := sesCopy.DB("artistview").C("artistviews")
	n, err := AMPc.Find(nil).Count()
	CheckError(err, "ArtistCount: artist count db search has fucked up")
	ArtCount = strconv.Itoa(n)
	return
}

//AlbumCount exported
func AlbumCount() (AlbCount string) {
	sesCopy := DBcon()
	defer sesCopy.Close()
	AMPc := sesCopy.DB("albview").C("albview")
	n, err := AMPc.Find(nil).Count()
	CheckError(err, "AlbumCount: Albumcount has fucked up")
	AlbCount = strconv.Itoa(n)
	return
}

//TitleCount exported
func TitleCount() (TitleCount string) {
	sesCopy := DBcon()
	defer sesCopy.Close()
	AMPc := sesCopy.DB("maindb").C("maindb")
	n, err := AMPc.Find(nil).Count()
	CheckError(err, "TitleCount: Titlecount has fucked up")
	TitleCount = strconv.Itoa(n)
	return
}

func calcBytes(fb int64) string {
	if fb >= 1099511627776 {
		tera := fb / 1099511627776
		tb1 := strconv.FormatInt(tera, 10)
		tb2 := tb1 + " TB"
		return tb2
	} else if fb >= 1073741824 {
		giga := fb / 1073741824
		gb1 := strconv.FormatInt(giga, 10)
		gb2 := gb1 + " GB"
		return gb2
	} else if fb >= 1048576 {
		mega := fb / 1048576
		mb1 := strconv.FormatInt(mega, 10)
		mb2 := mb1 + " MB"
		return mb2
	} else if fb >= 1024 {
		kilo := fb / 1024
		kb1 := strconv.FormatInt(kilo, 10)
		kb2 := kb1 + " KB"
		return kb2
	} else {
		foo := strconv.FormatInt(fb, 10)
		bt := foo + " bytes"
		return bt
	}
}

//FileSizeF exported
func FileSizeF() {
	FS := getFS()
	artc := ArtistCount()
	albc := AlbumCount()
	titc := TitleCount()
	Totals := map[string]string{"ArtistCount": artc,
		"AlbumCount":    albc,
		"TitleCount":    titc,
		"FileSizeTotal": FS,
	}
	sesCopy := DBcon()
	defer sesCopy.Close()
	DBSTATSc := sesCopy.DB("goampgo").C("dbstats")
	DBSTATSc.Insert(Totals)
	fmt.Printf("Size on disk MP3's %v \n", FS)
	fmt.Printf("ArtCount %v \n", artc)
	fmt.Printf("AlbCount %v \n", albc)
	fmt.Printf("TitleCount %v \n", titc)
}
