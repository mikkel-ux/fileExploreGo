<script lang="ts">
  export let direction: "left" | "right" = "left";
  export let disabled: boolean = false;
</script>

<button
  class="p-2 rounded-lg bg-gray-700 text-white hover:bg-gray-600 disabled:opacity-50"
  on:click
  on:mouseover
  on:mouseenter
  on:mouseleave
  {disabled}
  {...$$restProps}
>
  <svg
    class="w-6 h-6"
    fill="none"
    stroke="currentColor"
    viewBox="0 0 24 24"
    xmlns="http://www.w3.org/2000/svg"
  >
    {#if direction === "left"}
      <path
        stroke-linecap="round"
        stroke-linejoin="round"
        stroke-width="2"
        d="M15 19l-7-7 7-7"
      ></path>
    {:else}
      <path
        stroke-linecap="round"
        stroke-linejoin="round"
        stroke-width="2"
        d="M9 5l7 7-7 7"
      ></path>
    {/if}
  </svg>
</button>
