<script lang="ts">
	import { onMount } from 'svelte';
	import type { ChatRoom } from '$lib/interfaces.ts';
	import { GetChatRooms } from '$lib/chat-rooms.ts';

	let rooms: ChatRoom[] = [];

	onMount(async () => {
		rooms = await GetChatRooms();
		if (!rooms) rooms = [];
	});
</script>

<h1>Your Chat Rooms:</h1>
<div>
	{#each rooms as room (room.id)}
		<div class="room-container">
			<a class="room" href="/chats/{room.id}">
				<div>{room.title}</div>
				<span>{room.description ? room.description : 'No description'}</span>
			</a>
		</div>
		<br />
	{:else}
		<div><span>Nothing to see, bitch</span></div>
	{/each}
</div>

<a href="/chats/new">Create New Room</a>

<style>
	.room-container {
		display: flex;
		flex-direction: column;
	}

	.room {
		border-radius: 10px;
		border: #000 solid 1px;
		width: fit-content;
		min-width: 425px;
		padding: 15px 21px;
	}

	.room span {
		color: #b0b0b0;
		font-size: 12px;
	}

	.room {
		text-decoration: none;
		color: inherit;
	}
</style>
