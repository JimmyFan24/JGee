package JGee

import (
	"fmt"
	"log"
	"time"
)

func Logger() HandlerFunc {
	return func(c *Context) {
		// Start timer
		fmt.Println("this is logger handlerfunc ")
		t := time.Now()
		// Process request
		c.Next()
		// Calculate resolution time
		fmt.Println("this is only logger after next")
		log.Printf("[%d] %s in %v", c.StatusCode, c.Req.RequestURI, time.Since(t))
	}
}
