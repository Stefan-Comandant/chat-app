<script lang="ts">
	import { page } from '$app/stores';
	import { onMount } from 'svelte';
	import type { ChatRoom, Message } from '$lib/interfaces.ts';
	import { GetRoom, FetchMessages } from "$lib/chat-rooms.ts"

	const id: string = $page.params.id;
	let currentRoom: ChatRoom = {};
	let messages: Message[] = [];
	let socket: WebSocket;
	let msg = '';

	
	onMount(async () => {
		socket = new WebSocket(`ws://localhost:7000/socket/${id}`);
		socket.onopen = () => {
			socket.onmessage = (event) => {
				messages = [...messages, JSON.parse(event.data)];
			};
		};
		currentRoom = await GetRoom(id);
		messages = await FetchMessages(currentRoom.messages);
		if (!messages) messages = [];
	});

</script>

<div>
	<h2>Room {currentRoom.title}</h2>
	<div>
		{#each messages as message (message.id)}
			<div>
				{message.text}
			</div>
		{/each}
	</div>
	<input type="text" bind:value={msg} />
	<button
		type="button"
		on:click={() => {
			const data = JSON.stringify({ text: msg, toid: parseInt(id) });
			socket.send(data);
			msg = '';
		}}>Send</button
	>
</div>
