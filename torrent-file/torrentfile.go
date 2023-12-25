package torrentfile

import (
	"bytes"
	"crypto/sha1"
	"fmt"
	"os"
	"github.com/jackpal/bencode-go"
)

//bInfo is the information dictionary in the torrent file, this has all the metadata
type bInfo struct{
	pieces	string  // this is a base-64 encoded string of all piece hashes put together
	piecesize int //this is the size of each piece in bytes
	length	int  //total length of the file in bytes
	name string //file name
}

//bTorrent represents the top-level structure of the torrent file
type bTorrent struct {
	url string //URL of the tracker
	info 	bInfo //information dictionary that has all the metadata
}

//torrentfile (self explainatory), holds the parsed information of the torrent file
type torrentfile struct {
	url			string //The tracker url
	infohash	[20]byte //a hash of the bencoded information dictionary
	piecehash	[][20]byte //an array of 20-byte hashes for each piece
	piecesize	int //piece size in bytes
	length		int //total file length in bytes
	name		string // file name
}

//Opens a torrent file from a given path and parses it's content
func Open(path string) (torrentfile, error){
	file, err := os.Open(path) //opens the torrent file
	if err != nil{
		return torrentfile{}, err //return an empty torrentfile and the error if opening fails
	}
	defer file.Close() //making sure that the file is closed

	var bt bTorrent //creating an empty bTorrent struct to hold the data
		
	err = bencode.Unmarshal(file, bt) //attempting to decode the torrent using bencode
	if err != nil {
		return torrentfile{}, err // return an empty torrentfile and the error if decoding fails
	}
	return torrentfile{},nil //a placeholder until a "totorrentfile" function is implemented
}

//hash calculates the SHA1 hash of the bInfo struct
func (i *bInfo) hash() ([20]byte, error){
	var buf bytes.Buffer //a buffer is created to hold the bencoded representation of the struct
	err := bencode.Marshal(&buf, *i) //the struct is encoded using bencode
	if err != nil{
		return [20]byte{},err //return empty byte array and error if encoding fails.
	}
	h := sha1.Sum(buf.Bytes()) //calculate the SHA-1 hash of the encoded data
	return h, nil //Return the calculated SHA-1 hash
}
