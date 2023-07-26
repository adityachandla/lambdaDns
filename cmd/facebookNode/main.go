package main

import (
	"context"
	"encoding/json"
	"strings"

	"github.com/adityachandla/lambdaDns/utils"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

const nodeId = 22

func processRequest(req *utils.UrlRequest) utils.Response {
	parts := strings.Split(req.Url, ".")
	if len(parts) == 2 {
		return utils.Response{Status: utils.FETCHED, NodeId: nodeId}
	}
	return utils.Response{Status: utils.FAILED, NodeId: nodeId}
}

func HandleRequest(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	var urlRequest utils.UrlRequest
	json.Unmarshal([]byte(request.Body), &urlRequest)
	responseBody := processRequest(&urlRequest)
	marshaledBody, err := json.Marshal(responseBody)
	utils.Check(err)
	return events.APIGatewayProxyResponse{Body: string(marshaledBody), StatusCode: 200}, nil
}

func main() {
	lambda.Start(HandleRequest)
}
