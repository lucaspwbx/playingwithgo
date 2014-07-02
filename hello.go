package main

import "github.com/gin-gonic/gin"

type LoginJSON struct {
  User string `json:"user" binding:"required"`
  Password string `json:"password" binding:"required"`
}

type Retorno struct {
  Message string `json:"message"`
  Status int `json:"status"`
}

var (
  teste gin.HandlerFunc
  posting gin.HandlerFunc
  patching gin.HandlerFunc
)

func init() {
  patching = func(c *gin.Context) {
    c.String(200, "patching")
  }
  teste = func(c *gin.Context) {
    var json LoginJSON
    if c.EnsureBody(&json) {
      if json.User=="manu" && json.Password=="123" {
	c.JSON(200, gin.H{"status": "you are logged in"})
      } else {
	c.JSON(401, gin.H{"status": "unauthorized"})
      }
    }
  }
  posting = func(c *gin.Context) {
    c.String(200,"posting")
  }
}

func main() {
  r := gin.Default()
  r.GET("/someJSON", func(c *gin.Context) {
    c.JSON(200, gin.H{"message": "hey", "status": 200})
  })
  r.GET("/moreJSON", func(c *gin.Context) {
    msg := Retorno{Message: "Retornando", Status: 200}
    c.JSON(200, msg)
  })
  r.GET("/user/:name", func(c *gin.Context) {
    name := c.Params.ByName("name")
    message := "Hello " + name
    c.String(200, message)
  })

  r.POST("/login", teste)
  r.POST("/somePost", posting)
  r.PATCH("/somePatch", patching)

  r.Run(":8080")
}
