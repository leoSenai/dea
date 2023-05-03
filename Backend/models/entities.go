package models

type Pessoa struct {
	IdPessoa        int64  `json:"idpessoa"`
	Nome            string `json:"nome"`
	DataNascimento  string `json:"dataNascimento"`
	NumeroDocumento string `json:"numeroDocumento"`
	TipoDocumento   string `json:"tipoDocumento"`
	Senha           string `json:"senha"`
	Salt            string `json:"salt"`
}
