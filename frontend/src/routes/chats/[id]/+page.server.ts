import { FetchConfig, type ChatRoom, type Message, type User } from '$lib/interfaces.ts';

export const load = async ({ fetch, params }: any) => {
	let messages: { status?: string; response?: Message[] } = {
		status: '',
		response: []
	};

	let members: { status?: string; response?: User[] } = {
		status: '',
		response: []
	};

	let room: { status?: string; response?: ChatRoom } = {
		status: '',
		response: {}
	};

	messages = await fetch(`http://localhost:9000/api/rooms/${params.id}/messages`, FetchConfig).then(
		(res: Response) => res.json()
	);
	if (!messages || messages.status !== 'success') messages.response = [];

	members = await fetch(`http://localhost:9000/api/rooms/${params.id}/members`, FetchConfig).then(
		(res: Response) => res.json()
	);
	if (!members || members.status !== 'success') members.response = [];

	room = await fetch(`http://localhost:9000/api/rooms/${params.id}`, FetchConfig).then(
		(res: Response) => res.json()
	);
	if (!room || room.status !== 'success') room.response = {};

	return {
		messages: messages.response,
		members: members.response,
		room: room.response
	};
};
