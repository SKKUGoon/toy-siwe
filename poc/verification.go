package poc

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/spruceid/siwe-go"
	"log"
	"net/http"
)

type siweMessage struct {
	Message   string `json:"message"`
	Signature string `json:"signature"`
}

func siweNonce(c *gin.Context) {
	c.Header("Content-Type", "text/plain")
	c.String(http.StatusOK, siwe.GenerateNonce())
}

func siweVerifyMsg(c *gin.Context) {
	// message + signature
	var vmsg siweMessage

	err := c.ShouldBindJSON(&vmsg)
	if err != nil {
		log.Println("siwe json parse error", err)
		return
	}
	// Parse received message
	pmsg, err := siwe.ParseMessage(vmsg.Message)
	if err != nil {
		log.Println("siwe parse message", err)
		return
	}
	// If requested from domain must check
	sampleDomain := "localhost:8080"
	publicKey, err := pmsg.Verify(vmsg.Signature, &sampleDomain, nil, nil)
	if err != nil {
		log.Println("siwe verification error", err)
		return
	}
	fmt.Println(publicKey)
	// message validate
	// siwe.Message.Verify() <- method
	//siwe.Message{}.Verify()
}
