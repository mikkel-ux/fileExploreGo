<script lang="ts">
  import { flip } from "svelte/animate";
  import { dndzone, dragHandle, dragHandleZone } from "svelte-dnd-action";

  type itemType = { id: number; name: string; isActive: boolean };

  let flipDurationMs: number = 150;

  let isDraggingActive: boolean = $state(false);
  let dropInSplitView: boolean = $state(false);

  let box1Ratio: number = $state(0.5);
  let isDragging: boolean = $state(false);
  let items: itemType[] = $state([
    { id: 1, name: "Item 1", isActive: false },
    { id: 2, name: "Item 2", isActive: true },
  ]);

  let items2: itemType[] = $state([]);

  $effect(() => {
    if (items.length === 1) {
      box1Ratio = 1;
    } else if (items.length === 2) {
      box1Ratio = 0.5;
    }
  });

  const startDragging = (e: MouseEvent) => {
    isDragging = true;
    document.body.style.cursor = "col-resize";
    window.addEventListener("mousemove", onDrag);
    window.addEventListener("mouseup", stopDragging);
  };

  const stopDragging = () => {
    isDragging = false;
    document.body.style.cursor = "";
    window.removeEventListener("mousemove", onDrag);
    window.removeEventListener("mouseup", stopDragging);
  };

  const onDrag = (e: MouseEvent) => {
    if (!isDragging) return;
    const container = document.getElementById("resize-container");
    if (!container) return;
    const rect = container.getBoundingClientRect();
    const newRatio = (e.clientX - rect.left) / rect.width;
    box1Ratio = Math.min(0.9, Math.max(0.1, newRatio));
  };

  function handleFooSort(e: any) {
    items = e.detail.items;
  }

  function handleFoo3Sort(e: any) {
    isDraggingActive = false;
    dropInSplitView = items.length === 2 || false;
    items = e.detail.items;
  }

  function handleFoo2Sort(e: any) {
    items2 = e.detail.items;
  }

  const transformDraggedElement = (
    draggedEl: HTMLElement,
    data: any,
    index: any
  ) => {
    draggedEl.style.opacity = "0.5";
    const draggedItem = data.id;
    const draggedItemIndex = items.findIndex((item) => item.id === draggedItem);
    if (draggedItemIndex !== -1) {
      isDraggingActive = items[draggedItemIndex].isActive;
    }
  };
</script>

<section
  class="flex flex-row p-2 gap-5 border-2 border-gray-400 rounded-lg h-40 w-full max-w-xl"
  use:dndzone={{
    items: items2,
    centreDraggedOnCursor: true,
    type: "tabs",
    dropFromOthersDisabled: isDraggingActive,
    /* @ts-ignore */
    transformDraggedElement,
  }}
  onconsider={handleFoo2Sort}
  onfinalize={handleFoo2Sort}
>
  {#each items2 as item (item.id)}
    <div class="h-full border-2 border-gray-200 relative group w-full">
      <div class="flex items-center justify-center h-full bg-gray-200">
        <span class="text-black">{item.name}</span>
      </div>
    </div>
  {/each}
</section>

<br />
<section
  id="resize-container"
  class="relative border-2 border-gray-400 rounded-lg h-40 w-full max-w-xl flex flex-row p-0.5"
  use:dragHandleZone={{
    items,
    flipDurationMs: flipDurationMs,
    type: "tabs",
    dropFromOthersDisabled: dropInSplitView,
    /* @ts-ignore */
    transformDraggedElement,
  }}
  onconsider={handleFooSort}
  onfinalize={handleFoo3Sort}
>
  {#each items as item, index (item.id)}
    <div
      class="h-full"
      style={`width: ${index === 0 ? box1Ratio * 100 : (1 - box1Ratio) * 100}%`}
      animate:flip={{ duration: flipDurationMs }}
    >
      <div
        class={`flex items-center justify-center h-full ${index === 0 ? "bg-gray-200" : "bg-gray-300"}`}
      >
        <span class="text-black">{item.id}</span>
        <span
          use:dragHandle
          aria-label="drag-handle for {item.id}"
          class="ml-2 cursor-move text-black">::</span
        >
      </div>
    </div>
  {/each}

  <div class="absolute top-0 left-0 h-full w-full pointer-events-none">
    {#if items.length > 1}
      <button
        type="button"
        class="absolute top-0 h-full w-2 cursor-col-resize bg-[#ccc] hover:bg-[#999] appearance-none border-none p-0 z-10 pointer-events-auto"
        style={`left: calc(${box1Ratio * 100}% - 1px);`}
        onmousedown={startDragging}
        aria-label="Resize views"
        onkeydown={(e) => {
          if (e.key === "ArrowLeft")
            box1Ratio = Math.max(0.1, box1Ratio - 0.01);
          if (e.key === "ArrowRight")
            box1Ratio = Math.min(0.9, box1Ratio + 0.01);
        }}
      ></button>
    {/if}
  </div>
</section>
