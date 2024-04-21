<script lang="ts">
	import { onMount } from 'svelte';
	import { type ChatRoom, type MessageDate, type User } from '$lib/interfaces.ts';
	import { GetChatRooms } from '$lib/chat-rooms.ts';
	import { page } from '$app/stores';

	let rooms: ChatRoom[] = [];
	let dialog: any = {};
	let modal: any = {};
	let selectMode = 'view';
	let selectedRoom: ChatRoom = { admins: [], owner: '' };

	onMount(async () => {
		rooms = await GetChatRooms();
		if (!rooms) rooms = [];
		dialog.close();
	});

	function formatDate(dateStr: string): MessageDate {
		if (!dateStr) return { ofYear: '', ofDay: '' };
		const date = new Date(dateStr);
		const hour = date.getHours() > 12 ? date.getHours() - 12 : date.getHours();
		const minute = date.getMinutes() > 9 ? date.getMinutes() : '0' + date.getMinutes();
		const meridian = date.getHours() > 12 ? 'PM' : 'AM';
		const day = date.getDate() > 9 ? date.getDate() : '0' + date.getDate();
		const month = date.getMonth() > 9 ? date.getMonth() : '0' + date.getMonth();
		const year = date.getFullYear();

		const yearDate = `${day}-${month}-${year}`;
		const time = `${hour}:${minute} ${meridian}`;

		return {
			ofYear: yearDate,
			ofDay: time
		};
	}
</script>

<div>
	{#each rooms as room (room.id)}
		<div class="room-container">
			<a
				class="room"
				href="/chats/{room.id}"
				on:contextmenu|preventDefault={async (event) => {
					dialog.show();

					selectedRoom = room;

					dialog.style.top = event.pageY + 'px';
					dialog.style.left = event.pageX + 'px';
				}}
			>
				<div>{room.title}</div>
				<span>{room.description ? room.description : 'No description'}</span>
			</a>
		</div>
		<br />
	{:else}
		<div><span>Nothing to see, bitch</span></div>
	{/each}
</div>

<dialog class="popup" bind:this={dialog}>
	<div>
		{#if $page.data.USER.id && selectedRoom.admins?.includes($page.data.USER.id)}
			<button on:click={() => modal.showModal()}>Edit Room</button>
		{/if}
		{#if $page.data.USER.id && selectedRoom.owner === $page.data.USER.id}
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

<a href="/chats/new">Create New Room</a>

<style>
	@import '../../lib/css/chats.css';
</style>
