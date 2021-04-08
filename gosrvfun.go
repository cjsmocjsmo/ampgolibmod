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
	"log"
	"math/rand"
	"os"
	"path"
	"strconv"
	"strings"
	// "time"
	// "crypto/sha512"
	// "encoding/hex"
	// "gopkg.in/mgo.v2"
	// "gopkg.in/mgo.v2/bson"
	"github.com/globalsign/mgo"
	"github.com/globalsign/mgo/bson"
	// mathrand "math/rand"
)

//Plist exported
type Plist struct {
	PLName string              `bson:"PLName"`
	PLId   string              `bson:"PLId"`
	Songs  []map[string]string `bson:"Songs"`
}

//Imgfa exported
type Imgfa struct {
	Album   string
	HSImage string
	Songs   []map[string]string
}

//Ralbinfo exported
type Ralbinfo struct {
	Songs   []map[string]string `bson:"songs"`
	HSImage string
}

//Voodoo exported
type Voodoo struct {
	Playlists []map[string]string
}

//SFDBcon exported
func SFDBcon() *mgo.Session {
	s, err := mgo.Dial(os.Getenv("AMP_AMPDB_ADDR"))
	if err != nil {
		log.Println("Session creation dial error")
		log.Println(err)
	}
	log.Println("Session Connection to db established")
	return s
}

//OffSet exported
// func OffSet() (OffSet int) {
// 	OfSt := os.Getenv("AMPGO_OFFSET")
// 	OffSet, _ = strconv.Atoi(OfSt)
// 	return
// }

//GetInitArtistInfo exported
func GetInitArtistInfo() (ArtVieW []ArtVIEW) {
	ofset := OffSet
	ses := SFDBcon()
	defer ses.Close()
	AMPc := ses.DB("artistview").C("artistviews")
	b1 := bson.M{"_id": 0}
	err := AMPc.Find(nil).Select(b1).Sort("artist").Limit(ofset).All(&ArtVieW)
	if err != nil {
		log.Println("find one has failed")
		log.Println(err)
	}
	log.Println("GArtView is complete")
	return
}

//GetInitAlbumInfo exported
func GetInitAlbumInfo() (ALBView []AlbvieW) {
	ofset := OffSet
	ses := SFDBcon()
	defer ses.Close()
	ALBc := ses.DB("albview").C("albview")
	b1 := bson.M{"_id": 0}
	err := ALBc.Find(nil).Select(b1).Sort("album").Limit(ofset).All(&ALBView)
	if err != nil {
		log.Println("initial album info has fucked up")
		log.Println(err)
	}
	log.Println("GInitialAlbumInfo is complete")
	return
}

//GetInitSongInfo exported
func GetInitSongInfo() (TView []map[string]string) {
	ofset := OffSet
	ses := SFDBcon()
	defer ses.Close()
	MAINc := ses.DB("maindb").C("maindb")
	b1 := bson.M{"_id": 0, "artist": 1, "title": 1, "fileID": 1}
	err := MAINc.Find(nil).Select(b1).Sort("Title").Limit(ofset).All(&TView)
	if err != nil {
		log.Println("intial song info fucked up")
		log.Println(err)
	}
	log.Println("GInitialSongInfo is complete")
	return
}

//GImageSongForAlbum exported
func GImageSongForAlbum(albid string) (IMGfa Imgfa) {
	ses := SFDBcon()
	defer ses.Close()
	ALBc := ses.DB("albview").C("albview")
	b1 := bson.M{"albumID": albid}
	b2 := bson.M{"_id": 0, "album": 1, "songs": 1, "hsimage": 1}
	err := ALBc.Find(b1).Select(b2).One(&IMGfa)
	if err != nil {
		log.Println("gimage song for album fucked up")
		log.Println(err)
	}
	return
}

//GStats exported
func GStats() (STAt map[string]string) {
	ses := SFDBcon()
	defer ses.Close()
	STATc := ses.DB("goampgo").C("dbstats")
	b1 := bson.M{"_id": 0}
	err := STATc.Find(nil).Select(b1).One(&STAt)
	if err != nil {
		log.Println("stats has fucked up")
		log.Println(err)
	}
	log.Println("GStats is complete")
	return
}

