package main

import (
	"fmt"
	"github.com/gobuffalo/packr"
	"github.com/hoozecn/angen"
	"github.com/skratchdot/open-golang/open"
	"net/http"
	"strconv"
	"time"
)

func handleIndex(w http.ResponseWriter, r *http.Request) {
	box := packr.NewBox("./templates")
	w.Header().Set("Content-Type", "text/html")
	w.WriteHeader(200)
	w.Write([]byte(box.String("index.html")))
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
