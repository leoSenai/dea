package repository

import (
	"api/db"
	"api/models"
)

func PostProximity(proximity models.Proximity) (err error) {
	conn, err := db.OpenConnection()

	if err != nil {
		return
	}

	conn.Create(proximity)

	return
}
