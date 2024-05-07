import { FetchConfig, type ChatRoom, type User } from '$lib/interfaces.ts';

export const load = async ({ fetch, parent }: any) => {
	await parent();
	let groups: { status?: string; response?: ChatRoom[] } = {
		status: '',
		response: []
	};

	let chats: { status?: string; response?: User[] } = {
		status: '',
		response: []
	};

	let users: { status?: string; response?: User[] } = {
		status: '',
		response: []
	};

	groups = await fetch('http://localhost:9000/api/rooms/broadcast', FetchConfig).then(
		(res: Response) => res.json()
	);
	if (!groups || groups.status !== 'success') groups.response = [];

	chats = await fetch('http://localhost:9000/api/rooms/direct', FetchConfig).then((res: Response) =>
		res.json()
	);
	if (!chats || chats.status !== 'success') chats.response = [];

	users = await fetch('http://localhost:9000/api/users', FetchConfig).then((res: Response) =>
		res.json()
	);
	if (!users || users.status !== 'success') users.response = [];

	return {
		groups: groups.response,
		chats: chats.response,
		users: users.response
	};
};
