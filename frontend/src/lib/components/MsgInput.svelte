<script lang="ts">
	import { page } from '$app/stores';

	export let id = '';
	let msg = '';
	export let socket: WebSocket;

	/*
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
*/
</script>

<div class="msg-input">
	<textarea bind:value={msg} spellcheck="false" />
	<button
		type="button"
		on:click={() => {
			if (!msg.length) return;
			const data = JSON.stringify({
				text: msg,
				toid: parseInt(id),
				type: $page.route.id?.includes('/direct') ? 'direct' : 'broadcast'
			});
			socket.send(data);

			msg = '';
		}}>Send</button
	>
</div>

<style>
	@import '../css/messages/input.css';
</style>
