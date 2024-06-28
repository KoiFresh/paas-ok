
function useRemoteSlicer() {
	const config = useRuntimeConfig()

	const isSlicing = ref(false);

	async function slice(file: File): Promise<RemoteSlicerResult> {
		isSlicing.value = true;

		const form = new FormData();
		form.append('files', file);

		try {
			const response = await fetch(`${config.public.apiBaseUrl}/cura:slice`, {
				method: 'POST',
				body: form
			});

			return response.json();
		} catch (error) {
			throw "An error occurred while slicing the file. Try again later."
		}
		finally {
			isSlicing.value = false;
		}
	}

	return { slice, isSlicing };
}
export { useRemoteSlicer };
