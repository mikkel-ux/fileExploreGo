<script lang="ts">
	import { onMount } from 'svelte';
	/* import { invoke } from "@tauri-apps/api/core"; */
	import { GetPath, GetFiles } from '../lib/wailsjs/go/goFiles/Files';
	import { getPath, setPath } from '../stores/pathStore';
	import { Searchdir } from '../lib/wailsjs/go/goFiles/fuzzySearch';

	type dirType = {
		name: string;
	};

	let dirs = $state<any[]>([]);
	let showDirs = $state<boolean>(false);
	let path = $state<string>('C:/Users/rumbo/.testFoulderForFE');
	let initialized = $state<boolean>(false);

	async function searchFile() {
		/* TODO search files */
		if (!initialized || path.trim() === '') return;
		/* const foo: string[] = await invoke('search_files', { path }); */
		const foo = await Searchdir(path);
		dirs = foo;
		showDirs = true;
		console.log('searching for', path);
	}

	$effect(() => {
		searchFile();
	});

	onMount(async () => {
		/* TODO get path */
		/* path = await invoke('get_path', { path: 'home' }); */
		setPath(path);
		console.log('path', path);
		initialized = true;
	});

	const test = async () => {
		let test: String[] = [];
		/* TODO get foulders */
		/* try {
			test = await invoke('get_foulders', { path });
		} catch (error) {
			console.log('error', error);
		} */
		const prettyPath = test.map((item) => {
			return item.replace(/\\/g, '/');
		});
		console.log('test', prettyPath);
	};
</script>

<div class="relative w-full">
	<input
		placeholder="Enter a name..."
		bind:value={path}
		class="border p-2 rounded-lg bg-gray-100 text-black w-full"
		autocomplete="off"
		onfocus={() => (showDirs = true)}
		onblur={() => setTimeout(() => (showDirs = false), 50)}
	/>
	{#if showDirs && dirs.length}
		<div
			class="absolute left-0 right-0 bg-gray-100 max-h-30 overflow-y-auto border rounded-lg shadow-lg z-10"
		>
			{#each dirs as dir}
				<div class="p-2 border-b border-gray-300 hover:bg-gray-200 cursor-pointer text-black">
					<p>{dir.path}</p>
				</div>
			{/each}
		</div>
	{/if}
</div>
