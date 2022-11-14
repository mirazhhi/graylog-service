package main

import (
    "fmt"
    "github.com/gin-gonic/gin"
    "net/http"
)

func main() {
    r := gin.Default() // создаем экземпляр


    r.GET("/ping", func(c *gin.Context) {
        c.JSON(http.StatusOK, gin.H{
          "message": "pong",
        })
    })
    r.Run()

    


    fmt.Println("Hello")
}