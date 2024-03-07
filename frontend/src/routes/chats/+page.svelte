<script lang="ts">
	import { onMount } from "svelte"
	import type { ChatRoom } from '$lib/interfaces.ts';

	let rooms: ChatRoom[] = [];

	async function GetChatRooms() {
		const response = await fetch('http://localhost:7000/rooms', {
			method: 'GET',
			credentials: 'include',
			headers: {
				'Content-Type': 'application/json'
			}
		}).then((res) => res.json());
		if (response.status === "success") return response.response
		return []
	}

	onMount(async () => {
		rooms = await GetChatRooms()
		if (!rooms) rooms = []
	})

	let info : ChatRoom = {
		members: [9, 10],
		admins: [9, 10],
		owner: 9,
		messages: [],
	}

	async function AddChatRoom() {
		const response = await fetch("http://localhost:7000/rooms", {
			method: "POST",
			body: JSON.stringify(info),
			credentials: "include",
			headers: {
				"Content-Type": "application/json"
			}
		}).then(res => res.json())
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

<form on:submit|preventDefault={AddChatRoom}>
	<input type="text" placeholder="Name" bind:value={info.title} />
	<input type="text" placeholder="Description" bind:value={info.description} />


	<button type="submit">Add</button>
</form>
