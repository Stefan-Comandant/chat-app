<script lang="ts">
	import { page } from '$app/stores';
	import { onMount } from 'svelte';
	import type { ChatRoom, Message, User } from '$lib/interfaces.ts';
	import { loading, settings } from '../../../../stores.ts';
	import { formatDate } from '$lib/users.ts';
	import MsgInput from '$lib/components/MsgInput.svelte';
	import MsgsContainer from '$lib/components/MsgsContainer.svelte';

	const id: string = $page.params.id;
	let currentRoom: ChatRoom = {};
	let currentRoomMembers: User[] = [];
	let messages: Message[] = [];
	let socket: WebSocket;
	let dates: Map<number, string> = new Map();

	$: darkMode = !$settings.LightMode;

	onMount(() => {
		socket = new WebSocket(`ws://localhost:7000/api/socket/${id}`);
		socket.onopen = () => {
			socket.onmessage = ({ data }: { data: string }) => {
				let message = JSON.parse(data);
				message = computeDateDivider(message, messages.length - 1);
				messages = [...messages, message];
			};
		};

		messages = $page.data.messages;
		currentRoomMembers = $page.data.members;
		currentRoom = $page.data.room;
		$loading.goPast = true;

		messages = messages.map(computeDateDivider);
	});

	function computeDateDivider(msg: Message, i: number): Message {
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
	<MsgsContainer {messages} {currentRoomMembers} />
	<MsgInput {id} {socket} />
</div>

<style>
	@import '../../../../lib/css/chat.css';
</style>
