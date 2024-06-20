
import * as cdk from 'aws-cdk-lib';
import { id } from '../utils/id';

class Bucket extends cdk.aws_s3.Bucket {
	static named = id('bucket');

	constructor(scope: cdk.Stack, args?: { id?: string }) {
		super(scope, args?.id ?? Bucket.named.id(), {
			removalPolicy: cdk.RemovalPolicy.DESTROY,
			autoDeleteObjects: true,
		});
	}
}

export { Bucket };
