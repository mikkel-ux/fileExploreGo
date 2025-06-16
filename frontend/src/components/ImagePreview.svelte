<script lang="ts">
  import type { FileDataType } from "../../type";
  interface Props {
    file: FileDataType;
    autoplay?: boolean;
    getImageUrl: (path: string) => string;
  }
  let { file, autoplay, getImageUrl }: Props = $props();
</script>

{#if file.extension.toLowerCase() === "gif"}
  {#if autoplay}
    <img
      src={getImageUrl(file.path)}
      alt="autoplay gif"
      class="max-h-40 sm:max-h-60 md:max-h-80 lg:max-h-[32rem] w-auto"
    />
  {:else if file.base64}
    <img
      src={`data:image/png;base64, ${file.base64}`}
      alt="first frame of gif"
      class="max-h-40 sm:max-h-60 md:max-h-80 lg:max-h-[32rem] w-auto"
    />
  {:else}
    <p class="text-gray-400">No preview available</p>
  {/if}
{:else}
  <img
    src={getImageUrl(file.path)}
    alt="preview"
    class="max-h-40 sm:max-h-60 md:max-h-80 lg:max-h-[32rem] w-auto"
  />
{/if}
