<script lang="ts">
	import type { User } from '$lib/interfaces.ts';
	import { Register } from '$lib/authentication.js';

	let info: User = {
		id: 0,
		username: '',
		about: '',
		email: '',
		password: '',
		currency: '',
		balance: 0
	};
</script>

<form on:submit|preventDefault={() => Register(info)}>
	<div>
		<input type="text" placeholder="Enter your username" bind:value={info.username} />
	</div>
	<div>
		<input type="text" placeholder="Enter your email" bind:value={info.email} />
	</div>
	<div>
		<input type="text" placeholder="Enter your password" bind:value={info.password} />
	</div>
	<button type="submit">Submit</button>
</form>

<style>
	form {
		display: flex;
		flex-direction: column;
		align-items: center;
		gap: 24px;
		height: fit-content;
		height: 426px;
		width: 350px;
		border: #000 solid 1px;
		padding: 24px 15px;
		border-radius: 25px;
	}

	form div input {
		border-radius: 25px;
		border: black solid 1px;
		width: 320px;
		height: 43px;
		font-size: 20px;
		padding-left: 10px;
	}

	form button {
		background: #fff;
		border: #005fcf solid 2px;
		cursor: pointer;
		border-radius: 25px;
		width: 240px;
		height: 43px;
		color: #005fcf;
		font-size: 25px;
		transition:
			background 0.5s,
			color 0.5s;
		/* box-shadow: 1px 1px 1px rgba(0, 0, 0, 0.1); */
		margin-top: auto;
	}

	form button:focus,
	form button:hover {
		background: #005fcf;
		border-color: #fff;
		color: #fff;
	}
</style>
