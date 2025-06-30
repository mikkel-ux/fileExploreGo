<script lang="ts">
	import type { FileDataType } from '../../../type';
	interface Props {
		file: FileDataType;
		autoplay?: boolean;
	}
	let { file, autoplay }: Props = $props();
</script>

{#if file.extension.toLowerCase() === '.gif'}
	{#if autoplay}
		<img
			src={file.base64}
			alt="autoplay gif"
			class="max-h-40 sm:max-h-60 md:max-h-80 lg:max-h-[32rem] w-auto"
		/>
	{:else if file.firstFrame}
		<img
			src={file.firstFrame}
			alt="first frame of gif"
			class="max-h-40 sm:max-h-60 md:max-h-80 lg:max-h-[32rem] w-auto"
		/>
	{:else}
		<p class="text-gray-400">No preview available</p>
	{/if}
{:else}
	<img
		src={file.base64}
		alt="preview"
		class="max-h-40 sm:max-h-60 md:max-h-80 lg:max-h-[32rem] w-auto"
	/>
{/if}
