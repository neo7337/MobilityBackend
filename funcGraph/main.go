package main

import (
	"bytes"
	"encoding/json"
	"graph/handler"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

// Response ->
type Response events.APIGatewayProxyResponse

func main() {
	lambda.Start(GraphHandler)
}

// GraphHandler =>
func GraphHandler(request events.APIGatewayProxyRequest) (Response, error) {
	param1 := request.QueryStringParameters["country"]
	param2 := request.QueryStringParameters["region"]

	finalList := handler.APIHandler(param1, param2)
	var buf bytes.Buffer
	body, err := json.Marshal(finalList)
	if err != nil {
		return Response{StatusCode: 404}, err
	}
	json.HTMLEscape(&buf, body)
	resp := Response{
		StatusCode:      200,
		IsBase64Encoded: false,
		Body:            buf.String(),
		Headers: map[string]string{
			"Content-Type":                "application/json",
			"X-MyCompany-Func-Reply":      "hello-handler",
			"Access-Control-Allow-Origin": "*",
		},
	}
	return resp, nil
}
