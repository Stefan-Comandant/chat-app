import { FetchConfig, type MessageDate, type User } from '$lib/interfaces.ts';

export async function GetProfileByID(id: string = ''): Promise<User> {
	const response = await fetch(`/api/users/${id}`, FetchConfig);
	const data = await response.json();

	if (data.status === 'success') return data.response;
	return {};
}

export async function GetProfiles(): Promise<User[]> {
	const response = await fetch('/api/users', FetchConfig);
	const data = await response.json();

	if (data.status === 'success') return data.response;
	return [];
}

export async function GetUsers(): Promise<User[]> {
	const response = await fetch(`/api/users`, FetchConfig);
	const data = await response.json();

	if (data.status === 'success') return data.response;
	return [];
}

export async function GetUserData(): Promise<User> {
	const response = await fetch('/api/getUserData', FetchConfig);
	const data = await response.json();

	if (data.status === 'success') return data.response;
	return {};
}

export async function GetRoomMembers(id: string): Promise<User[]> {
	const response = await fetch(`/api/rooms/${id}/members`, FetchConfig);
	const data = await response.json();

	return data.response;
}

// Function to get the moment of the day and of the year when the message was sent
export function formatDate(dateStr: string): MessageDate {
	if (!dateStr) return { ofYear: '', ofDay: '' };

	// Set pointers in time
	const date = new Date(dateStr);
	const todayDate = new Date();
	const { hour, minute, meridian, day, month, year } = getDateValues(date);
	const time = `${hour}:${minute} ${meridian}`;
	const today = `${getDateValues(todayDate).day}-${getDateValues(todayDate).month}-${getDateValues(todayDate).year}`;
	const yearDate = `${day}-${month}-${year}`;

	return {
		ofDay: time,
		ofYear: today === yearDate ? 'Today' : yearDate
	};
}

function getDateValues(date: Date) {
	return {
		minute: date.getMinutes() > 9 ? date.getMinutes() : '0' + date.getMinutes(),
		hour: date.getHours() > 12 ? date.getHours() - 12 : date.getHours(),
		day: date.getDate() > 9 ? date.getDate() : '0' + date.getDate(),
		month: date.getMonth() > 9 ? date.getMonth() : '0' + date.getMonth(),
		year: date.getFullYear(),
		meridian: date.getHours() > 12 ? 'PM' : 'AM'
	};
}

export function GetUsername(from: string = '', members: User[] = []): string {
	const member = members.filter((member) => member.id === from)[0];
	if (!member || !member.username) return '';
	return member.username;
}

export function getProfilePicture(id: string = '', members: User[] = []): string {
	const result = members.filter((member) => member.id === id)[0];
	if (!result || !result.profilepicture) return '';
	return result.profilepicture;
}
export function getPeer(users: User[] = [], members: string[] = [], USER: User = {}): User {
	return users.filter((user) => {
		return user.id == members.filter((member) => member != USER.id)[0];
	})[0];
}
