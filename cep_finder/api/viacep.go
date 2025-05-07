package api

import (
	"cep-finder/models"
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

// Função para consultar a ViaCEP.

func BuscarViaCep(ctx context.Context, cep string, ch chan<- models.Endereco) {
	url := fmt.Sprintf("https://viacep.com.br/ws/%s/json/", cep)

	req, _ := http.NewRequestWithContext(ctx, "GET", url, nil)
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Println("Erro ao consultar o ViaCep:", err)
		return
	}
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)

	var apiResponse models.ViaCepResponse
	if err := json.Unmarshal(body, &apiResponse); err != nil {
		fmt.Println("Erro ao decodificar a resposta da ViaCep:", err)
		return
	}

	ch <- models.Endereco{
		Cep:    apiResponse.Cep,
		Estado: apiResponse.Uf,
		Cidade: apiResponse.Localidade,
		Bairro: apiResponse.Bairro,
		Rua:    apiResponse.Logradouro,
		Fonte:  "ViaCEP",
	}
}
