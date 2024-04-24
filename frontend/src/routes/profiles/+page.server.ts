import { FetchConfig, type User } from '$lib/interfaces.ts';

export const load = async ({ parent, fetch }: any) => {
	await parent();
	let profiles: { status?: string; response?: User[] } = {};

	profiles = await fetch(`http://localhost:9000/api/users`, FetchConfig).then((res: Response) =>
		res.json()
	);

	if (!profiles || profiles.status !== 'success') profiles.response = [];

	return {
		profiles: profiles.response
	};
};
