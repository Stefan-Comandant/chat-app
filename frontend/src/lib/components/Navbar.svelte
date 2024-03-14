<script lang="ts">
	export let page: any = {};

	console.log(page);
</script>

<nav>
	<a href="/" class:active={page.route.id === '/'}>Home</a>
	<a href="/login" class:active={page.route.id === '/login'}>Login</a>
	<a href="/register" class:active={page.route.id === '/register'}>Register</a>
	<a href="/chats" class:active={page.route.id.includes('/chats')}>Chats</a>
	<a href="/profiles" class:active={page.route.id.includes('/profiles')}>Profiles</a>
</nav>

<style>
	nav {
		display: flex;
		justify-content: space-around;
		position: fixed;
		left: 0;
		right: 0;
	}

	nav a {
		text-decoration: none;
		color: #000;
		border-radius: 20px;
		padding: 10px 20px;
		font-family: Arial;
	}

	.active {
		background: #d6d6d6;
	}
</style>
