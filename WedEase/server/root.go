package server

import (
	"sanjay/WedEase/database"
	"sanjay/WedEase/providers"
)

type Server struct {
	Middelware providers.Middelware
	Database   database.DBHelper
}
