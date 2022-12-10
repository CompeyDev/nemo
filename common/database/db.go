package database

import (
	"context"
	"fmt"

	"github.com/CompeyDev/nemo/common/logger"
	"github.com/CompeyDev/nemo/db"
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

func CheckInstanceExistence(id string) bool {
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

func GetConnectedInstances() (int, error) {
	client, connectErr := ConnectDB()

	if connectErr != nil {
		DisconnectDB(client)
		logger.CustomError("DB_Manager", "Failed to initialize connection with SQLite database.")
		logger.CustomError("DB_MANAGER", fmt.Sprint(connectErr))
		// A -1 return value signifies an error.
		return -1, connectErr
	}

	ctx := context.Background()

	query, err := client.PayloadClient.FindMany().Exec(ctx)

	if err != nil {
		DisconnectDB(client)
		logger.CustomError("DB_MANAGER", "Failed to fetch connected instances information.")
		logger.CustomError("DB_MANAGER", err.Error())
		return -1, err
	}

	return len(query), nil
}

func AddToPayloadQueue(taskType string, payloadId string) error {
	client, connectErr := ConnectDB()

	if connectErr != nil {
		DisconnectDB(client)
		logger.CustomError("DB_Manager", "Failed to initialize connection with SQLite database.")
		logger.CustomError("DB_MANAGER", fmt.Sprint(connectErr))
		return connectErr
	}

	ctx := context.Background()

	_, err := client.PayloadQueue.CreateOne(
		db.PayloadQueue.AssocID.Set(payloadId),
		db.PayloadQueue.Type.Set(taskType),
		db.PayloadQueue.IsCompleted.Set(false),
	).Exec(ctx)

	if err != nil {
		DisconnectDB(client)
		logger.CustomError("DB_MANAGER", "Failed to add to payload tasks queue.")
		logger.CustomError("DB_MANAGER", err.Error())
		return err
	}

	return nil
}

func GetQueue() ([]db.PayloadQueueModel, error) {
	client, connectErr := ConnectDB()

	if connectErr != nil {
		DisconnectDB(client)
		logger.CustomError("DB_Manager", "Failed to initialize connection with SQLite database.")
		logger.CustomError("DB_MANAGER", fmt.Sprint(connectErr))
		return nil, connectErr
	}

	ctx := context.Background()

	query, err := client.PayloadQueue.FindMany().Exec(ctx)

	if err != nil {
		DisconnectDB(client)
		logger.CustomError("DB_MANAGER", "Failed to fetch active queue information.")
		logger.CustomError("DB_MANAGER", err.Error())
		return nil, err
	}

	return query, nil
}
