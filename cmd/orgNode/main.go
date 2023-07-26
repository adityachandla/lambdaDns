package main

import (
	"context"
	"encoding/json"
	"log"
	"strings"

	"github.com/adityachandla/lambdaDns/utils"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

const nodeId = 12

var mapper = map[string]string{
	"google":   "googleNode",
	"facebook": "facebookNode",
}

func processRequest(req *utils.UrlRequest) utils.Response {
	log.Printf("Processing %s at orgNode", req.Url)
	parts := strings.Split(req.Url, ".")
	if len(parts) == 1 {
		return utils.Response{Status: utils.FETCHED, NodeId: nodeId}
	}
	lastPart := parts[len(parts)-1]
	remaining := strings.Join(parts[:len(parts)-1], ".")
	for k, v := range mapper {
		if k == lastPart {
			response := utils.GetRedirectResponse(v, remaining)
			return utils.Response{Status: utils.REDIRECT, Redirect: response, NodeId: nodeId}
		}
	}
	return utils.Response{Status: utils.FAILED, NodeId: nodeId}
}

func HandleRequest(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	var urlRequest utils.UrlRequest
	json.Unmarshal([]byte(request.Body), &urlRequest)
	responseBody := processRequest(&urlRequest)
	marshaledBody, err := json.Marshal(&responseBody)
	utils.Check(err)
	return events.APIGatewayProxyResponse{Body: string(marshaledBody), StatusCode: 200}, nil
}

func main() {
	lambda.Start(HandleRequest)
}
