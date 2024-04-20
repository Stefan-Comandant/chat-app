export interface User {
	id?: number;
  profilepicture?: string,
	username?: string;
	about?: string;
	email?: string;
	password?: string;
	currency?: string;
	balance?: number;
}

export interface Message {
	id?: number;
	text: string
	sentat?: string;
	fromid?: number;
	toid?: number;
}

export interface ChatRoom {
	id?: number;
	title?: string;
	createdat?: string;
	description?: string;
	members?: number[]
	admins?: number[]
	owner?: number;
	ownerData?: User;
	messages?: number[]
}

export interface HTTPResponse {
  status?: string;
  response?: any;
}

export interface VerificationSession {
  id?: number;
  code?: string;
  userid?: number; 
}

export interface MessageDate {
	ofYear?: string;
	ofDay?: string;
}

export const FetchConfig : any = {
	method: "GET",
	credentials: "include",
	headers: {
		"Content-Type": "application/json",
	},
}

export interface Setting{
	LightMode: boolean;
}