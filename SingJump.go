package main

import (
	"fmt"
	"net"
	"net/http"
	"strings"
	"encoding/json"

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

type SystemSet struct {
	Volume int `json:"setSysVolume"`
}

type Play struct {
}

type Sys struct {
}

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

func audioSetVol(vol int) error {
	err := volume.SetVolume(vol)
	if err != nil {
		fmt.Println("vol set err: ", err)
	}
	return err
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
	ip := strings.Split(addr.String(), ":")[0]

	return ip
}

/**
 * @note web api define
 */
func (p Play) Pause(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	fmt.Fprintln(w, "{\"play\": \"pause\"}")
	audioPlay(AudioPlay)
}

func (p Play) Next(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	fmt.Fprintln(w, "{\"play\": \"next\"}")
	audioPlay(AudioNext)
}

func (p Play) Previous(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	fmt.Fprintln(w, "{\"play\": \"previous\"}")
	audioPlay(AudioPrev)
}

func (p Play) Like(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	keyOperate(KeyLike)
	fmt.Fprintln(w, "{\"play\": \"like\"}")
}

func (p Play) VolumeUp(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	audioPlay(AudioVolUp)
	vol := audioGetVol()
	fmt.Fprintln(w, "{\"play\": \"volup\", \"volume\": ", vol, "}")
}

func (p Play) VolumeDown(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	audioPlay(AudioVolDown)
	vol := audioGetVol()
	fmt.Fprintln(w, "{\"play\": \"voldown\", \"volume\": ", vol, "}")
}

func (p Play) Lyric(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	keyOperate(KeyLyric)
	fmt.Fprintln(w, "{\"play\": \"lyric\"}")
}

func (p Play) Mute(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	audioPlay(AudioMute)
	vol := audioGetVol()
	if audioGetMuteStat() == true {
		vol = 0
	}
	fmt.Fprintln(w, "{\"play\": \"mute\", \"volume\": ", vol, "}")
}

func (s Sys) sysGetVol(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	vol := audioGetVol()
	fmt.Fprintln(w, "{\"volume\": ", vol, "}")
}

func (s Sys) sysSetVol(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	var parm SystemSet
	err := json.NewDecoder(r.Body).Decode(&parm)
	if err != nil {
		http.Error(w, "{\"error\":\"Error JSON format\"", http.StatusBadRequest)
		return
	}

	err = audioSetVol(parm.Volume)
	if err != nil {
		fmt.Fprintln(w, "{\"error\": \"invalid volume range\"}")
		return
	}

	fmt.Fprintln(w, "{\"volume\": ", parm.Volume, "}")
}

func main() {
	ip := getOutBoundIP()
	fmt.Println("IP address:\t", ip, "\nPort:\t\t", "18890")

	router := httprouter.New()
	pl := Play{}
	s := Sys{}
	router.GET("/play/pause", pl.Pause)
	router.GET("/play/next", pl.Next)
	router.GET("/play/previous", pl.Previous)
	router.GET("/play/volup", pl.VolumeUp)
	router.GET("/play/voldown", pl.VolumeDown)
	router.GET("/play/like", pl.Like)
	router.GET("/play/lyric", pl.Lyric)
	router.GET("/play/mute", pl.Mute)
	router.GET("/sys/getvol", s.sysGetVol)
	router.POST("/sys/setvol", s.sysSetVol)

	err := http.ListenAndServe(":18890", router)
	if err != nil {
		panic(err)
	}
}
