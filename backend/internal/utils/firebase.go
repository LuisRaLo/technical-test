package utils

import (
	"context"
	"encoding/json"
	"fmt"
	"os"

	firebase "firebase.google.com/go"
	"firebase.google.com/go/auth"
	"go.uber.org/zap"

	"google.golang.org/api/option"
)

type ResultFirebase struct {
	TokenData *auth.Token `json:"token"`
}

type FirebaseConfig struct {
	Type                    string `json:"type"`
	ProjectID               string `json:"project_id"`
	PrivateKeyID            string `json:"private_key_id"`
	PrivateKey              string `json:"private_key"`
	ClientEmail             string `json:"client_email"`
	ClientID                string `json:"client_id"`
	AuthURI                 string `json:"auth_uri"`
	TokenURI                string `json:"token_uri"`
	AuthProviderX509CertUrl string `json:"auth_provider_x509_cert_url"`
	ClientX509CertUrl       string `json:"client_x509_cert_url"`
	UniverseDomain          string `json:"universe_domain"`
}

// GetFirebaseSession gets a firebase session
func GetFirebaseSession(context context.Context, logger *zap.SugaredLogger) (*auth.Client, error) {
	configData, err := firebaseConfigData(logger)
	if err != nil {
		logger.Errorln("error getting config data: %v", err)
		return nil, err
	}

	opt := option.WithCredentialsJSON(configData)

	app, err := firebase.NewApp(context, nil, opt)
	if err != nil {
		logger.Errorln("error initializing app: %v", err)
		return nil, err
	}

	client, err := app.Auth(context)
	if err != nil {
		logger.Errorln("error getting Auth client: %v", err)
		return nil, err
	}

	return client, nil
}

// GetUserDataByToken gets user data by token
func GetUserDataByToken(ctx context.Context, logger *zap.SugaredLogger, token string) (ResultFirebase, error) {
	var result ResultFirebase = ResultFirebase{}

	client, err := GetFirebaseSession(ctx, logger)
	if err != nil {
		logger.Errorln("error getting Auth client: %v", err)
		return result, err
	}

	tokenData, err := client.VerifyIDToken(ctx, token)
	if err != nil {
		logger.Errorln("error verifying token: %v", err.Error())
		return result, fmt.Errorf("token is invalid or expired, please login again")
	}

	result.TokenData = tokenData

	return result, nil
}

func firebaseConfigData(logger *zap.SugaredLogger) ([]byte, error) {
	var config FirebaseConfig = FirebaseConfig{

		Type:                    os.Getenv("FIREBASE_TYPE"),
		ProjectID:               os.Getenv("FIREBASE_PROJECT_ID"),
		PrivateKeyID:            os.Getenv("FIREBASE_PRIVATE_KEY_ID"),
		PrivateKey:              os.Getenv("FIREBASE_PRIVATE_KEY"),
		ClientEmail:             os.Getenv("FIREBASE_CLIENT_EMAIL"),
		ClientID:                os.Getenv("FIREBASE_CLIENT_ID"),
		AuthURI:                 os.Getenv("FIREBASE_AUTH_URI"),
		TokenURI:                os.Getenv("FIREBASE_TOKEN_URI"),
		AuthProviderX509CertUrl: os.Getenv("FIREBASE_AUTH_PROVIDER_X509_CERT_URL"),
		ClientX509CertUrl:       os.Getenv("FIREBASE_CLIENT_X509_CERT_URL"),
		UniverseDomain:          os.Getenv("FIREBASE_UNIVERSE_DOMAIN"),
	}

	configData, err := json.Marshal(config)
	if err != nil {
		logger.Errorln("error marshalling config data: %v", err)
		return nil, err
	}

	return configData, nil
}