//ArtistAlpha exported
func ArtistAlpha() (ARDist []string) {
	ses := SFDBcon()
	defer ses.Close()
	ARTVc := ses.DB("artistview").C("artistviews")
	err := ARTVc.Find(nil).Distinct("page", &ARDist)
	if err != nil {
		log.Println("artist alpha has fucked up")
		log.Println(err)
	}
	log.Println("ArtistAlpha is complete")
	return
}

// AlbumAlpha exported
func AlbumAlpha() (ALDist []string) {
	ses := SFDBcon()
	defer ses.Close()
	ALBVc := ses.DB("albview").C("albview")
	err := ALBVc.Find(nil).Distinct("page", &ALDist)
	if err != nil {
		log.Println("album alpha fucked up")
		log.Println(err)
	}
	log.Println("AlbumAlpha is complete")
	return
}

//TitleAlpha exported
func TitleAlpha() (TDist []string) {
	ses := SFDBcon()
	defer ses.Close()
	MAINc := ses.DB("maindb").C("maindb")
	err := MAINc.Find(nil).Distinct("page", &TDist)
	if err != nil {
		log.Println("title alpha fucked up")
		log.Println(err)
	}
	log.Println("TitleAlpha is complete")
	return
}

// func getPicIndex(albl []map[string]string) (NumPics []int) {
// 	for i := range albl {
// 		NumPics = append(NumPics, i)
// 	}
// 	return
// }

// func shuffleList(num []int) []int {
// 	dest := make([]int, len(num))
// 	perm := rand.Perm(len(num))
// 	for i, v := range perm {
// 		dest[v] = num[i]
// 	}
// 	return dest
// }

// func frandom(min, max int) int {
// 	rand.Seed(time.Now().UTC().Unix())
// 	return rand.Intn(max-min) + min
// }

//Rai exported
// var Rai Ralbinfo

//RAI exported
// var RAI []Ralbinfo
// func getAlbInfo(sll []int, rp []map[string]string) (RAI []Ralbinfo) {
// 	RAI = nil
// 	Lsll := len(sll)
// 	myrand := frandom(0, Lsll)
// 	rprand := rp[myrand]
// 	ses := SFDBcon()
// 	defer ses.Close()
// 	MAINc := ses.DB("goampgo").C("albumviews")
// 	for _, z := range rprand {
// 		b1 := bson.M{"albumiD": z}
// 		b2 := bson.M{"_id": 0, "songs": 1, "hsimage": 1}
// 		MAINc.Find(b1).Select(b2).One(&Rai)
// 		RAI = append(RAI, Rai)
// 	}
// 	return
// }

//RPics exported
// func RPics() []Ralbinfo {
// 	ses := SFDBcon()
// 	defer ses.Close()
// 	RANc := ses.DB("goampgo").C("randompics")
// 	var RP []map[string]string
// 	RP = nil
// 	err := RANc.Find(nil).Select(bson.M{"_id": 0}).All(&RP)
// 	if err != nil {
// 		fmt.Println("RP has fucked up")
// 		fmt.Println(err)
// 	}
// 	PI := getPicIndex(RP)
// 	SL1 := shuffleList(PI)
// 	SL2 := shuffleList(SL1)
// 	randalb := getAlbInfo(SL2, RP)
// 	log.Println("RPics is complete")
// 	return randalb
// }

//RamdomAlbPicPlay exported
func RamdomAlbPicPlay(fid string) map[string]string {
	ses := SFDBcon()
	defer ses.Close()
	MAINc := ses.DB("goampgo").C("maindb")
	b1 := bson.M{"FileId": fid}
	b2 := bson.M{"_id": 0, "Dirpath": 1, "Filename": 1, "ImgInfo": 1, "Album": 1, "Title": 1}
	var RAPP map[string]string
	err := MAINc.Find(b1).Select(b2).One(&RAPP)
	if err != nil {
		fmt.Println("RAPP has fucked up")
	}
	CLIAc := ses.DB("goampgo").C("CliaArgs")
	b3 := bson.M{"_id": 0, "mediapath": 1, "httppath": 1, "catname": 1}
	var CLI map[string]string
	err = CLIAc.Find(nil).Select(b3).One(&CLI)
	if err != nil {
		fmt.Println(err)
	}
	zoo := strings.Split(RAPP["Dirpath"], CLI["mediapath"])
	_, boo1 := path.Split(RAPP["Filename"])
	HTTPMPath := "/" + CLI["catname"] + zoo[1] + boo1
	var IInfo map[string]string
	IMGc := ses.DB("goimggo").C("imgdb")
	b5 := bson.M{"ImgId": RAPP["ImgInfo"]}
	b6 := bson.M{"_id": 0}
	err = IMGc.Find(b5).Select(b6).One(&IInfo)
	if err != nil {
		NAPc := ses.DB("goimggo").C("goimgdb")
		err = NAPc.Find(b5).Select(b6).One(&IInfo)
	}
	Result := map[string]string{"SongId": fid,
		"HLImage":   IInfo["LB64Image"],
		"MusicPath": HTTPMPath,
		"Album":     RAPP["Album"],
		"Title":     RAPP["Title"],
	}
	return Result
}

