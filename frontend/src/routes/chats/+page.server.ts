import { FetchConfig, type ChatRoom } from '$lib/interfaces.ts';

export const load = async ({ fetch, parent }: any) => {
	await parent();
	let rooms: { status?: string; response?: ChatRoom[] } = {
		status: '',
		response: []
	};

	rooms = await fetch('http://localhost:9000/api/rooms', FetchConfig).then((res: Response) =>
		res.json()
	);
	if (!rooms || rooms.status !== 'success') rooms.response = [];

	return {
		rooms: rooms.response
	};
};
