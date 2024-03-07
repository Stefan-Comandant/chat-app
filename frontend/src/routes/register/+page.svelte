<script>
	let info1 = {
		Username: '',
		About: '',
		Email: '',
		Password: '',
		Balance: 0,
		Currency: 'RON'
	};
	

	async function Register() {
		const response = await fetch('http://localhost:7000/register', {
			method: 'POST',
			body: JSON.stringify(info1),
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

<form
	on:submit|preventDefault={Register}
>
	<input type="text" placeholder="Username" bind:value={info1.Username} />
	<input type="text" placeholder="Email" bind:value={info1.Email} />
	<input type="text" placeholder="Password" bind:value={info1.Password} />
	<button type="submit">Submit</button>
</form>

<input type="text" bind:value={code} />
<button	type="button" on:click={VerifyWithCode}>Verify</button>
