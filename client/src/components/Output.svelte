<script lang="ts">
	import { toast } from '$lib/toastStore';

	let {
		fullShortUrl = '',
		shortCode = '',
		originalUrl = ''
	} = $props<{
		fullShortUrl?: string;
		shortCode?: string;
		originalUrl?: string;
	}>();

	let copied = $state(false);
	let showQr = $state(false);

	const copyToClipboard = async () => {
		if (!fullShortUrl) return;
		try {
			await navigator.clipboard.writeText(fullShortUrl);
			copied = true;
			toast.add('Copied to clipboard!', 'success');
			setTimeout(() => {
				copied = false;
			}, 2500);
		} catch {
			toast.add('Failed to copy to clipboard', 'error');
		}
	};

	// Generate QR Code URL via public API image
	let qrImageUrl = $derived(
		fullShortUrl
			? `https://api.qrserver.com/v1/create-qr-code/?size=200x200&data=${encodeURIComponent(fullShortUrl)}`
			: ''
	);
</script>

{#if fullShortUrl}
	<div class="w-full bg-slate-900 border border-amber-500/30 rounded-2xl p-6 sm:p-8 shadow-2xl animate-fade-in flex flex-col gap-6">
		<!-- Success Badge Header -->
		<div class="flex items-center justify-between border-b border-slate-800 pb-4">
			<div class="flex items-center gap-2 text-emerald-400 font-semibold text-sm">
				<span class="w-2.5 h-2.5 rounded-full bg-emerald-400 animate-pulse"></span>
				<span>URL Shortened Successfully!</span>
			</div>
			<button
				type="button"
				onclick={() => (showQr = !showQr)}
				class="text-xs text-amber-400 hover:text-amber-300 font-medium flex items-center gap-1.5 px-3 py-1.5 rounded-lg bg-amber-500/10 border border-amber-500/20 transition-colors"
			>
				<svg class="w-4 h-4" fill="none" viewBox="0 0 24 24" stroke="currentColor">
					<path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 4v1m6 11h2m-6 0h-2v4m0-11v3m0 0h.01M12 12h4.01M16 20h4M4 12h4m12 0h.01M5 8h2a1 1 0 001-1V5a1 1 0 00-1-1H5a1 1 0 00-1 1v2a1 1 0 001 1zm12 0h2a1 1 0 001-1V5a1 1 0 00-1-1h-2a1 1 0 00-1 1v2a1 1 0 001 1zM5 20h2a1 1 0 001-1v-2a1 1 0 00-1-1H5a1 1 0 00-1 1v2a1 1 0 001 1z" />
				</svg>
				<span>{showQr ? 'Hide QR Code' : 'Show QR Code'}</span>
			</button>
		</div>

		<!-- Shortened Link Box -->
		<div class="flex flex-col sm:flex-row items-stretch sm:items-center justify-between gap-3 bg-slate-950 p-4 rounded-xl border border-slate-800">
			<div class="flex flex-col overflow-hidden min-w-0 pr-2">
				<span class="text-xs text-slate-400 font-medium">Your Short URL:</span>
				<a
					href={fullShortUrl}
					target="_blank"
					rel="noopener noreferrer"
					class="text-amber-400 font-bold text-base sm:text-lg hover:underline truncate"
				>
					{fullShortUrl}
				</a>
			</div>

			<div class="flex items-center gap-2 shrink-0">
				<button
					type="button"
					onclick={copyToClipboard}
					class="flex-1 sm:flex-initial px-4 py-2.5 rounded-lg text-xs font-bold transition-all flex items-center justify-center gap-1.5
					{copied
						? 'bg-emerald-500 text-slate-950'
						: 'bg-amber-500 hover:bg-amber-400 text-slate-950'}"
				>
					{#if copied}
						<svg class="w-4 h-4" fill="none" viewBox="0 0 24 24" stroke="currentColor">
							<path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M5 13l4 4L19 7" />
						</svg>
						<span>Copied!</span>
					{:else}
						<svg class="w-4 h-4" fill="none" viewBox="0 0 24 24" stroke="currentColor">
							<path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M8 5H6a2 2 0 00-2 2v12a2 2 0 002 2h10a2 2 0 002-2v-1M8 5a2 2 0 002 2h2a2 2 0 002-2M8 5a2 2 0 012-2h2a2 2 0 012 2m0 0h2a2 2 0 012 2v3m2 4H10m0 0l3-3m-3 3l3 3" />
						</svg>
						<span>Copy Link</span>
					{/if}
				</button>

				<a
					href={fullShortUrl}
					target="_blank"
					rel="noopener noreferrer"
					class="p-2.5 rounded-lg bg-slate-800 hover:bg-slate-700 text-slate-300 transition-colors flex items-center justify-center"
					title="Test Link in New Tab"
					aria-label="Open link in new tab"
				>
					<svg class="w-4 h-4" fill="none" viewBox="0 0 24 24" stroke="currentColor">
						<path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M10 6H6a2 2 0 00-2 2v10a2 2 0 002 2h10a2 2 0 002-2v-4M14 4h6m0 0v6m0-6L10 14" />
					</svg>
				</a>
			</div>
		</div>

		<!-- Destination Info -->
		{#if originalUrl}
			<div class="flex items-center gap-2 text-xs text-slate-400 overflow-hidden">
				<span class="font-medium shrink-0">Redirects to:</span>
				<span class="truncate text-slate-300 font-mono bg-slate-950 px-2 py-1 rounded border border-slate-800">
					{originalUrl}
				</span>
			</div>
		{/if}

		<!-- Optional QR Code Display -->
		{#if showQr && qrImageUrl}
			<div class="flex flex-col items-center justify-center p-6 bg-slate-950 border border-slate-800 rounded-xl gap-3">
				<img src={qrImageUrl} alt="QR Code for short URL" class="w-44 h-44 rounded-lg bg-white p-2" />
				<span class="text-xs text-slate-400 font-medium">Scan QR code to open link</span>
			</div>
		{/if}
	</div>
{/if}

<style>
	@keyframes fadeIn {
		from { opacity: 0; transform: translateY(10px); }
		to { opacity: 1; transform: translateY(0); }
	}
	.animate-fade-in {
		animation: fadeIn 0.3s cubic-bezier(0.16, 1, 0.3, 1) forwards;
	}
</style>
