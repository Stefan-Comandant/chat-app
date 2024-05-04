import { FetchConfig, type User } from '$lib/interfaces.ts';

export const load = async ({ fetch, params, parent }: any) => {
	await parent();
	let profile: { status?: string; response?: User } = {
		status: '',
		response: {}
	};

	profile = await fetch(`http://localhost:9000/api/users/${params.id}`, FetchConfig).then(
		(res: Response) => res.json()
	);
	if (!profile || profile.status !== 'success')
		profile.response = {
			id: '',
			username: '',
			about: '',
			email: '',
			password: '',
			currency: '',
			balance: 0
		};

	return {
		profile: profile.response
	};
};
