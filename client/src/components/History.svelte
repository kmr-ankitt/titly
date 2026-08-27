<script lang="ts">
	import { onMount } from 'svelte';
	import { getHistory, removeHistoryItem, clearHistory, type HistoryItem } from '$lib/history';
	import { toast } from '$lib/toastStore';

	let historyItems = $state<HistoryItem[]>([]);

	export const refresh = () => {
		historyItems = getHistory();
	};

	onMount(() => {
		refresh();
	});

	const copyLink = async (url: string) => {
		try {
			await navigator.clipboard.writeText(url);
			toast.add('Link copied to clipboard!', 'success');
		} catch {
			toast.add('Failed to copy link', 'error');
		}
	};

	const handleDelete = (id: string) => {
		historyItems = removeHistoryItem(id);
		toast.add('Item removed from history', 'info');
	};

	const handleClearAll = () => {
		historyItems = clearHistory();
		toast.add('History cleared', 'info');
	};

	const formatDate = (timestamp: number) => {
		return new Date(timestamp).toLocaleDateString(undefined, {
			month: 'short',
			day: 'numeric',
			hour: '2-digit',
			minute: '2-digit'
		});
	};
</script>

{#if historyItems.length > 0}
	<div class="w-full bg-slate-900 border border-slate-800 rounded-2xl p-6 sm:p-8 shadow-2xl flex flex-col gap-5">
		<div class="flex items-center justify-between border-b border-slate-800 pb-4">
			<div class="flex items-center gap-2">
				<svg class="w-5 h-5 text-amber-400" fill="none" viewBox="0 0 24 24" stroke="currentColor">
					<path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 8v4l3 3m6-3a9 9 0 11-18 0 9 9 0 0118 0z" />
				</svg>
				<h2 class="text-base font-bold text-slate-200">Recent Shortened Links</h2>
				<span class="text-xs px-2 py-0.5 rounded-full bg-slate-800 text-slate-400 font-semibold">
					{historyItems.length}
				</span>
			</div>

			<button
				type="button"
				onclick={handleClearAll}
				class="text-xs text-rose-400 hover:text-rose-300 font-medium hover:underline"
			>
				Clear All
			</button>
		</div>

		<div class="flex flex-col gap-3 max-h-80 overflow-y-auto pr-1">
			{#each historyItems as item (item.id)}
				<div class="flex flex-col sm:flex-row items-start sm:items-center justify-between gap-3 p-4 bg-slate-950 border border-slate-800/80 rounded-xl hover:border-slate-700 transition-colors">
					<div class="flex flex-col min-w-0 pr-2">
						<div class="flex items-center gap-2">
							<a
								href={item.fullShortUrl}
								target="_blank"
								rel="noopener noreferrer"
								class="text-amber-400 font-bold text-sm hover:underline truncate"
							>
								{item.fullShortUrl}
							</a>
							<span class="text-[10px] text-slate-500 shrink-0 font-medium">
								{formatDate(item.createdAt)}
							</span>
						</div>
						<span class="text-xs text-slate-400 truncate mt-0.5 font-mono">
							{item.longUrl}
						</span>
					</div>

					<div class="flex items-center gap-2 shrink-0 self-end sm:self-auto">
						<button
							type="button"
							onclick={() => copyLink(item.fullShortUrl)}
							class="px-3 py-1.5 rounded-lg bg-slate-800 hover:bg-slate-700 text-slate-200 text-xs font-semibold transition-colors flex items-center gap-1.5"
						>
							<svg class="w-3.5 h-3.5" fill="none" viewBox="0 0 24 24" stroke="currentColor">
								<path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M8 5H6a2 2 0 00-2 2v12a2 2 0 002 2h10a2 2 0 002-2v-1M8 5a2 2 0 002 2h2a2 2 0 002-2M8 5a2 2 0 012-2h2a2 2 0 012 2m0 0h2a2 2 0 012 2v3m2 4H10m0 0l3-3m-3 3l3 3" />
							</svg>
							<span>Copy</span>
						</button>

						<button
							type="button"
							onclick={() => handleDelete(item.id)}
							class="p-1.5 rounded-lg text-slate-500 hover:text-rose-400 hover:bg-rose-950/30 transition-colors"
							aria-label="Remove item"
						>
							<svg class="w-4 h-4" fill="none" viewBox="0 0 24 24" stroke="currentColor">
								<path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M19 7l-.867 12.142A2 2 0 0116.138 21H7.862a2 2 0 01-1.995-1.858L5 7m5 4v6m4-6v6m1-10V4a1 1 0 00-1-1h-4a1 1 0 00-1 1v3M4 7h16" />
							</svg>
						</button>
					</div>
				</div>
			{/each}
		</div>
	</div>
{/if}
