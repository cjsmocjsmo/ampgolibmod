////////////////////////////////////////////////////////////////////////////////
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
	"os"
	"path/filepath"
	// "io"
	"fmt"
	"bufio"
	"encoding/base64"
	"github.com/disintegration/imaging"
	"strconv"
	"strings"
)

func resizeImage(infile string, outfile string) {
	pic, err := imaging.Open(infile)
	if err != nil {
		fmt.Println(infile)
		fmt.Println("this is file Open error noartthumb")
		print(err)
		// os.Remove(infile)
		
	}
	// CheckError(err, "\n resizeImage: this is file Open error noartthumb \n")
	sjImage := imaging.Resize(pic, 100, 0, imaging.Lanczos)
	err = imaging.Save(sjImage, outfile)
	CheckError(err, "resizeImage: image save has fucked up")
}

func createB64Image(img string) (b64i string) {
	ltn, err := os.Open(img)
	CheckError(err, "createB64Image: image open has fucked up")
	defer ltn.Close()
	ltnInfo, _ := ltn.Stat()
	var Lsize int64 = ltnInfo.Size()
	lbuf := make([]byte, Lsize)
	ltnReader := bufio.NewReader(ltn)
	ltnReader.Read(lbuf)
	lb64img := base64.StdEncoding.EncodeToString(lbuf)
	b64i = "data:image/png;base64," + lb64img
	return
}

func gSize(ppath string) (Size string) {
	stn, err := os.Open(ppath)
	CheckError(err, "gSize: file open has fucked up")
	defer stn.Close()
	stnInfo, _ := stn.Stat()
	var Siz int64 = stnInfo.Size()
	Size = strconv.FormatInt(Siz, 10)
	return
}

func splitpath(path string) (dir string, artist string, album string, ext string) {
	ext = filepath.Ext(path)
	dir = filepath.Dir(path)
	sptext := strings.Split(path, ".")
	dirp := strings.Split(sptext[0], "/")
	blob := strings.Split(dirp[4], "_-_")
	artist = blob[0]
	album = blob[1]
	return
}

//Gopher exported
type Gopher struct {
	Dirpath  string `bson:"dirpath"`
	Filename string `bson:"filename"`
	ImgID    string `bson:"imgID"`
	Size     string `bson:"size"`
	B64Image string `bson:"b64Image"`
	Ext      string `bson:"ext"`
}

//UnknownJpg exported
func UnknownJpg(img string) {
	tdir := os.Getenv("AMPGO_TEMP_DIRPATH") // /root/fsData/tmp
	t, _ := UUID()
	outfile := tdir + "/" + t + ".jpg"
	resizeImage(img, outfile)
	b64out := createB64Image(outfile)
	var gopher Gopher
	gopher.Dirpath = filepath.Dir(img)
	gopher.Filename = img
	gopher.ImgID, _ = UUID()
	gopher.Size = gSize(img)
	gopher.B64Image = b64out
	gopher.Ext = filepath.Ext(img)
	go func() {
		ses := DBcon()
		defer ses.Close()
		uimg := ses.DB("unknownjpg").C("meta")
		uimg.Insert(gopher)
	}()
	os.Remove(outfile)
}

//Gofer exported
type Gofer struct {
	Dirpath  string `bson:"dirpath"`
	Filename string `bson:"filename"`
	ImgID    string `bson:"imgID"`
	Size     string `bson:"size"`
	B64Image string `bson:"b64Image"`
	Ext      string `bson:"ext"`
	Artist   string `bson:"artist"`
	Album    string `bson:"album"`
	Title    string `bson:"title"`
}

//Thumbnails exported
func Thumbnails(img string) Gofer {
	tmpdir := os.Getenv("AMPGO_TEMP_DIRPATH") // /root/fsData/tmp
	t, _ := UUID()
	outfile := tmpdir + "/" + t + ".jpg"
	resizeImage(img, outfile)
	b64out := createB64Image(outfile)
	var gopher Gofer
	dir, artist, album, ext := splitpath(img)
	gopher.Dirpath = dir
	gopher.Filename = img
	gopher.ImgID, _ = UUID()
	gopher.Size = gSize(img)
	gopher.B64Image = b64out
	gopher.Ext = ext
	gopher.Artist = artist
	gopher.Album = album

	ses := DBcon()
	defer ses.Close()
	uimg := ses.DB("thumbnails").C("meta")
	uimg.Insert(gopher)

	go func(outfile string) {
		os.Remove(outfile)
	}(outfile)

	return gopher
}
