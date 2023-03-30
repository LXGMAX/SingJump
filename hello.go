package main

import (
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/go-vgo/robotgo"
)

func ctrlAltKeyDown() {
	robotgo.KeyTap("")
}

func playPause(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	fmt.Fprintln(w, "{\"play\": \"pause\"}")
}

func playNext(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	fmt.Fprintln(w, "{\"play\": \"next\"}")
}

func playPrevious(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	fmt.Fprintln(w, "{\"play\": \"previous\"}")
}

func playLike(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	fmt.Fprintln(w, "{\"play\": \"like\"}")
}

func playVolumeUp(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	fmt.Fprintln(w, "{\"play\": \"volup\"}")
}

func playVolumeDown(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	fmt.Fprintln(w, "{\"play\": \"voldown\"}")
}

func playLyric(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	fmt.Fprintln(w, "{\"play\": \"lyric\"}")
}

func main() {
	//http.HandleFunc("/", handler)

	router := httprouter.New()
	router.POST("/play/pause", playPause)
	router.POST("/play/next", playNext)
	router.POST("/play/previous", playPrevious)
	router.POST("/play/like", playLike)
	router.POST("/play/volup", playVolumeUp)
	router.POST("/play/voldown", playVolumeDown)
	router.POST("/play/lyric", playLyric)

	err := http.ListenAndServe(":8890", router)
	if err != nil {
		panic(err)
	}
}
