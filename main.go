package main

import (
	"flag"
	"github.com/axengine/cache"
	"github.com/gin-contrib/cache/persistence"
	"github.com/gin-gonic/gin"
	"log"
	"strconv"
	"time"
)

func PanicIfError(err error) {
	if err != nil {
		log.Panicln(err)
	}
}
func main() {
	cliPath := flag.String("source", "./source", "source 仓库路径")
	port := flag.Int("port", 10001, "listening port")
	flag.Parse()
	BasePath = *cliPath
	if *port <= 1 || *port > 65535 {
		log.Fatalln("invalid port")
	}

	PanicIfError(UpdateData())

	cacheTime := time.Minute * 5
	store := persistence.NewInMemoryStore(time.Minute * 5)

	router := gin.Default()
	router.GET("/metadata", cache.CachePage(store, cacheTime, APIMetadata))
	router.GET("/problem-set-list", cache.CachePage(store, cacheTime, APIProblemSetList))
	router.GET("/problem-set/:problemset/metadata", cache.CachePage(store, cacheTime, APIProblemSetMetadata))
	router.GET("/problem-list/:problemset/:page", cache.CachePage(store, cacheTime, APIProblemSetPage))
	router.GET("/problem/:problemset/:problem", cache.CachePage(store, cacheTime, APIProblem))
	PanicIfError(router.Run("0.0.0.0:" + strconv.Itoa(*port)))
}
