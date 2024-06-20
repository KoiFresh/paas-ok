
import * as cdk from 'aws-cdk-lib';
import { v4 as uuidv4 } from 'uuid';

class Stack extends cdk.Stack {
	private _app: cdk.App;

	constructor(app: cdk.App, id: string, props?: cdk.StackProps) {
		super(app, id, props);
		this._app = app;
	}

	public outputs(args: { [key: string]: string }) {
		for (const [key, value] of Object.entries(args)) {
			new cdk.CfnOutput(this, uuidv4(), {
				key: key,
				value: value,
			});
		}
	}

	public synth() {
		this._app.synth();
	}
}

export { Stack };
