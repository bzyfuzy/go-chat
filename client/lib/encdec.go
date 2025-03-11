package lib

import (
	"bufio"
	"bytes"
	"compress/gzip"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/pion/webrtc/v3"
)

func Encode(obj interface{}) (string, error) {
	b, err := json.Marshal(obj)
	if err != nil {
		return "", err
	}
	var data bytes.Buffer
	gz, err := gzip.NewWriterLevel(&data, gzip.BestCompression)
	if err != nil {
		panic(err)
	}
	if _, err := gz.Write(b); err != nil {
		panic(err)
	}
	if err := gz.Close(); err != nil {
		panic(err)
	}
	sdp := base64.StdEncoding.EncodeToString(data.Bytes())
	return sdp, nil
}

func Decode(in string, obj interface{}) error {
	b, err := base64.StdEncoding.DecodeString(in)
	if err != nil {
		return err
	}

	br := bytes.NewReader(b)
	gz, err := gzip.NewReader(br)
	if err != nil {
		panic(err)
	}
	data, err := io.ReadAll(gz)
	if err != nil {
		panic(err)
	}
	if err := gz.Close(); err != nil {
		panic(err)
	}

	return json.Unmarshal(data, obj)
}

func ReadSDP() (string, error) {
	r := bufio.NewReader(os.Stdin)
	var in string
	for {
		var err error
		in, err = r.ReadString('\n')
		if err != io.EOF {
			if err != nil {
				return "", err
			}
		}
		in = strings.TrimSpace(in)
		if len(in) > 0 {
			break
		}
	}

	fmt.Println("")
	return in, nil
}

func SDPPrompt() (webrtc.SessionDescription, error) {
	fmt.Println("Paste the remote SDP: ")

	answer := webrtc.SessionDescription{}
	for {
		text, err := ReadSDP()
		if err != nil {
			return answer, err
		}
		sdp := text
		if err := Decode(sdp, &answer); err == nil {
			break
		}
		fmt.Println("Invalid SDP. Enter again.")
	}
	return answer, nil
}
