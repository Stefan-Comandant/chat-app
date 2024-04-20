<script lang="ts">
	import type { User, ChatRoom, HTTPResponse } from '$lib/interfaces.ts';
	import { AddChatRoom } from '$lib/chat-rooms.ts';
	import { onMount, createEventDispatcher } from 'svelte';
	import EditButton from '../buttons/Edit-Button.svelte';
	import { GetUsers } from '$lib/users.ts';

	const dispatcher = createEventDispatcher();

	let users: User[] = [];
	let openModal = false;

	onMount(async () => {
		users = await GetUsers();
		if (!users) users = [];
	});

	let info: ChatRoom = {
		title: '',
		members: [],
		admins: []
	};

	function AddMember(event: any, id: number) {
		if (!event) return;

		const target = event.target;

		if (target.checked === true) {
			info.members = [...info.members, id];
		} else {
			info.members = info.members.filter((member) => member != id);
		}
	}

	function AddAdmin(event: any, id: number) {
		if (!event) return;
		const target = event.target;

		if (target.checked === true) {
			info.admins = [...info.admins, id];
		} else {
			info.admins = info.admins.filter((admin) => admin != id);
		}
	}

	let response: HTTPResponse = {
		response: ''
	};
</script>

<form
	on:submit|preventDefault={async () => {
		if (!info.title.length) return;

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
						<input type="checkbox" on:input={(event) => AddMember(event, user.id)} />
					</div>
					<div class="check">
						<div>Admin</div>
						<input type="checkbox" on:input={(event) => AddAdmin(event, user.id)} />
					</div>
				</div>
			</div>
		{/each}
		<div></div>
	</div>
</dialog>

<style>
	@import '../../css/authentication.css';

	.profile-picture {
		width: 40px;
		height: 40px;
		border-radius: 50%;
	}

	.members-display {
		width: fit-content;
		height: fit-content;
	}

	.member {
		display: flex;
		align-items: center;
		gap: 10px;
		font-size: 22px;
		margin-bottom: 10px;
	}

	.members-container {
		border: #000 solid 1px;
		border-radius: 10px;
		min-width: 316px;
		min-height: 161px;
		padding: 8px 16px;
		overflow-y: scroll;
	}

	dialog {
		width: 427px;
		height: 427px;
		top: 50%;
		right: -50%;
		transform: translate(-50%, -50%);
		border-radius: 25px;
		border: 1px solid red;
	}

	.account {
		border-radius: 20px;
		display: flex;
		align-items: center;
		justify-content: space-between;
		margin-bottom: 10px;
	}

	.details {
		border-radius: 20px;
		min-height: 58px;
		min-width: 287px;
		border: 1px solid #000;
		display: flex;
		gap: 10px;
		align-items: center;
		padding: 18px 16px;
	}

	.details div {
		font-size: 20px;
		display: flex;
		height: 55px;
		flex-direction: column;
		justify-content: space-between;
	}

	.details img {
		width: 55px;
		height: 55px;
	}

	.check {
		display: flex;
		align-items: center;
		justify-content: space-between;
	}

	.checks {
		height: 58px;
		display: flex;
		flex-direction: column;
		justify-content: space-between;
	}
	.edit-btn {
		transform: scale(0.4);
		background: none;
		border: none;
		cursor: pointer;
		width: fit-content;
		height: fit-content;
	}
	.members-title {
		display: flex;
		align-items: center;
		justify-content: center;
		height: 40px;
	}
</style>
