<script lang="ts">
	export let page: any = {};
</script>

{#if page.data.USER.username !== 'Guest'}
	<nav>
		<a href="/" class:active={page.route.id === '/'}>Home</a>
		<a href="/login" class:active={page.route.id === '/login'}>Login</a>
		<a href="/register" class:active={page.route.id === '/register'}>Register</a>
		<a href="/chats" class:active={page.route.id.includes('/chats')}>Chats</a>
		<a href="/profiles" class:active={page.route.id.includes('/profiles')}>Profiles</a>
		<a href="/settings" class:active={page.route.id.includes('/settings')}>Settings</a>
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
