import { APIGatewayProxyEvent, APIGatewayProxyResult } from "aws-lambda";
import { MultipartRequest, parse } from "lambda-multipart-parser"
import { S3, DynamoDB, SQS } from "aws-sdk";
import { randomUUID } from "crypto";

const { S3_BUCKET_NAME, DYNAMODB_TABLE_NAME, QUEUE_URL, LOCALSTACK_HOSTNAME } = process.env;
const endpoint = LOCALSTACK_HOSTNAME ? `http://${LOCALSTACK_HOSTNAME}:4566` : undefined;

const s3 = new S3({ endpoint: endpoint });
const dynamodb = new DynamoDB({ endpoint: endpoint });
const sqs = new SQS({ endpoint: endpoint });

async function handler(request: APIGatewayProxyEvent): Promise<APIGatewayProxyResult> {
	let formdata: MultipartRequest
	try {
		formdata = await parse(request);
	} catch (error) {
		return {
			statusCode: 400,
			body: JSON.stringify({ "error": "Invalid request", message: error }),
			headers: {
				"Access-Control-Allow-Origin": "*",
			},
		}
	}

	const message = formdata["message"];
	const email = formdata["email"];
	const files = formdata["files"];

	const orderId = randomUUID();
	const createdAt = new Date().toISOString();

	await Promise.all([
		Promise.all(files.map(async (file) => {
			s3.putObject({
				Bucket: S3_BUCKET_NAME,
				Key: `${orderId}/${file.filename}`,
				Body: file.content,
			}).promise();
		})),

		dynamodb.putItem({
			TableName: DYNAMODB_TABLE_NAME,
			Item: {
				"id": { S: orderId },
				"message": { S: message },
				"email": { S: email },
				"status": { S: "pending" },
			},
		}).promise(),

		sqs.sendMessage({
			QueueUrl: QUEUE_URL,
			MessageBody: JSON.stringify({ "id": orderId, "createdAt": createdAt }),
		}).promise()
	]);

	return {
		statusCode: 200,
		body: JSON.stringify({ "id": orderId }),
		headers: {
			"Access-Control-Allow-Origin": "*",
		},
	}
}

export { handler };
