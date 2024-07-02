
function useRemoteSlicer() {
	const config = useRuntimeConfig()

	const progress = ref(false);

	async function slice(file: File, options: SliceOptions): Promise<RemoteSlicerResult> {
		progress.value = true;

		const form = new FormData();
		form.append('files', file);
		form.append('color', options.color);
		form.append('quality', options.quality);
		form.append('material', options.material);
		form.append('infill', options.infill.toFixed(0));

		const response = await fetch(`${config.public.apiBaseUrl}/cura:slice`, {
			method: 'POST',
			body: form
		}).catch(() => {
			throw "An error occurred while slicing the file. Try again later."
		}).finally(() => {
			progress.value = false;
		});

		if (!response.ok) {
			throw await response.text();
		}

		return response.json();
	}

	return { slice, progress };
}
export { useRemoteSlicer };
