package main

import (
	"flag"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/qingw1230/study-im-server/internal/api/router"
	"github.com/qingw1230/study-im-server/pkg/common/log"
)

func main() {
	gin.SetMode(gin.ReleaseMode)
	r := router.Router()
	log.NewPrivateLog("api")
	ginPort := flag.Int("port", 10000, "default 10000 as port")
	flag.Parse()
	r.Run(":" + strconv.Itoa(*ginPort))
}
