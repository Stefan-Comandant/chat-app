<script lang="ts">
	import { page } from '$app/stores';
	import { onMount } from 'svelte';
	import type { User } from '$lib/interfaces.ts';
	import { loading } from '../../../stores.ts';

	let profile: User = {
		id: '',
		username: '',
		about: '',
		email: '',
		password: '',
		currency: '',
		balance: 0
	};

	onMount(async () => {
		profile = $page.data.profile;
		$loading.goPast = true;
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
