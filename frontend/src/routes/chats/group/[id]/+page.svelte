<script lang="ts">
	import { page } from '$app/stores';
	import { onMount } from 'svelte';
	import type { ChatRoom, Message, User } from '$lib/interfaces.ts';
	import { loading, settings } from '../../../../stores.ts';
	import MsgInput from '$lib/components/MsgInput.svelte';
	import MsgsContainer from '$lib/components/MsgsContainer.svelte';

	const ROOM_ID: string = $page.params.id;
	let currentRoom: ChatRoom = {};
	let currentRoomMembers: User[] = [];
	let messages: Message[] = [];
	let socket: WebSocket;

	$: darkMode = !$settings.LightMode;

	onMount(() => {
		socket = new WebSocket(`ws://localhost:7000/api/socket/${ROOM_ID}`);
		socket.onopen = () => {
			socket.onmessage = ({ data }: { data: string }) => {
				let message = JSON.parse(data);
				message.shortened = message.text.length > 1400;
				messages = [...messages, message];
			};
		};

		messages = $page.data.messages;
		currentRoomMembers = $page.data.members;
		currentRoom = $page.data.room;
		Object.freeze(currentRoom);
		$loading.goPast = true;
	});
</script>

<svelte:head>
	<title>{currentRoom.title}</title>
</svelte:head>

<div class="container" class:dark={!!darkMode}>
	<div class="room-title">
		<img src={currentRoom.profilepicture} alt="pfp" class="room-picture" />
		<span>{currentRoom.title}</span>
	</div>
	<MsgsContainer {messages} {currentRoomMembers} />
	<MsgInput id={ROOM_ID} {socket} />
</div>

<style>
	@import '../../../../lib/css/messages/chat.css';
</style>
