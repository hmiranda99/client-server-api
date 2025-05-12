package server

import (
	"context"
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"time"

	_ "modernc.org/sqlite"
)

type Cotacao struct {
	Bid string `json:"bid"`
}

type ResponseAPI struct {
	USDBRL Cotacao `json:"USDBRL"`
}

func StartServer() {
	http.HandleFunc("/cotacao", handlerCotacao)
	log.Println("Servidor rodando na porta 8080...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func handlerCotacao(w http.ResponseWriter, r *http.Request) {
	ctxAPI, cancelAPI := context.WithTimeout(r.Context(), 1000*time.Millisecond)
	defer cancelAPI()

	req, err := http.NewRequestWithContext(ctxAPI, http.MethodGet, "https://economia.awesomeapi.com.br/json/last/USD-BRL", nil)
	if err != nil {
		http.Error(w, "Erro ao criar request externa", http.StatusInternalServerError)
		log.Println("Erro request:", err)
		return
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		http.Error(w, "Erro ao buscar cotação", http.StatusGatewayTimeout)
		log.Println("Erro client:", err)
		return
	}
	defer resp.Body.Close()

	var data ResponseAPI
	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		http.Error(w, "Erro ao decodificar JSON", http.StatusInternalServerError)
		log.Println("Erro decode:", err)
		return
	}

	go salvarCotacao(data.USDBRL.Bid)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"bid": data.USDBRL.Bid})
}

func salvarCotacao(bid string) {
	db, err := sql.Open("sqlite", "./cotacoes.db")
	if err != nil {
		log.Println("Erro ao abrir DB:", err)
		return
	}
	defer db.Close()

	createTable := `CREATE TABLE IF NOT EXISTS cotacoes (id INTEGER PRIMARY KEY, bid TEXT);`
	if _, err := db.Exec(createTable); err != nil {
		log.Println("Erro ao criar tabela:", err)
		return
	}

	insert := `INSERT INTO cotacoes (bid) VALUES (?)`
	if _, err := db.Exec(insert, bid); err != nil {
		log.Println("Erro ao inserir cotação:", err)
	}
}
