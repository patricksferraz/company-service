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

	"github.com/c-4u/company-service/application/kafka"
	"github.com/c-4u/company-service/application/rest"
	"github.com/c-4u/company-service/infrastructure/db"
	"github.com/c-4u/company-service/infrastructure/external"
	"github.com/c-4u/company-service/infrastructure/external/topic"
	"github.com/c-4u/company-service/utils"
	ckafka "github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/spf13/cobra"
)

// NewRestCmd represents the rest command
func NewRestCmd() *cobra.Command {
	var restPort int
	var dsn string
	var dsnType string
	var servers string
	var groupId string
	var restCmd = &cobra.Command{
		Use:   "rest",
		Short: "Run rest Service",

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

			deliveryChan := make(chan ckafka.Event)
			k, err := external.NewKafka(servers, groupId, []string{topic.NEW_EMPLOYEE}, deliveryChan)
			if err != nil {
				log.Fatal(err)
			}

			go k.DeliveryReport()
			go kafka.StartKafkaProcessor(database, servers, groupId, k)
			rest.StartRestServer(database, authConn, k, restPort)
		},
	}

	dDsn := os.Getenv("DSN")
	sDsnType := os.Getenv("DSN_TYPE")
	dServers := utils.GetEnv("KAFKA_BOOTSTRAP_SERVERS", "kafka:9094")
	dGroupId := utils.GetEnv("KAFKA_CONSUMER_GROUP_ID", "company-service")

	restCmd.Flags().StringVarP(&dsn, "dsn", "d", dDsn, "dsn")
	restCmd.Flags().StringVarP(&dsnType, "dsnType", "t", sDsnType, "dsn type")
	restCmd.Flags().StringVarP(&servers, "servers", "s", dServers, "kafka servers")
	restCmd.Flags().StringVarP(&groupId, "groupId", "i", dGroupId, "kafka group id")
	restCmd.Flags().IntVarP(&restPort, "port", "p", 8080, "rest server port")

	return restCmd
}

func init() {
	rootCmd.AddCommand(NewRestCmd())

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// restCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// restCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
