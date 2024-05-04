export interface User {
	id?: string;
	profilepicture?: string;
	username?: string;
	about?: string;
	email?: string;
	password?: string;
	currency?: string;
	balance?: number;
}

export interface Message {
	id?: number;
	text: string;
	sentat?: string;
	from?: string;
	to?: string;
	shortened?: boolean;
}

export interface ChatRoom {
	id?: string;
	title?: string;
	createdat?: string;
	description?: string;
	members?: string[];
	admins?: string[];
	owner?: string;
	ownerData?: User;
	messages?: number[];
}

export interface HTTPResponse {
	status?: string;
	response?: any;
	id?: string;
}

export interface VerificationSession {
	id?: string;
	code?: string;
	userid?: string;
}

export interface MessageDate {
	ofYear?: string;
	ofDay?: string;
}

export const FetchConfig: any = {
	method: 'GET',
	credentials: 'include',
	headers: {
		'Content-Type': 'application/json'
	}
};

export interface Setting {
	LightMode: boolean;
}
