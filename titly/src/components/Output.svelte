<script>
	export let shortUrl;
</script>

<div class="flex items-center justify-center">
	{#if shortUrl}
		<p class="mt-4 text-green-600 font-semibold">
			Shortened URL: <a
				href={shortUrl}
				class="underline text-yellow-500"
				target="_blank"
				rel="noopener noreferrer">{shortUrl}</a
			>
		</p>
	{/if}
</div>
