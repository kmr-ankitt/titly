<script lang="ts">
	import Form from '../components/Form.svelte';
	import Output from '../components/Output.svelte';
	import History from '../components/History.svelte';
	import { addHistoryItem } from '$lib/history';

	let currentFullShortUrl = $state('');
	let currentShortCode = $state('');
	let currentOriginalUrl = $state('');

	let historyRef: { refresh: () => void } | undefined = $state(undefined);

	const handleShortenSuccess = (shortCode: string, fullUrl: string, originalUrl: string) => {
		currentShortCode = shortCode;
		currentFullShortUrl = fullUrl;
		currentOriginalUrl = originalUrl;

		// Save to local history
		addHistoryItem({
			shortUrl: shortCode,
			fullShortUrl: fullUrl,
			longUrl: originalUrl
		});

		// Refresh history component
		if (historyRef && typeof historyRef.refresh === 'function') {
			historyRef.refresh();
		}
	};
</script>

<section class="flex flex-col items-center gap-10 py-6 sm:py-12">
	<!-- Hero Section -->
	<div class="text-center max-w-2xl flex flex-col items-center gap-4">
		<div class="inline-flex items-center gap-2 px-3 py-1 rounded-full bg-amber-500/10 border border-amber-500/20 text-amber-400 text-xs font-bold tracking-wide uppercase">
			<span class="w-1.5 h-1.5 rounded-full bg-amber-400"></span>
			Fast, Reliable & Modern
		</div>

		<h1 class="text-4xl sm:text-6xl font-black tracking-tight text-white leading-tight">
			Shorten Links with <span class="text-transparent bg-clip-text bg-gradient-to-r from-amber-400 via-orange-400 to-amber-500">Instant Speed</span>
		</h1>

		<p class="text-base sm:text-lg text-slate-400 font-normal leading-relaxed">
			Transform long, cluttered URLs into clean, memorable links with real-time redirection and solid reliability.
		</p>
	</div>

	<!-- Form Card -->
	<div class="w-full max-w-3xl flex flex-col gap-6">
		<Form onShortenSuccess={handleShortenSuccess} />

		<!-- Shortened Link Result Card -->
		{#if currentFullShortUrl}
			<Output
				fullShortUrl={currentFullShortUrl}
				shortCode={currentShortCode}
				originalUrl={currentOriginalUrl}
			/>
		{/if}

		<!-- Recent Links History -->
		<History bind:this={historyRef} />
	</div>
</section>
