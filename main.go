package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type AccountPassword struct {
	WebSerive   string `json:web_service`
	Password    string `json:passowrd`
	DateCreated string `json:date_created`
}

var passwords = []AccountPassword{
	{WebSerive: "facebook", Password: "123", DateCreated: "dsa1"},
	{WebSerive: "twitter", Password: "123", DateCreated: "dsa2"},
	{WebSerive: "google", Password: "123", DateCreated: "dsa111"},
}

func getAllPasswords(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, passwords)

}

func addPassword(c *gin.Context) {
	var newPassword AccountPassword

	if err := c.BindJSON(&newPassword); err != nil {
		fmt.Println("error parsing json")
		return
	}

	passwords = append(passwords, newPassword)
	c.IndentedJSON(http.StatusCreated, newPassword)
}

func getPasswordByWebsite(c *gin.Context) {
	website := c.Param("website")

	for _, a := range passwords {
		if a.WebSerive == website {
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}

	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "website not found"})
}

func main() {
	fmt.Println("Start bitch")
	router := gin.Default()
	router.GET("/api/password", getAllPasswords)
	router.POST("api/password", addPassword)
	router.GET("api/password/:website", getPasswordByWebsite)
	router.Run("localhost:1111")
}
