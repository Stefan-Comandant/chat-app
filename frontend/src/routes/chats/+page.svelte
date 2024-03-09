<script lang="ts">
	import { onMount } from "svelte"
	import type { ChatRoom, User } from '$lib/interfaces.ts';
	import { AddChatRoom, GetChatRooms, GetUsers } from "$lib/chat-rooms.ts"

	let users: User[] = []
	let rooms: ChatRoom[] = [];

	onMount(async () => {
		rooms = await GetChatRooms()
		if (!rooms) rooms = []
		users = await GetUsers()
		if (!users) users = []
	})

	let info : ChatRoom = {
		members: [],
		admins: [],
		messages: [],
	}

	function AddMember(event: any, id: number) {
		if (!event) return

		const target = event.target

		if (target.checked === true) {
			info.members.push(id)
		} else {
			info.members = info.members.filter(member => member != id)
		}
	}

	function AddAdmin(event: any, id: number) {
		if (!event) return
		const target = event.target

		if (target.checked === true) {
			info.admins.push(id)
		} else {
			info.admins = info.admins.filter(admin => admin != id)
		}
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

<form on:submit|preventDefault={async () => {
		const room = await AddChatRoom(info)
		if (room) rooms = [...rooms, room]
	}
}>
	<input type="text" placeholder="Name" bind:value={info.title} />
	<input type="text" placeholder="Description" bind:value={info.description} />
	<br />
	<span>Members:</span>
	<ul>
		{#each users as user (user.id)}
			<li>
				<span>
					<span>{user.username}</span>
					<input type="checkbox"on:input={(event) => AddMember(event, user.id)}/>
				</span>
				<br />
				<span>
					<span>Admin</span>
					<input type="checkbox" on:input={(event) => AddAdmin(event, user.id)}/>
				</span>
			</li>
		{/each}
	</ul>
	<button type="button">Add members</button>

	<br />
	<button type="submit">Create chat room</button>
</form>
