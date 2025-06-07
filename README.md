# CEP Finder - Desafio Multithreading Golang

[![Go Version](https://img.shields.io/badge/Go-1.24.0-blue.svg)](https://golang.org/)
[![License](https://img.shields.io/badge/License-MIT-green.svg)](LICENSE)

Este projeto é parte da **atividade 02 da pós de Golang** e demonstra o uso de **multithreading** e **consultas simultâneas a APIs** para buscar informações de CEP de forma eficiente.

## 📖 Sobre o Projeto

O **CEP Finder** realiza consultas simultâneas a duas APIs de CEP diferentes usando goroutines e seleciona a resposta mais rápida, implementando conceitos avançados de concorrência em Go.

### 🎯 Características Principais

- **Consultas Simultâneas**: Realiza requisições paralelas para duas APIs diferentes
- **Timeout Controlado**: Limite de 1 segundo para as requisições
- **Seleção por Performance**: Aceita apenas a resposta mais rápida
- **Clean Architecture**: Código organizado em camadas (models, api, main)
- **Tratamento de Erro**: Gerenciamento adequado de timeouts e erros de rede

### 🔌 APIs Utilizadas

1. **BrasilAPI**: `https://brasilapi.com.br/api/cep/v1/{cep}`
2. **ViaCEP**: `https://viacep.com.br/ws/{cep}/json/`

## 🏗️ Estrutura do Projeto

```
cep-finder/
├── main.go                 # Ponto de entrada da aplicação
├── go.mod                  # Configuração do módulo Go
├── api/
│   ├── brasilapi.go       # Cliente para BrasilAPI
│   └── viacep.go          # Cliente para ViaCEP  
└── models/
    └── endereco.go        # Estruturas de dados
```

### 📁 Descrição dos Arquivos

| Arquivo | Função |
|---------|--------|
| `main.go` | Coordena o fluxo principal, gerencia goroutines e timeout |
| `models/endereco.go` | Define estruturas para as respostas das APIs e modelo interno |
| `api/brasilapi.go` | Implementa consulta à BrasilAPI com context e channels |
| `api/viacep.go` | Implementa consulta à ViaCEP com context e channels |

## 🚀 Como Executar

### ✅ Pré-requisitos

- **Go 1.24.0+** instalado
- Conexão com a internet para acessar as APIs

### 📥 Instalação

1. **Clone o repositório**:
```bash
git clone https://github.com/danielencestari/pos_go02.git
cd pos_go02/cep-finder
```

2. **Verifique o módulo Go**:
```bash
go mod tidy
```

### ▶️ Executando a Aplicação

```bash
go run main.go
```

### 📱 Uso

1. Execute o programa
2. Digite um CEP válido quando solicitado (exemplo: `01153000`)
3. Aguarde o resultado da API mais rápida
4. O programa exibirá:
   - Qual API respondeu primeiro
   - Dados completos do endereço (CEP, Estado, Cidade, Bairro, Rua)

### 💡 Exemplo de Uso

```bash
$ go run main.go
Digite o CEP: 01153000
Resposta mais rápida (BrasilAPI): {Cep:01153-000 Estado:SP Cidade:São Paulo Bairro:Barra Funda Rua:Rua Vitorino Carmilo Fonte:BrasilAPI}
CEP:  01153-000
Estado:  SP
Cidade:  São Paulo
Bairro:  Barra Funda
Rua:  Rua Vitorino Carmilo
```

## 🧪 Testando

### Testes Manuais

1. **CEP Válido**: Teste com CEPs conhecidos (ex: `01153000`, `20040020`)
2. **CEP Inválido**: Teste com CEPs inexistentes para verificar tratamento de erro
3. **Timeout**: Simule conexão lenta para testar o timeout de 1 segundo

### Casos de Teste Sugeridos

| CEP | Localização Esperada | Teste |
|-----|---------------------|-------|
| `01153000` | São Paulo - SP | CEP válido comum |
| `20040020` | Rio de Janeiro - RJ | CEP válido alternativo |
| `00000000` | - | CEP inválido |
| `123` | - | Formato inválido |

## 🔧 Tecnologias e Conceitos

### 🛠️ Stack Tecnológica

- **Go 1.24.0**: Linguagem principal
- **Context**: Controle de timeout e cancelamento
- **Goroutines**: Concorrência e paralelismo  
- **Channels**: Comunicação entre goroutines
- **HTTP Client**: Requisições às APIs REST
- **JSON**: Serialização/deserialização de dados

### 📚 Conceitos Aplicados

- **Multithreading**: Execução simultânea de requisições
- **Context with Timeout**: Controle de tempo limite
- **Channel Communication**: Comunicação segura entre goroutines
- **Select Statement**: Seleção não-bloqueante de operações
- **Clean Architecture**: Separação de responsabilidades
- **Error Handling**: Tratamento adequado de erros

## 🏛️ Arquitetura

### 🔄 Fluxo de Execução

1. **Input**: Usuário fornece o CEP
2. **Context**: Criação de contexto com timeout de 1s
3. **Goroutines**: Duas goroutines executam consultas simultâneas
4. **Channels**: Resultados enviados via channel com buffer
5. **Select**: Primeira resposta é aceita, outras descartadas
6. **Output**: Exibição dos dados da API mais rápida

### 🎯 Padrões Utilizados

- **Separation of Concerns**: Cada pacote tem responsabilidade específica
- **Dependency Injection**: Context passado como parâmetro
- **Channel Pattern**: Comunicação via channels tipados
- **Timeout Pattern**: Controle de tempo com context.WithTimeout

## 🚨 Tratamento de Erros

O sistema trata os seguintes cenários:

- ✅ **Timeout**: Limite de 1 segundo respeitado
- ✅ **Erro de Rede**: Conexão indisponível
- ✅ **JSON Inválido**: Resposta malformada das APIs
- ✅ **CEP Inexistente**: APIs retornam erro 404

## 🔍 Monitoramento e Debug

### Logs Disponíveis

- Erros de consulta às APIs
- Erros de decodificação JSON
- Timeout de requisições

### Como Debugar

1. Adicione logs adicionais nas funções das APIs
2. Verifique a conectividade com as APIs
3. Teste com CEPs conhecidamente válidos
4. Monitore o tempo de resposta das APIs

## 📈 Performance

### Benchmarks Esperados

- **Tempo médio**: < 500ms para CEPs válidos
- **Timeout**: 1s máximo garantido
- **Concorrência**: 2 requisições simultâneas
- **Memoria**: Uso mínimo com channels bufferizados

## 🤝 Contribuição

Para contribuir com o projeto:

1. Fork o repositório
2. Crie uma branch para sua feature (`git checkout -b feature/nova-feature`)
3. Commit suas mudanças (`git commit -am 'Adiciona nova feature'`)
4. Push para a branch (`git push origin feature/nova-feature`)
3. Abra um Pull Request

## 📄 Licença

Este projeto é parte de uma atividade acadêmica da pós-graduação em Golang.

## 🔗 Links Úteis

- [Repositório GitHub](https://github.com/danielencestari/pos_go02)
- [BrasilAPI](https://brasilapi.com.br)
- [ViaCEP](https://viacep.com.br)
- [Documentação Go](https://golang.org/doc/)
- [Go Concurrency Patterns](https://golang.org/doc/effective_go.html#concurrency)

---

**Desenvolvido por**: [Daniel Encestari](https://github.com/danielencestari)  
**Curso**: Pós-graduação em Golang  
**Atividade**: Desafio 02 - Multithreading e APIs 
