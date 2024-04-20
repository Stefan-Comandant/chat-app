<script lang="ts">
	import type { User, HTTPResponse } from '$lib/interfaces.ts';
	import { createEventDispatcher } from 'svelte';

	const dispatch = createEventDispatcher();

	let info: User = {
		email: '',
		password: ''
	};

	export let response: HTTPResponse;
</script>

<form
	on:submit|preventDefault={async () => {
		dispatch('login', info);
	}}
>
	<div>
		<input type="text" placeholder="Enter your email" bind:value={info.email} />
	</div>
	<div>
		<input type="text" placeholder="Enter your password" bind:value={info.password} />
	</div>
	<button type="submit">Submit</button>
	<span class:error={response.status === 'error'} class:success={response.status === 'success'}
		>{response.response}</span
	>
</form>

<style>
	@import '../../css/authentication.css';
</style>
