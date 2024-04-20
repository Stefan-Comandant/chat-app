<script lang="ts">
	import { VerifyWithCode, Login } from '$lib/authentication.ts';
	import LoginForm from '$lib/components/forms/Login-Form.svelte';
	import type { VerificationSession, HTTPResponse } from '$lib/interfaces.ts';

	let response: HTTPResponse = {
		response: ''
	};

	let verification: VerificationSession = {};
	let modal: HTMLDialogElement;
	let container: HTMLDivElement;
</script>

<div bind:this={container} class="container">
	<LoginForm
		{response}
		on:login={async (event) => {
			if (!event.detail.email || !event.detail.password) {
				response.status = 'error';
				response.response = "You can't submit an empty form";
				return;
			}
			response = await Login(event.detail);
			modal.show();
			container.style.display = 'none';
			verification.userid = response.id;
		}}
	/>
</div>

<dialog bind:this={modal}>
	<div>
		<div>
			<label for="code-input">Enter the verification code sent through your email</label>
			<input type="text" bind:value={verification.code} id="code-input" maxlength="8" />
		</div>
		<button
			type="button"
			on:click={async () => {
				response = await VerifyWithCode(verification);
			}}>Verify</button
		>
		<span class:error={response.status === 'error'} class:success={response.status === 'success'}
			>{response.response}</span
		>
	</div>
</dialog>

<style>
	.container {
		position: absolute;
		left: 50%;
		top: 50%;
		transform: translate(-50%, -50%);
	}

	dialog {
		border: 1px solid #b0b0b0;
		border-radius: 25px;
		min-height: 100px;
		height: 200px;
		position: absolute;
		top: 50%;
		bottom: 50%;
	}

	dialog > div,
	dialog > div > div {
		display: flex;
		flex-direction: column;
		justify-content: space-around;
		align-items: center;
	}

	dialog > div > div {
		gap: 10px;
	}

	dialog > div {
		gap: 50px;
	}

	dialog > div > div > input {
		border-radius: 20px;
		padding: 10px 20px;
		font-size: 20px;
		max-width: 160px;
		border: 1px solid #b0b0b0;
	}

	dialog > div > button {
		border: 1px solid #005fcf;
		background: #fff;
		color: #005fcf;
		font-size: 25px;
		transition:
			background 0.5s,
			color 0.5s;
		cursor: pointer;
		border-radius: 20px;
		min-width: 200px;
		padding: 10px 0;
	}

	dialog > div > button:focus,
	dialog > div > button:hover {
		background: #005fcf;
		border-color: #fff;
		color: #fff;
	}
</style>
