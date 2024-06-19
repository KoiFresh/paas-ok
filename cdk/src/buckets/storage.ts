import { RemovalPolicy, Stack } from "aws-cdk-lib";
import { Bucket } from "aws-cdk-lib/aws-s3";


function StorageBucket(stack: Stack): Bucket {
	const bucket = new Bucket(stack, 'storage', {
		removalPolicy: RemovalPolicy.DESTROY,
		autoDeleteObjects: true,
	});

	return bucket;
}

export { StorageBucket }