<script lang="ts">
	import { goto } from '$app/navigation';
	import { Login } from '$lib/authentication.ts';
	import LoginForm from '$lib/components/forms/Login-Form.svelte';
	import type { HTTPResponse } from '$lib/interfaces.ts';

	let response: HTTPResponse = {
		response: ''
	};
</script>

<div class="container">
	<LoginForm
		{response}
		on:login={async (event) => {
			if (!event.detail.email || !event.detail.password) {
				response.status = 'error';
				response.response = "You can't submit an empty form";
				return;
			}
			response = await Login(event.detail);
			goto(`/authorize/${response.id}`);
		}}
	/>
</div>

<style>
	.container {
		position: absolute;
		left: 50%;
		top: 50%;
		transform: translate(-50%, -50%);
		width: fit-content;
		height: fit-content;
	}
</style>
