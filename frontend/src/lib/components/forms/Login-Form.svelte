<script lang="ts">
	import type { User, HTTPResponse } from '$lib/interfaces.ts';
	import { createEventDispatcher, onMount } from 'svelte';
	import { loading, settings } from '../../../stores.ts';
	import LoadingCircle from '../LoadingCircle.svelte';

	const dispatch = createEventDispatcher();

	let info: User = {
		email: '',
		password: ''
	};
	$: darkMode = !$settings.LightMode;
	onMount(() => {
		$loading.goPast = true;
	});
	export let response: HTTPResponse;

	let isLoading = false;
</script>

<form
	class:dark={!!darkMode}
	on:submit|preventDefault={async () => {
		isLoading = true;
		dispatch('login', info);
	}}
>
	<div>
		<input type="email" placeholder="Enter your email" bind:value={info.email} />
	</div>
	<div>
		<input type="text" placeholder="Enter your password" bind:value={info.password} />
	</div>
	<button type="submit" style="pointer-events: {isLoading ? 'none' : 'auto'}; ">
		{#if isLoading}
			<LoadingCircle />
		{:else}
			Submit
		{/if}
	</button>
	<span class:error={response.status === 'error'} class:success={response.status === 'success'}
		>{response.response}</span
	>
</form>

<style>
	@import '../../css/authentication.css';
</style>
