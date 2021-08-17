package main

import (
	"net/http"

	"github.com/STulling/Biermuur_go/audio"
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

func main() {
	router := gin.Default()
	router.GET("/play/:name", play)
	go audioPlayer.Start()

	router.Run("localhost:1337")
}
