import { FetchConfig, type ChatRoom, type User } from '$lib/interfaces.ts';

export const load = async ({ fetch, parent }: any) => {
	await parent();
	let rooms: { status?: string; response?: ChatRoom[] } = {
		status: '',
		response: []
	};

	let chats: { status?: string; response?: User[] } = {
		status: '',
		response: []
	};

	rooms = await fetch('http://localhost:9000/api/rooms/broadcast', FetchConfig).then(
		(res: Response) => res.json()
	);
	if (!rooms || rooms.status !== 'success') rooms.response = [];

	chats = await fetch('http://localhost:9000/api/rooms/direct', FetchConfig).then((res: Response) =>
		res.json()
	);
	if (!chats || chats.status !== 'success') chats.response = [];

	return {
		rooms: rooms.response,
		chats: chats.response
	};
};
