<script lang="ts">
	import type { User } from "$lib/interfaces.ts"
	import { Login, VerifyWithCode } from "$lib/authentication.ts" 
	import LoginForm from "$lib/components/forms/Login-Form.svelte"

	let code = ""

	let info: User = {
		Email: "",
		Password: "",
	}

</script>

<LoginForm />

<input type="text" bind:value={code} />
<button	type="button" on:click={() => VerifyWithCode(code)}>Verify</button>

