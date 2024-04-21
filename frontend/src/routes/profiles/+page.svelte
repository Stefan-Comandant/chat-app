<script lang="ts">
	import { onMount } from 'svelte';
	import type { User } from '$lib/interfaces.ts';
	import { GetProfiles } from '$lib/users.ts';

	let profiles: User[] = [];

	onMount(async () => {
		profiles = await GetProfiles();
		if (!profiles) profiles = [];
	});
</script>

<h1>Other Users</h1>
<div>
	{#each profiles as profile (profile.id)}
		<div>
			<a class="account" href="/profiles/{profile.id}">
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
