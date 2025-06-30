<script lang="ts">
	import type { FileDataType } from '../../../type';
	import { Folder, Inspect, FileText, File } from '@lucide/svelte';

	interface Props {
		file: FileDataType;
		isImage: (file: FileDataType) => boolean;
	}

	let { file, isImage }: Props = $props();
</script>

{#if file.type === 'dir'}
	<Folder size="100%" />
{:else if file.extension.toLowerCase() === '.txt'}
	<FileText size="100%" />
{:else if isImage(file)}
	{#if file.extension === '.gif'}
		<img src={file.firstFrame} alt="first frame of gif" class="object-contain h-full w-full p-1" />
	{:else}
		<img src={file.base64} alt="preview" class="object-contain h-full w-full" />
	{/if}
{:else if file.type === 'file'}
	<File size="100%" />
{:else}
	<p>no image</p>
{/if}
