package repository

import (
	"api/db"
	"api/models"
	"log"
)

func GetAnamneseHasAskingByAnamneseId(id int64) (anamneseHasAsking models.AnamneseHasAsking, err error) {
	conn, err := db.OpenConnection()
	if err != nil {
		return
	}

	row := conn.First(&anamneseHasAsking, id)
	log.Printf("row: %v", row)

	return
}

func GetAnamneseHasAskingByAskingId(id int64) (anamneseHasAsking models.AnamneseHasAsking, err error) {
	conn, err := db.OpenConnection()
	if err != nil {
		return
	}

	row := conn.First(&anamneseHasAsking, id)
	log.Printf("row: %v", row)

	return
}

func GetAllAnamneseHasAsking() (anamneseHasAskings []models.AnamneseHasAsking, err error) {
	conn, err := db.OpenConnection()
	if err != nil {
		return
	}

	rows := conn.Find(&anamneseHasAskings)
	log.Printf("rows: %v", rows)

	return
}

func PostAnamneseHasAsking(anamneseHasAskingPost models.AnamneseHasAsking) (anamneseHasAskingBack models.AnamneseHasAsking, err error) {
	conn, err := db.OpenConnection()
	if err != nil {
		return
	}

	row := conn.Create(&anamneseHasAskingPost)
	log.Printf("row: %v", row)
	conn.First(&anamneseHasAskingBack, anamneseHasAskingPost.IdAnamnese)

	return
}

func DeleteAnamneseHasAsking(anamneseHasAskingPost models.AnamneseHasAsking) (anamneseHasAskingBack models.AnamneseHasAsking, err error) {
	conn, err := db.OpenConnection()
	if err != nil {
		return
	}

	anamneseHasAskingBack = anamneseHasAskingPost
	anamneseHasAskingBack.IdAnamnese = 0

	if anamneseHasAskingPost.IdAnamnese != 0 {
		row := conn.Table("anamnese").Where("idanamnese = ?", anamneseHasAskingPost.IdAnamnese).Updates(&anamneseHasAskingBack)
		log.Printf("row: %v", row)
		conn.First(&anamneseHasAskingBack, anamneseHasAskingPost.IdAnamnese)
	}

	return
}
