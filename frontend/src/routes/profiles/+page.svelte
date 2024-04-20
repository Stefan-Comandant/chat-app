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
	.account {
		display: flex;
		text-decoration: none;
		color: inherit;
		width: fit-content;
		padding: 10px 15px;
		gap: 10px;
		border-radius: 25px;
	}

	.account > div > span {
		color: #a0a0a0;
	}
	.account > div {
		display: flex;
		flex-direction: column;
		justify-content: space-around;
	}

	.account:hover {
		border: #000 1px solid;
	}

	.account > img {
		width: 55px;
		height: 55px;
		border-radius: 50%;
	}
</style>
