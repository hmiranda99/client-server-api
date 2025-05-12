package client

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

func StartClient() {
	ctx, cancel := context.WithTimeout(context.Background(), 300*time.Millisecond)
	defer cancel()

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, "http://localhost:8080/cotacao", nil)
	if err != nil {
		log.Fatal("Erro ao criar request:", err)
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatal("Erro ao fazer request:", err)
	}
	defer resp.Body.Close()

	log.Println("Status da resposta:", resp.Status)
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal("Erro ao ler corpo da resposta:", err)
	}

	log.Println("Corpo da resposta:", string(body))

	var data map[string]interface{}
	if err := json.Unmarshal(body, &data); err != nil {
		log.Fatal("Erro ao decodificar JSON:", err)
	}

	bid := data["bid"].(string)

	err = ioutil.WriteFile("cotacao.txt", []byte("Dólar: "+bid), 0644)
	if err != nil {
		log.Fatal("Erro ao salvar arquivo:", err)
	}

	log.Println("Cotação salva com sucesso:", bid)
}
