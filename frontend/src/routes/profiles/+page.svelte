<script lang="ts">
	import { onMount } from "svelte"
	import type { User } from "$lib/interfaces.ts"
	import { FetchConfig } from "$lib/interfaces.ts"

	let profiles : User[] = []

	async function GetProfiles() {
		const data = await fetch("http://localhost:7000/users", FetchConfig).then(res => res.json())
		if (data.status === "success") return data.response
		return []
	}

	onMount(async () => {
		profiles = await GetProfiles()
		if (!profiles) profiles = []
	})
</script>

<h1>List of user:</h1>
<ul>
	{#each profiles as profile (profile.id)}
		<li>
			<a href="/profiles/{profile.id}">
				<h3>{profile.username}</h3>
			</a>
		</li>
	{/each}
</ul>
