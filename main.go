package main

import (
	"context"
	"flag"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	ginadapter "github.com/awslabs/aws-lambda-go-api-proxy/gin"
	"github.com/cakoshakib/curate-backend/handlers"
	"github.com/gin-gonic/gin"
)

var ginLambda *ginadapter.GinLambda

func Handler(ctx context.Context, req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	// If no name is provided in the HTTP request body, throw an error
	return ginLambda.ProxyWithContext(ctx, req)
}

func main() {

	router := gin.Default()
	router.GET("/hello", handlers.HelloHandler)

	// dev mode then start locally
	dev := flag.Bool("dev", false, "")
	flag.Parse()
	if *dev {
		router.Run()
		return
	}

	ginLambda = ginadapter.New(router)
	lambda.Start(Handler)
}
