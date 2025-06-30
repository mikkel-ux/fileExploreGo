<script lang="ts">
	import { onMount } from 'svelte';
	import { goto } from '$app/navigation';
	import { updateHistory, getCurrentPath } from '../stores/tabsStore';
	import { secsectFile, selectedFile, removeSelectedFile } from '../stores/tabsStore';
	import { GetFiles, OpenFile, GetBase64OfImage } from '../lib/wailsjs/go/goFiles/Files';

	import { isImage } from '../functions/checkFileExtension';
	import RenderIcon from '../components/RenderIcon.svelte';
	import type { FileDataType } from '../../../type';

	import BigItemListTest from '../components/tests/BigItemListTest.svelte';
	import LazyLoadingTest from '../components/tests/LazyLoadingTest.svelte';
	import { Grid } from 'svelte-virtual';
	import { string } from '@tensorflow/tfjs';

	let files = $state<FileDataType[]>([]);
	let clickTimer = $state<NodeJS.Timeout | null>(null);
	let enterTimer = $state<NodeJS.Timeout | null>(null);
	let items = [...Array(1000).keys()];

	let test = $state<number>(0);

	onMount(async () => {
		GetFiles(getCurrentPath())
			.then((result: any) => {
				files = result;
				console.log('files', files);
			})
			.catch((error) => {
				console.error('Error fetching files:', error);
			});
	});

	const handleClick = (file: FileDataType) => {
		if (clickTimer) {
			clearTimeout(clickTimer);
			clickTimer = null;
			openFile(file);
		} else {
			clickTimer = setTimeout(() => {
				secsectFile(file);
				clickTimer = null;
			}, 200);
		}
	};

	const handleKeyDown = (e: KeyboardEvent, file: FileDataType) => {
		if (e.key === 'Enter') {
			e.preventDefault();
			if (enterTimer) {
				clearTimeout(enterTimer);
				enterTimer = null;
				openFile(file);
			} else {
				enterTimer = setTimeout(() => {
					secsectFile(file);
					enterTimer = null;
				}, 200);
			}
		}
	};

	const openFile = async (file: FileDataType) => {
		if (file.type === 'dir') {
			updateHistory(file.path);
			removeSelectedFile();
			return;
		}
		try {
			await OpenFile(file.path);
		} catch (error) {
			return;
		}
	};
</script>

<section class="w-full h-full overflow-y-auto overflow-x-hidden">
	<div
		class="grid grid-cols-5 grid-rows-[repeat(1fr)] gap-2 justify-items-center
  text-center pl-10 pr-10 pb-5 pt-2 select-none"
	>
		{#each files as file}
			<div
				onclick={() => handleClick(file)}
				role="button"
				tabindex="0"
				class:opacity-50={file.isHidden}
				class:bg-highlight={$selectedFile && $selectedFile.name === file.name}
				class="ffContainer w-[15cqw] h-[20cqh] rounded-lg flex flex-col items-center justify-center hover:bg-secondary-bg
        transition-bg ease-in-out gap-2 overflow-hidden"
				onkeydown={(e) => {
					handleKeyDown(e, file);
				}}
			>
				<div class="h-[60%] w-full flex items-center justify-center overflow-hidden">
					<RenderIcon {file} {isImage} />
				</div>

				<p class="truncate w-full">
					{file.name}
				</p>
			</div>
		{/each}
	</div>
</section>

<!-- <section class="w-full h-full">
  <LazyLoadingTest>
    {#snippet children({ item }: any)}
      <BigItemListTest {item} />
    {/snippet}
  </LazyLoadingTest>
</section> -->

<!-- <section class="w-full h-full" bind:clientHeight={test}>
  <Grid itemCount={items.length} itemHeight={50} itemWidth={60} height={test}>
    {#snippet item({ index, style })}
      <div {style}>
        {items[index]}
      </div>
    {/snippet}
  </Grid>
</section> -->

<!-- <section class="w-full h-full">
  <LazyLoadingTest />
</section> -->

<style>
	/*   @container (min-width: 1200px) {
    section {
      grid-template-columns: repeat(5, 1fr);
      div {
        height: 15cqh;
        width: 15cqw;
      }
    }
  } */
	section {
		container-type: inline-size;
		container-name: main;
	}

	@container (max-width: 600px) {
		div {
			grid-template-columns: repeat(3, 1fr);
			.ffContainer {
				height: 20cqh;
				width: 25cqw;
			}
		}
	}
	/* pl-10 pr-10 */
	@container (max-width: 400px) {
		div {
			grid-template-columns: repeat(2, 1fr);
			padding-left: 0;
			padding-right: 0;
			.ffContainer {
				width: 30cqw;
			}
		}
	}

	@container (max-width: 270px) {
		div {
			grid-template-columns: 1fr;
			.ffContainer {
				width: 50cqw;
			}
		}
	}

	section::-webkit-scrollbar {
		width: 10px;
	}

	section::-webkit-scrollbar-track {
		background: #fff;
		border-radius: 10px;
	}

	section::-webkit-scrollbar-thumb {
		background: #888;
		border-radius: 10px;
	}

	section::-webkit-scrollbar-thumb:hover {
		background: #555;
	}
</style>
