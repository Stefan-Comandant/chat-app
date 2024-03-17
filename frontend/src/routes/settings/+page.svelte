<script lang="ts">
	import type { Setting } from '$lib/interfaces.ts';

	let settings: string | Setting | null = window.localStorage.getItem('settings');

	if (!settings) {
		settings = {
			LightMode: false
		};
	} else {
		settings = JSON.parse(settings);
	}

	function UpdateSetting() {
		localStorage.setItem('settings', JSON.stringify(settings));
	}
</script>

<div>
	<label for="color-theme">Color Theme</label>
	<select id="color-theme">
		<option value="dark">Dark Mode</option>
		<option value="light">Light Mode</option>
	</select>
</div>

<style></style>
