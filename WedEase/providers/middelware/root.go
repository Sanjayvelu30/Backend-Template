package middelware

import "sanjay/WedEase/database"

type Middelware struct {
	DBHelper database.DBHelperProvider
}

func NewMiddelware(dbHelper database.DBHelperProvider) *Middelware {
	return &Middelware{
		DBHelper: dbHelper,
	}
}
