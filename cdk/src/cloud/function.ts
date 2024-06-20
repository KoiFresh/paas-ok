import { Stack } from "aws-cdk-lib";
import { AwsIntegration, LambdaIntegration } from "aws-cdk-lib/aws-apigateway";
import { v4 as uuidv4 } from 'uuid';
import { Runtime } from "aws-cdk-lib/aws-lambda";
import { NodejsFunction } from "aws-cdk-lib/aws-lambda-nodejs";
import { id } from "../utils/id";

class Function extends NodejsFunction {
	static named = id('function');

	constructor(stack: Stack, args: {
		id?: string, src: string, lockfile: string, environment?: { [key: string]: string }
	}) {
		super(stack, args.id ?? Function.named.id(), {
			entry: args.src,
			runtime: Runtime.NODEJS_20_X,
			depsLockFilePath: args.lockfile,
			environment: args.environment,
		});
	}

	public asIntegration(): AwsIntegration {
		return new LambdaIntegration(this);
	}
}

export { Function };
