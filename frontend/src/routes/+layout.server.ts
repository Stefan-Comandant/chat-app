import { FetchConfig } from '$lib/interfaces.ts';
import { redirect } from '@sveltejs/kit';

export const load = async ({ cookies, route, fetch }: any) => {
	const cookie = cookies.get('session_cookie');
	let userData = {
		status: '',
		response: {
			username: 'Guest',
			email: 'example@gmail.com',
			password: 'guest',
			profilepicture: ''
		}
	};
	if (!cookie) {
		if (!['/login', '/register'].includes(route.id)) redirect(303, '/login');
		return { USER: userData.response };
	}

	userData = await fetch('http://localhost:9000/api/getUserData', FetchConfig).then(
		(res: Response) => res.json()
	);
	if (!userData || userData.status !== 'success')
		userData.response = {
			username: 'Guest',
			email: 'example@gmail.com',
			password: 'guest',
			profilepicture: ''
		};
	return {
		USER: userData.response
	};
};
