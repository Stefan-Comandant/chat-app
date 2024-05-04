<script lang="ts">
	import type { Setting } from '$lib/interfaces.ts';
	import { onMount } from 'svelte';
	import { loading, settings } from '../../stores.ts';

	let setting: Setting;

	function UpdateSetting(setting: Setting) {
		localStorage.setItem('settings', JSON.stringify(setting));

		settings.update(() => setting);
	}

	onMount(() => {
		setting = JSON.parse(String(localStorage.getItem('settings')));
		if (!settings)
			setting = {
				LightMode: false
			};
		settings.set(setting);
		colorTheme = setting.LightMode ? 'light' : 'dark';
		$loading.goPast = true;
	});

	let colorTheme = 'dark';
</script>

<div>
	<label for="color-theme">Color Theme</label>
	<select
		id="color-theme"
		bind:value={colorTheme}
		on:change={() => {
			setting.LightMode = colorTheme === 'light';
			UpdateSetting(setting);
		}}
	>
		<option value="dark">Dark Mode</option>
		<option value="light">Light Mode</option>
	</select>
</div>

<style></style>
