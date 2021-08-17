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
	audioPlayer.Play(name)
	c.String(http.StatusOK, "OK")
}

func main() {
	//router := gin.Default()
	//router.GET("/play/:name", play)

	//router.Run("localhost:1337")

	audioPlayer.Play("good.mp3")
}
