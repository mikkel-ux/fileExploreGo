<script lang="ts">
  import { fade, fly } from "svelte/transition";
  import { removeSelectedFile, selectedFile } from "../stores/tabsStore";
  import { convertFileSrc } from "@tauri-apps/api/core";
  import { isImage } from "../functions/checkFileExtension";
  import { Folder, FileText } from "@lucide/svelte";
  import { invoke } from "@tauri-apps/api/core";
  import { onMount } from "svelte";
  import ImagePreview from "./ImagePreview.svelte";

  let autoplay = $state<boolean>(false);

  const close = () => {
    removeSelectedFile();
  };

  const getImageUrl = (path: string) => {
    return convertFileSrc(path);
  };

  const openFile = async (path: string) => {
    try {
      await invoke("open_in_default_app", { path });
    } catch (error) {
      console.error("Error opening file:", error);
    }
  };
</script>

<section
  in:fly={{ x: 100, duration: 200 }}
  out:fade={{ duration: 200 }}
  class="w-2xl h-full flex flex-col gap-4 p-2 bg-secondary-bg transition-colors text-white overflow-y-auto"
>
  <div class="flex justify-between">
    <p class="text-2xl font-bold">Details</p>
    <button class="font-bold hover:opacity-50" onclick={() => close()}>
      x
    </button>
  </div>
  <div
    class="border-b-4 border-t-4 border-gray-500 flex flex-col justify-center items-center p-2 gap-2"
  >
    {#if $selectedFile}
      <!-- Folder -->
      {#if $selectedFile.type === "folder"}
        <Folder size="50%" />
      {:else if $selectedFile.extension.toLowerCase() === "txt"}
        <!-- File Icon -->
        <FileText size="50%" />
      {:else if isImage($selectedFile)}
        <ImagePreview file={$selectedFile} {autoplay} {getImageUrl} />

        <!-- PDF -->
      {:else if $selectedFile.extension.toLowerCase() === "pdf"}
        <iframe
          src={getImageUrl($selectedFile.path)}
          class="w-full rounded-lg h-60"
          scrolling="auto"
          title="PDF Preview"
        ></iframe>
      {:else}
        <p class="text-gray-400">No preview available</p>
      {/if}
      <!-- {:else}
      <p class="text-gray-400">No file selected</p>
      -->
    {/if}

    <p class="text-sm text-center break-all">
      {$selectedFile ? $selectedFile.name : "No file selected"}
    </p>
  </div>

  <div class="grid grid-cols-2 gap-4 text-sm border-b-4 border-gray-500">
    <div>
      <p class="text-gray-400">Type</p>
      <p class="font-semibold">
        {$selectedFile?.type === "folder"
          ? "Folder"
          : ($selectedFile?.extension ?? "–")}
      </p>
    </div>

    <div>
      <p class="text-gray-400">Size</p>
      <p class="font-semibold">{$selectedFile?.size ?? "–"}</p>
    </div>

    <div class="col-span-2">
      <p class="text-gray-400">Location</p>
      <p class="break-words font-semibold">{$selectedFile?.path ?? "–"}</p>
    </div>

    <div>
      <p class="text-gray-400">Created</p>
      <p class="font-semibold">{$selectedFile?.created ?? "–"}</p>
    </div>

    <div>
      <p class="text-gray-400">Modified</p>
      <p class="font-semibold">{$selectedFile?.modified ?? "–"}</p>
    </div>

    <div class="col-span-2 mb-2">
      <p class="text-gray-400">Last Accessed</p>
      <p class="font-semibold">{$selectedFile?.accessed ?? "–"}</p>
    </div>
  </div>

  <div class="flex justify-end">
    <button
      onclick={() => openFile($selectedFile?.path ?? "")}
      class="w-full bg-folder text-black rounded-lg hover:bg-fuchsia-50
    "
    >
      open
    </button>
  </div>
</section>
