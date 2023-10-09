package repository

import (
	"api/db"
	"api/models"
	"log"
)

func GetAnamneseHasAskingByAnamneseId(id int64) (anamneseHasAsking []models.AnamneseHasAsking, err error) {
	conn, err := db.GetDB()
	if err != nil {
		return
	}

	row := conn.First(&anamneseHasAsking, "anamnese_idanamnese = ?", id)
	log.Printf("row: %v", row)

	return
}

func GetAnamneseHasAskingByAskingId(id int64) (anamneseHasAsking []models.AnamneseHasAsking, err error) {
	conn, err := db.GetDB()
	if err != nil {
		return
	}

	row := conn.First(&anamneseHasAsking, "pergunta_idpergunta = ?", id)
	log.Printf("row: %v", row)

	return
}

func GetAllAnamneseHasAsking() (anamneseHasAskings []models.AnamneseHasAsking, err error) {
	conn, err := db.GetDB()
	if err != nil {
		return
	}

	rows := conn.Find(&anamneseHasAskings)
	log.Printf("rows: %v", rows)

	return
}

func PostAnamneseHasAsking(anamneseHasAskingPost models.AnamneseHasAsking) (anamneseHasAskingBack models.AnamneseHasAsking, err error) {
	conn, err := db.GetDB()
	if err != nil {
		return
	}

	row := conn.Create(&anamneseHasAskingPost)
	log.Printf("row: %v", row)
	conn.First(&anamneseHasAskingBack, anamneseHasAskingPost.IdAnamnese)

	return
}
