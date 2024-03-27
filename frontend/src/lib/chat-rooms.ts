import type { ChatRoom } from "$lib/interfaces.ts"
import { FetchConfig } from "$lib/interfaces.ts"

export async function AddChatRoom(info: ChatRoom) {
		let response = await fetch("/api/rooms", { ...FetchConfig, method: "POST", body: JSON.stringify(info)});

    if (response.ok) response = await response.json();
    else response = JSON.parse(await response.text());

		console.log(response)
    return response
}

export async function GetChatRooms() {
		let response = await fetch('/api/rooms', FetchConfig);
	  
	  if (response.ok) response = await response.json();
    else response = JSON.parse(await response.text());

    if (response.status === "success") return response.response
		return []
}

export async function GetRoom(id: string) {
		let response = await fetch(`/api/rooms/${id}`, FetchConfig);
    
    if (response.ok) response = await response.json();
    else response = JSON.parse(await response.text());

		if (response.status === 'success') return response.response;
		return {};
}

export async function GetUsers() {
	let response = await fetch(`/api/users`, FetchConfig);

  if (response.ok) response = await response.json();
  else response = JSON.parse(await response.text());

	if (response.status === 'success') return response.response;
	return []
}

export async function FetchMessages(id: number[] | undefined) {
		const body = JSON.stringify(id);
		let response= await fetch('/api/messages', { ...FetchConfig, method: "PUT", body: body});

    if (response.ok) response = await response.json();
    else response = JSON.parse(await response.text());

		if (response.status === 'success') return response.response;
		return [];
}

export async function GetUserData() {
	let response = await fetch("/api/getUserData", FetchConfig)

  if (response.ok) response = await response.json();
  else response = JSON.parse(await response.text());

	console.log(response)

	if (response.status === "success") return response.response
	return {}
}
