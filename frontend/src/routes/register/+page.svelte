<script lang="ts">
	import { VerifyWithCode, Register } from '$lib/authentication.ts';
	import RegisterForm from '$lib/components/forms/Register-Form.svelte';
	import type { VerificationSession, HTTPResponse } from '$lib/interfaces.ts';

	let verification: VerificationSession = {
		userid: 30
	};

	let response: HTTPResponse = {
		response: ''
	};
</script>

<div class="body">
	<div class="container">
		<RegisterForm
			{response}
			on:register={async (event) => {
				response = await Register(event.detail);

				verification.userid = response.id;
			}}
		/>
	</div>

	<input type="text" bind:value={verification.code} />
	<button
		type="button"
		on:click={async () => {
			response = await VerifyWithCode(verification);
		}}>Verify</button
	>
</div>

<style>
	.body {
		height: 100dvh;
		width: 100%;
		display: flex;
		flex-direction: column;
		justify-content: center;
		align-items: center;
	}
</style>
