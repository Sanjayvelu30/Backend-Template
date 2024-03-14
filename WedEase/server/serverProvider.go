package server

import (
	"sanjay/WedEase/database"
	"sanjay/WedEase/providers/middelware"
)

// Initialize Server with configuration.
func ServerInit() *Server {
	dbHelper := database.Connect()
	middleware := middelware.NewMiddelware(dbHelper)

	return &Server{
		Middelware: middleware,
		Database:   *dbHelper,
	}

}
