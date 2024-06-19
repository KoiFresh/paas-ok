import { Stack } from "aws-cdk-lib";
import { RestApi } from "aws-cdk-lib/aws-apigateway";
import { ApiGateway } from "aws-cdk-lib/aws-events-targets";

function Gateway(stack: Stack): RestApi {
	const gateway = new RestApi(stack, 'gateway')

	return gateway;
}

export { Gateway }