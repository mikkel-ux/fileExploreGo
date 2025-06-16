<script lang="ts">
  import { dndzone } from "svelte-dnd-action";
  import { flip } from "svelte/animate";
  const flipDurationMs = 200;
  function handlezone1Sort(e: any) {
    zone1Items = e.detail.items;
  }

  function handlezone2Sort(e: any) {
    zone2Items = e.detail.items;
  }
  let zone1Items = [
    { id: 1, name: "Item1 1" },
    { id: 2, name: "Item1 2" },
    { id: 3, name: "Item1 3" },
  ];

  let zone2Items: { id: number; name: string }[] = [];

  let type = "foo";
</script>

<div class="flex flex-row gap-4 p-4">
  <section
    use:dndzone={{ items: zone1Items, flipDurationMs, type }}
    on:consider={handlezone1Sort}
    on:finalize={handlezone1Sort}
    class="w-48 p-4 border-2 border-gray-400 rounded-lg min-h-32"
  >
    {#each zone1Items as item (item.id)}
      <div
        class="p-4 bg-gray-200 m-2 rounded-lg text-black"
        draggable="true"
        animate:flip={{ duration: flipDurationMs }}
      >
        {item.name}
      </div>
    {/each}
  </section>

  <section
    use:dndzone={{ items: zone2Items, flipDurationMs, type }}
    on:consider={handlezone2Sort}
    on:finalize={handlezone2Sort}
    class="w-48 p-4 border-2 border-gray-400 rounded-lg min-h-32"
  >
    {#each zone2Items as item (item.id)}
      <div
        class="p-4 bg-gray-200 m-2 rounded-lg text-black test"
        draggable="true"
        animate:flip={{ duration: flipDurationMs }}
      >
        {item.name}
      </div>
    {/each}
  </section>
</div>

<style>
  .test {
    cursor: pointer !important;
  }
</style>
