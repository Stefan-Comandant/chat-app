<script langs="ts">
	let code = ""

	let info2 = {
		Email: "",
		Password: "",
	}

	async function Login() {
		const response = await fetch('http://localhost:7000/login', {
			method: 'POST',
			body: JSON.stringify(info2),
			credentials: 'include',
			headers: {
				'Content-Type': 'application/json'
			}
		}).then(res => res.json());
		console.log(response)
	}

	async function VerifyWithCode(){	
		await fetch(`http://localhost:7000/code/${code}`);
	}

</script>

<form on:submit|preventDefault={Login}>
	<input type="text" placeholder="Email" bind:value={info2.Email} />
	<input type="text" placeholder="Password" bind:value={info2.Password} />
	<button type="submit">Submit</button>
</form>

<input type="text" bind:value={code} />
<button	type="button" on:click={VerifyWithCode}>Verify</button>

