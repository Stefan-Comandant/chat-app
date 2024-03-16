<script lang="ts">
	import { page } from '$app/stores';
	import Navbar from '$lib/components/Navbar.svelte';
</script>

<header>
	<Navbar page={$page} />
</header>

<main>
	<slot />
</main>

<style>
	* {
		font-family: Arial;
		box-sizing: border-box;
		padding: 0;
		margin: 0;
	}
</style>
