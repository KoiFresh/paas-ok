import { APIGatewayProxyEvent, APIGatewayProxyResult } from "aws-lambda";

async function handler(request: APIGatewayProxyEvent): Promise<APIGatewayProxyResult> {
	return {
		statusCode: 200,
		body: "pongx",
		headers: {
			"Access-Control-Allow-Origin": "*",
		},
	}
}

export { handler };
