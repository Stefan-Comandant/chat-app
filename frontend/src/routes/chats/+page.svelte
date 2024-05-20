<script lang="ts">
	import { type ChatRoom, type MessageDate, type User } from '$lib/interfaces.ts';
	import { page } from '$app/stores';
	import { onMount } from 'svelte';
	import { loading, settings } from '../../stores.ts';
	import { formatDate, getPeer } from '$lib/users.ts';
	import { AddChatRoom } from '$lib/chat-rooms.ts';
	import { goto } from '$app/navigation';

	let groups: ChatRoom[] = [];
	let chats: ChatRoom[] = [];
	let USER: User = {};
	let users: User[] = [];

	let dialog: HTMLDialogElement;
	let modal: HTMLDialogElement;
	let usersModal: HTMLDialogElement;
	let selectMode = 'view';
	let selectedRoom: ChatRoom = { admins: [], owner: '' };

	onMount(() => {
		groups = $page.data.groups;
		chats = $page.data.chats;
		USER = $page.data.USER;
		users = $page.data.users;
		$loading.goPast = true;
	});

	$: darkMode = !$settings.LightMode;

	let convos = 'rooms';
</script>

<div>
	<button
		class="convo-switch {convos == 'rooms' ? 'active' : ''}"
		on:click={() => {
			convos = 'rooms';
		}}>Rooms</button
	>
	<button
		class="convo-switch {convos == 'chats' ? 'active' : ''}"
		on:click={() => {
			convos = 'chats';
		}}>Chats</button
	>
	{#if convos == 'rooms'}
		{#each groups as group (group.id)}
			<div class="room-container" class:dark={!!darkMode}>
				<a
					class="room"
					class:dark={!!darkMode}
					href="/chats/group/{group.id}"
					on:contextmenu|preventDefault={async (event) => {
						dialog.show();

						selectedRoom = group;

						dialog.style.top = event.pageY + 'px';
						dialog.style.left = event.pageX + 'px';
					}}
				>
					<div class="pfp-input-container">
						<img src={group.profilepicture} alt="pfp" class="room-picture" />
					</div>
					<div>
						<div>{group.title}</div>
						<span>{group.description ? group.description : 'No description'}</span>
					</div>
				</a>
			</div>
			<br />
		{/each}
	{:else}
		{#each chats as chat (chat.id)}
			<div class="room-container" class:dark={!!darkMode}>
				<a
					class="room"
					class:dark={!!darkMode}
					href="/chats/direct/{chat.id}"
					style="flex-direction: column; align-items: flex-start; gap"
				>
					<div>{getPeer(users, chat.members, USER).username}</div>
					<span
						>{getPeer(users, chat.members, USER).about
							? getPeer(users, chat.members, USER).about
							: 'Masturbez!'}</span
					>
				</a>
			</div>
			<br />
		{/each}{/if}
</div>

<dialog class="popup" bind:this={dialog}>
	<div>
		{#if USER.id && selectedRoom.admins?.includes(USER.id)}
			<button on:click={() => modal.showModal()}>Edit Room</button>
		{/if}
		{#if USER.id && selectedRoom.owner === USER.id}
			<button>Delete Room</button>
		{/if}
		<button>Quit Room</button>
	</div>
</dialog>

{#if selectedRoom}
	<dialog bind:this={modal}>
		{#if selectMode === 'view'}
			<span>Title</span>
			<div>{selectedRoom.title}</div>
		{:else}
			<label for="title-input">Title</label>
			<input type="text" id="title-input" />
		{/if}
		{#if selectMode === 'view' && selectedRoom.description}
			<span>Description</span>
			<div>{selectedRoom.description}</div>
		{:else if selectedRoom.description}
			<label for="title-input">Description</label>
			<input type="text" id="title-input" />
		{/if}
		<span>Created At</span>
		<div>
			{formatDate(String(selectedRoom.createdat)).ofDay} on {formatDate(
				String(selectedRoom.createdat)
			).ofYear}
		</div>
		<span>Owner</span>
		<div>{selectedRoom.ownerData?.username}</div>
	</dialog>
{/if}

{#if convos == 'rooms'}
	<a href="/chats/new">Create New Room</a>
{:else}
	<button
		type="button"
		on:click={() => {
			usersModal.open ? usersModal.close() : usersModal.showModal();
		}}>New Chat</button
	>
{/if}

<dialog bind:this={usersModal} class:dark={!!darkMode}>
	<div style="display: flex; flex-direction: column; gap: 1.2rem">
		{#each [...users] as user (user.id)}
			<div class="account">
				<div class="details">
					<img class="profile-picture" alt="Pfp" src={user.profilepicture} />
					<div>
						<div>{user.username}</div>
						<span>{user.about}</span>
					</div>
				</div>
				<div>
					<button
						type="button"
						on:click={async () => {
							const response = await AddChatRoom({
								title: 'Direct Conversation',
								description: '',
								members: [user.id ? user.id : ''],
								admins: [],
								messages: [],
								type: 'direct'
							});

							if (response.status === 'success') {
								goto(`/chats/direct/${response.id}`);
							}
						}}>New Chat</button
					>
				</div>
			</div>
		{/each}
	</div>
</dialog>

<style>
	@import '../../lib/css/chats.css';
	@import '../../lib/css/users/dialog.css';
	@import '../../lib/css/img-preview.css';
</style>
