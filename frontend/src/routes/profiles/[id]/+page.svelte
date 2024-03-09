<script lang="ts">
	import { page } from "$app/stores"
	import { onMount } from "svelte"
	import type { User } from "$lib/interfaces.ts"

	let profile : User = {}
	const id: string = $page.params.id
	console.log(id)

	async function GetProfileByID() {
		const data = await fetch(`http://localhost:7000/users/${id}`, {
			method: "GET",
			credentials: "include",
			headers: {
				"Content-Type": "application/json"
			}
		}).then(res => res.json())
		if (data.status === "success") return data.response
		return {}
	}

	onMount(async () => {
		profile = await GetProfileByID()
		if (!profile) profile = {}
	})
</script>

<div>
	<h2>{profile.username}</h2>
	<p>{profile.about}</p>
</div>
