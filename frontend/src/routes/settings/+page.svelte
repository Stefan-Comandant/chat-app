<script lang="ts">
	import type { Setting } from '$lib/interfaces.ts';
	import { onMount } from 'svelte';

	let settings: Setting;

	function UpdateSetting(setting: Setting) {
		localStorage.setItem('settings', JSON.stringify(setting));
	}

	onMount(() => {
		settings = JSON.parse(String(localStorage.getItem('settings')));
		if (!settings)
			settings = {
				LightMode: false
			};
		console.log(localStorage);
	});
</script>

<div>
	<label for="color-theme">Color Theme</label>
	<select id="color-theme">
		<option value="dark">Dark Mode</option>
		<option value="light">Light Mode</option>
	</select>
</div>

<style></style>
