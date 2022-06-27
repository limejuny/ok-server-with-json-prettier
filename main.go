package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	g := r.Group("")
	{
		g.POST("/hooks", echo)
	}

	r.Run()
}

func echo(c *gin.Context) {
	var data map[string]interface{}

	// save the body to a buffer
	body, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		fmt.Println(err)
		c.JSON(500, gin.H{
			"error": err,
		})
		return
	}
	// parse the body to a map
	err = json.Unmarshal(body, &data)
	if err != nil {
		fmt.Println(string(body))
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	b, _ := json.MarshalIndent(data, "", "  ")
	fmt.Println(string(b))

	c.JSON(200, data)
}
