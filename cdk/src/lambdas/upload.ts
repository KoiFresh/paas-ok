import { Stack } from "aws-cdk-lib";
import { Function, Runtime, Code } from "aws-cdk-lib/aws-lambda"
import { Bucket } from "aws-cdk-lib/aws-s3";

function UploadLambda(stack: Stack, args: { bucket: Bucket }): Function {
	const lambda = new Function(stack, 'upload', {
		runtime: Runtime.PROVIDED_AL2023,
		code: Code.fromAsset('../service/upload.zip'),
		handler: 'bootstrap',
		environment: {
			S3_BUCKET_NAME: args.bucket.bucketName
		}
	})

	return lambda;
}

export { UploadLambda }
