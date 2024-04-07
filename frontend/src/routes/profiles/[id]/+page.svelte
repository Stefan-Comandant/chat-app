<script lang="ts">
	import { page } from '$app/stores';
	import { onMount } from 'svelte';
	import type { User } from '$lib/interfaces.ts';
	import { FetchConfig } from '$lib/interfaces.ts';
  import { GetProfileByID } from "$lib/users.ts"

	let profile: User = {
		id: 0,
		username: '',
		about: '',
		email: '',
		password: '',
		currency: '',
		balance: 0
	};
	const id: string = $page.params.id;

	onMount(async () => {
		profile = await GetProfileByID(id);
		if (!profile)
			profile = {
				id: 0,
				username: '',
				about: '',
				email: '',
				password: '',
				currency: '',
				balance: 0
			};
	});

</script>

<div class="container">
	<img src={profile.profilepicture} alt="Pfp" />
	<div>
		<div>{profile.username}</div>
		<span>{profile.about}</span>
	</div>
</div>

<style>
	.container img {
		border-radius: 50%;
		width: 200px;
		height: 200px;
	}

	.container {
		display: flex;
		gap: 20px;
		align-items: center;
	}

	.container div span {
		color: #a0a0a0;
	}

	.container > div > div {
		font-weight: 700;
		font-size: 40px;
	}

	.container > div {
		display: flex;
		height: 100px;
		width: fit-content;
		flex-direction: column;
		justify-content: space-around;
		align-items: flex-start;
		padding: 10px 20px;
		border-radius: 25px;
	}
</style>