//PathArt exported
// func PathArt(fid string) map[string]string {
// 	return RamdomAlbPicPlay(fid)
// }

//SongInfo exported
func SongInfo(pagenum string) (SIS []map[string]string) {
	ses := SFDBcon()
	defer ses.Close()
	MAINc := ses.DB("maindb").C("maindb")
	b1 := bson.M{"page": pagenum}
	b2 := bson.M{"_id": 0, "title": 1, "ileID": 1, "artist": 1}
	err := MAINc.Find(b1).Select(b2).All(&SIS)
	if err != nil {
		log.Println("song info has fucked up")
		log.Println(err)
	}
	log.Println("SongInfo is complete")
	return
}

//AlbumInfo exported
func AlbumInfo(pagenum string) (AI []AlbvieW) {
	ses := SFDBcon()
	defer ses.Close()
	ALBVc := ses.DB("albview").C("albview")
	b1 := bson.M{"page": pagenum}
	b2 := bson.M{"_id": 0, "artist": 1, "artistID": 1, "album": 1, "albumID": 1, "hsimage": 1, "songs": 1, "numsongs": 1}
	err := ALBVc.Find(b1).Select(b2).All(&AI)
	if err != nil {
		log.Println("AlbumInfo has fucked up")
		log.Println(err)
	}
	log.Println("AlbumInfo is complete")
	return
}

//ArtistInfo exported
func ArtistInfo(pagenum string) (ARTI []ArtVIEW) {
	ses := SFDBcon()
	defer ses.Close()
	ARTc := ses.DB("artistview").C("artistviews")
	b1 := bson.M{"page": pagenum}
	b2 := bson.M{"_id": 0, "artist": 1, "artistID": 1, "albums": 1, "page": 1}
	err := ARTc.Find(b1).Select(b2).All(&ARTI)
	if err != nil {
		log.Println("ArtistInfo has fucked up")
		log.Println(err)
	}
	log.Println("ArtistInfo is complete")
	return
}

//SongSearch exported
// func SongSearch(searchv string) []map[string]string {
// 	ses := SFDBcon()
// 	defer ses.Close()
// 	MAINc := ses.DB("goampgo").C("maindb")
// 	var SOS []map[string]string
// 	b1 := bson.M{"$text": bson.M{"$search": searchv}}
// 	b2 := bson.M{"_id": 0, "Artist": 1, "FileId": 1, "Title": 1}
// 	err := MAINc.Find(b1).Select(b2).All(&SOS)
// 	if err != nil {
// 		log.Println("artsearch has fucked up")
// 		log.Println(err)
// 	}
// 	log.Println("SongSearch is complete")
// 	return SOS
// }

//AlbumSearch exported
// func AlbumSearch(searchv string) []AlbvieW{
// 	ses := SFDBcon()
// 	defer ses.Close()
// 	ALBVc := ses.DB("goampgo").C("albumviews")
// 	b1 := bson.M{"$text": bson.M{"$search": searchv}}
// 	b2 := bson.M{"_id": 0}
// 	var ALBS []AlbvieW
// 	err := ALBVc.Find(b1).Select(b2).All(&ALBS)
// 	if err != nil {
// 		log.Println("AlbumText Search has fucked up")
// 		log.Println(err)
// 	}
// 	log.Println("AlbumSearch is compete")
// 	return ALBS
// }

