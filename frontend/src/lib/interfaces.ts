export type User = {
	id?: string;
	profilepicture?: string;
	username?: string;
	about?: string;
	email?: string;
	password?: string;
	currency?: string;
	balance?: number;
};

export type Message = {
	id?: number;
	text: string;
	sentat?: string;
	from?: string;
	to?: string;
	shortened?: boolean;
};

export type ChatRoom = {
	id?: string;
	title?: string;
	profilepicture?: string;
	createdat?: string;
	description?: string;
	members?: string[];
	admins?: string[];
	owner?: string;
	ownerData?: User;
	messages?: number[];
	type?: string;
};

export type HTTPResponse = {
	status?: string;
	response?: any;
	id?: string;
};

export type VerificationSession = {
	id?: string;
	code?: string;
	userid?: string;
};

export type MessageDate = {
	ofYear?: string;
	ofDay?: string;
};

export const FetchConfig: any = {
	method: 'GET',
	credentials: 'include',
	headers: {
		'Content-Type': 'application/json'
	}
};

export type Setting = {
	LightMode: boolean;
};
