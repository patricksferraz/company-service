package kafka

import (
	"fmt"

	"github.com/c-4u/company-service/domain/service"
	"github.com/c-4u/company-service/infrastructure/db"
	"github.com/c-4u/company-service/infrastructure/external"
	"github.com/c-4u/company-service/infrastructure/repository"
)

func StartKafkaServer(database *db.Postgres, kafkaConsumer *external.KafkaConsumer, kafkaProducer *external.KafkaProducer) {
	repository := repository.NewRepository(database, kafkaProducer)
	service := service.NewService(repository)

	fmt.Println("kafka consumer has been started")
	processor := NewKafkaProcessor(service, kafkaConsumer)
	processor.Consume()
}