//ArtistSearch exported
// func ArtistSearch(searchv string) []ArtVIEW{
// 	var ARTS []ArtVIEW
// 	ses := SFDBcon()
// 	defer ses.Close()
// 	ARTVc := ses.DB("goampgo").C("artistviews")
// 	b1 := bson.M{"$text": bson.M{"$search": searchv}}
// 	b2 := bson.M{"_id": 0}
// 	err := ARTVc.Find(b1).Select(b2).All(&ARTS)
// 	if err != nil {
// 		log.Println("ArtistText Search has fucked up")
// 		log.Println(err)
// 	}
// 	log.Println("ArtistSearch is complete")
// 	return ARTS
// }

//PlaylistCheck exported
func PlaylistCheck() int {
	ses := SFDBcon()
	defer ses.Close()
	result := 0
	pldb := "playlistsdb"
	coll, _ := ses.DB("goampgo").CollectionNames()
	for _, col := range coll {
		comp := strings.EqualFold(col, pldb)
		if comp == true {
			result = 1
		}
	}
	log.Println("PlaylistCheck is complete")
	return result
}

//AllPlayLists exported
func AllPlayLists() []Plist {
	var PLS []Plist
	ses := SFDBcon()
	defer ses.Close()
	PLSc := ses.DB("goampgo").C("playlistsdb")
	b1 := bson.M{"_id": 0}
	err := PLSc.Find(nil).Select(b1).All(&PLS)
	if err != nil {
		log.Println("Pls has fucked up")
		log.Println(err)
	}
	log.Println("AllPlayLists is complete")
	return PLS
}

//UUID exported
// func UUID() (UUID string) {
// 	aSeed := time.Now()
// 	aseed := aSeed.UnixNano()
// 	mathrand.Seed(aseed)
// 	u := mathrand.Int63n(aseed)
// 	p := strconv.FormatInt(u, 10)
// 	return p
// }

//AddPlayListNameToDB exported
func AddPlayListNameToDB(pln string) []Plist {
	var APlS Plist
	APlS.PLName = pln
	APlS.PLId, _ = UUID()
	APlS.Songs = nil
	ses := SFDBcon()
	defer ses.Close()
	PLSc := ses.DB("goampgo").C("playlistsdb")
	PLSc.Insert(APlS)
	AllPls := AllPlayLists()
	log.Println("AddPlayListNameToDB is complete")
	return AllPls
}

// AddSongsToPlistDB exported
func AddSongsToPlistDB(sn string, sid string, plid string) string {
	var ASPLS Plist
	ses := SFDBcon()
	defer ses.Close()
	MAINc := ses.DB("goampgo").C("maindb")
	b1 := bson.M{"FileId": sid}
	b2 := bson.M{"_id": 0}
	var SongInfo map[string]string
	SongInfo = make(map[string]string)
	err := MAINc.Find(b1).Select(b2).One(&SongInfo)
	if err != nil {
		log.Println("songinfo has fucked up")
		log.Println(err)
	}
	PLSc := ses.DB("goampgo").C("playlistsdb")
	b3 := bson.M{"PLId": plid}
	err = PLSc.Find(b3).Select(b2).One(&ASPLS)
	if err != nil {
		log.Println("ASPLS has fucked up")
		log.Println(err)
	}
	newMap := map[string]string{"Title": SongInfo["Title"], "FileId": SongInfo["FileId"]}
	ASPLS.Songs = append(ASPLS.Songs, newMap)
	err = PLSc.Update(b3, ASPLS)
	if err != nil {
		log.Println("PLSc update has fucked up")
		log.Println(err)
	}
	log.Println("AddSongsToPlistDB is complete")
	return "whoot"
}

//AllPlaylistSongsFromDB exported
func AllPlaylistSongsFromDB(plid string) Plist {
	var AA Plist
	ses := SFDBcon()
	defer ses.Close()
	PLSc := ses.DB("goampgo").C("playlistsdb")
	b1 := bson.M{"PLId": plid}
	b2 := bson.M{"_id": 0}
	err := PLSc.Find(b1).Select(b2).One(&AA)
	if err != nil {
		log.Println("AllPlaylistSongsFromDB has fucked up")
		log.Println(err)
	}
	log.Println("AllPlaylistSongsFromDB is complete")
	return AA
}

