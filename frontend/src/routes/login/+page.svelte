<script lang="ts">
	import type { User } from "$lib/interfaces.ts"
	import { Login, VerifyWithCode } from "$lib/authentication.ts" 

	let code = ""

	let info: User = {
		Email: "",
		Password: "",
	}

</script>

<form on:submit|preventDefault={() => Login(info)}>
	<input type="text" placeholder="Email" bind:value={info.Email} />
	<input type="text" placeholder="Password" bind:value={info.Password} />
	<button type="submit">Submit</button>
</form>

<input type="text" bind:value={code} />
<button	type="button" on:click={() => VerifyWithCode(code)}>Verify</button>

