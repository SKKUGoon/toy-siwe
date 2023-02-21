package poc

import (
	"github.com/gin-gonic/gin"
	"toy-siwe/util"
)

type WS struct {
	Conn *gin.Engine
}

func WebServerStartUp() WS {
	router := gin.Default()
	router.Use(util.CORSMiddleware())

	src := WS{
		Conn: router,
	}
	return src
}

func (w WS) WebServerSIWE() WS {
	login := w.Conn.Group("/api/login")
	login.GET("/nonce", siweNonce)
	login.POST("/verify", siweVerifyMsg)
	return w
}

//func (w WS) WebServerSession() WS {
//	login := w.Conn.Group("/api/session")
//	login.GET("/login", func(c *gin.Context) {
//		session := sessions.Default(c)
//
//		iu
//	})
//	return w
//}
