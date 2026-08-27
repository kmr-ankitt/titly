<script lang="ts">
	import { sendUrl } from '$lib/api';
	import { toast } from '$lib/toastStore';
	import { z, ZodError } from 'zod';

	let { onShortenSuccess } = $props<{
		onShortenSuccess: (shortCode: string, fullUrl: string, originalUrl: string) => void;
	}>();

	let rawUrl = $state('');
	let isLoading = $state(false);
	let validationError = $state('');
	let apiError = $state('');

	const urlSchema = z.string().trim().refine(
		(val) => {
			let testVal = val;
			if (!/^https?:\/\//i.test(testVal)) {
				testVal = 'https://' + testVal;
			}
			try {
				const parsed = new URL(testVal);
				return parsed.hostname.includes('.');
			} catch {
				return false;
			}
		},
		{
			message: 'Please enter a valid URL (e.g. example.com or https://example.com)'
		}
	);

	const handleSubmit = async (event: SubmitEvent) => {
		event.preventDefault();
		validationError = '';
		apiError = '';

		let trimmed = rawUrl.trim();
		if (!trimmed) {
			validationError = 'URL cannot be empty';
			return;
		}

		// Auto-prefix protocol if user omitted it
		if (!/^https?:\/\//i.test(trimmed)) {
			trimmed = 'https://' + trimmed;
		}

		try {
			urlSchema.parse(trimmed);
		} catch (e) {
			if (e instanceof ZodError) {
				validationError = e.errors[0]?.message || 'Invalid URL format';
				return;
			}
		}

		isLoading = true;
		try {
			const res = await sendUrl(trimmed);
			const shortCode = res.short_url;
			const origin = window.location.origin;
			const fullShortUrl = `${origin}/${shortCode}`;

			onShortenSuccess(shortCode, fullShortUrl, trimmed);
			toast.add('URL shortened successfully!', 'success');
		} catch (err: any) {
			apiError = err?.message || 'Failed to generate short URL. Please try again.';
			toast.add(apiError, 'error');
		} finally {
			isLoading = false;
		}
	};
</script>

<div class="w-full bg-slate-900 border border-slate-800 rounded-2xl p-6 sm:p-8 shadow-2xl">
	<form onsubmit={handleSubmit} class="flex flex-col gap-4">
		<div class="flex flex-col gap-2">
			<label for="url-input" class="text-sm font-bold text-slate-300 flex items-center justify-between">
				<span>Destination URL</span>
				<span class="text-xs font-normal text-slate-500">Supports HTTP / HTTPS</span>
			</label>
			
			<div class="relative flex items-center">
				<div class="absolute left-4 text-slate-500 pointer-events-none">
					<svg class="w-5 h-5" fill="none" viewBox="0 0 24 24" stroke="currentColor">
						<path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M13.828 10.172a4 4 0 00-5.656 0l-4 4a4 4 0 105.656 5.656l1.102-1.101m-.758-4.899a4 4 0 005.656 0l4-4a4 4 0 00-5.656-5.656l-1.1 1.1" />
					</svg>
				</div>
				<input
					id="url-input"
					type="text"
					bind:value={rawUrl}
					placeholder="https://example.com/very-long-link-path"
					disabled={isLoading}
					class="w-full pl-12 pr-4 py-3.5 bg-slate-950 border border-slate-700/80 rounded-xl text-slate-100 placeholder-slate-500 text-sm sm:text-base focus:outline-none focus:border-amber-500 focus:ring-1 focus:ring-amber-500 transition-colors disabled:opacity-60"
				/>
				{#if rawUrl}
					<button
						type="button"
						onclick={() => { rawUrl = ''; validationError = ''; apiError = ''; }}
						class="absolute right-4 text-slate-500 hover:text-slate-300 transition-colors p-1"
						aria-label="Clear input"
					>
						<svg class="w-4 h-4" fill="none" viewBox="0 0 24 24" stroke="currentColor">
							<path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M6 18L18 6M6 6l12 12" />
						</svg>
					</button>
				{/if}
			</div>

			<!-- Inline Validation Error -->
			{#if validationError}
				<div class="flex items-center gap-1.5 text-xs text-rose-400 font-medium mt-1">
					<svg class="w-4 h-4 shrink-0" fill="none" viewBox="0 0 24 24" stroke="currentColor">
						<path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 8v4m0 4h.01M21 12a9 9 0 11-18 0 9 9 0 0118 0z" />
					</svg>
					<span>{validationError}</span>
				</div>
			{/if}
		</div>

		<!-- Server API Error Alert -->
		{#if apiError}
			<div class="p-4 rounded-xl bg-rose-950/40 border border-rose-800/60 text-rose-300 text-xs sm:text-sm flex items-start gap-3">
				<svg class="w-5 h-5 text-rose-400 shrink-0 mt-0.5" fill="none" viewBox="0 0 24 24" stroke="currentColor">
					<path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 9v2m0 4h.01m-6.938 4h13.856c1.54 0 2.502-1.667 1.732-3L13.732 4c-.77-1.333-2.694-1.333-3.464 0L3.34 16c-.77 1.333.192 3 1.732 3z" />
				</svg>
				<div class="flex-1">
					<p class="font-semibold text-rose-200">Shortening Failed</p>
					<p class="mt-0.5 text-rose-300/90">{apiError}</p>
				</div>
				<button
					type="button"
					onclick={() => { apiError = ''; }}
					class="text-rose-400 hover:text-white"
					aria-label="Dismiss error"
				>
					<svg class="w-4 h-4" fill="none" viewBox="0 0 24 24" stroke="currentColor">
						<path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M6 18L18 6M6 6l12 12" />
					</svg>
				</button>
			</div>
		{/if}

		<!-- Submit Button -->
		<button
			type="submit"
			disabled={isLoading}
			class="w-full py-3.5 px-6 rounded-xl bg-amber-500 hover:bg-amber-400 active:bg-amber-600 text-slate-950 font-bold text-sm sm:text-base shadow-lg shadow-amber-500/10 flex items-center justify-center gap-2 transition-all disabled:opacity-60 disabled:cursor-not-allowed mt-2"
		>
			{#if isLoading}
				<svg class="w-5 h-5 animate-spin text-slate-950" fill="none" viewBox="0 0 24 24">
					<circle class="opacity-25" cx="12" cy="12" r="10" stroke="currentColor" stroke-width="4"></circle>
					<path class="opacity-75" fill="currentColor" d="M4 12a8 8 0 018-8V0C5.373 0 0 5.373 0 12h4zm2 5.291A7.962 7.962 0 014 12H0c0 3.042 1.135 5.824 3 7.938l3-2.647z"></path>
				</svg>
				<span>Shortening URL...</span>
			{:else}
				<span>Shorten URL</span>
				<svg class="w-4 h-4" fill="none" viewBox="0 0 24 24" stroke="currentColor">
					<path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M14 5l7 7m0 0l-7 7m7-7H3" />
				</svg>
			{/if}
		</button>
	</form>
</div>
