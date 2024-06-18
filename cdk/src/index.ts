import { App, Stack, StackProps } from 'aws-cdk-lib';
import { Construct } from 'constructs';

class PaasOkStack extends Stack {
	constructor(scope: Construct, id: string, props?: StackProps) {

		super(scope, id);
		// The code that defines your stack goes here
	}
}

const app = new App();
new PaasOkStack(app, 'PaasOkStack');
