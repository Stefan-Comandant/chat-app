import type { Setting } from '$lib/interfaces.ts';
import { writable } from 'svelte/store';

const value: Setting = {
	LightMode: false
};

export const settings = writable(value);
export const loading = writable({ goPast: false, loading: false });
