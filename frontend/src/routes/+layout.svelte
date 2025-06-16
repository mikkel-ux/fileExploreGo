<!-- https://wails.io/docs/guides/sveltekit/ -->
<!-- https://v3alpha.wails.io/guides/customising-windows/ -->
<script lang="ts">
	import '../app.css';

	let { children } = $props();
</script>

{@render children()}
