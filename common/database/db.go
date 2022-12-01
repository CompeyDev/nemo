package database

import (
	"context"
	"fmt"

	"github.com/CompeyDev/nemo/common/logger"
	"github.com/CompeyDev/nemo/prisma/db"
)

func ConnectDB() (*db.PrismaClient, error) {
	client := db.NewClient()
	if err := client.Prisma.Connect(); err != nil {
		return nil, err
	}

	defer func() {
		if err := client.Prisma.Disconnect(); err != nil {
			logger.CustomError("DB_Manager", fmt.Sprintf("Failed to open SQLite session with error %s", err.Error()))
		}
	}()

	return client, nil
}

func CreatePayloadInstance(id string, name string) error {
	client, connectErr := ConnectDB()

	if connectErr != nil {
		logger.CustomError("DB_Manager", "Failed to initialize connection with SQLite database.")
	}

	ctx := context.Background()

	_, err := client.PayloadClient.CreateOne(
		db.PayloadClient.ID.Set(id),
		db.PayloadClient.Name.Set(name),
	).Exec(ctx)

	if err != nil {
		return err
	}

	return nil

}

func UpdateCheckInTime(id string) error {
	client, connectErr := ConnectDB()

	if connectErr != nil {
		logger.CustomError("DB_Manager", "Failed to initialize connection with SQLite database.")
	}

	ctx := context.Background()

	CheckInTime, err := client.PayloadClient.FindUnique(db.PayloadClient.ID.Equals(id)).Update(
		db.PayloadClient.ID.Set(id),
	).Exec(ctx)

	if err != nil {
		return err
	}

	logger.CustomInfo("DB_Manager", fmt.Sprintf("Payload last checked in at %s", CheckInTime))

	return nil

}

func CheckInstanceExistence(id string) (bool, error) {
	client, connectErr := ConnectDB()

	if connectErr != nil {
		logger.CustomError("DB_Manager", "Failed to initialize connection with SQLite database.")
	}

	ctx := context.Background()

	query, err := client.PayloadClient.FindUnique(db.PayloadClient.ID.Equals(id)).Exec(ctx)

	if err != nil {
		return false, err
	}

	return (query == nil), nil
}
