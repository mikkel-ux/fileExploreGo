<script lang="ts">
  import { onMount } from "svelte";

  let widthRatio: number = $state(200);
  let isResizing = false;

  const startResize = (e: MouseEvent) => {
    isResizing = true;
    e.preventDefault();
  };

  const stopResize = () => {
    isResizing = false;
  };

  const resize = (e: MouseEvent) => {
    if (isResizing) {
      widthRatio = Math.min(500, Math.max(100, e.clientX));
    }
  };

  onMount(() => {
    window.addEventListener("mousemove", resize);
    window.addEventListener("mouseup", stopResize);

    return () => {
      window.removeEventListener("mousemove", resize);
      window.removeEventListener("mouseup", stopResize);
    };
  });
</script>

<div
  style={`width: ${widthRatio}px`}
  class="flex h-full relative flex-shrink-0 flex-grow-0"
>
  <div class="flex-grow bg-gray-800 text-white">
    <p>Home</p>
  </div>

  <button
    class="resize-handle absolute top-0 right-0 h-full w-2 cursor-col-resize"
    onmousedown={startResize}
    aria-label="Resize sidebar"
    onkeydown={(e) => {
      if (e.key === "ArrowLeft") widthRatio = Math.max(100, widthRatio - 10);
      if (e.key === "ArrowRight") widthRatio = Math.min(800, widthRatio + 10);
    }}
  ></button>
</div>

<style>
  .resize-handle {
    width: 5px;
    cursor: ew-resize;
    background-color: rgba(255, 255, 255, 0.15);
  }

  .resize-handle:hover {
    background-color: rgba(255, 255, 255, 0.25);
  }

  .resize-handle:focus {
    outline: none;
    box-shadow:
      0 0 0 2px rgba(255, 255, 255, 0.5),
      0 0 6px 2px rgba(0, 123, 255, 0.4);
    background-color: rgba(255, 255, 255, 0.15);
  }
</style>
