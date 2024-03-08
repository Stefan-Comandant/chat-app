<script lang="ts">
	import { page } from '$app/stores';
	import { onMount } from 'svelte';
	import type { ChatRoom, Message } from '$lib/interfaces.ts';
	import { GetRoom } from "$lib/chat-rooms.ts"

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

	async function FetchMessages(id: number[]) {
		const body = JSON.stringify(id);
		const data = await fetch('http://localhost:7000/messages', {
			method: 'PUT',
			body: body,
			credentials: 'include',
			headers: {
				'Content-Type': 'application/json'
			}
		}).then((res) => res.json());
		if (data.status === 'success') return data.response;
		return [];
	}
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
	<input type="text" id="msg_input" bind:value={msg} />
	<button
		type="button"
		on:click={() => {
			const data = JSON.stringify({ text: msg, toid: parseInt(id) });
			socket.send(data);
			msg = '';
		}}>Send</button
	>
</div>
