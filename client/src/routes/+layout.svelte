<script lang="ts">
	import '../app.css';
	import Navbar from '../components/Navbar.svelte';
	import Toast from '../components/Toast.svelte';

	let { children } = $props();
</script>

<svelte:head>
	<title>Titly — Blazing Fast URL Shortener</title>
	<meta name="description" content="Shorten long URLs instantly with Titly's fast, reliable shortener." />
</svelte:head>

<div class="min-h-screen flex flex-col bg-slate-950 text-slate-100 selection:bg-amber-500 selection:text-slate-950">
	<!-- Navbar -->
	<Navbar />

	<!-- Main Content area -->
	<main class="flex-1 w-full max-w-5xl mx-auto px-4 sm:px-6 py-8">
		{@render children()}
	</main>

	<!-- Footer -->
	<footer class="w-full border-t border-slate-900 bg-slate-950 py-6 text-center text-xs text-slate-500">
		<div class="max-w-5xl mx-auto px-4 flex flex-col sm:flex-row items-center justify-between gap-3">
			<p>© {new Date().getFullYear()} Titly. All rights reserved.</p>
			<p class="flex items-center gap-1">
				<span>Powered by</span>
				<span class="text-amber-400 font-semibold">Svelte 5</span>
				<span>&</span>
				<span class="text-sky-400 font-semibold">Go Gin</span>
			</p>
		</div>
	</footer>

	<!-- Floating Toast Notifications -->
	<Toast />
</div>
