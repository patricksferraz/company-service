/*
Copyright © 2021 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"log"
	"os"

	"github.com/c-4u/company-service/application/grpc"
	"github.com/c-4u/company-service/application/kafka"
	"github.com/c-4u/company-service/infrastructure/db"
	"github.com/c-4u/company-service/infrastructure/external"
	"github.com/c-4u/company-service/infrastructure/external/topic"
	"github.com/c-4u/company-service/utils"
	ckafka "github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/spf13/cobra"
)

// grpcCmd represents the grpc command
func NewGrpcCmd() *cobra.Command {
	var grpcPort int
	var dsn string
	var dsnType string
	var servers string
	var groupId string

	grpcCmd := &cobra.Command{
		Use:   "grpc",
		Short: "Run gRPC Service",

		Run: func(cmd *cobra.Command, args []string) {
			database, err := db.NewPostgres(dsnType, dsn)
			if err != nil {
				log.Fatal(err)
			}

			if utils.GetEnv("DB_DEBUG", "false") == "true" {
				database.Debug(true)
			}

			if utils.GetEnv("DB_MIGRATE", "false") == "true" {
				database.Migrate()
			}
			defer database.Db.Close()

			authServiceAddr := os.Getenv("AUTH_SERVICE_ADDR")
			authConn, err := external.GrpcClient(authServiceAddr)
			if err != nil {
				log.Fatal(err)
			}
			defer authConn.Close()

			kc, err := external.NewKafkaConsumer(servers, groupId, topic.CONSUMER_TOPICS)
			if err != nil {
				log.Fatal("cannot start kafka consumer", err)
			}

			deliveryChan := make(chan ckafka.Event)
			kp, err := external.NewKafkaProducer(servers, deliveryChan)
			if err != nil {
				log.Fatal("cannot start kafka producer", err)
			}

			go kp.DeliveryReport()
			go kafka.StartKafkaServer(database, kc, kp)
			grpc.StartGrpcServer(database, authConn, kp, grpcPort)
		},
	}

	dDsn := os.Getenv("DSN")
	sDsnType := os.Getenv("DSN_TYPE")
	dServers := utils.GetEnv("KAFKA_BOOTSTRAP_SERVERS", "kafka:9094")
	dGroupId := utils.GetEnv("KAFKA_CONSUMER_GROUP_ID", "employee-service")

	grpcCmd.Flags().StringVarP(&dsn, "dsn", "d", dDsn, "dsn")
	grpcCmd.Flags().StringVarP(&dsnType, "dsnType", "t", sDsnType, "dsn type")
	grpcCmd.Flags().StringVarP(&servers, "servers", "s", dServers, "kafka servers")
	grpcCmd.Flags().StringVarP(&groupId, "groupId", "i", dGroupId, "kafka group id")
	grpcCmd.Flags().IntVarP(&grpcPort, "port", "p", 50051, "gRPC Server port")

	return grpcCmd
}

func init() {
	rootCmd.AddCommand(NewGrpcCmd())

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// grpcCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// grpcCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
