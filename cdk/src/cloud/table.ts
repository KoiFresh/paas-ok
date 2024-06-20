import * as cdk from "aws-cdk-lib";
import { v4 as uuidv4 } from "uuid";
import { id } from "../utils/id";

class Table extends cdk.aws_dynamodb.Table {
	static named = id('table');

	constructor(scope: cdk.Stack, args?: { partitionKey?: cdk.aws_dynamodb.Attribute }) {
		super(scope, Table.named.id(), {
			partitionKey: args?.partitionKey ?? { name: 'id', type: cdk.aws_dynamodb.AttributeType.STRING },
			removalPolicy: cdk.RemovalPolicy.DESTROY,
		});
	}
}

export { Table };
