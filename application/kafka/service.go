package kafka

import (
	"context"
	"fmt"

	ckafka "github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/patricksferraz/company-service/application/kafka/schema"
	"github.com/patricksferraz/company-service/domain/service"
	"github.com/patricksferraz/company-service/infrastructure/external"
	"github.com/patricksferraz/company-service/infrastructure/external/topic"
)

type KafkaProcessor struct {
	Service *service.Service
	Kc      *external.KafkaConsumer
}

func NewKafkaProcessor(service *service.Service, kafkaConsumer *external.KafkaConsumer) *KafkaProcessor {
	return &KafkaProcessor{
		Service: service,
		Kc:      kafkaConsumer,
	}
}

func (p *KafkaProcessor) Consume() {
	p.Kc.Consumer.SubscribeTopics(p.Kc.ConsumerTopics, nil)
	for {
		msg, err := p.Kc.Consumer.ReadMessage(-1)
		if err == nil {
			// fmt.Println(string(msg.Value))
			p.processMessage(msg)
		}
	}
}

func (p *KafkaProcessor) processMessage(msg *ckafka.Message) {
	switch _topic := *msg.TopicPartition.Topic; _topic {
	case topic.NEW_EMPLOYEE:
		// TODO: add fault tolerance
		err := p.createEmployee(msg)
		if err != nil {
			fmt.Println("creation error ", err)
		}
	default:
		fmt.Println("not a valid topic", string(msg.Value))
	}
}

func (p *KafkaProcessor) createEmployee(msg *ckafka.Message) error {
	employeeEvent := schema.NewEmployeeEvent()
	err := employeeEvent.ParseJson(msg.Value)
	if err != nil {
		return err
	}

	err = p.Service.CreateEmployee(context.TODO(), employeeEvent.Employee.ID)
	if err != nil {
		return err
	}

	return nil
}
