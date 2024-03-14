<script lang="ts">
	import { onMount } from 'svelte';
	import type { ChatRoom } from '$lib/interfaces.ts';
	import { AddChatRoom, GetChatRooms } from '$lib/chat-rooms.ts';
	import ChatForm from '$lib/components/forms/Chat-Form.svelte';

	let rooms: ChatRoom[] = [];

	onMount(async () => {
		rooms = await GetChatRooms();
		if (!rooms) rooms = [];
	});
</script>

<h1>Your Chat Rooms:</h1>
<div>
	{#each rooms as room (room.id)}
		<a href="/chats/{room.id}">
			<div>
				<span>Name: {room.title}</span>
				<br />
				<span>Description: {room.description}</span>
			</div>
		</a>
		<br />
	{:else}
		<div><span>Nothing to see, bitch</span></div>
	{/each}
</div>

<ChatForm
	on:addChatRoom={async (event) => {
		const room = await AddChatRoom(event.detail);
		if (room) rooms = [...rooms, room];
	}}
/>
