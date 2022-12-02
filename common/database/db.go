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

	return client, nil
}

func DisconnectDB(client *db.PrismaClient) {
	defer func() {
		if err := client.Prisma.Disconnect(); err != nil {
			logger.CustomError("DB_Manager", fmt.Sprintf("Failed to close SQLite session with error %s", err.Error()))
		}
	}()
}

func CreatePayloadInstance(id string, name string) error {
	client, connectErr := ConnectDB()

	if connectErr != nil {
		DisconnectDB(client)
		logger.CustomError("DB_Manager", "Failed to initialize connection with SQLite database.")
		logger.CustomError("DB_MANAGER", connectErr.Error())
	}

	ctx := context.Background()

	_, err := client.PayloadClient.CreateOne(
		db.PayloadClient.ID.Set(id),
		db.PayloadClient.Name.Set(name),
	).Exec(ctx)

	if err != nil {
		DisconnectDB(client)
		return err
	}

	DisconnectDB(client)

	return nil

}

func UpdateCheckInTime(id string) error {
	client, connectErr := ConnectDB()

	if connectErr != nil {
		logger.CustomError("DB_Manager", "Failed to initialize connection with SQLite database.")
		logger.CustomError("DB_MANAGER", connectErr.Error())
		DisconnectDB(client)
	}

	ctx := context.Background()

	CheckInTime, err := client.PayloadClient.FindUnique(db.PayloadClient.ID.Equals(id)).Update(
		db.PayloadClient.ID.Set(id),
	).Exec(ctx)

	if err != nil {
		DisconnectDB(client)
		return err
	}

	logger.CustomInfo("DB_Manager", fmt.Sprintf("Payload last checked in at %s", CheckInTime))

	DisconnectDB(client)

	return nil

}

func CheckInstanceExistence(id string) (bool) {
	client, connectErr := ConnectDB()

	if connectErr != nil {
		DisconnectDB(client)
		logger.CustomError("DB_Manager", "Failed to initialize connection with SQLite database.")
		logger.CustomError("DB_MANAGER", fmt.Sprint(connectErr))
	}

	ctx := context.Background()

	query, err := client.PayloadClient.FindUnique(db.PayloadClient.ID.Equals(id)).Exec(ctx)

	if err != nil {
		DisconnectDB(client)
		return false
	}

	DisconnectDB(client)

	return (query == nil)
}

func Test() {
	client, connectErr := ConnectDB()

	if connectErr != nil {
		DisconnectDB(client)
		logger.CustomError("DB_Manager", "Failed to initialize connection with SQLite database.")
		logger.CustomError("DB_MANAGER", fmt.Sprint(connectErr))
	}

	ctx := context.Background()

	CreatePayloadInstance("id_testing", "name")

	d, e := client.PayloadClient.FindMany().Exec(ctx)

	if e != nil {
		fmt.Println(e)
	}

	fmt.Println(d)
}