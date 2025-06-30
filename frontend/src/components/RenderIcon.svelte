<script lang="ts">
	import type { FileDataType } from '../../../type';
	import { Folder, Inspect, FileText, File } from '@lucide/svelte';

	interface Props {
		file: FileDataType;
		getBase64Image: (path: string) => string;
		isImage: (file: FileDataType) => boolean;
	}

	let { file, getBase64Image, isImage }: Props = $props();
</script>

{#if file.type === 'folder'}
	<Folder size="100%" />
{:else if file.extension.toLowerCase() === 'txt'}
	<FileText size="100%" />
{:else if isImage(file)}
	{#if file.extension === 'gif' && file.base64}
		<img
			src={`data:image/png;base64, ${file.base64}`}
			alt="first frame of gif"
			class="object-contain h-full w-full p-1"
		/>
	{:else}
		<img src={getBase64Image(file.path)} alt="preview" class="object-contain h-full w-full" />
	{/if}
{:else if file.type === 'file'}
	<File size="100%" />
{:else}
	<p>no image</p>
{/if}
