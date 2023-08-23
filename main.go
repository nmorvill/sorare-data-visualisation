package main

import (
	"fmt"
	"net/http"
	"sg/internal/players"
	"sg/internal/sorare"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	runServer(collectData())
}

func runServer(p []sorare.PlayerName) {
	r := gin.Default()

	r.Use(cors.Default())
	r.Static("/css", "./web/style")
	r.GET("/player/:slug", func(c *gin.Context) {
		c.Data(http.StatusOK, "text/html; charset=utf-8", players.RenderPlayerPlage(c.Param("slug")))
	})
	r.GET("/player-infos/:slug", func(c *gin.Context) {
		c.Data(http.StatusOK, "text/html; charset=utf-8", players.RenderGeneralPlayerPage(c.Param("slug")))
	})
	r.POST("/player-search", func(c *gin.Context) {
		search := c.PostForm("search")
		if len(search) > 0 {
			c.Data(http.StatusOK, "text/html; charset=utf-8", players.RenderSearchResults(search, &p))
		}
	})
	r.Run()
}

func collectData() []sorare.PlayerName {
	start := time.Now()
	players := sorare.GetAllPlayersNames()
	fmt.Printf("%d players got in %s\n", len(players), time.Since(start).String())
	return players
}
