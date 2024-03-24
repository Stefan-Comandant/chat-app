import type { ChatRoom } from "$lib/interfaces.ts"
import { FetchConfig } from "$lib/interfaces.ts"

export async function AddChatRoom(info: ChatRoom) {
		const response = await fetch("/api/rooms", { ...FetchConfig, method: "POST", body: JSON.stringify(info)}).then(res => res.json())

		console.log(response)
		if (response.status === "success") return response.response
		return {}
}

export async function GetChatRooms() {
		const response = await fetch('/api/rooms', FetchConfig).then((res) => res.json());
		if (response.status === "success") return response.response
		return []
}

export async function GetRoom(id: string) {
		const response = await fetch(`/api/rooms/${id}`, FetchConfig).then((res) => res.json());
		if (response.status === 'success') return response.response;
		return {};
}

export async function GetUsers() {
	const response = await fetch(`/api/users`, FetchConfig).then((res) => res.json());
	if (response.status === 'success') return response.response;
	return []
}

export async function FetchMessages(id: number[] | undefined) {
		const body = JSON.stringify(id);
		const data = await fetch('/api/messages', { ...FetchConfig, method: "PUT", body: body}).then((res) => res.json());

		if (data.status === 'success') return data.response;
		return [];
}

export async function GetUserData() {
	const data = await fetch("/api/getUserData", FetchConfig).then(res => res.json())
	
	console.log(data)

	if (data.status === "success") return data.response
	return {}
}
