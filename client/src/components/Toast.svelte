<script lang="ts">
	import { toast } from '$lib/toastStore';
</script>

<div class="fixed bottom-6 right-6 z-50 flex flex-col gap-2 max-w-sm w-full px-4 pointer-events-none">
	{#each $toast as t (t.id)}
		<div
			class="pointer-events-auto flex items-center justify-between p-4 rounded-xl border shadow-xl transition-all duration-300 animate-slide-up
			{t.type === 'success'
				? 'bg-slate-900 border-emerald-500/40 text-emerald-400'
				: t.type === 'error'
					? 'bg-slate-900 border-rose-500/40 text-rose-400'
					: 'bg-slate-900 border-blue-500/40 text-blue-400'}"
		>
			<div class="flex items-center gap-3">
				{#if t.type === 'success'}
					<svg class="w-5 h-5 shrink-0" fill="none" viewBox="0 0 24 24" stroke="currentColor">
						<path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M5 13l4 4L19 7" />
					</svg>
				{:else if t.type === 'error'}
					<svg class="w-5 h-5 shrink-0" fill="none" viewBox="0 0 24 24" stroke="currentColor">
						<path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 9v2m0 4h.01m-6.938 4h13.856c1.54 0 2.502-1.667 1.732-3L13.732 4c-.77-1.333-2.694-1.333-3.464 0L3.34 16c-.77 1.333.192 3 1.732 3z" />
					</svg>
				{:else}
					<svg class="w-5 h-5 shrink-0" fill="none" viewBox="0 0 24 24" stroke="currentColor">
						<path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M13 16h-1v-4h-1m1-4h.01M21 12a9 9 0 11-18 0 9 9 0 0118 0z" />
					</svg>
				{/if}
				<span class="text-sm font-medium">{t.message}</span>
			</div>
			<button
				onclick={() => toast.remove(t.id)}
				class="text-slate-400 hover:text-white transition-colors p-1"
				aria-label="Close notification"
			>
				<svg class="w-4 h-4" fill="none" viewBox="0 0 24 24" stroke="currentColor">
					<path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M6 18L18 6M6 6l12 12" />
				</svg>
			</button>
		</div>
	{/each}
</div>

<style>
	@keyframes slideUp {
		from {
			opacity: 0;
			transform: translateY(16px);
		}
		to {
			opacity: 1;
			transform: translateY(0);
		}
	}
	.animate-slide-up {
		animation: slideUp 0.25s cubic-bezier(0.16, 1, 0.3, 1) forwards;
	}
</style>
