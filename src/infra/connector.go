package infra

import (
	"context"
	"log"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func Connect() (*mongo.Client, error) {
	mongoURI := os.Getenv("MONGO_URI")
	if mongoURI == "" {
		log.Println("Aviso: MONGO_URI não definida no ambiente. Verifique seu .env.")
		return nil, os.ErrNotExist
	}
	clientOptions := options.Client().ApplyURI(mongoURI)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatalf("Falha ao conectar ao MongoDB: %v", err)
	}
	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatalf("Falha ao pingar o MongoDB: %v", err)
	}
	log.Println("Infra: Conexão com o MongoDB estabelecida e ping bem-sucedido!")
	return client, nil
}
