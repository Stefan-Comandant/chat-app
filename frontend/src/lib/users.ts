import { FetchConfig } from "$lib/interfaces.ts"

export async function GetProfileByID(id: number) {
  const response = await fetch(`/api/users/${id}`, FetchConfig);
	const data = await response.json()
  
  return data.response;
}

export async function GetProfiles() {
		const response = await fetch('/api/users', FetchConfig);
    const data = await response.json()
    
    return data.response
}

export async function GetUsers() {
	const response = await fetch(`/api/users`, FetchConfig);
  const data = await response.json()
  
	return data.response
}

export async function GetUserData() {
	const response = await fetch("/api/getUserData", FetchConfig)
  const data = await response.json()

	return data.response
}
