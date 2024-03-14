<script lang="ts">
	import { page } from '$app/stores';
	import { onMount } from 'svelte';
	import type { User } from '$lib/interfaces.ts';
	import { FetchConfig } from '$lib/interfaces.ts';

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

	async function GetProfileByID() {
		const data = await fetch(`http://localhost:7000/users/${id}`, FetchConfig).then((res) =>
			res.json()
		);
		if (data.status === 'success') return data.response;
		return {};
	}

	onMount(async () => {
		profile = await GetProfileByID();
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

<div>
	<h2>{profile.username}</h2>
	<p>{profile.about}</p>
</div>
