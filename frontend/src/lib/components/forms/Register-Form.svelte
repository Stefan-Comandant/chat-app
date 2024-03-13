<script lang="ts">
	import type { User } from '$lib/interfaces.ts';
	import { Register } from '$lib/authentication.js';

	let info: User = {
		id: 0,
		username: '',
		about: '',
		email: '',
		password: '',
		currency: '',
		balance: 0
	};
</script>

<form on:submit|preventDefault={() => Register(info)}>
	<div>
		<input type="text" placeholder="Enter your username" bind:value={info.username} />
	</div>
	<div>
		<input type="text" placeholder="Enter your email" bind:value={info.email} />
	</div>
	<div>
		<input type="text" placeholder="Enter your password" bind:value={info.password} />
	</div>
	<button type="submit">Submit</button>
</form>

<style>
	@import '../../css/authentication.css';
</style>
