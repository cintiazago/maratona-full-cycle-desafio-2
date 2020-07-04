package actions

import (
    "net/http"
    
	"github.com/gobuffalo/buffalo"
)

// HelloMessage default implementation.
func HelloMessage(c buffalo.Context) error {
	return c.Render(http.StatusOK, r.HTML("hello/message.html"))
}

