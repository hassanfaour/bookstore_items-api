package app

import (
	"github.com/hassanfaour/bookstore_items-api/controllers/ping"
	"github.com/hassanfaour/bookstore_items-api/controllers/users"
)
func mapUrls(){
	router.GET("/ping", ping.Ping)
	router.GET("/users/:user_id", users.GetUser)
	router.POST("/users", users.CreateUser)
}
