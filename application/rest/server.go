package rest

import (
	"fmt"
	"log"

	_ "github.com/c-4u/company-service/application/rest/docs"
	_service "github.com/c-4u/company-service/domain/service"
	"github.com/c-4u/company-service/infrastructure/db"
	"github.com/c-4u/company-service/infrastructure/external"
	"github.com/c-4u/company-service/infrastructure/repository"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"go.elastic.co/apm/module/apmgin"
	"google.golang.org/grpc"
)

// @title Company Swagger API
// @version 1.0
// @description Swagger API for Golang Project Company.
// @termsOfService http://swagger.io/terms/

// @contact.name Coding4u
// @contact.email contato@coding4u.com.br

// @BasePath /api/v1
// @query.collection.format multi

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
func StartRestServer(database *db.Postgres, authConn *grpc.ClientConn, kafkaProducer *external.KafkaProducer, port int) {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	r.Use(cors.New(cors.Config{
		AllowMethods:     []string{"POST", "OPTIONS", "GET", "PUT"},
		AllowHeaders:     []string{"Content-Type", "Content-Length", "Accept-Encoding", "X-CSRF-Token", "Authorization", "Accept", "Origin", "Cache-Control", "X-Requested-With"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowAllOrigins:  true,
		AllowCredentials: true,
	}))
	r.Use(apmgin.Middleware(r))

	authService := external.NewAuthClient(authConn)
	authMiddlerare := NewAuthMiddleware(authService)
	repository := repository.NewRepository(database, kafkaProducer)
	service := _service.NewService(repository)
	restService := NewRestService(service)

	v1 := r.Group("api/v1")
	{
		v1.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
		authorized := v1.Group("/companies", authMiddlerare.Require())
		{
			authorized.POST("", restService.CreateCompany)
			authorized.GET("", restService.SearchCompanies)
			authorized.GET("/:company_id", restService.FindCompany)
			authorized.PUT("/:company_id", restService.UpdateCompany)

			employees := authorized.Group("/:company_id/employees")
			{
				employees.POST("/", restService.AddEmployeeToCompany)
				employees.POST("/:employee_id/work-scales", restService.AddWorkScaleToEmployee)
			}

			workScales := authorized.Group("/:company_id/work-scales")
			{
				workScales.POST("", restService.CreateWorkScale)
				workScales.GET("", restService.SearchWorkScales)
				workScales.GET("/:work_scale_id", restService.FindWorkScale)

				clocks := workScales.Group("/:work_scale_id/clocks")
				{
					clocks.POST("", restService.AddTimeToWorkScale)
					clocks.GET("/:clock_id", restService.FindClock)
					clocks.DELETE("/:clock_id", restService.DeleteClock)
					clocks.PUT("/:clock_id", restService.UpdateClock)
				}
			}
		}
	}

	addr := fmt.Sprintf("0.0.0.0:%d", port)
	err := r.Run(addr)
	if err != nil {
		log.Fatal("cannot start rest server", err)
	}

	log.Printf("rest server has been started on port %d", port)
}
