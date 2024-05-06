import { FetchConfig, type ChatRoom, type Message, type User } from '$lib/interfaces.ts';

export const load = async ({ fetch, params, parent }: any) => {
	const data = await parent();
	let messages: { status?: string; response?: Message[] } = {
		status: '',
		response: []
	};

	let peer: { status?: string; response?: User } = {
		status: '',
		response: {}
	};

	messages = await fetch(
		`http://localhost:9000/api/room/${params.id}/direct/messages`,
		FetchConfig
	).then((res: Response) => res.json());

	if (!messages || messages.status !== 'success') messages.response = [];

	peer = await fetch(`http://localhost:9000/api/users/${params.id}`, FetchConfig).then(
		(res: Response) => res.json()
	);

	if (!peer || peer.status !== 'success') peer.response = {};

	return {
		messages: messages.response,
		USER: data.USER,
		peer: peer.response
	};
};
