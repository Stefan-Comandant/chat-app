<script lang="ts">
	import { page } from '$app/stores';
	import Navbar from '$lib/components/Navbar.svelte';
	import { onDestroy, onMount } from 'svelte';
	import { store } from '../stores.ts';
	import type { Setting } from '$lib/interfaces.ts';

	let unsubscribe = () => {};
	let settings: Setting = {
		LightMode: true
	};

	onMount(() => {
		const result = localStorage.getItem('settings');
		if (result) {
			settings = JSON.parse(String(result));
		}
		if ($store !== settings) {
			store.update(() => settings);
		}
		unsubscribe = store.subscribe((value) => {
			const backgroundColor = value.LightMode ? '#ffffff' : '#121212';

			document.body.style.backgroundColor = backgroundColor;
		});
	});

	onDestroy(() => {
		unsubscribe();
	});
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

	header {
		height: 60px;
		z-index: 1000;
		position: fixed;
		top: 0;
		left: 0;
		right: 0;
	}

	main {
		padding-top: 60px;
	}

	:global(body) {
		background: #121212;
	}
</style>
