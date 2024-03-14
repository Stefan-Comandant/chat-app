<script lang="ts">
	import type { User, ChatRoom } from '$lib/interfaces.ts';
	import { GetUsers } from '$lib/chat-rooms.ts';
	import { onMount, createEventDispatcher } from 'svelte';

	const dispatcher = createEventDispatcher();

	let users: User[] = [];

	onMount(async () => {
		users = await GetUsers();
		if (!users) users = [];
	});

	let info: ChatRoom = {
		id: 0,
		title: '',
		createdat: '',
		description: '',
		owner: 0,
		members: [],
		admins: [],
		messages: []
	};

	function AddMember(event: any, id: number) {
		if (!event) return;

		const target = event.target;

		if (target.checked === true) {
			info.members.push(id);
		} else {
			info.members = info.members.filter((member) => member != id);
		}
	}

	function AddAdmin(event: any, id: number) {
		if (!event) return;
		const target = event.target;

		if (target.checked === true) {
			info.admins.push(id);
		} else {
			info.admins = info.admins.filter((admin) => admin != id);
		}
	}
</script>

<form
	on:submit|preventDefault={() => {
		dispatcher('addChatRoom', info);
	}}
>
	<input type="text" placeholder="Name" bind:value={info.title} />
	<input type="text" placeholder="Description" bind:value={info.description} />
	<br />
	<span>Members:</span>
	<ul>
		{#each users as user (user.id)}
			<li>
				<span>
					<span>{user.username}</span>
					<input type="checkbox" on:input={(event) => AddMember(event, user.id)} />
				</span>
				<br />
				<span>
					<span>Admin</span>
					<input type="checkbox" on:input={(event) => AddAdmin(event, user.id)} />
				</span>
			</li>
		{/each}
	</ul>
	<br />
	<button type="submit">Create chat room</button>
</form>
