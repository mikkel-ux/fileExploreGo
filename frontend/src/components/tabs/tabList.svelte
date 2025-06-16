<script lang="ts">
  import { flip } from "svelte/animate";
  import { fade, fly } from "svelte/transition";
  /* import { tabs, isDragging } from "../../stores/test-TabsStore"; */
  import { tabsStore, isDragging, activeTabId } from "../../stores/tabsStore";

  let { onTabClick = () => {}, onTabClose = () => {} } = $props();
</script>

{#if !$isDragging}
  {#each $tabsStore as tab (tab.id)}
    <div
      in:fly={{ y: 10, duration: 150 }}
      out:fade={{ duration: 150 }}
      class={`flex items-center p-2 rounded-t-lg h-7 w-32 ${tab.id === $activeTabId ? "bg-highlight" : "bg-secondary-bg"} tab`}
      role="button"
      tabindex="0"
      onclick={() => onTabClick(tab.id)}
      onkeydown={(e) => {
        if (e.key === "Enter" || e.key === " ") {
          e.preventDefault();
          onTabClick(tab.id);
        }
      }}
    >
      <span class="overflow-hidden text-ellipsis whitespace-nowrap flex-1">
        {#if tab.layout === "single"}
          {tab.view[0].title}
        {:else if tab.layout === "split"}
          {tab.view[0].title} | {tab.view[1].title}
        {/if}
      </span>
      <button
        class="ml-5 z-10"
        onclick={(e) => {
          e.stopPropagation();
          onTabClose(tab.id);
        }}
      >
        X
      </button>
    </div>
  {/each}
{:else}
  {#each $tabsStore as tab (tab.id)}
    <div
      animate:flip={{ duration: 100 }}
      class={`flex items-center p-2 rounded-t-lg h-7 w-32 ${tab.id === $activeTabId ? "bg-highlight" : "bg-secondary-bg"} tab`}
      role="button"
      tabindex="0"
      onclick={() => onTabClick(tab.id)}
      onkeydown={(e) => {
        if (e.key === "Enter" || e.key === " ") {
          e.preventDefault();
          onTabClick(tab.id);
        }
      }}
    >
      <span class="overflow-hidden text-ellipsis whitespace-nowrap flex-1">
        {#if tab.layout === "single"}
          {tab.view[0].title}
        {:else if tab.layout === "split"}
          {tab.view[0].title} | {tab.view[1].title}
        {/if}
      </span>
      <button
        class="ml-5 z-10"
        onclick={(e) => {
          e.stopPropagation();
          onTabClose(tab.id);
        }}
      >
        X
      </button>
    </div>
  {/each}
{/if}

<style>
  .tab {
    cursor: default !important;
  }
</style>
