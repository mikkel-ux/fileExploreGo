<script>
  let { value = $bindable() } = $props();
  const test = () => {
    value += 1;
    console.log("in caild", value);
  };
</script>

<button class="bg-blue-500 text-white p-2 rounded-lg mt-4" onclick={test}>
  +1
</button>
