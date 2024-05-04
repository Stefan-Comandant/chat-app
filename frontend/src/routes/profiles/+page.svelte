<script lang="ts">
	import { onMount } from 'svelte';
	import type { User } from '$lib/interfaces.ts';
	import { page } from '$app/stores';
	import { loading, settings } from '../../stores.ts';

	let profiles: User[] = [];

	onMount(() => {
		profiles = $page.data.profiles;
		$loading.goPast = true;
	});

	$: darkMode = !$settings.LightMode;
</script>

<h1 class:dark={!!darkMode}>Other Users</h1>
<div>
	{#each profiles as profile (profile.id)}
		<div>
			<a class="account" href="/profiles/{profile.id}" class:dark={!!darkMode}>
				<img src={profile.profilepicture} alt="Pfp" />
				<div>
					<div>{profile.username}</div>
					<span>{profile.about}</span>
				</div>
			</a>
		</div>
	{/each}
</div>

<style>
	@import '../../lib/css/profiles.css';
</style>
