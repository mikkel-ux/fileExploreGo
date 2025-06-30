<script lang="ts">
	import { Greet } from '../lib/wailsjs/go/main/App';
	import { SayHi } from '../lib/wailsjs/go/main/Utils';
	import { GetFiles, GetDefaultDirs } from '../lib/wailsjs/go/goFiles/Files';

	import { onMount } from 'svelte';
	import { goto } from '$app/navigation';
	import { addToHistory } from '../stores/historyStore';
	import { getPath, setPath } from '../stores/pathStore';
	import { selectedFile } from '../stores/tabsStore';
	import TabsDrag from '../components/tests/TabsDrag.svelte';
	import Resize from '../components/tests/resize.svelte';
	import Test from '../components/tests/Test.svelte';
	import Sidebar from '../components/Sidebar.svelte';
	import View from './View.svelte';
	import FileFoulderInfo from '../components/FileFoulderInfo.svelte';

	import ImageTest from '../components/tests/ImageTest.svelte';

	import type { FileDataType } from '../../../type';

	import Button from '../components/Button.svelte';
	import ArrowButton from '../components/ArrowButton.svelte';

	import PropsTest from '../components/PropsTest.svelte';

	let value = $state(0);

	const navigateToPage = () => {
		goto('/expolorer');
		addToHistory('/expolorer');
	};

	/* type FileData = {
		name: string;
		path: string;
		size: string;
		extension: string;
		created: string;
		modified: string;
		accessed: string;
		fileType: string;
		isHidden: boolean;
		isReadOnly: boolean;
		base64?: string;
	};

	interface FileDataMap {
		home: string;
		desktop: string;
		documents?: string;
		downloads?: string;
		pictures?: string;
		music?: string;
		videos?: string;
		[key: string]: string | undefined;
	}

	function handleClick() {
		GetDefaultDirs()
			.then((data: FileDataMap) => {
				console.log(data.home);
			})
			.catch((error: Error) => {
				console.log('Error in Test:', error);
			});
	} */
</script>

<section class="w-full h-full flex flex-row min-h-0">
	<Sidebar />
	<View />
	{#if $selectedFile}
		<FileFoulderInfo />
	{/if}
</section>
