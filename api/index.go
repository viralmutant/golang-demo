package handler

import (
	"net/http"

	. "github.com/tbxark/g4vercel"
)

func Handler(w http.ResponseWriter, r *http.Request) {
    // Create a new instance of the g4vercel server
    server := New()

    // Define the handler for the path "/"
    server.get("/", func(ctx *Context) {
        // Send a JSON response
        ctx.JSON(http.StatusOK, H{
            "message": "Hello from Mittal",
        })
    })

    // Handle the request
    server.Handle(w, r)
}