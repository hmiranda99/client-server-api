# Desafio: Cotação do Dólar com Go

Este projeto contém dois programas escritos em Go: um servidor (`server.go`) e um cliente (`client.go`) que se comunicam via HTTP para obter a cotação do dólar, salvá-la em banco de dados e armazená-la em arquivo.

---

## 📦 Estrutura do Projeto

```
cotacao/
├── client.go       # Cliente HTTP que solicita a cotação e salva em um arquivo
├── server.go       # Servidor que busca a cotação na API externa e salva no SQLite
└── cotacao.txt     # Arquivo criado pelo client.go com o valor do dólar
```

---

## 🚀 Como rodar o projeto

### 1. Pré-requisitos

- Go 1.18 ou superior instalado  
- Conexão com a internet (para acessar a API de câmbio)

### 2. Clone o repositório

```bash
git clone https://github.com/seu-usuario/cotacao-go.git
cd cotacao-go
```

### 3. Execute o servidor

```bash
go run server.go
```

O servidor será iniciado em `http://localhost:8080`.

### 4. Em outro terminal, execute o cliente

```bash
go run client.go
```

O cliente fará uma requisição para o endpoint `/cotacao`, salvará o valor atual do dólar no arquivo `cotacao.txt`, e imprimirá no terminal:

```
Cotação salva com sucesso: 5.43
```

---

## 🧠 Tecnologias e conceitos utilizados

- `net/http` — servidor e cliente HTTP
- `context` — timeouts para chamadas externas e banco
- `encoding/json` — manipulação de JSON
- `database/sql` com `github.com/mattn/go-sqlite3` — persistência em banco de dados SQLite
- `os` e `io/ioutil` — manipulação de arquivos

---

## ⚠️ Timeouts

- Chamada à API externa: **1000ms**
- Gravação no banco de dados: **10ms**
- Requisição do client para o server: **300ms**

Erros de timeout são registrados no log para fins de debug.

---

## 📁 Exemplo do arquivo gerado

`cotacao.txt`:

```
Dólar: 5.43
```

---

## 🗃 Banco de dados

- Nome: `cotacoes.db`
- Tabela: `cotacoes (id INTEGER PRIMARY KEY, bid TEXT)`

---

## ✅ Melhorias possíveis

- Adicionar testes automatizados
- Criar interface para consultar histórico de cotações
- Exportar os dados para CSV

---

## 📄 Licença

MIT
