<script lang="ts">
	import { page } from '$app/stores';
	import { onMount } from 'svelte';
	import type { ChatRoom, Message, MessageDate, User } from '$lib/interfaces.ts';
	import { loading, settings } from '../../../stores.ts';

	const id: string = $page.params.id;
	let currentRoom: ChatRoom = {};
	let currentRoomMembers: User[] = [];
	let msg = '';
	let messages: Message[] = [];
	let socket: WebSocket;
	let dates: Map<number, string> = new Map();

	// Function to get the moment of the day and of the year when the message was sent
	function formatDate(dateStr: string): MessageDate {
		if (!dateStr) return { ofYear: '', ofDay: '' };

		// Set pointers in time
		const date = new Date(dateStr);
		const todayDate = new Date();
		const { hour, minute, meridian, day, month, year } = getDateValues(date);
		const time = `${hour}:${minute} ${meridian}`;
		const today = `${getDateValues(todayDate).day}-${getDateValues(todayDate).month}-${getDateValues(todayDate).year}`;
		const yearDate = `${day}-${month}-${year}`;

		return {
			ofDay: time,
			ofYear: today === yearDate ? 'Today' : yearDate
		};
	}

	function getDateValues(date: Date) {
		return {
			minute: date.getMinutes() > 9 ? date.getMinutes() : '0' + date.getMinutes(),
			hour: date.getHours() > 12 ? date.getHours() - 12 : date.getHours(),
			day: date.getDate() > 9 ? date.getDate() : '0' + date.getDate(),
			month: date.getMonth() > 9 ? date.getMonth() : '0' + date.getMonth(),
			year: date.getFullYear(),
			meridian: date.getHours() > 12 ? 'PM' : 'AM'
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

	$: darkMode = !$settings.LightMode;
	let msgContainer: HTMLDivElement;
	let showBtn = false;

	onMount(() => {
		socket = new WebSocket(`ws://localhost:7000/api/socket/${id}`);
		socket.onopen = () => {
			socket.onmessage = ({ data }: { data: string }) => {
				messages = [...messages, JSON.parse(data)];
				computeDateDivider(JSON.parse(data), messages.length - 1);
			};
		};

		messages = $page.data.messages;
		currentRoomMembers = $page.data.members;
		currentRoom = $page.data.room;
		$loading.goPast = true;
		msgContainer.onscroll = () => {
			if (msgContainer.scrollTop < msgContainer.scrollHeight - msgContainer.clientHeight) {
				showBtn = true;
			} else showBtn = false;
		};

		messages = messages.map(computeDateDivider);
	});

	function computeDateDivider(msg: Message, i: number) {
		const formattedDate = formatDate(String(msg.sentat));
		if (![...dates.values()].includes(String(formattedDate.ofYear))) {
			dates.set(i, String(formattedDate.ofYear));
		}

		msg.shortened = msg.text.length > 1400;

		return msg;
	}
</script>

<svelte:head>
	<title>{currentRoom.title}</title>
</svelte:head>

<div class="container" class:dark={!!darkMode}>
	<div class="room-title">
		<span>{currentRoom.title}</span>
	</div>
	<div class="msg-container" bind:this={msgContainer}>
		{#each messages as message, index (message.id)}
			{#if dates.get(index)}
				<div class="date-display">
					{dates.get(index)}
				</div>
			{/if}
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
							><a href="/profiles/{message.from}">{GetUsername(message.from, currentRoomMembers)}</a
							></span
						>
					{/if}
					<div class={message.shortened ? 'shortened' : ''}>
						{message.shortened ? message.text.split('').slice(0, 1400).join('') : message.text}
						{#if message.shortened}
							<button type="button" class="show-more" on:click={() => (message.shortened = false)}
								>Show more</button
							>
						{/if}
					</div>
					<span>{formatDate(String(message.sentat)).ofDay}</span>
				</div>
			</div>
		{/each}
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
		{#if showBtn}
			<button
				type="button"
				on:click={() => {
					msgContainer.scrollTo({
						top: msgContainer.scrollHeight
					});
				}}>↓</button
			>
		{/if}
	</div>
</div>

<style>
	@import '../../../lib/css/chat.css';
</style>
