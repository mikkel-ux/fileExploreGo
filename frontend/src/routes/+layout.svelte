<!-- https://wails.io/docs/guides/sveltekit/ -->
<!-- https://v3alpha.wails.io/guides/customising-windows/ -->
<script lang="ts">
	let { children } = $props();

	import { goto } from '$app/navigation';
	/* import {
    goBack,
    goForward,
    canGoBack,
    canGoForward,
    getHistory,
  } from "../stores/historyStore"; */
	import { goBack, goForward, canGoBack, canGoForward, getHistory } from '../stores/tabsStore';
	import '../app.css';
	import Search from '../components/search.svelte';
	import ArrowButton from '../components/ArrowButton.svelte';
	import TitleBar from '../components/TitleBar.svelte';

	let canGoBackValue = $state(false);
	let canGoForwardValue = $state(false);

	$effect(() => {
		canGoBackValue = $canGoBack;
		canGoForwardValue = $canGoForward;
	});

	const navigateBack = () => {
		if (canGoBackValue) {
			const path = goBack();
			goto(path);
		}
	};

	const navigateForward = () => {
		if (canGoForwardValue) {
			const path = goForward();
			goto(path);
		}
	};
</script>

<div class="h-full w-full flex flex-col min-h-0">
	<TitleBar />
	<header
		class="bg-gray-800 text-white p-4 grid grid-cols-[auto_1fr] gap-4 border-b-4 border-black"
	>
		<section class="flex gap-2">
			<ArrowButton direction="left" disabled={!canGoBackValue} on:click={() => navigateBack()}
			></ArrowButton>

			<ArrowButton
				direction="right"
				disabled={!canGoForwardValue}
				on:click={() => navigateForward()}
			></ArrowButton>
		</section>
		<Search />
	</header>

	{@render children()}
</div>
