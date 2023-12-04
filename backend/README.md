## What is Go?

```go
package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default() // gin.Default() returns an Engine instance with the Logger and Recovery middleware already attached.
	r.GET("/ping", func(c *gin.Context) { // gin.Context is the most important part of gin. It allows us to pass variables between middleware, manage the flow, validate the JSON of a request and render a JSON response for example.
		c.JSON(http.StatusOK, gin.H{ //gin.H is a shortcut for map[string]interface{}
			"message": "pong",
		})
	})

	r.Run(":8080")
}
```

## What is Gin?

Gin is a web framework written in Go (Golang). It features a martini-like API with much better performance, up to 40 times faster thanks to httprouter. If you need performance and good productivity, you will love Gin.

## Done Till Now

1. Student login, register and all working
2. Hostel, admin models created