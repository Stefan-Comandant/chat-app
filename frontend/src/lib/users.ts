import { FetchConfig, type User } from "$lib/interfaces.ts"

export async function GetProfileByID(id: number = 0): Promise<User> {
  const response = await fetch(`/api/users/${id}`, FetchConfig);
	const data = await response.json()
  
  if (data.status === "success") return data.response;
  return {};
}

export async function GetProfiles(): Promise<User[]> {
	const response = await fetch('/api/users', FetchConfig);
  const data = await response.json();

  if (data.status === "success") return data.response
	return []
}

export async function GetUsers(): Promise<User[]> {
	const response = await fetch(`/api/users`, FetchConfig);
  const data = await response.json()
  
  if (data.status === "success") return data.response
	return []
}

export async function GetUserData(): Promise<User> {
	const response = await fetch("/api/getUserData", FetchConfig)
  const data = await response.json()

  if (data.status === "success") return data.response
	return {}
}