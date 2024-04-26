<script lang="ts">
	import { store } from '../../stores.ts';

	export let page: any = {};
	$: darkMode = !$store.LightMode;

	const options = [
		{
			route: '/',
			text: 'Home'
		},
		{
			route: '/login',
			text: 'Login'
		},
		{
			route: '/register',
			text: 'Register'
		},
		{
			route: '/chats',
			text: 'Chats'
		},
		{
			route: '/profiles',
			text: 'Profiles'
		},
		{
			route: '/settings',
			text: 'Settings'
		}
	];
</script>

{#if Object.keys(page.data.USER).length}
	<nav class:dark={!!darkMode}>
		{#each options as option}
			<a class:dark={!!darkMode} href={option.route} class:active={page.route.id === option.route}
				>{option.text}</a
			>
		{/each}
		<img src={page.data.USER.profilepicture} alt="Pfp" />
	</nav>
{/if}

<style>
	img {
		width: 42px;
		height: 42px;
		border-radius: 50%;
	}

	nav {
		min-height: fit-content;
		height: 60px;
		width: 100%;
		display: flex;
		justify-content: space-around;
		align-items: center;
		border-bottom: 1px solid #b0b0b0;
		overflow-x: scroll;
		z-index: 1000;
		background: #fff;
	}

	nav::-webkit-scrollbar {
		display: none;
	}

	nav.dark {
		border-bottom: 2px solid #d6d6d6;
		background: #02042a;
	}

	nav a {
		text-decoration: none;
		color: #000;
		border-radius: 20px;
		padding: 10px 20px;
		height: fit-content;
	}

	nav a.dark {
		color: #fff;
	}

	.active {
		background: #d6d6d6;
	}

	.active.dark {
		background: rgba(255, 255, 255, 0.2);
	}
</style>
