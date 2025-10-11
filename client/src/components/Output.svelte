<script>
	export let shortUrl;
</script>

<div
	class="flex items-center pl-4 absolute bg-gray-50 rounded-xl sm:static p-5 w-full"
>
	{#if shortUrl}
		<p class="mt-4 text-green-600 font-semibold text-center sm:text-left">
			Shortened URL: <a
				href={shortUrl}
				class="underline text-yellow-500 break-words"
				target="_blank"
				rel="noopener noreferrer">{shortUrl}</a
			>
		</p>
	{/if}
</div>
