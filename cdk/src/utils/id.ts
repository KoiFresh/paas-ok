
class Named {
	private _counter: number = 0;
	private _prefix: string;

	constructor(prefix: string) {
		this._prefix = prefix;
	}

	public id() {
		this._counter++;
		return `${this._prefix}-${this._counter}`;
	};
}

function id(prefix: string) {
	return new Named(prefix);
}

export { id, Named };
