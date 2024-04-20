<script lang="ts">
	import { page } from '$app/stores';
	import { onMount } from 'svelte';
	import type { User } from '$lib/interfaces.ts';
	import { GetProfileByID } from '$lib/users.ts';

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
		profile = await GetProfileByID(parseInt(id));
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
	@import '../../../lib/css/profile.css';
</style>
