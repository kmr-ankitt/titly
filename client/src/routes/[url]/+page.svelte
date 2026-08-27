<script lang="ts">
	import { onMount } from 'svelte';
	import type { RedirectLoadData } from './+page';

	let { data } = $props<{ data: RedirectLoadData }>();

	let countdown = $state(2);
	let isRedirecting = $state(false);

	const executeRedirect = () => {
		const destination = data.targetUrl || data.backendUrl;
		if (destination) {
			isRedirecting = true;
			window.location.href = destination;
		}
	};

	onMount(() => {
		// If valid redirect target exists, trigger redirect countdown
		if (!data.errorStatus) {
			const timer = setInterval(() => {
				if (countdown > 1) {
					countdown--;
				} else {
					clearInterval(timer);
					executeRedirect();
				}
			}, 800);

			return () => clearInterval(timer);
		}
	});
</script>

<svelte:head>
	<title>{data.errorStatus ? '404 - Link Not Found' : 'Redirecting... — Titly'}</title>
</svelte:head>

<div class="max-w-xl mx-auto py-12 flex flex-col items-center justify-center">
	{#if !data.errorStatus}
		<!-- Redirection Progress Card -->
		<div class="w-full bg-slate-900 border border-slate-800 rounded-2xl p-8 shadow-2xl flex flex-col items-center text-center gap-6">
			<div class="w-16 h-16 rounded-2xl bg-amber-500/10 border border-amber-500/20 flex items-center justify-center text-amber-400">
				<svg class="w-8 h-8 animate-bounce" fill="none" viewBox="0 0 24 24" stroke="currentColor">
					<path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M14 5l7 7m0 0l-7 7m7-7H3" />
				</svg>
			</div>

			<div class="flex flex-col gap-2">
				<span class="text-xs uppercase font-bold tracking-wider text-amber-400">Redirecting Link</span>
				<h1 class="text-2xl font-black text-white">Taking you to your destination...</h1>
				<p class="text-sm text-slate-400">
					Short code: <span class="font-mono text-amber-300 font-bold">/{data.shortCode}</span>
				</p>
			</div>

			{#if data.targetUrl}
				<div class="w-full bg-slate-950 p-4 rounded-xl border border-slate-800 flex flex-col gap-1 text-left">
					<span class="text-xs text-slate-500 font-semibold">Target URL:</span>
					<span class="text-sm font-mono text-slate-200 truncate">{data.targetUrl}</span>
				</div>
			{/if}

			<div class="flex flex-col items-center gap-3 w-full">
				<button
					type="button"
					onclick={executeRedirect}
					class="w-full py-3 px-6 rounded-xl bg-amber-500 hover:bg-amber-400 text-slate-950 font-bold text-sm transition-all flex items-center justify-center gap-2"
				>
					{#if isRedirecting}
						<svg class="w-4 h-4 animate-spin text-slate-950" fill="none" viewBox="0 0 24 24">
							<circle class="opacity-25" cx="12" cy="12" r="10" stroke="currentColor" stroke-width="4"></circle>
							<path class="opacity-75" fill="currentColor" d="M4 12a8 8 0 018-8V0C5.373 0 0 5.373 0 12h4zm2 5.291A7.962 7.962 0 014 12H0c0 3.042 1.135 5.824 3 7.938l3-2.647z"></path>
						</svg>
						<span>Redirecting now...</span>
					{:else}
						<span>Continue immediately ({countdown}s)</span>
						<svg class="w-4 h-4" fill="none" viewBox="0 0 24 24" stroke="currentColor">
							<path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M14 5l7 7m0 0l-7 7m7-7H3" />
						</svg>
					{/if}
				</button>

				<a href="/" class="text-xs text-slate-400 hover:text-white transition-colors">
					Cancel and return to home
				</a>
			</div>
		</div>
	{:else}
		<!-- Error / 404 Card -->
		<div class="w-full bg-slate-900 border border-rose-500/30 rounded-2xl p-8 shadow-2xl flex flex-col items-center text-center gap-6">
			<div class="w-16 h-16 rounded-2xl bg-rose-500/10 border border-rose-500/20 flex items-center justify-center text-rose-400">
				<svg class="w-8 h-8" fill="none" viewBox="0 0 24 24" stroke="currentColor">
					<path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 9v2m0 4h.01m-6.938 4h13.856c1.54 0 2.502-1.667 1.732-3L13.732 4c-.77-1.333-2.694-1.333-3.464 0L3.34 16c-.77 1.333.192 3 1.732 3z" />
				</svg>
			</div>

			<div class="flex flex-col gap-2">
				<span class="text-xs uppercase font-bold tracking-wider text-rose-400">
					Error {data.errorStatus}
				</span>
				<h1 class="text-2xl font-black text-white">Short URL Not Found</h1>
				<p class="text-sm text-slate-400 max-w-sm">
					{data.errorMessage || `The short link "/${data.shortCode}" does not exist or has expired.`}
				</p>
			</div>

			<div class="w-full bg-slate-950 p-4 rounded-xl border border-slate-800 text-xs text-slate-400 font-mono">
				Requested path: <span class="text-rose-300">/{data.shortCode}</span>
			</div>

			<div class="flex flex-col sm:flex-row items-center gap-3 w-full">
				<a
					href="/"
					class="w-full py-3 px-6 rounded-xl bg-amber-500 hover:bg-amber-400 text-slate-950 font-bold text-sm transition-all flex items-center justify-center gap-2"
				>
					<svg class="w-4 h-4" fill="none" viewBox="0 0 24 24" stroke="currentColor">
						<path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M3 12l2-2m0 0l7-7 7 7M5 10v10a1 1 0 001 1h3m10-11l2 2m-2-2v10a1 1 0 01-1 1h-3m-6 0a1 1 0 001-1v-4a1 1 0 011-1h2a1 1 0 011 1v4a1 1 0 001 1m-6 0h6" />
					</svg>
					<span>Back to Home</span>
				</a>
			</div>
		</div>
	{/if}
</div>
