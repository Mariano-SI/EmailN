package main

import (
	"emailn/internal/domain/campaign"
	"emailn/internal/infrastructure/database"
	"emailn/internal/infrastructure/mail"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	if os.Getenv("ENV") != "production" {
		if err := godotenv.Load("../../.env"); err != nil {
			log.Fatal("Error loading .env file")
		}
	}

	db := database.NewDb()
	campaignRepository := database.CampaignRepository{Db: db}
	campaignService := campaign.ServiceImp{Repository: &campaignRepository, SendMail: mail.SendMail}
	campaigns, err := campaignRepository.GetCampaignsToBeSent()
	if err != nil {
		fmt.Println(err.Error())
	}

	for _, campaign := range campaigns {
		campaignService.SendEmailAndUpdateStatus(&campaign)
	}
}
