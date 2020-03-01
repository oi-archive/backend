package main

import (
	"flag"
	"github.com/axengine/cache"
	"github.com/gin-contrib/cache/persistence"
	"github.com/gin-gonic/gin"
	"log"
	"time"
)

func PanicIfError(err error) {
	if err != nil {
		log.Panicln(err)
	}
}
func main() {
	cliPath := flag.String("source", "./source", "source 仓库路径")
	flag.Parse()
	BasePath = *cliPath

	PanicIfError(UpdateData())

	cacheTime := time.Minute * 5
	store := persistence.NewInMemoryStore(time.Minute * 5)

	router := gin.Default()
	router.GET("/metadata", cache.CachePage(store, cacheTime, APIMetadata))
	router.GET("/problem-set-list", cache.CachePage(store, cacheTime, APIProblemSetList))
	router.GET("/problem-set/:problemset/metadata", cache.CachePage(store, cacheTime, APIProblemSetMetadata))
	router.GET("/problem-list/:problemset/:page", cache.CachePage(store, cacheTime, APIProblemSetPage))
	router.GET("/problem/:problemset/:problem", cache.CachePage(store, cacheTime, APIProblem))
	PanicIfError(router.Run("0.0.0.0:10001"))
}
