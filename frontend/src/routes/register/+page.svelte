<script lang="ts">
	import { goto } from '$app/navigation';
	import { Register } from '$lib/authentication.ts';
	import RegisterForm from '$lib/components/forms/Register-Form.svelte';
	import type { HTTPResponse } from '$lib/interfaces.ts';

	let response: HTTPResponse = {
		response: ''
	};
</script>

<div class="container">
	<RegisterForm
		{response}
		on:register={async (event) => {
			response = await Register(event.detail);

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
