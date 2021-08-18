package main

import (
	"net/http"

	"github.com/STulling/Biermuur_go/audio"
	"github.com/STulling/Biermuur_go/displaycontroller"
	"github.com/STulling/Biermuur_go/musicio"
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

func main() {
	router := gin.Default()
	router.GET("/api/songs/play/:name", play)
	router.GET("/api/songs/add/:name", add)
	router.GET("/api/DJ/:action", setAction)
	router.GET("/api/songs", list)
	router.GET("/api/playlists", listPlaylists)
	router.GET("/api/playlists/play/:name", playPlaylist)
	go audioPlayer.Start()

	displaycontroller.SetCallback("debug")
	go displaycontroller.RunDisplayPipe()

	router.Run("localhost:1337")
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