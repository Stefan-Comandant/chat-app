<script lang="ts">
	import { VerifyWithCode, Login } from '$lib/authentication.ts';
	import LoginForm from '$lib/components/forms/Login-Form.svelte';
  import type { VerificationSession, HTTPResponse } from "$lib/interfaces.ts"

  let response: HTTPResponse = {
    response: "",
  }
  
	let verification: VerificationSession = {};
</script>

<div class="body">
	<LoginForm response={response} on:login={async (event) => {
    response = await Login(event.detail)

    verification.userid = response.id;
  }}/>

	<div>
		<input type="text" bind:value={verification.code} />
		<button type="button" on:click={() => VerifyWithCode(verification)}>Verify</button>
	</div>
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
