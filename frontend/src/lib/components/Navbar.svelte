<script lang="ts">
	export let page: any = {};

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
	<nav>
		{#each options as option}
			<a href={option.route} class:active={page.route.id === option.route}>{option.text}</a>
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
		display: flex;
		justify-content: space-around;
		align-items: center;
		position: fixed;
		left: 0;
		right: 0;
		height: fit-content;
		padding: 10px 0;
		top: 0;
		border-bottom: 1px solid #b0b0b0;
	}

	nav a {
		text-decoration: none;
		color: #000;
		border-radius: 20px;
		padding: 10px 20px;
		height: fit-content;
	}

	.active {
		background: #d6d6d6;
	}
</style>
