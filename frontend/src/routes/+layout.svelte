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
