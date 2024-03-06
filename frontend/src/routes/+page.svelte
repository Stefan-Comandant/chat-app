<script>
	let info1 = {
		Username: '',
		About: '',
		Email: '',
		Password: '',
		Balance: 0,
		Currency: 'RON'
	};
	
	let info2 = {
		Email: "",
		Password: "",
	}
	let code = '';
</script>

<h1>Welcome to your library project</h1>
<p>Create your package using @sveltejs/package and preview/showcase your work with SvelteKit</p>
<p>Visit <a href="https://kit.svelte.dev">kit.svelte.dev</a> to read the documentation</p>

<form
	on:submit|preventDefault={async () => {
		const response = await fetch('http://localhost:7000/register', {
			method: 'POST',
			body: JSON.stringify(info1),
			credentials: 'include',
			headers: {
				'Content-Type': 'application/json'
			}
		}).then(res => res.json());
		console.log(response)
	}}
>
	<input type="text" placeholder="Username" bind:value={info1.Username} />
	<input type="text" placeholder="Email" bind:value={info1.Email} />
	<input type="text" placeholder="Password" bind:value={info1.Password} />
	<button type="submit">Submit</button>
</form>

<form on:submit|preventDefault={async () => {
	const response = await fetch('http://localhost:7000/login', {
		method: 'POST',
		body: JSON.stringify(info2),
		credentials: 'include',
		headers: {
			'Content-Type': 'application/json'
		}
	}).then(res => res.json());
	console.log(response)

}}>
	<input type="text" placeholder="Email" bind:value={info2.Email} />
	<input type="text" placeholder="Password" bind:value={info2.Password} />
	<button type="submit">Submit</button>
</form>

<input type="text" bind:value={code} />
<button
	type="button"
	on:click={async () => {
		await fetch(`http://localhost:7000/code/${code}`);
	}}
>Verify</button>

<button type="button" on:click={async () =>{
	const response = await fetch("http://localhost:7000/logout", {
		method: "GET",
		credentials: "include",
		headers: {
			"Content-Type": "application/json"
		}
	}).then(res => res.json())
	console.log(response)
}}>Logout</button>

<a href="/chats">Chats</a>
