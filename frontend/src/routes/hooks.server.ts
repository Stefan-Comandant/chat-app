import type { Handle, MaybePromise } from '@sveltejs/kit';

export const handle: Handle = (async ({ event, resolve }) => {
	const response: MaybePromise<Response> = await resolve(event);

	return response;
}) satisfies Handle;
