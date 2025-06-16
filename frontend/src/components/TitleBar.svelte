<script lang="ts">
  import { getCurrentWindow } from "@tauri-apps/api/window";
  import { dndzone } from "svelte-dnd-action";
  import TabList from "./tabs/tabList.svelte";
  /* import {
    tabs,
    nextTabId,
    activeTabId,
    isDragging,
  } from "../stores/test-TabsStore"; */
  import {
    tabsStore,
    activeTabId,
    isDragging,
    addTab,
    setActiveTab,
    removeTab,
  } from "../stores/tabsStore";
  import { tick } from "svelte";
  import { get } from "svelte/store";

  const appWindow = getCurrentWindow();

  const minimize = () => {
    appWindow.minimize();
  };

  const toggleMaximize = () => {
    appWindow.toggleMaximize();
  };

  const close = () => {
    appWindow.close();
  };

  const addNewTab = () => {
    addTab();
  };

  const closeTab = async (id: number) => {
    removeTab(id);
    await tick();
  };

  const handleKeyDown = (event: KeyboardEvent) => {
    if (event.ctrlKey && event.key.toLocaleLowerCase() === "t") {
      addNewTab();
    }
    if (event.ctrlKey && event.key.toLocaleLowerCase() === "w") {
      const activeTab = $tabsStore.find((tab) => tab.id === $activeTabId);
      if (activeTab) {
        closeTab(activeTab.id);
      }
    }
  };

  const activeTab = (id: number) => {
    setActiveTab(id);
  };

  function handleSort(e: any) {
    isDragging.set(e.type === "finalize" ? false : true);
    tabsStore.update(() => e.detail.items);
    console.log("items", $tabsStore);
  }
</script>

<svelte:window onkeydown={handleKeyDown} />

<section
  data-tauri-drag-region
  class="col-span-2 row-start-1 row-end-1 flex pr-1 gap-5"
>
  <div
    class="flex-1 gap-2 flex items-end overflow-hidden pt-3"
    data-tauri-drag-region
  >
    <div
      onwheel={(e) => {
        e.preventDefault();
        const scrollAmount = e.deltaY > 0 ? 100 : -100;
        e.currentTarget.scrollBy({ left: scrollAmount, behavior: "smooth" });
      }}
      use:dndzone={{
        items: $tabsStore,
        type: "tabs",
        flipDurationMs: 150,
        dropTargetStyle: {},
      }}
      onconsider={handleSort}
      onfinalize={handleSort}
      class="flex gap-1.5 items-center overflow-x-auto overflow-y-hidden tab-scrollbar min-w-2 min-h-7"
    >
      <TabList onTabClick={activeTab} onTabClose={closeTab} />
    </div>
    <button
      class="bg-secondary-bg text-white p-2 rounded-t-lg h-7 w-7 flex items-center justify-center justify-self-center
      hover:bg-highlight hover:opacity-90 transition-opacity transition-highlight ease-in-out"
      onclick={addNewTab}>+</button
    >
  </div>

  <div class="flex gap-2 items-center">
    <button onclick={minimize} class="hover:bg-gray-200">_</button>
    <button onclick={toggleMaximize} class="hover:bg-gray-200">[]</button>
    <button onclick={close} class="hover:bg-gray-200">X</button>
  </div>
</section>

<style>
  .tab-scrollbar {
    overflow-x: auto;
    transition: width 0.15s ease;
  }

  .tab-scrollbar::-webkit-scrollbar {
    height: 5px;
  }

  .tab-scrollbar::-webkit-scrollbar-track {
    background: #fff;
    border-radius: 10px;
  }

  .tab-scrollbar::-webkit-scrollbar-thumb {
    background: #888;
    border-radius: 10px;
  }

  .tab-scrollbar::-webkit-scrollbar-thumb:hover {
    background: #555;
  }

  /* .tab-scrollbar {
    transition: opacity 0.3s ease;
  }

  .tab-scrollbar::-webkit-scrollbar {
    opacity: 1;
    transition: opacity 0.3s ease;
  } */
</style>
