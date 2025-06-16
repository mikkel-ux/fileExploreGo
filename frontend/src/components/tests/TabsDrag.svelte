<script lang="ts">
  import { dndzone } from "svelte-dnd-action";
  import { flip } from "svelte/animate";

  type Tab = { id: number; name: string };

  let tabs: Tab[] = [
    { id: 1, name: "Tab 1" },
    { id: 2, name: "Tab 2" },
    { id: 3, name: "Tab 3" },
  ];

  function handleSort(e: any) {
    tabs = e.detail.items;
  }
</script>

<div
  class="border-2 border-gray-400 rounded-lg min-h-20 p-4 min-w-32 flex flex-row gap-4"
  use:dndzone={{
    items: tabs,
    flipDurationMs: 150,
    type: "tabs",
    dropTargetStyle: {},
  }}
  on:consider={handleSort}
  on:finalize={handleSort}
>
  {#each tabs as tab (tab.id)}
    <div
      animate:flip={{ duration: 150 }}
      class="p-4 bg-gray-200 m-2 rounded-lg text-black test"
      style="cursor: pointer;"
      tabindex="0"
      role="button"
      data-id={tab.id}
      on:click={() => console.log("Clicked tab", tab.id)}
      on:keydown={(e) => {
        if (e.key === "Enter" || e.key === " ") {
          console.log("Clicked tab", tab.id);
        }
      }}
    >
      {tab.name}
    </div>
  {/each}
</div>

<style>
  .test {
    background-color: #f0f0f0;
    border-radius: 8px;
    padding: 10px;
    margin: 5px;
    cursor: pointer !important;
    transition: background-color 0.3s ease;
  }
</style>
