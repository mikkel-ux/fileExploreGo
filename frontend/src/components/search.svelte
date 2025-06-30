<script lang="ts">
	import { onMount } from 'svelte';
	import { GetPath, GetFiles } from '../lib/wailsjs/go/goFiles/Files';
	import { getPath, setPath } from '../stores/pathStore';
	import { Searchdir } from '../lib/wailsjs/go/goFiles/fuzzySearch';

	type dirType = {
		name: string;
		path: string;
		points: number;
	};

	let dirs = $state<dirType[]>([]);
	let showDirs = $state<boolean>(false);
	let path = $state<string>('C:/Users/rumbo/.testFoulderForFE');
	let initialized = $state<boolean>(false);

	$effect(() => {
		if (initialized && path.trim() !== '') {
			searchFile();
		}
	});

	onMount(async () => {
		/* GetPath('home')
			.then((res) => {
				path = res;
			})
			.catch((err) => {
				console.error('Error getting path:', err);
			});
		setPath(path); */
		initialized = true;
	});

	async function searchFile() {
		if (!initialized || path.trim() === '') return;
		Searchdir(path)
			.then((res) => {
				dirs = res;
				console.log('dirs', dirs);
			})
			.catch((err) => {
				console.error('Error searching files:', err);
			});
		/* showDirs = true; */
		console.log('searching for', path);
	}

	function changePath(newPath: string) {
		console.log('Saving path:', newPath);
	}
</script>

<div class="relative w-full">
	<input
		placeholder="Enter a name..."
		bind:value={path}
		class="border p-2 rounded-lg bg-gray-100 text-black w-full"
		autocomplete="off"
		onfocus={() => (showDirs = true)}
		onblur={() => setTimeout(() => (showDirs = false), 100)}
	/>
	{#if showDirs && dirs.length}
		<div
			class="absolute left-0 right-0 bg-gray-100 max-h-30 overflow-y-auto border rounded-lg shadow-lg z-10"
		>
			{#each dirs as dir}
				<div
					class="p-2 border-b border-gray-300 hover:bg-gray-200 cursor-pointer text-black"
					role="button"
					tabindex="0"
					onclick={() => changePath(dir.path)}
					onkeydown={(e) => {
						if (e.key === 'Enter') {
							changePath(dir.path);
						}
					}}
				>
					<p>{dir.path}</p>
				</div>
			{/each}
		</div>
	{/if}
</div>
