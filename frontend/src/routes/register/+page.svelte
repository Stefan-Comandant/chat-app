<script lang="ts">
	import { VerifyWithCode } from '$lib/authentication.ts';
	import RegisterForm from '$lib/components/forms/Register-Form.svelte';

	let code = '';
</script>

<div class="body">
	<div class="container">
		<RegisterForm />
	</div>

	<input type="text" bind:value={code} />
	<button type="button" on:click={() => VerifyWithCode(code)}>Verify</button>
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
