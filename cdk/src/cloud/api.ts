import { Stack } from "aws-cdk-lib";
import { IResource, RestApi } from "aws-cdk-lib/aws-apigateway";
import { Function } from "./function"
import { id } from "../utils/id";

class Api extends RestApi {
	static named = id('api');

	constructor(stack: Stack, args: { cors: boolean }) {
		super(stack, Api.named.id());

		if (args.cors) {
			this.root.addCorsPreflight({
				allowOrigins: ['*'],
			});
		}
	}

	public method(method: "GET" | "POST" | "PUT", path: string, func: Function) {
		const parts = path.split('/').filter(part => part);

		let ressource: IResource = this.root;
		for (const part of parts) {
			ressource = ressource.getResource(part) ?? ressource.addResource(part, {

			});
		}

		ressource.addMethod(method, func.asIntegration());
	}
}

export { Api };
