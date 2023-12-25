package torrentfile

import(
	"fmt"
	"bytes"
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
	defer file.Close()

	bt := bTorrent{}
	err := bencode.Unmarshal(file, &bt)
	if err != nil {
		return torrentfile{}, err
	}

}
