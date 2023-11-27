package repository

import (
	"api/db"
	"api/models"
	"log"
)

func GetCboById(id int64) (cbo models.Cbo, err error) {
	conn, err := db.GetDB()
	if err != nil {
		return
	}

	row := conn.First(&cbo, id)
	log.Printf("row: %v", row)

	return
}

func GetCboByCode(code int64) (cbo models.Cbo, err error) {
	conn, err := db.GetDB()
	if err != nil {
		return
	}

	row := conn.Where("codigo = ?", code).Find(&cbo)
	log.Printf("row: %v", row)

	return
}

func GetAllCbo() (cbos []models.Cbo, err error) {
	conn, err := db.GetDB()
	if err != nil {
		return
	}

	rows := conn.Find(&cbos)
	log.Printf("rows: %v", rows)

	return
}

func PostCbo(cboPost models.Cbo) (cboBack models.Cbo, err error) {
	conn, err := db.GetDB()
	if err != nil {
		return
	}

	row := conn.Create(&cboPost)
	log.Printf("row: %v", row)

	conn.First(&cboBack, cboPost.IdCbo)

	return
}

func PutCbo(cboPut models.Cbo) (cboBack models.Cbo, err error) {
	conn, err := db.GetDB()
	if err != nil {
		return
	}

	cboBack = cboPut
	cboBack.IdCbo = 0

	if cboPut.IdCbo != 0 {
		row := conn.Table("cbo").Where("idcbo = ?", cboPut.IdCbo).Updates(&cboBack)
		log.Printf("row: %v", row)
	} else {
		conn.Omit("idcbo").Create(&models.Cbo{
			Desc: cboPut.Desc,
			Code: cboPut.Code,
		})
	}

	conn.First(&cboBack, cboPut.IdCbo)

	return
}
