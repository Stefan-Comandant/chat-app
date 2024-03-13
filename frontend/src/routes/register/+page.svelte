<script lang="ts">
	import { VerifyWithCode } from '$lib/authentication.ts';
	import RegisterForm from '$lib/components/forms/Register-Form.svelte';

	let code = '';
</script>

<div class="container">
	<RegisterForm />
</div>

<input type="text" bind:value={code} />
<button type="button" on:click={() => VerifyWithCode(code)}>Verify</button>
