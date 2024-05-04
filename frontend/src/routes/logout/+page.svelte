<script>
	import { Logout } from '$lib/authentication.ts';
	import { onMount } from 'svelte';
	import { loading } from '../../stores.ts';

	onMount(() => {
		$loading.goPast = true;
	});
</script>

<button type="button" on:click={Logout}>Logout</button>
