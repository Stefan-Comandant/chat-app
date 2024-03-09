<script lang="ts">
	import type { User } from "$lib/interfaces.ts"
	import { Register, VerifyWithCode } from "$lib/authentication.ts"
	import RegisterForm from "$lib/components/forms/Register-Form.svelte"

	let code = ""

	let info: User = {
		Username: '',
		About: '',
		Email: '',
		Password: '',
		Balance: 0,
		Currency: 'RON'
	};
	
</script>

<RegisterForm />

<input type="text" bind:value={code} />
<button	type="button" on:click={() => VerifyWithCode(code)}>Verify</button>
