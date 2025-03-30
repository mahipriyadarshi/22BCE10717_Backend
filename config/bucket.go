package config

import (
	"context"
	"log"
	"os"
	"github.com/joho/godotenv"


	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

var S3Client *s3.Client

func InitAWS() {
	cfg, err := config.LoadDefaultConfig(context.TODO(),
		config.WithRegion(os.Getenv("AWS_REGION")),
	)
	if err != nil {
		log.Fatal("AWS config error:", err)
	}
	S3Client = s3.NewFromConfig(cfg)
	log.Println("AWS S3 initialized")
}
