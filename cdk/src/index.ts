import { App, Stack, StackProps } from 'aws-cdk-lib';
import { Construct } from 'constructs';
import { UploadLambda } from './lambdas/upload';
import { StorageBucket } from './buckets/storage';
import { Gateway } from './gateways/gateway';
import { LambdaIntegration } from 'aws-cdk-lib/aws-apigateway';

class PaasOkStack extends Stack {
	constructor(scope: Construct, id: string, props?: StackProps) {
		super(scope, id);
		// The code that defines your stack goes here

		const storage = StorageBucket(this);
		const upload = UploadLambda(this, {
			bucket: storage
		});
		const gateway = Gateway(this);
		//gateway.root.addResource('upload').addMethod('POST', new LambdaIntegration(upload));
		const paths = { test: gateway.root.addResource('test') };

		paths.test.addMethod('GET', new LambdaIntegration(upload))
		paths.test.addCorsPreflight({
			allowOrigins: ['*'],
		})
	}
}

const app = new App();
new PaasOkStack(app, 'PaasOkStack');
