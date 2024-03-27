<script lang="ts">
	import type { User, HTTPResponse } from '$lib/interfaces.ts';
	import { Login } from '$lib/authentication.ts';

	let info: User = {
		username: '',
		email: '',
		password: '',
	};

  let response : HTTPResponse = {
    response: "",
  }

</script>

<form on:submit|preventDefault={async () => {
  response = await Login(info)
}}>
	<div>
		<input type="text" placeholder="Enter your email" bind:value={info.email} />
	</div>
	<div>
		<input type="text" placeholder="Enter your password" bind:value={info.password} />
	</div>
	<button type="submit">Submit</button>
  <span class:error={response.status === "error"} class:success={response.status === "success"}>{response.response}</span>
</form>

<style>
	@import '../../css/authentication.css';

  .error{
    color: red;
  }

  .success{
    color: lightgreen;
  }
</style>
