<script lang="ts">
	import type { User, HTTPResponse } from '$lib/interfaces.ts';
	import { Register } from '$lib/authentication.ts';
	import { createEventDispatcher } from 'svelte';

	const dispatch = createEventDispatcher();

	let info: User = {
		username: '',
		email: '',
		password: ''
	};

	export let response: HTTPResponse = {};

	let fileInput: any;

	let showImage = false;
	let image;
</script>

<form
	on:submit|preventDefault={() => {
		dispatch('register', info);
	}}
>
	<div>
		<input type="text" placeholder="Enter your username" bind:value={info.username} />
	</div>
	<div class="pfp-input-container">
		<label for="file-input">Enter a profile picture</label>
		<input
			id="file-input"
			type="file"
			accept=".jpg, .jpeg, .png"
			bind:this={fileInput}
			on:change={(event) => {
				const file = fileInput.files[0];

				if (file) showImage = true;

				const reader = new FileReader();
				reader.readAsDataURL(file);
				reader.addEventListener('load', () => {
					const result = reader.result;
					image.setAttribute('src', result);
					info.profilepicture = result;
				});
			}}
		/>
		{#if showImage}
			<img bind:this={image} />
		{/if}
	</div>
	<div>
		<input type="text" placeholder="Enter your email" bind:value={info.email} />
	</div>
	<div>
		<input type="text" placeholder="Enter your password" bind:value={info.password} />
	</div>
	<button type="submit">Submit</button>
	<span class:error={response.status === 'error'} class:success={response.status === 'success'}
		>{response.response}</span
	>
</form>

<style>
	@import '../../css/authentication.css';
	.error {
		color: red;
	}

	.success {
		color: lightgreen;
	}

	.pfp-input-container {
		display: flex;
		justify-content: space-between;
		align-items: center;
		gap: 20px;
	}

	.pfp-input-container label {
		font-size: 20px;
		color: #a0a0a0;
	}

	.pfp-input-container input {
		display: none;
	}

	.pfp-input-container img {
		width: 60px;
		height: 60px;
		border: 1px solid black;
		border-radius: 50%;
	}
</style>
