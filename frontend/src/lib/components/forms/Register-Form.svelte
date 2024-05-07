<script lang="ts">
	import type { User, HTTPResponse } from '$lib/interfaces.ts';
	import { createEventDispatcher, onMount } from 'svelte';
	import { loading, settings } from '../../../stores.ts';

	const dispatch = createEventDispatcher();

	let info: User = {
		username: '',
		email: '',
		password: ''
	};

	export let response: HTTPResponse = {};
	$: darkMode = !$settings.LightMode;

	let fileInput: any;

	let showImage = false;
	let image: HTMLImageElement;

	onMount(() => {
		$loading.goPast = true;
	});
</script>

<form
	class:dark={darkMode}
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
					image.setAttribute('src', String(result));
					info.profilepicture = String(result);
				});
			}}
		/>
		{#if showImage}
			<img bind:this={image} alt="pfp" />
		{/if}
	</div>
	<div>
		<input type="email" placeholder="Enter your email" bind:value={info.email} />
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
	@import '../../css/img-preview.css';
</style>
