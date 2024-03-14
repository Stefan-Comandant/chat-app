export interface User {
	id: number;
	username: string;
	about: string;
	email: string;
	password: string;
	currency: string;
	balance: number;
}

export interface Message {
	id: number;
	text: string
	sentat: string;
	fromid: number;
	toid: number;
}

export interface ChatRoom {
	id: number;
	title: string;
	createdat: string;
	description: string;
	members: number[]
	admins: number[]
	owner: number
	messages: number[]
}

export let FetchConfig : any = {
	method: "GET",
	credentials: "include",
	headers: {
		"Content-Type": "application/json",
	},
}
