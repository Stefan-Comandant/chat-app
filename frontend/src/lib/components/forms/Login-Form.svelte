<script lang="ts">
	import type { User } from "$lib/interfaces.ts"
	import { Register } from "$lib/authentication.ts"

	let info : User = {}
</script>

<form on:submit|preventDefault={() => Login(info)}>
	<input type="text" placeholder="Email" bind:value={info.Email} />
	<input type="text" placeholder="Password" bind:value={info.Password} />
	<button type="submit">Submit</button>
</form>
