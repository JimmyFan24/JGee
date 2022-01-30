package main

import (
	"JGee/JGee"
	"net/http"
)

func main()  {

	r := JGee.New()
	v1 := r.Group("/v1")
	{
		/*v1.GET("/", func(c *JGee.Context) {
			// expect /hello?name=geektutu
			c.HTML(http.StatusOK, "<h1>Hello Gee</h1>")
		})*/
		v1.GET("/hello/:name", func(c *JGee.Context) {
			// expect /hello?name=geektutu
			c.String(http.StatusOK, "hello %s, you're at %s\n", c.Query("name"), c.Path)
		})
	}
	r.Run(":9999")
	//fmt.Println(ParsePattern("/abc/*122/*abc"))
}
