<script lang="ts">
	import ChatForm from '$lib/components/forms/Chat-Form.svelte';
	import { AddChatRoom } from '$lib/chat-rooms.ts';
</script>

<ChatForm
	on:addChatRoom={async (event) => {
		await AddChatRoom(event.detail);
	}}
/>
