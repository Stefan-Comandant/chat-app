<script lang="ts">
	import type { Setting } from '$lib/interfaces.ts';
	import { onMount } from 'svelte';
	import { store } from '../../stores.ts';

	let settings: Setting;

	function UpdateSetting(setting: Setting) {
		localStorage.setItem('settings', JSON.stringify(setting));

		store.update(() => setting);
	}

	onMount(() => {
		settings = JSON.parse(String(localStorage.getItem('settings')));
		if (!settings)
			settings = {
				LightMode: false
			};
		store.set(settings);
		colorTheme = settings.LightMode ? 'light' : 'dark';
	});

	let colorTheme = 'dark';
</script>

<div>
	<label for="color-theme">Color Theme</label>
	<select
		id="color-theme"
		bind:value={colorTheme}
		on:change={() => {
			settings.LightMode = colorTheme === 'light';
			UpdateSetting(settings);
		}}
	>
		<option value="dark">Dark Mode</option>
		<option value="light">Light Mode</option>
	</select>
</div>

<style></style>
