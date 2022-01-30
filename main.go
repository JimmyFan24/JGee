package main

import (
	"JGee/JGee"
	"fmt"
	"log"
	"net/http"
	"time"
)

func onlyForV2() JGee.HandlerFunc {
	return func(c *JGee.Context) {
		fmt.Println("this is onlyForV2 func and c.index = ")
		// Start timer
		t := time.Now()
		// if a server error occurred
		//c.Fail(500, "Internal Server Error")
		// Calculate resolution time
		c.Next()
		log.Printf("this is only ForV2 after next [%d] %s in %v for group v2", c.StatusCode, c.Req.RequestURI, time.Since(t))
	}
}

func main() {

	r := JGee.New()
	r.Use(JGee.Logger()) // global midlleware
	r.GET("/", func(c *JGee.Context) {
		c.HTML(http.StatusOK, "<h1>Hello Gee</h1>")
	})
	//r.Use(JGee.Logger())
	/*r.GET("/", func(c *JGee.Context) {
		c.HTML(http.StatusOK, "<h1>Hello Gee</h1>")
	})*/

	v2 := r.Group("/v2")
	v2.Use(onlyForV2())
	{
		/*v1.GET("/", func(c *JGee.Context) {
			// expect /hello?name=geektutu
			c.HTML(http.StatusOK, "<h1>Hello Gee</h1>")
		})*/
		v2.GET("/hello/:name", func(c *JGee.Context) {
			//expect /hello?name=geektutu
			fmt.Println("this is /hello/:name")
			c.String(http.StatusOK, "hello %s, you're at %s\n", c.Query("name"), c.Path)

		})
	}
	//v1 := r.Group("/v1")
	//{
	/*v1.GET("/", func(c *JGee.Context) {
		// expect /hello?name=geektutu
		c.HTML(http.StatusOK, "<h1>Hello Gee</h1>")
	})*/
	//	v1.GET("/hello/", func(c *JGee.Context) {
	// expect /hello?name=geektutu
	//	c.String(http.StatusOK, "hello %s, you're at %s\n", c.Query("name"), c.Path)
	//	})
	//	}*/
	r.Run(":9999")
	//fmt.Println(ParsePattern("/abc/*122/*abc"))
}
