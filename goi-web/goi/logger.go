package goi

import (
	"log"
	"time"
)

func Logger() HandlerFunc {
	return func(c *Context) {
		//start time
		t := time.Now()
		//process request
		c.Next()
		//calculate resolution time
		log.Printf("[%d] %s in %v", c.StatusCode, c.Req.RequestURI, time.Since(t))
	}
}
