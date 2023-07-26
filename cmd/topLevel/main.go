package main

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"

	"github.com/adityachandla/lambdaDns/utils"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

const nodeId = 0

var mapper = map[string]string{
	"org": "orgNode",
	"com": "comNode",
}

func processRequest(req *utils.UrlRequest) utils.Response {
	if strings.Index(req.Url, "https://") != -1 {
		req.Url = strings.TrimPrefix(req.Url, "https://")
	}
	if strings.Index(req.Url, "www.") != -1 {
		req.Url = strings.TrimPrefix(req.Url, "www.")
	}
	parts := strings.Split(req.Url, ".")
	lastPart := parts[len(parts)-1]
	remaining := strings.Join(parts[:len(parts)-1], ".")
	for k, v := range mapper {
		if k == lastPart {
			response := utils.GetRedirectResponse(v, remaining)
			fmt.Printf("Invoking %s with %s\n", v, remaining)
			return utils.Response{Status: utils.REDIRECT, Redirect: response, NodeId: nodeId}
		}
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
