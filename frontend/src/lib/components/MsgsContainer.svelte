<script lang="ts">
	import { page } from '$app/stores';
	import type { Message, User } from '$lib/interfaces.ts';
	import { GetUsername, formatDate, getProfilePicture } from '$lib/users.ts';

	export let messages: Message[] = [];
	const dates: Map<number, string> = new Map();
	export let currentRoomMembers: User[] = [];

	function computeDateDivider(msg: Message, i: number): Message {
		const formattedDate = formatDate(String(msg.sentat));
		if (![...dates.values()].includes(String(formattedDate.ofYear))) {
			dates.set(i, String(formattedDate.ofYear));
		}

		msg.shortened = msg.text.length > 1400;
		return msg;
	}

	$: messages = messages.map(computeDateDivider);
</script>

<div class="msg-container">
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
				<div class:shortened={message.shortened}>
					{message.shortened ? message.text.split('').slice(0, 1400).join('') : message.text}
					{#if message.shortened}
						<button
							type="button"
							class="show-more"
							on:click={() => (message = { ...message, shortened: false })}>Show more</button
						>
					{/if}
				</div>
				<span>{formatDate(String(message.sentat)).ofDay}</span>
			</div>
		</div>
	{/each}
</div>

<style>
	@import '../css/messages/container.css';
</style>
