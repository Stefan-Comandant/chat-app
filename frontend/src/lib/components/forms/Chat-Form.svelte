<script lang="ts">
	import type { User, ChatRoom, HTTPResponse } from '$lib/interfaces.ts';
	import { AddChatRoom } from '$lib/chat-rooms.ts';
	import { onMount } from 'svelte';
	import EditButton from '../buttons/Edit-Button.svelte';
	import { GetUsers } from '$lib/users.ts';
	let users: User[] = [];
	let openModal = false;

	onMount(async () => {
		users = await GetUsers();
		if (!users) users = [];
	});

	let info: ChatRoom = {
		members: [],
		admins: []
	};

	function AddMember(event: any, id: string = '') {
		console.log(id);
		if (!event || !id) return;

		const target = event.target;

		if (target.checked === true) {
			info.members = [...info.members, id];
		} else {
			info.members = info.members?.filter((member) => member != id);
		}
	}

	function AddAdmin(event: any, id: string = '') {
		if (!event || !id) return;
		const target = event.target;

		if (target.checked === true) {
			info.admins = [...info.admins, id];
		} else {
			info.admins = info.admins?.filter((admin) => admin != id);
		}
	}

	let response: HTTPResponse = {
		response: ''
	};
</script>

<form
	on:submit|preventDefault={async () => {
		if (!info.title?.length) return;

		response = await AddChatRoom(info);
	}}
>
	<div>
		<input type="text" placeholder="Name" bind:value={info.title} />
	</div>
	<div>
		<input type="text" placeholder="Description" bind:value={info.description} />
	</div>
	<div class="members-display">
		<div class="members-title">
			<span>Members</span>
			<button class="edit-btn" type="button" on:click={() => (openModal = !openModal)}
				><EditButton /></button
			>
		</div>
		<div class="members-container">
			{#each info.members as member (member)}
				<div class="member">
					<img
						class="profile-picture"
						src={users.filter((user) => user.id === member)[0].profilepicture}
						alt="Pfp"
					/>
					<span>{users.filter((user) => user.id === member)[0].username}</span>
				</div>
			{/each}
		</div>
	</div>
	<button type="submit">Submit</button>
	<span class:error={response.status === 'error'} class:success={response.status === 'success'}
		>{response.response}</span
	>
</form>

<dialog open={openModal}>
	<div>
		{#each users as user (user.id)}
			<div class="account">
				<div class="details">
					<img class="profile-picture" alt="Pfp" src={user.profilepicture} />
					<div>
						<div>{user.username}</div>
						<span>{user.about}</span>
					</div>
				</div>

				<div class="checks">
					<div class="check">
						<div>Member</div>
						<input
							type="checkbox"
							on:input={(event) => {
								AddMember(event, user.id);
								console.log(event, user.id);
							}}
						/>
					</div>
					<div class="check">
						<div>Admin</div>
						<input
							type="checkbox"
							on:input={(event) => {
								AddAdmin(event, user.id);
								console.log(event, user.id);
							}}
						/>
					</div>
				</div>
			</div>
		{/each}
		<div></div>
	</div>
</dialog>

<style>
	@import '../../css/authentication.css';
	@import '../../css/new-chat.css';
</style>
