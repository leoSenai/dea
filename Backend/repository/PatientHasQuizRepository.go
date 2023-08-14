package repository

import "api/db"

func VerifyResponseQuizById(idQuiz int64) (hasQuiz bool, err error) {
	conn, err := db.OpenConnection()
	if err != nil {
		return
	}

	var count int64
	row := conn.Table("patient").Where("respondido_em IS NOT NULL AND idquiz = ?", idQuiz).Count(&count)
	if row.Error != nil {
		err = row.Error
		return
	}

	if count > 0 {
		hasQuiz = true
	}

	return
}
