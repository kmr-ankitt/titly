<script lang="ts">
	import { sendUrl } from '$lib/api';
	import { z, ZodError } from 'zod';
	import Output from './Output.svelte';

	let url = $state('');
	let error = $state('');

	const urlSchema = z.string().refine(
		(val) => {
			try {
				new URL(val);
				return true;
			} catch {
				return false;
			}
		},
		{
			message: 'Invalid URL format. Ensure it is a valid URL starting with http:// or https://'
		}
	);

	let shortUrl = $state('');

	const handleSubmit = async (event: Event) => {
		event.preventDefault();

		try {
			urlSchema.parse(url);
			error = '';
			const { short_url } = await sendUrl(url);
			shortUrl = short_url;
		} catch (e) {
			if (e instanceof ZodError) {
				error = e.errors[0]?.message || 'URL must start with http:// or https://';
			} else {
				error = 'Unknown error';
			}
		}
	};
</script>

<div
	class="relative w-full max-w-3xl mx-auto mt-10 aspect-video bg-white border border-gray-300 rounded-xl shadow-lg sm:aspect-auto sm:mt-6 sm:p-5 sm:max-w-3xl"
>
	<form
		onsubmit={handleSubmit}
		class="absolute inset-0 flex flex-col justify-center gap-4 p-6 bg-gray-50 rounded-xl sm:static sm:p-4"
	>
		<label for="url" class="font-bold text-lg text-gray-700 sm:text-base">URL:</label>
		<input
			id="url"
			type="text"
			bind:value={url}
			placeholder="https://www.netflix.com"
			required
			class="p-3 border border-yellow-300 rounded focus:outline-none focus:ring-2 focus:ring-yellow-500 sm:p-2"
		/>
		{#if error}
			<p class="text-red-500 text-sm sm:text-xs">{error}</p>
		{/if}
		<button
			type="submit"
			class="px-6 py-3 bg-yellow-500 text-white rounded hover:bg-yellow-400 focus:outline-none focus:ring-2 focus:ring-yellow-600 sm:px-4 sm:py-2"
		>
			Submit
		</button>
	</form>
	<Output {shortUrl} />
</div>
