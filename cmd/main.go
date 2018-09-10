package main

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"github.com/hoozecn/angen"
	"github.com/skratchdot/open-golang/open"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"
)

// Generated from templates/index.html
var _templatesIndexHtml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xcc\x58\xcf\x6f\xdc\xc4\x17\xbf\xe7\xaf\x78\x9d\xf6\x9b\x78\xb3\xfe\xb5\xfb\x45\xa8\xea\xda\x8b\x50\x11\x52\x85\x28\x87\xb6\x12\x52\x55\xa5\x5e\xef\xec\xee\xa4\xf6\x8c\x63\x8f\x37\x09\xed\x1e\x90\x90\x0a\x42\x55\x11\xa2\x8a\x7a\x42\x70\x29\x1c\xda\x72\x00\x2e\xa8\xff\x0d\x49\x9a\xff\x02\xcd\x8c\xed\xb5\xbd\xeb\x2c\x6d\x5a\x60\x2f\x19\xcf\xbc\xf7\x99\xcf\x7b\xef\x33\x6f\x1c\x3b\xe7\x86\xcc\xe7\xfb\x11\x86\x09\x0f\x83\xfe\x9a\x23\xfe\x40\xe0\xd1\xb1\x8b\x3e\x9b\x18\x3e\x45\x62\x0e\x7b\xc3\xfe\x1a\x00\x80\x13\x62\xee\x81\x3f\xf1\xe2\x04\x73\x17\xdd\xb8\xfe\xa1\x71\x11\x95\x97\xa8\x17\x62\x17\x4d\x09\xde\x8d\x58\xcc\x91\x5c\x51\x3f\x9f\x51\x8e\x29\x77\xd1\x2e\x19\xf2\x89\x3b\xc4\x53\xe2\x63\x43\x3e\xe8\x90\x26\x38\x36\x12\xdf\x0b\xbc\x41\x80\x5d\xca\x74\x20\x94\x70\xe2\x05\x72\x12\xbb\x1d\xd3\xd6\x21\xf4\xf6\x48\x98\x86\x95\x29\x42\xab\x53\x15\x32\x13\xce\x23\x03\xef\xa4\x64\xea\xa2\x4f\x8d\x1b\xef\x1b\x97\x59\x18\x79\x9c\x0c\x02\x8c\xe6\x7c\x08\x76\xf1\x70\x8c\x73\x4f\x4e\x78\x80\xfb\xc7\xdf\x7d\x7f\xf4\xe5\x37\x87\x8f\x7f\x82\x69\xc7\xb4\x1d\x4b\xcd\x2a\x8b\x80\xd0\x3b\x30\x89\xf1\xc8\x45\x62\x87\xe4\x92\x65\xf9\x43\x6a\x0e\x18\xe3\x7e\x92\x98\x3e\x0b\x2d\x31\x4e\x78\xec\x45\xd6\x3b\x66\xc7\xec\x58\x7e\x92\xcc\xe7\xcc\x90\x50\xd3\x4f\x12\x04\x31\x0e\x5c\x94\xf0\xfd\x00\x27\x13\x8c\xb9\xc8\xb5\xa5\x92\xed\x0c\xd8\x70\xbf\xbf\xe6\x0c\xc9\x14\xfc\xc0\x4b\x12\x17\x09\xc2\x1e\xa1\x38\x46\x20\x5d\x5c\x14\x7a\xf1\x98\x50\x83\xb3\xe8\x12\x74\xed\x68\x2f\x0f\x61\xc4\xe2\xb0\x5f\xa4\xbe\x8c\x21\x56\x8c\x71\xcc\xd2\x08\xf5\x4b\xb5\x11\x41\x79\x03\x1c\xc0\x88\xc5\x2e\xa2\x5b\xa2\xc0\xa8\x7f\xf8\xf4\xe0\xe8\xf9\x83\x4b\xe0\x58\x72\xb1\xe6\x40\x68\x94\xf2\x0a\xae\x20\x18\xb3\x00\x81\xd0\x93\x8b\x68\x1a\x0e\x04\x59\x25\x89\x0c\x14\xc8\x70\x3e\x9e\x7a\x41\x8a\x5d\xd4\x29\x71\x71\xac\x21\x99\x9e\x89\x3b\x4d\x43\xd4\x3f\x7a\xf4\xcb\xe1\xd3\x83\x37\x46\x5d\x60\x66\xcc\xe5\x30\x23\xde\x5d\x24\x7e\x16\xe6\x9c\x71\x2f\x90\xdc\x4f\xee\x3f\x7c\x1b\xdc\xd5\x06\x45\xda\x6d\xdb\xb6\xdf\x64\x04\xd1\x16\x1b\x8d\x12\x21\xe3\xc3\xcf\x1f\x1e\x3f\xf9\xe3\xf5\xa2\x10\x54\x0b\xa4\x9c\xab\x8d\xac\x53\x88\x0e\x52\xce\x19\xcd\xe1\x06\x9c\xc2\x80\x53\x23\x8a\x49\xe8\xc5\xfb\x0a\xd1\xf3\x39\x61\x74\x8b\xe3\x84\xa3\xfe\xd1\x6f\x5f\xbf\x7c\xfe\x08\xb4\xae\x52\xb8\x0e\x5d\xa5\x97\x96\x63\x29\xa8\xfe\xeb\x60\x47\x31\x1b\xa6\x3e\x46\x59\xef\x38\x13\x56\x8c\x65\x1e\x4f\xee\x3f\x38\x7e\xf1\xac\xc8\xe6\x1c\x72\x55\x91\x56\x74\x88\x65\xd5\xab\x2d\x89\xdf\xe1\x57\x3f\xff\xf9\xe2\xf1\xe1\xc3\xe7\x2a\xa2\x42\x98\x49\xe4\xd1\xac\x4e\x99\x66\x1d\x4b\xcc\xd5\xc0\xeb\x95\x3f\xeb\xd9\x5e\xc2\xf0\xe4\xc7\x2f\x5e\x3e\xfb\xe1\xe8\xe0\xf7\x93\x83\x5f\x6b\xc4\x48\x88\x5f\x9d\x97\x63\xa9\xde\x99\x4b\xcc\x49\xfc\x98\x44\x5c\x2d\xfa\x8c\x26\x1c\x2e\xa8\xee\x05\x2e\x0c\x99\x9f\x86\x98\x72\x73\x27\xc5\xf1\xfe\x35\x1c\x60\x9f\xb3\x58\x43\xe7\xb3\xfe\xd6\xea\x55\xdd\x68\x1a\x9e\xee\x25\x8e\x6a\xdd\x49\x26\xb8\xd9\x6d\xe3\x7c\x66\xb2\x51\x75\x8c\x56\x3b\x46\x4d\x8e\x24\xc4\x2b\xfc\x48\x88\x17\xdc\xd4\x89\x3d\xdd\x51\xd9\x08\x57\xe9\x3b\x4a\xa9\x94\x3b\x8c\x31\xff\x44\x2e\x7d\x84\xf7\x35\xea\xeb\x40\xc3\x16\xdc\x2d\x0a\x14\x63\x9e\xc6\x14\x90\x72\xdf\x42\xd0\x06\xea\x43\x1b\x90\x1a\x86\x8a\xc9\xac\x09\x74\x09\x62\x80\x39\x10\x8e\x45\x3d\x02\xe6\x7b\xc1\x35\xce\x62\x6f\x8c\xcd\x31\xe6\x57\x38\x0e\xb5\x65\x84\xb2\x80\xc5\x8f\x8c\x40\x13\xee\x65\xc8\x1c\x76\x0a\x2e\x44\xe2\xfd\xe8\x0a\xe5\xca\xa8\x57\xb1\x11\xbe\xd3\xba\x63\x29\xca\x69\x65\x61\xb6\x36\x1f\xd5\xf3\x61\x2f\x8f\x3b\xa9\xc5\xad\x83\xca\x5b\x25\xfe\x72\xd0\xc9\x29\x41\xeb\x80\x44\x8e\x33\x84\xa5\xfb\x8d\x3c\x51\x5f\xe2\x05\x1a\x4d\x2b\x19\x11\x91\x0a\xcd\x3b\x60\xd7\xe3\xcd\x22\x30\x3a\xf3\xdc\xcc\x00\x07\x09\x2e\x9c\x5c\xd7\x05\x1b\xee\xdd\x83\xfc\xa1\xd3\x80\xb1\x08\x51\x35\x1b\xb1\x18\xb4\xa9\x17\x03\x01\x57\x82\x19\xd0\xe9\x01\x81\xbe\x2b\xff\x1a\xc6\xb2\x5a\x08\xbb\x4d\x17\x48\xaf\xb1\x18\x35\x16\x34\x6d\x50\xa1\xaf\x51\x1d\x96\xa9\x79\x9e\xb6\x50\x88\xb8\x05\x16\x68\xa5\xb9\x16\x6c\x96\x33\x9b\xab\x6f\x51\xe4\x54\x4b\xf0\x8e\x2e\x00\xaa\xf2\x8e\xe5\x61\xbc\x79\x6b\x1e\x82\xcc\x84\x54\x3e\xb8\x60\x8b\x1c\x38\xf0\xb1\xc7\x27\x66\xc4\x76\x05\x88\x19\x60\x3a\x16\x6f\xe3\xb4\xd5\x03\xd2\x6e\x2f\x53\x77\x02\x2e\x20\xd4\x5b\xcc\xb0\x58\xdc\x56\xb8\xdb\xe0\x00\xed\xc1\xf6\x22\x82\xf8\x25\xd0\x76\x21\xc1\x3b\x37\xe7\x87\x04\xac\x06\x1e\xa2\x56\x60\xc0\x76\xab\x75\xab\xb7\x00\x24\xa2\x20\xf0\xbf\x55\xae\x4d\x25\xcc\x4a\x61\x46\x69\x32\xd1\x92\x92\xd9\xe2\x41\x8b\x31\x6f\xea\x30\xd7\x45\x0b\x55\x87\x85\x2e\x29\x72\xc1\xad\xfb\xae\x0e\xd4\x17\x45\x2d\xa6\x3a\xb6\x72\xda\x04\x8d\xfa\xb0\x09\x94\x42\x1b\x3a\xd5\x3a\xab\xee\xca\xe8\xe5\x89\x47\xc7\xa2\x2b\x17\x7b\x6b\xe5\xdd\x94\x59\x71\x31\x15\x99\xcd\xee\x2a\x53\xbe\x45\x95\x42\xcc\xed\xd5\x8d\x54\x36\xa7\x69\xb8\x60\x2d\x4a\x3b\xbf\x88\xca\xd6\x72\xae\xb0\xaf\x1c\xfc\x73\x24\xb9\xea\x5d\xd5\x14\x81\x16\xac\xaf\x43\x31\x23\xdb\xc4\xfa\x7a\xce\xb7\x2f\x0e\xba\x7c\x14\x74\xc4\x53\x5d\x36\xb2\xdd\x66\xce\x72\xcb\xd6\xd2\x33\x5b\x50\x94\x6f\xb6\xf5\xba\x2f\x48\x79\x7e\x45\xce\xeb\x28\x19\xe9\x8a\x4a\x4d\x39\xca\xa5\xb8\xe5\x4a\xb7\xcb\x4a\x27\x9f\xa5\x54\xf8\xc8\xd2\x87\xde\x9e\x66\xeb\xd9\x98\x50\x2d\xa7\x61\x14\xe8\x3a\x14\x61\x2e\x43\xcb\xee\xe7\x1c\x76\x13\x6c\x53\xc4\xdb\xf9\xff\xc5\x5e\x35\xc8\xec\x2e\x37\x09\xa5\x38\xbe\x8e\xf7\xb8\x3c\xbb\xd0\xce\x31\xda\x80\x8e\x9f\x7c\x5b\x3b\xcc\xf9\x9b\x43\xc5\xab\x4c\x7c\x91\x6f\x6b\x01\x41\xcd\x2b\x61\x48\xaa\x6a\xa2\x57\xeb\x9f\xb3\x8c\x70\x2e\x53\x46\xfd\x5c\xe7\xb9\xe4\x7b\xb9\x81\x10\xe6\x69\xeb\x8a\x74\xb3\x45\x41\xaa\x64\xf2\x6f\x1c\x26\x65\x5d\x12\xd2\xdc\xa1\x9a\xb7\x37\x79\xa0\xc4\x53\xb1\xe5\xb2\x13\x96\x2c\x57\xb3\xde\xcc\xae\x56\xf4\x3c\xdb\x5a\xab\xa1\xc6\x73\x83\x6c\xa2\xf1\x3d\xb8\xf6\xef\x54\x4b\x54\x2c\x20\xfe\x9d\x4a\xc1\x70\x39\x02\x6c\x46\x31\x9e\x62\xca\x3f\xc0\x23\x2f\x0d\xb8\xb6\xa4\x3a\x6f\xb3\x9a\xff\x68\x73\x54\x4f\x6a\xc3\xfe\x62\x25\x45\x8f\x78\xc5\x26\xd5\x58\xfc\xa2\x47\x65\x80\xed\x7c\x63\xbd\xa9\x65\x36\xcb\xa2\x3a\xbf\x4b\xe8\x90\xed\x9a\x2c\xc2\x54\xbb\x6d\x0d\xd9\x2e\x0d\x98\x37\x7c\x8f\x52\xf7\xc2\x5d\x89\x34\x5b\xa7\xbe\x1c\x0b\xf4\xd9\xba\x6c\x75\xf2\x59\xee\x3f\x5b\x57\x8c\xdc\x0b\x77\xd5\x60\x76\xbb\x49\x77\xab\x64\x26\xbf\x08\x9c\x51\x63\xcd\xd1\x74\x45\x18\xdd\x82\x7e\x91\xb5\xae\x0e\xdd\x56\x4e\x7a\x05\xd5\x8d\xf3\xe5\x8f\x02\x1b\xff\xc9\xf3\xf0\xd6\xae\xfe\x46\x6d\xda\x7f\xbb\x01\xad\xad\x39\x56\xfe\x6f\xbc\x63\x65\xdf\x55\x2d\xf9\xb9\xfb\xaf\x00\x00\x00\xff\xff\x82\x71\xf0\x03\xfe\x16\x00\x00")

