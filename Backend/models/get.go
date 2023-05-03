package models

import (
	"api/db"
	"log"
)

func Get(id int64) (pessoa Pessoa, err error) {
	conn, err := db.OpenConnection()
	if err != nil {
		return
	}

	// defer conn.Close()

	row := conn.First(&pessoa, id)
	log.Printf("row: %v", row)
	// row := conn.QueryRow(`SELECT * FROM pessoa WHERE idpessoa=$1`, id)

	// err = row.Scan(&pessoa.IdPessoa, &pessoa.Nome, &pessoa.DataNascimento, &pessoa.NumeroDocumento, &pessoa.TipoDocumento, &pessoa.Senha, &pessoa.Salt)

	return
}

func Hello() {
	log.Printf("Hello Word")
}
