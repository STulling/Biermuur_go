package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/STulling/Biermuur_go/audio"
	"github.com/STulling/Biermuur_go/display"
	"github.com/STulling/Biermuur_go/displaycontroller"
	"github.com/STulling/Biermuur_go/musicio"
	"github.com/STulling/Biermuur_go/musicio/playlists"
	"github.com/gin-gonic/gin"
)

var (
	audioPlayer = audio.CreateAudioPlayer()
)

func play(c *gin.Context) {
	name := c.Param("name")
	go audio.MusicQueue.AddSong(name)
	c.String(http.StatusOK, "OK")
}

func list(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, musicio.ListSongs())
}

func listPlaylists(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, musicio.ListSongs())
}

func add(c *gin.Context) {
	name := c.Param("name")
	go musicio.AddSong(name)
	c.String(http.StatusOK, "OK")
}

func setAction(c *gin.Context) {
	action := c.Param("action")
	go displaycontroller.SetCallback(action)
	c.String(http.StatusOK, "OK")
}

func playPlaylist(c *gin.Context) {
	name := c.Param("name")
	go playlists.PlayPlaylist(name)
	c.String(http.StatusOK, "OK")
}

func simpleAction(c *gin.Context) {
	switch action := c.Param("action"); action {
	case "clear":
		audio.MusicQueue.Clear()
		displaycontroller.SetCallback("clear")
	default:
		fmt.Printf("Unknown action: %s.\n", action)
	}
	c.String(http.StatusOK, "OK")
}

func main() {
	router := gin.Default()
	router.GET("/api/songs/play/:name", play)
	router.GET("/api/songs/add/:name", add)
	router.GET("/api/DJ/:action", setAction)
	router.GET("/api/songs", list)
	router.GET("/api/playlists", listPlaylists)
	router.GET("/api/playlists/play/:name", playPlaylist)
	router.GET("/api/common/:action", simpleAction)
	display.Init()
	time.Sleep(time.Second * 3)
	displaycontroller.SetCallback("clear")
	go displaycontroller.RunDisplayPipe()
	go audioPlayer.Start()

	router.Run("0.0.0.0:1337")
}

/*
api.add_resource(SongAdder, '/api/songs/add/<string:song_name>')
api.add_resource(SongModifier, '/api/songs/<string:song_name>')
api.add_resource(CommonControls, '/api/common/<string:action>')
api.add_resource(Settings, '/api/settings/<string:setting>')
api.add_resource(DJControls, '/api/DJ/<string:action>')
api.add_resource(PlaylistControls, '/api/playlists/<string:action>/<string:playlist_name>')
api.add_resource(PlaylistLister, '/api/playlists')
*/
