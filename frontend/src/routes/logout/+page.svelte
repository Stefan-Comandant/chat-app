<script>
	async function Logout() {
		const response = await fetch("http://localhost:7000/logout", {
			method: "GET",
			credentials: "include",
			headers: {
				"Content-Type": "application/json"
			}
		}).then(res => res.json())
		console.log(response)
	}
</script>

<button type="button" on:click={Logout}>Logout</button>