//CreatePlayerPlaylist exported
func CreatePlayerPlaylist(plid string) Voodoo {
	var AASPLS Plist
	ses := SFDBcon()
	defer ses.Close()
	CLIc := ses.DB("goampgo").C("CliaArgs")
	var clia map[string]string
	clia = make(map[string]string)
	err := CLIc.Find(nil).Select(bson.M{"_id": 0}).One(&clia)
	if err != nil {
		log.Println("clic has fucked up")
		log.Println(err)
	}
	mediapath := clia["mediapath"]
	catname := clia["catname"]
	PLSc := ses.DB("goampgo").C("playlistsdb")
	b1 := bson.M{"PLId": plid}
	b2 := bson.M{"_id": 0}
	err = PLSc.Find(b1).Select(b2).One(&AASPLS)
	if err != nil {
		log.Println("playlistdb error has fucked up")
		log.Println(err)
	}
	var apls map[string]string
	apls = make(map[string]string)
	var allpls []map[string]string
	for _, pl := range AASPLS.Songs {
		MAINc := ses.DB("goampgo").C("maindb")
		b3 := bson.M{"FileId": pl["FileId"]}
		b4 := bson.M{"_id": 0, "Title": 1, "Album": 1, "ImgInfo": 1, "Dirpath": 1, "Filename": 1}
		MAINc.Find(b3).Select(b4).One(&apls)
		ssplit := strings.Split(apls["Filename"], mediapath)
		psplit := strings.Split(ssplit[1], ".mp3")
		newpath := "/" + catname + psplit[0]
		var lb64 map[string]string = make(map[string]string)
		b5 := bson.M{"ImgId": apls["ImgInfo"]}
		b6 := bson.M{"_id": 0}
		IMGc := ses.DB("goimggo").C("imgdb")
		err = IMGc.Find(b5).Select(b6).One(&lb64)
		if err != nil {
			IMGcc := ses.DB("goimggo").C("goimgdb")
			IMGcc.Find(nil).Select(b6).One(&lb64)
		}
		result := map[string]string{"Title": apls["Title"], "Album": apls["Album"], "HLImage": lb64["LB64Image"], "PlayPath": newpath}
		allpls = append(allpls, result)
	}
	var VD Voodoo
	VD.Playlists = allpls
	log.Println("CreatePlayerPlaylist is complete")
	return VD
}

//DropCollection exported
func DropCollection(col string) {
	ses := SFDBcon()
	defer ses.Close()
	CTDc := ses.DB("goampgo").C(col)
	CTDc.DropCollection()
}

//GetSongList exported
func GetSongList() []map[string]string {
	ses := SFDBcon()
	defer ses.Close()
	MAINc := ses.DB("goampgo").C("maindb")
	var allids []map[string]string
	err := MAINc.Find(nil).Select(bson.M{"_id": 0, "FileId": 1}).All(&allids)
	if err != nil {
		fmt.Println(err)
	}
	return allids
}

//GetAlbumIndex exported
func GetAlbumIndex(sl []map[string]string) []int {
	var NumPics []int
	for i := range sl {
		NumPics = append(NumPics, i)
	}
	return NumPics
}

//ShuffleList exported
func ShuffleList(num []int) []int {
	dest := make([]int, len(num))
	perm := rand.Perm(len(num))
	for i, v := range perm {
		dest[v] = num[i]
	}
	return dest
}

//GetNewList exported
func GetNewList(d []int, a []map[string]string) []map[string]string {
	var NewList []map[string]string
	for _, v := range d {
		NewList = append(NewList, a[v])
	}
	return NewList
}

//ChunckIt exported
func ChunckIt(nl []map[string]string, plc string) {
	myplc, _ := strconv.Atoi(plc)
	var outSlice []string
	count := 0
	for _, v := range nl {
		count++
		if count == myplc {
			outSlice = append(outSlice, v["FileId"])
			ses := SFDBcon()
			defer ses.Close()
			RSONc := ses.DB("goampgo").C("randomsongs")
			RSONc.Insert(outSlice)
			count = 0
			outSlice = nil
		} else if count < myplc {
			outSlice = append(outSlice, v["FileId"])
		} else {
			fmt.Println("end of loop")
		}
	}
}

