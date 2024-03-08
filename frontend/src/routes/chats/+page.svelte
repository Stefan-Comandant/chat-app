<script lang="ts">
	import { onMount } from "svelte"
	import type { ChatRoom, User } from '$lib/interfaces.ts';
	import { AddChatRoom, GetChatRooms } from "$lib/chat-rooms.ts"

	let users: User[] = []
	let rooms: ChatRoom[] = [];

	onMount(async () => {
		rooms = await GetChatRooms()
		if (!rooms) rooms = []
	})

	let info : ChatRoom = {
		members: [],
		admins: [],
		messages: [],
	}
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

<form on:submit|preventDefault={() => AddChatRoom(info)}>
	<input type="text" placeholder="Name" bind:value={info.title} />
	<input type="text" placeholder="Description" bind:value={info.description} />
	<ul>
		{#each users as user}
			<li>
				<span>
					{user.username}
				</span>
				<br />
				<label for="admin-checkbox">Admin</label>
				<input type="checkbox" id="admin-checkbox"/>
			</li>
		{/each}
	</ul>
	<button type="button">Add members</button>

	<br />
	<button type="submit">Create chat room</button>
</form>
