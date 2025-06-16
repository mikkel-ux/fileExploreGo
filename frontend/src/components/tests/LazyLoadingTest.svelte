<script lang="ts">
  import { onMount, onDestroy } from "svelte";
  import BigItemListTest from "./BigItemListTest.svelte";

  let items = [...Array(100).keys()];
  let visible = new Set<number>();
  let observer: IntersectionObserver;

  onMount(() => {
    observer = new IntersectionObserver(
      (entries) => {
        for (const entry of entries) {
          if (entry.isIntersecting) {
            const index = Number(entry.target.getAttribute("data-index"));
            if (!isNaN(index)) {
              visible.add(index);
              observer.unobserve(entry.target);
            }
          }
        }
      },
      { rootMargin: "200px" }
    );

    requestAnimationFrame(() => {
      document.querySelectorAll("[data-index]").forEach((el) => {
        observer.observe(el);
      });
    });
  });

  onDestroy(() => {
    if (observer) {
      observer.disconnect();
    }
  });
</script>

<ul class="h-full w-full overflow-auto">
  {#each items as item, i}
    {#if visible.has(i)}
      <BigItemListTest {item} />
    {:else}
      <li data-index={i} class="p-4 border-b border-gray-200">Loading…</li>
    {/if}
  {/each}
</ul>
