<script lang="ts">
	import { settings, loading } from '../../stores.ts';

	export let page: any = {};
	$: darkMode = !$settings.LightMode;

	const options = [
		{
			route: '/chats',
			text: 'Chats'
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
			route: '/profiles',
			text: 'Profiles'
		},
		{
			route: '/settings',
			text: 'Settings'
		}
	];

	let intervalMs = 200;
	let multiplier = 0;
	loading.subscribe((value) => {
		if (value.loading) {
			const interval = setInterval(() => {
				if (multiplier !== 6 || value.goPast) multiplier++;
				if ($loading.goPast) intervalMs = 100;

				if (multiplier == 12) {
					clearInterval(interval);
					multiplier = 0;
					$loading.goPast = false;
					$loading.loading = false;
				}
			}, intervalMs);
		}
	});
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
{#if $loading.loading}
	<div class="bar">
		<div class="line" style="width: {10 * multiplier}%" />
	</div>
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
		background: #4a90e2;
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
		background: #00bfa5;
	}

	.active.dark {
		background: rgba(255, 255, 255, 0.2);
	}

	.bar {
		width: 100%;
		height: 4px;
	}

	.line {
		width: 0;
		max-width: 100%;
		height: 4px;
		background: red;
		transition: all 800ms;
	}
</style>
