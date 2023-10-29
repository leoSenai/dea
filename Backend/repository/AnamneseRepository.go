package repository

import (
	"api/db"
	"api/models"
	"log"
)

func GetAnamneseById(id int64) (anamnese models.Anamnese, err error) {
	conn, err := db.GetDB()
	if err != nil {
		return
	}

	row := conn.First(&anamnese, id)
	log.Printf("row: %v", row)

	return
}

func GetAnamneseByIdUserPatient(idUser int64, idPatient int64) (anamnese models.Anamnese, err error) {
	conn, err := db.GetDB()
	if err != nil {
		return
	}

	row := conn.Where("usuario_idusuario = ? AND paciente_idpaciente = ?", idUser, idPatient).Find(&anamnese)
	log.Printf("row: %v", row)

	return
}

func GetAllAnamnese() (anamneses []models.Anamnese, err error) {
	conn, err := db.GetDB()
	if err != nil {
		return
	}

	rows := conn.Find(&anamneses)
	log.Printf("rows: %v", rows)

	return
}

func PostAnamnese(anamnesePost models.Anamnese) (anamneseBack models.Anamnese, err error) {
	conn, err := db.GetDB()
	if err != nil {
		return
	}

	row := conn.Create(&anamnesePost)
	log.Printf("row: %v", row)
	conn.First(&anamneseBack, anamnesePost.IdAnamnese)

	return
}

func PutAnamnese(anamnesePut models.Anamnese) (anamneseBack models.Anamnese, err error) {
	conn, err := db.GetDB()
	if err != nil {
		return
	}

	anamneseBack = anamnesePut
	anamneseBack.IdAnamnese = 0

	if anamnesePut.IdAnamnese != 0 {
		row := conn.Table("anamnese").Where("idanamnese = ?", anamnesePut.IdAnamnese).Updates(&anamneseBack)
		log.Printf("row: %v", row)
		conn.First(&anamneseBack, anamnesePut.IdAnamnese)
	}

	return
}
