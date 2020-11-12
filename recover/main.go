package main

import (
	"net/http"

	"gee"
)

func main() {
	r := gee.Default()
	r.GET("/", func(c *gee.Context) {
		c.String(http.StatusOK, "Hello va1ner\n")
	})

	r.GET("/panic", func(c *gee.Context) {
		names := []string{"va1ner"}
		c.String(http.StatusOK, names[100])
	})

	r.Run(":9999")
}
