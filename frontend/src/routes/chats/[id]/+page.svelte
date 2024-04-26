<script lang="ts">
	import { page } from '$app/stores';
	import { onMount } from 'svelte';
	import type { ChatRoom, Message, MessageDate, User } from '$lib/interfaces.ts';
	import { FetchConfig } from '$lib/interfaces.ts';
	import { store } from '../../../stores.ts';

	const id: string = $page.params.id;
	let currentRoom: ChatRoom = {};
	let currentRoomMembers: User[] = [];
	let msg = '';
	let messages: Message[] = [];
	let socket: WebSocket;
	let datesGroup: string[] = [];

	onMount(async () => {
		socket = new WebSocket(`ws://localhost:7000/api/socket/${id}`);
		socket.onopen = () => {
			socket.onmessage = (event) => {
				messages = [...messages, JSON.parse(event.data)];
				datesGroup = [];
			};
		};
		messages = $page.data.messages;
		currentRoomMembers = $page.data.members;
		currentRoom = $page.data.room;
	});

	function formatDate(dateStr: string, goal: string): MessageDate {
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

		if (datesGroup.indexOf(yearDate) == -1) {
			if (goal !== 'time') datesGroup = [...datesGroup, yearDate];

			return {
				ofDay: time,
				ofYear: yearDate
			};
		}

		return {
			ofYear: '',
			ofDay: time
		};
	}

	function GetUsername(from: string = '', members: User[] = []): string {
		const member = members.filter((member) => member.id === from)[0];
		if (!member || !member.username) return '';
		return member.username;
	}

	function getProfilePicture(id: string = '', members: User[] = []): string {
		const result = members.filter((member) => member.id === id)[0];
		if (!result || !result.profilepicture) return '';
		return result.profilepicture;
	}

	$: darkMode = !$store.LightMode;
</script>

<svelte:head>
	<title>{currentRoom.title}</title>
</svelte:head>

<div class="container" class:dark={!!darkMode}>
	<div>
		<div class="room-title">
			<span>{currentRoom.title}</span>
		</div>
		<div class="msg-container">
			{#each messages as message (message.id)}
				<div>
					{#if message.from != $page.data.USER.id}
						<img
							class="msg-profile-picture"
							alt="pfp"
							src={getProfilePicture(message.from, currentRoomMembers)}
						/>
					{/if}
					<div class="msg-content" class:sent-by-me={$page.data.USER.id === message.from}>
						{#if message.from !== $page.data.USER.id}
							<span
								><a href="/profiles/{message.from}"
									>{GetUsername(message.from, currentRoomMembers)}</a
								></span
							>
						{/if}
						<div>
							{message.text}
						</div>
						<span>{formatDate(String(message.sentat), 'time').ofDay}</span>
					</div>
				</div>
				{#if formatDate(String(message.sentat), 'time').ofYear}
					<div
						style="display: {messages[messages.length - 1].id === message.id ? 'none' : 'auto'}"
						class="date-display"
					>
						{formatDate(String(message.sentat), 'date').ofYear}
					</div>
				{/if}
			{/each}
		</div>
	</div>

	<div class="msg-input">
		<textarea bind:value={msg} />
		<button
			type="button"
			on:click={() => {
				if (!msg.length) return;
				const data = JSON.stringify({ text: msg, toid: parseInt(id) });
				socket.send(data);

				msg = '';
			}}>Send</button
		>
	</div>
</div>

<style>
	@import '../../../lib/css/chat.css';
</style>
