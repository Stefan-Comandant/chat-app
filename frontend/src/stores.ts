import type { Setting } from '$lib/interfaces.ts';
import { writable } from 'svelte/store';

const value: Setting = {
	LightMode: false
};

export const store = writable(value);
