<script lang="ts">
  export let className = "";
</script>

<button
  class={`px-4
  py-2
  rounded-lg
  font-medium
  transition
  cursor-pointer
  ${className}
  `}
  on:click
  on:mouseover
  on:mouseenter
  on:mouseleave
  {...$$restProps}
>
  <slot />
</button>
