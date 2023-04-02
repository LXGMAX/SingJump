package main

import (
	"fmt"
	"net"
	"net/http"
	"strings"

	"github.com/julienschmidt/httprouter"
	"github.com/go-vgo/robotgo"
	"github.com/itchyny/volume-go"
)

const (
	KeyPause 		= "p"
	KeyNext  		= "right"
	KeyPrevious = "left"
	KeyVolDown 	= "down"
	KeyVolUp 		= "up"
	KeyLyric 		= "d"
	KeyLike 		= "l"
)

const (
	AudioMute		 = "audio_mute"
	AudioVolDown = "audio_vol_down"
	AudioVolUp   = "audio_vol_up"
	AudioPlay    = "audio_play"
	AudioPrev    = "audio_prev"
	AudioNext    = "audio_next"
)

func keyOperate(key string) {
	robotgo.KeyTap(key, "ctrl", "alt")
}

func audioPlay(play string) {
	robotgo.KeyTap(play)
}

func audioGetVol() int {
	volume, err := volume.GetVolume()
	if err != nil {
		fmt.Println("vol err: ", err)
	}
	return volume
}

func audioGetMuteStat() bool {
	mute, err := volume.GetMuted()
	if err != nil {
		fmt.Println("get mute err: ", err)
	}
	return mute
}

func getOutBoundIP() string {
	connect, err := net.Dial("udp", "8.8.4.4:53")
	if err != nil {
		fmt.Println(err)
		return ""
	}

	addr := connect.LocalAddr().(*net.UDPAddr)
	fmt.Println(addr.String())
	ip := strings.Split(addr.String(), ":")[0]

	return ip
}

func playPause(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	fmt.Fprintln(w, "{\"play\": \"pause\"}")
	audioPlay(AudioPlay)
}

func playNext(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	fmt.Fprintln(w, "{\"play\": \"next\"}")
	audioPlay(AudioNext)
}

func playPrevious(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	fmt.Fprintln(w, "{\"play\": \"previous\"}")
	audioPlay(AudioPrev)
}

func playLike(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	keyOperate(KeyLike)
	fmt.Fprintln(w, "{\"play\": \"like\"}")
}

func playVolumeUp(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	audioPlay(AudioVolUp)
	vol := audioGetVol()
	fmt.Fprintln(w, "{\"play\": \"volup\", \"volume\": ", vol, "}")
}

func playVolumeDown(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	audioPlay(AudioVolDown)
	vol := audioGetVol()
	fmt.Fprintln(w, "{\"play\": \"voldown\", \"volume\": ", vol, "}")
}

func playLyric(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	keyOperate(KeyLyric)
	fmt.Fprintln(w, "{\"play\": \"lyric\"}")
}

func playMute(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	audioPlay(AudioMute)
	vol := audioGetVol()
	if audioGetMuteStat() == true {
		vol = 0
	}
	fmt.Fprintln(w, "{\"play\": \"mute\", \"volume\": ", vol, "}")
}

func sysGetVol(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	vol := audioGetVol()
	fmt.Fprintln(w, "{\"volume\": ", vol, "}")
}

func main() {
	ip := getOutBoundIP()
	fmt.Println("IP address:\t", ip, "\nPort:\t\t", "18890")

	router := httprouter.New()
	router.GET("/play/pause", playPause)
	router.GET("/play/next", playNext)
	router.GET("/play/previous", playPrevious)
	router.GET("/play/volup", playVolumeUp)
	router.GET("/play/voldown", playVolumeDown)
	router.GET("/play/like", playLike)
	router.GET("/play/lyric", playLyric)
	router.GET("/play/mute", playMute)
	router.GET("/sys/getvol", sysGetVol)

	err := http.ListenAndServe(":18890", router)
	if err != nil {
		panic(err)
	}
}
