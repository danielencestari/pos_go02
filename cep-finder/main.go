package main

import (
	"cep-finder/api"
	"cep-finder/models"
	"context"
	"fmt"
	"time"
)

//Faz o input, inicia as concorrências, lida com o fluxo principal.

func main() {
	var cep string
	fmt.Print("Digite o CEP: ")
	fmt.Scanln(&cep)

	// cria o canal de comunicação entre as goroutines do tipo Endereço com um buffer de tamanho 2
	// aqui irei receber os resultados das chamadas para as duas API's de CEP que estão sendo feitas de forma concorrente
	// tem que ser um buffer 2 garante que teremos
	resultChan := make(chan models.Endereco, 2)

	// ctx de 1s para cancelar as goroutines caso o tempo limite seja excedido
	// defer cancel para garantir que o contexto será encerrado
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	// inicia as goroutines para buscar o CEP nas duas API's de forma concorrente e sem bloquear o fluxo principal
	// os resultados serão enviados para o canal resultChan
	go api.BuscarBrasilAPI(ctx, cep, resultChan)
	go api.BuscarViaCep(ctx, cep, resultChan)

	// select para receber o resultado mais rápido das duas API's
	// caso o tempo limite seja excedido, o select irá cair no caso de ctx.Done() e a mensagem "Tempo limite excedido." será exibida
	// caso o resultado seja recebido antes do tempo limite, o select irá cair no caso de resultChan e os dados do endereço serão exibidos
	select {
	case result := <-resultChan:
		fmt.Printf("Resposta mais rápida (%s): %+v\n", result.Fonte)
		fmt.Println("CEP: ", result.Cep)
		fmt.Println("Estado: ", result.Estado)
		fmt.Println("Cidade: ", result.Cidade)
		fmt.Println("Bairro: ", result.Bairro)
		fmt.Println("Rua: ", result.Rua)
	case <-ctx.Done():
		fmt.Println("Tempo limite excedido.")
	}
}
