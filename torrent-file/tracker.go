package torrentfile

import(
	"github.com/jackpal/bencode-go"
	"net/http"
	"net/url"
	"strconv"
	"time"
	"storrent/peers"
)

type trackresponse struct {
	Interval int `bencode:"interval"`
	Peers []string `bencode:"peers"`
}

func (t* torrentfile) buildurl (peerId [20]byte, port uint64) (string, error) {
	base, err := url.Parse(t.url)
	if err!= nil {
        return "", err
    }
	parameters := url.Values{
		"info_hash": []string{string(t.infohash[:])},
		"peer_id":   []string{string(peerId[:])},
		"port":      []string{strconv.Itoa(int(port))},
		"uploaded":  []string{"0"},
		"downloaded":[]string{"0"},
		"compact":   []string{"1"},
		"left":     []string{strconv.Itoa(t.length)},
	}
	base.RawQuery = parameters.Encode()
	return base.String(), nil
}
