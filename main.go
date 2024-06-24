package main

import (
  "net/http"

  "github.com/gin-gonic/gin"
)

type user struct {
  ID   string  `json:"id"`
  Name string  `json:"name"`
}

var users = []user {
  {ID: "1", Name: "Mike Linkey"},
  {ID: "2", Name: "Laurie Linkey"},
  {ID: "3", Name: "Josh Linkey"},
  {ID: "4", Name: "Travis Linkey"},
  {ID: "5", Name: "Bradley Linkey"},
}

func main() {
  router := gin.Default()
  router.GET("/", home)
  router.GET("/users", getUsers)

  router.Run("localhost:8080")
}

func getUsers(c *gin.Context) {
  c.IndentedJSON(http.StatusOK, users)
}

func home(c *gin.Context) {
  c.IndentedJSON(http.StatusOK, "Server is working!")
}

