<script>
	let info = {
		Username: '',
		Email: '',
		Password: '',
		Balance: 0,
		Currency: 'RON'
	};

	let code = '';
</script>

<h1>Welcome to your library project</h1>
<p>Create your package using @sveltejs/package and preview/showcase your work with SvelteKit</p>
<p>Visit <a href="https://kit.svelte.dev">kit.svelte.dev</a> to read the documentation</p>

<form
	on:submit|preventDefault={async () => {
		await fetch('http://localhost:7000/register', {
			method: 'POST',
			body: JSON.stringify(info),
			credentials: 'include',
			headers: {
				'Content-Type': 'application/json'
			}
		});
	}}
>
	<input type="text" placeholder="Username" bind:value={info.Username} />
	<input type="text" placeholder="Email" bind:value={info.Email} />
	<input type="text" placeholder="Password" bind:value={info.Password} />
	<button type="submit">Submit</button>
</form>

<input type="text" bind:value={code} />
<button
	type="button"
	on:click={async () => {
		const response = await fetch(`http://localhost:7000/code/${code}`).then((res) => res.json());
		console.log(response);
	}}>Verify</button
>
