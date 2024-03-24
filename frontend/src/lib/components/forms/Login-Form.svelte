<script lang="ts">
	import type { User } from '$lib/interfaces.ts';
	import { Login } from '$lib/authentication.ts';

	let info: User = {
		username: '',
		email: '',
		password: '',
	};
</script>

<form on:submit|preventDefault={() => Login(info)}>
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
