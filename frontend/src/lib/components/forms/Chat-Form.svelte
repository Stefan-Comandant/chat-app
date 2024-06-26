<script lang="ts">
	import type { User, ChatRoom, HTTPResponse } from '$lib/interfaces.ts';
	import { AddChatRoom } from '$lib/chat-rooms.ts';
	import { onMount } from 'svelte';
	import EditButton from '../buttons/Edit-Button.svelte';
	import { GetUsers } from '$lib/users.ts';
	import { loading, settings } from '../../../stores.ts';
	import LoadingCircle from '../LoadingCircle.svelte';
	let users: User[] = [];
	let openModal = false;

	onMount(async () => {
		users = await GetUsers();
		if (!users) users = [];
	});

	let info: ChatRoom = {
		members: [],
		admins: [],
		type: 'broadcast'
	};

	function AddMember(event: any, id: string = '') {
		if (!event || !id) return;

		const target = event.target;

		if (target.checked === true) {
			info.members = [...info.members, id];
		} else {
			info.members = info.members?.filter((member) => member != id);
		}
	}

	function AddAdmin(event: any, id: string = '') {
		if (!event || !id) return;
		const target = event.target;

		if (target.checked === true) {
			info.admins = [...info.admins, id];
		} else {
			info.admins = info.admins?.filter((admin) => admin != id);
		}
	}

	$: darkMode = !$settings.LightMode;

	let response: HTTPResponse = {
		response: ''
	};

	onMount(() => {
		$loading.goPast = true;
	});

	let image: HTMLImageElement;
	let fileInput: HTMLInputElement;
	let showImage = false;

	let isLoading = false;

	$: {
		if (isLoading) isLoading = Object.values(response).every((key) => key.length === 0);
	}
</script>

<form
	class:dark={!!darkMode}
	on:submit|preventDefault={async () => {
		if (!info.title?.length) return;
		response = {
			response: '',
			status: ''
		};

		isLoading = true;

		response = await AddChatRoom(info);
	}}
>
	<div>
		<input type="text" placeholder="Name" bind:value={info.title} />
	</div>
	<div class="pfp-input-container">
		<label for="file-input">Enter a profile picture</label>
		<input
			id="file-input"
			type="file"
			accept=".jpg, .jpeg, .png"
			bind:this={fileInput}
			on:change={() => {
				if (!fileInput || !fileInput.files) return;
				const file = fileInput.files[0];

				if (file) showImage = true;

				const reader = new FileReader();
				reader.readAsDataURL(file);
				reader.addEventListener('load', () => {
					const result = reader.result;
					image.setAttribute('src', String(result));
					info.profilepicture = String(result);
				});
			}}
		/>
		{#if showImage}
			<img bind:this={image} alt="pfp" />
		{/if}
	</div>
	<div>
		<input type="text" placeholder="Description" bind:value={info.description} />
	</div>
	<div class="members-display">
		<div class="members-title">
			<span>Members</span>
			<button class="edit-btn" type="button" on:click={() => (openModal = !openModal)}
				><EditButton stroke={darkMode ? '#ffffff' : '#000000'} /></button
			>
		</div>
		<div class="members-container">
			{#each [...info.members] as member (member)}
				<div class="member">
					<img
						class="profile-picture"
						src={users.filter((user) => user.id === member)[0].profilepicture}
						alt="Pfp"
					/>
					<span>{users.filter((user) => user.id === member)[0].username}</span>
				</div>
			{/each}
		</div>
	</div>
	<button type="submit" style="pointer-events: {isLoading ? 'none' : 'auto'};">
		{#if isLoading}
			<LoadingCircle />
		{:else}
			Submit
		{/if}</button
	>
	<span class:error={response.status === 'error'} class:success={response.status === 'success'}
		>{response.response}</span
	>
</form>

<dialog open={openModal} class:dark={!!darkMode}>
	<div>
		{#each users as user (user.id)}
			<div class="account">
				<div class="details" style="border: none;">
					<img class="profile-picture" alt="Pfp" src={user.profilepicture} />
					<div>
						<div>{user.username}</div>
						<span>{user.about}</span>
					</div>
				</div>

				<div class="checks">
					<div class="check">
						<div>Member</div>
						<input
							type="checkbox"
							on:input={(event) => {
								AddMember(event, user.id);
							}}
						/>
					</div>
					<div class="check">
						<div>Admin</div>
						<input
							type="checkbox"
							on:input={(event) => {
								AddAdmin(event, user.id);
							}}
						/>
					</div>
				</div>
			</div>
		{/each}
	</div>
</dialog>

<style>
	@import '../../css/authentication.css';
	@import '../../css/new-chat.css';
	@import '../../css/users/dialog.css';
	@import '../../css/img-preview.css';
</style>
