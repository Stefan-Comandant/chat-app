<script lang="ts">
	import { type ChatRoom, type MessageDate, type User } from '$lib/interfaces.ts';
	import { page } from '$app/stores';
	import { onMount } from 'svelte';
	import { loading, settings } from '../../stores.ts';
	import { GetUsername, formatDate, getPeer } from '$lib/users.ts';
	import ChatForm from '$lib/components/forms/Chat-Form.svelte';
	import { AddChatRoom } from '$lib/chat-rooms.ts';

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
					<div>{group.title}</div>
					<span>{group.description ? group.description : 'No description'}</span>
				</a>
			</div>
			<br />
		{/each}
	{:else}
		{#each chats as chat (chat.id)}
			<div class="room-container" class:dark={!!darkMode}>
				<a class="room" class:dark={!!darkMode} href="/chats/direct/{chat.id}">
					<div>{getPeer(users, USER).username}</div>
					<span>{getPeer(users, USER).about ? getPeer(users, USER).about : 'Masturbez!'}</span>
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
	<div>
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

							console.log(response);
						}}>New Chat</button
					>
				</div>
			</div>
		{/each}
	</div>
</dialog>

<style>
	@import '../../lib/css/chats.css';
	@import '../../lib/css/new-chat.css';
</style>
