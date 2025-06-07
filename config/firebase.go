package config

import (
	"context"
	"log"

	"cloud.google.com/go/firestore"
	firebase "firebase.google.com/go/v4"
	"google.golang.org/api/option"
)

// FirestoreClient inicializa y retorna un cliente de Firestore
func FirestoreClient() (*firestore.Client, error) {
	ctx := context.Background()
	// Ruta al archivo de credenciales proporcionado
	opt := option.WithCredentialsFile("./serviceAccountKey.json")
	app, err := firebase.NewApp(ctx, nil, opt)
	if err != nil {
		log.Fatalf("Error inicializando la app de Firebase: %v", err)
		return nil, err
	}

	client, err := app.Firestore(ctx)
	if err != nil {
		log.Fatalf("Error inicializando Firestore: %v", err)
		return nil, err
	}

	return client, nil
}