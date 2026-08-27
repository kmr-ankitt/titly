<script lang="ts">
	import { page } from '$app/state';
</script>

<svelte:head>
	<title>{page.status || 500} - Page Error</title>
</svelte:head>

<div class="max-w-md mx-auto py-16 flex flex-col items-center justify-center text-center gap-6">
	<div class="w-16 h-16 rounded-2xl bg-rose-500/10 border border-rose-500/20 flex items-center justify-center text-rose-400">
		<svg class="w-8 h-8" fill="none" viewBox="0 0 24 24" stroke="currentColor">
			<path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 8v4m0 4h.01M21 12a9 9 0 11-18 0 9 9 0 0118 0z" />
		</svg>
	</div>

	<div class="flex flex-col gap-2">
		<span class="text-xs uppercase font-bold tracking-wider text-rose-400">
			Error {page.status || 500}
		</span>
		<h1 class="text-3xl font-black text-white">
			{page.status === 404 ? 'Page Not Found' : 'Something Went Wrong'}
		</h1>
		<p class="text-sm text-slate-400">
			{page.error?.message || 'An unexpected error occurred while processing your request.'}
		</p>
	</div>

	<a
		href="/"
		class="w-full py-3 px-6 rounded-xl bg-amber-500 hover:bg-amber-400 text-slate-950 font-bold text-sm transition-all flex items-center justify-center gap-2"
	>
		<svg class="w-4 h-4" fill="none" viewBox="0 0 24 24" stroke="currentColor">
			<path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M3 12l2-2m0 0l7-7 7 7M5 10v10a1 1 0 001 1h3m10-11l2 2m-2-2v10a1 1 0 01-1 1h-3m-6 0a1 1 0 001-1v-4a1 1 0 011-1h2a1 1 0 011 1v4a1 1 0 001 1m-6 0h6" />
		</svg>
		<span>Return to Safety</span>
	</a>
</div>
