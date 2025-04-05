<script>
	import Form from '../components/Form.svelte';
</script>

<div class="absolute top-0 -z-10 h-full w-full bg-white">
	<div
		class="absolute bottom-auto left-auto right-0 top-0 h-[500px] w-[500px] -translate-x-[30%] translate-y-[20%] rounded-full bg-[rgba(255,165,0,0.5)] opacity-50 blur-[80px]"
	></div>
	<div
		class="absolute bottom-0 left-auto right-auto top-auto h-[400px] w-[400px] translate-x-[10%] -translate-y-[10%] rounded-full bg-[rgba(0,128,255,0.5)] opacity-50 blur-[60px]"
	></div>
	<div
		class="absolute bottom-auto left-0 right-auto top-auto h-[300px] w-[300px] -translate-x-[20%] translate-y-[30%] rounded-full bg-[rgba(128,0,255,0.5)] opacity-50 blur-[50px]"
	></div>
</div>
<section class="h-screen w-full flex flex-col items-center justify-center text-zinc-900 gap-5 p-4">
	<h1 class="text-4xl font-extrabold">Welcome to <span class="text-yellow-500">Titly</span></h1>
	<p class="text-lg">Blazing fast URL Shortner</p>
	<Form />
</section>
