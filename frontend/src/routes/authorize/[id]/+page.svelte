<script lang="ts">
	import { VerifyWithCode } from '$lib/authentication.ts';
	import { type HTTPResponse, type VerificationSession } from '$lib/interfaces.ts';
	import { onMount } from 'svelte';
	import { page } from '$app/stores';
	import { goto } from '$app/navigation';

	let response: HTTPResponse = {
		response: ''
	};

	let verification: VerificationSession = {};

	onMount(() => {
		verification.userid = parseInt($page.params.id);
	});
</script>

<div class="container">
	<div>
		<label for="code-input">Enter the verification code sent through your email</label>
		<input type="text" bind:value={verification.code} id="code-input" maxlength="8" />
	</div>
	<button
		type="button"
		on:click={async () => {
			response = await VerifyWithCode(verification);

			goto('/');
		}}>Verify</button
	>
	<span class:error={response.status === 'error'} class:success={response.status === 'success'}
		>{response.response}</span
	>
</div>

<style>
	@import '../../../lib/css/authorization.css';
</style>
