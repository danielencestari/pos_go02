package api

import (
	"cep-finder/models"
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

// Função para consultar a BrasilAPI.

func BuscarBrasilAPI(ctx context.Context, cep string, ch chan<- models.Endereco) {
	url := fmt.Sprintf("https://brasilapi.com.br/api/cep/v1/%s", cep)

	req, _ := http.NewRequestWithContext(ctx, "GET", url, nil)
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Println("Erro ao consultar a BrasilAPI:", err)
		return
	}
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)

	var apiResponse models.BrasilAPIResponse
	if err := json.Unmarshal(body, &apiResponse); err != nil {
		fmt.Println("Erro ao decodificar a resposta da BrasilAPI:", err)
		return
	}

	ch <- models.Endereco{
		Cep:    apiResponse.Cep,
		Estado: apiResponse.State,
		Cidade: apiResponse.City,
		Bairro: apiResponse.Neighborhood,
		Rua:    apiResponse.Street,
		Fonte:  "BrasilAPI",
	}
}
