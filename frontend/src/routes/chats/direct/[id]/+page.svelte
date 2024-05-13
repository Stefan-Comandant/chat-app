<script lang="ts">
	import { page } from '$app/stores';
	import { onMount } from 'svelte';
	import type { Message, User } from '$lib/interfaces.ts';
	import { loading, settings } from '../../../../stores.ts';
	import MsgInput from '$lib/components/MsgInput.svelte';
	import MsgsContainer from '$lib/components/MsgsContainer.svelte';

	const id = $page.params.id;
	let peer: User = {};
	let messages: Message[] = [];
	let socket: WebSocket;

	$: darkMode = !$settings.LightMode;

	onMount(() => {
		socket = new WebSocket(`ws://localhost:7000/api/socket/${id}`);
		socket.onopen = () => {
			socket.onmessage = ({ data }: { data: string }) => {
				let message = JSON.parse(data);
				message.shortened = message.text.length > 1400;
				messages = [...messages, message];
			};
		};

		messages = $page.data.messages;
		peer = $page.data.peer;
		$loading.goPast = true;
	});
</script>

<svelte:head>
	<title>{peer.username}</title>
</svelte:head>

<div class="container" class:dark={!!darkMode}>
	<div class="room-title">
		<img src={peer.profilepicture} alt="pfp" class="room-picture" />
		<span>{peer.username}</span>
	</div>
	{#if messages.length > 0}
		<MsgsContainer {messages} currentRoomMembers={[$page.data.USER, peer]} />
		<MsgInput {id} {socket} />
	{/if}
</div>

<style>
	@import '../../../../lib/css/chat.css';
</style>