//AddRandomPlaylist exported
// func AddRandomPlaylist(pln string, plc string) []Plist {
// 	var RSONG Plist
// 	DropCollection("randomsongs")
// 	SL := GetSongList()
// 	AI := GetAlbumIndex(SL)
// 	DEST1 := ShuffleList(AI)
// 	DEST2 := ShuffleList(DEST1)
// 	NL := GetNewList(DEST2, SL)
// 	ChunckIt(NL, plc)
// 	ses := SFDBcon()
// 	defer ses.Close()
// 	RSONc := ses.DB("goampgo").C("randomsongs")
// 	var rlist []map[string]string
// 	err := RSONc.Find(nil).Select(bson.M{"_id": 0}).All(&rlist)
// 	if err != nil {
// 		log.Println("randomsongs has fucked up")
// 		log.Println(err)
// 	}
// 	n, _ := RSONc.Count()
// 	myrand := frandom(0, n)
// 	rset := rlist[myrand]
// 	var SongI []map[string]string
// 	for _, S := range rset {
// 		MAINc := ses.DB("goampgo").C("maindb")
// 		var sinfo map[string]string
// 		sinfo = make(map[string]string)
// 		b1 := bson.M{"FileId": S}
// 		b2 := bson.M{"_id": 0, "Artist": 1, "Album": 1, "Title": 1, "FileId": 1}
// 		MAINc.Find(b1).Select(b2).One(&sinfo)
// 		SongI = append(SongI, sinfo)
// 	}
// 	RSONG.PLName = pln
// 	RSONG.PLId, _ = UUID()
// 	RSONG.Songs = SongI
// 	PLDBc := ses.DB("goampgo").C("playlistsdb")
// 	PLDBc.Insert(RSONG)
// 	APL := AllPlayLists()
// 	log.Println("AddRandomPlaylist is complete")
// 	return APL
// }

//DeletePlaylistFromDB exported
func DeletePlaylistFromDB(plid string) []map[string]string {
	ses := SFDBcon()
	defer ses.Close()
	PLDBc := ses.DB("goampgo").C("playlistsdb")
	err := PLDBc.Remove(bson.M{"PLId": plid})
	if err != nil {
		log.Printf("there was a problem removing %v playlist", plid)
		log.Println(err)
	}
	var Allplist []map[string]string
	b1 := bson.M{"_id": 0, "PLName": 1, "PLId": 1}
	err = PLDBc.Find(nil).Select(b1).All(&Allplist)
	if err != nil {
		fmt.Println(err)
	}
	log.Println("DeletePlaylistFromDB is complete")
	return Allplist
}

//DeleteSongFromPlaylist exported
func DeleteSongFromPlaylist(plname string, songid string) Plist {
	ses := SFDBcon()
	defer ses.Close()
	PLDBc := ses.DB("goampgo").C("playlistsdb")
	b1 := bson.M{"PLName": plname}
	b2 := bson.M{"_id": 0}
	var PLs Plist
	err := PLDBc.Find(b1).Select(b2).One(&PLs)
	if err != nil {
		log.Println("playlistsdb has fucked up")
		log.Println(err)
	}
	delindex := 0
	for I, V := range PLs.Songs {
		if V["FileId"] == songid {
			delindex = I
		}
	}
	PLs.Songs = append(PLs.Songs[:delindex], PLs.Songs[delindex+1:]...)
	var NPLs Plist
	NPLs.PLName = PLs.PLName
	NPLs.PLId = PLs.PLId
	NPLs.Songs = PLs.Songs
	b3 := bson.M{"PLName": PLs.PLName, "PLId": PLs.PLId, "Songs": PLs.Songs}
	err = PLDBc.Update(b1, b3)
	if err != nil {
		log.Println("del song from playlist fucked up")
		log.Println(err)
	}
	log.Println("DeleteSongFromPlaylist is complete")
	return NPLs
}

//GetUserFromDB exported
// func GetUserFromDB(un string, pw string) map[string]string {
// 	h := sha512.New()
// 	h.Write([]byte(pw))
// 	sha512hex := hex.EncodeToString(h.Sum(nil))
// 	ses := SFDBcon()
// 	defer ses.Close()
// 	USERc := ses.DB("goampgo").C("usercreds")
// 	UInfo := make(map[string]string)
// 	b1 := bson.M{"Username": un, "Password": sha512hex}
// 	b2 := bson.M{"_id":0}
// 	err := USERc.Find(b1).Select(b2).One(&UInfo)
// 	if err != nil {
// 		log.Println("user does not exist")
// 		log.Println(err)
// 	}
// 	return UInfo
// }
