<script lang="ts">
	import { VerifyWithCode } from '$lib/authentication.ts';
	import LoginForm from '$lib/components/forms/Login-Form.svelte';

	let code = '';
</script>

<div class="body">
	<LoginForm />

	<div>
		<input type="text" bind:value={code} />
		<button type="button" on:click={() => VerifyWithCode(code)}>Verify</button>
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
