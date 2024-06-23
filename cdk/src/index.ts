import { App, StackProps } from 'aws-cdk-lib';
import * as cloud from './cloud';

class AwsStack extends cloud.Stack {
	constructor(scope: App, id: string, props?: StackProps) {
		super(scope, id);

		const api = new cloud.Api(this, { cors: true });
		const bucket = new cloud.Bucket(this);
		const table = new cloud.Table(this);

		api.method("GET", "/order", new cloud.Function(this, {
			directory: "../service/functions/endpoints/order",
			environment: {
				S3_BUCKET_NAME: bucket.bucketName,
				DYNAMODB_TABLE_NAME: table.tableName,
			},
		}));

		this.outputs({
			"apiurl": api.url,
			"s3bucketname": bucket.bucketName,
			"dynamodbtablename": table.tableName,
		});
	}
}

const app = new App();
const stack = new AwsStack(app, 'paas-ok-stack');
stack.synth();
