package main

import (
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/go-vgo/robotgo"
)

func ctrlAltKeyDown() {
	robotgo.KeyTap("cmd")
}

func playPause(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	fmt.Fprintln(w, "{\"play\": \"pause\"}")
	ctrlAltKeyDown()
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
	router.GET("/play/pause", playPause)
	router.GET("/play/next", playNext)
	router.GET("/play/previous", playPrevious)
	router.GET("/play/like", playLike)
	router.GET("/play/volup", playVolumeUp)
	router.GET("/play/voldown", playVolumeDown)
	router.GET("/play/lyric", playLyric)

	err := http.ListenAndServe(":8890", router)
	if err != nil {
		panic(err)
	}
}
