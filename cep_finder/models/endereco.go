package models

// Definição de structs e tipos de dados globais. - structs para as respostas das APIs e struct comum para uso interno.
// Deixa claro o formato dos dados, centraliza alterações de modelos.

type BrasilAPIResponse struct {
	Cep          string `json:"cep,omitempty"`
	State        string `json:"state,omitempty"`
	City         string `json:"city,omitempty"`
	Neighborhood string `json:"neighborhood,omitempty"`
	Street       string `json:"street,omitempty"`
}

type ViaCepResponse struct {
	Cep        string `json:"cep,omitempty"`
	Uf         string `json:"uf,omitempty"`
	Localidade string `json:"localidade,omitempty"`
	Bairro     string `json:"bairro,omitempty"`
	Logradouro string `json:"logradouro,omitempty"`
}

type Endereco struct {
	Cep    string `json:"cep,omitempty"`
	Estado string `json:"estado,omitempty"`
	Cidade string `json:"cidade,omitempty"`
	Bairro string `json:"bairro,omitempty"`
	Rua    string `json:"rua,omitempty"`
	Fonte  string `json:"fonte,omitempty"`
}
