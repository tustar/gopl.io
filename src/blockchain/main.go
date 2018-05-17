package main

import (
	"database/sql"
	"fmt"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"net/http"
)

func main() {

	db, err := sql.Open("mysql",
		"root:tustar@tcp(127.0.0.1:3306)/blockchain?parseTime=true")
	if err != nil {
		log.Fatalln(err)
	}
	defer db.Close()

	db.SetMaxIdleConns(20)
	db.SetMaxOpenConns(20)

	if err := db.Ping(); err != nil {
		log.Fatalln(err)
	}

	router := gin.Default()
	router.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "It works")
	})

	v1 := router.Group("v1")
	{
		// user
		v1.POST("/user/login", login(db))

		v1.POST("/user/code", code(db))
	}

	//
	router.Run(":4000")
}

func login(db *sql.DB) func(c *gin.Context) {
	return func(c *gin.Context) {
		mobile := c.Request.FormValue("mobile")
		code := c.Request.FormValue("code")
		rs, err := db.Exec("INSERT INTO user(mobile, code) VALUES (?, ?)", mobile, code)
		if err != nil {
			log.Fatalln(err)
		}

		id, err := rs.LastInsertId()
		if err != nil {
			log.Fatalln(err)
		}
		fmt.Println("insert user Id {}", id)
		msg := fmt.Sprintf("insert successful %d", id)
		c.JSON(http.StatusOK, gin.H{"code": 200, "data": "", "msg": msg, "extra": ""})
	}
}

func code(db *sql.DB) func(c *gin.Context) {
	return func(c *gin.Context) {
		mobile := c.Request.FormValue("mobile")
		code := c.Request.FormValue("code")
		rs, err := db.Exec("INSERT INTO user(mobile, code) VALUES (?, ?)", mobile, code)
		if err != nil {
			log.Fatalln(err)
		}

		id, err := rs.LastInsertId()
		if err != nil {
			log.Fatalln(err)
		}
		fmt.Println("insert user Id {}", id)
		msg := fmt.Sprintf("insert successful %d", id)
		c.JSON(http.StatusOK, gin.H{"code": 200, "data": "", "msg": msg, "extra": ""})
	}
}