func handleIndex(w http.ResponseWriter, r *http.Request) {
	reader, _ := gzip.NewReader(bytes.NewReader(_templatesIndexHtml))
	html, _ := ioutil.ReadAll(reader)
	w.Header().Set("Content-Type", "text/html")
	w.WriteHeader(200)
	w.Write(html)
}

func handleGenerate(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	var nc, nn, maxCount uint64
	var err error

	if nc, err = strconv.ParseUint(query.Get("nc"), 10, 64); err != nil {
		w.WriteHeader(500)
		w.Write([]byte("参数错误 nc"))
		return
	}

	if nn, err = strconv.ParseUint(query.Get("nn"), 10, 64); err != nil {
		w.WriteHeader(500)
		w.Write([]byte("参数错误 nn"))
		return
	}

	if maxCount, err = strconv.ParseUint(query.Get("count"), 10, 64); err != nil {
		w.WriteHeader(500)
		w.Write([]byte("参数错误 count"))
		return
	}

	var offset uint64 = 0

	if query.Get("offset") != "" {
		if offset, err = strconv.ParseUint(query.Get("offset"), 10, 64); err != nil {
			w.WriteHeader(500)
			w.Write([]byte("参数错误 count"))
			return
		}
	}

	w.Header().Set("Content-Disposition", fmt.Sprintf(`attachment; filename="%d字母-%d数字.txt"`, nc, nn))
	w.WriteHeader(200)

	gen := angen.NewGenerator(int(nc), int(nn), offset, maxCount)
	ch, closer := gen.Gen()
	defer close(closer)

	for v := range ch {
		_, err := w.Write([]byte(*v))
		if err != nil {
			return
		}

		_, err = w.Write([]byte("\r\n"))
		if err != nil {
			return
		}
	}
}

// Server.
func main() {
	http.HandleFunc("/", handleIndex)
	http.HandleFunc("/download", handleGenerate)

	go func() {
		t := time.NewTimer(2 * time.Second)
		select {
		case <-t.C:
			{
				open.Run("http://127.0.0.1:9900")
			}
		}
	}()

	if err := http.ListenAndServe("0.0.0.0:9900", nil); err != nil {
		fmt.Printf("%s", err.Error())
	}
}
