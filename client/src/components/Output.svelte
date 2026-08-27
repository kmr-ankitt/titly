<script>
	export let new_url;
</script>

<div
	class="flex items-center pl-4 absolute bg-gray-50 rounded-xl sm:static p-5 w-full"
>
	{#if new_url}
		<p class="mt-4 text-green-600 font-semibold text-center sm:text-left">
			Shortened URL: <a
				href={new_url}
				class="underline text-yellow-500 break-words"
				target="_blank"
				rel="noopener noreferrer">{new_url}</a
			>
		</p>
	{/if}
</div>
