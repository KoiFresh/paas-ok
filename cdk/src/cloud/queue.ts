import * as  sqs from "aws-cdk-lib/aws-sqs";
import { Stack } from "aws-cdk-lib/core";
import { id } from "../utils/id";


class Queue extends sqs.Queue {
	static named = id("queue");

	constructor(stack: Stack) {
		super(stack, Queue.named.id());
	}
}

export { Queue };
