package kafka

import (
	"fmt"

	"github.com/c-4u/company-service/domain/service"
	"github.com/c-4u/company-service/infrastructure/db"
	"github.com/c-4u/company-service/infrastructure/external"
	"github.com/c-4u/company-service/infrastructure/repository"
)

func StartKafkaProcessor(database *db.Postgres, servers, groupId string, kafka *external.Kafka) {
	repository := repository.NewRepository(database, kafka)
	service := service.NewService(repository)

	fmt.Println("kafka consumer has been started")
	processor := NewKafkaProcessor(service, kafka)
	processor.Consume()
}
