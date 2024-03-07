import type { User } from "$lib/interfaces.ts"

export async function Login(info: User) {
		const response = await fetch('http://localhost:7000/login', {
			method: 'POST',
			body: JSON.stringify(info),
			credentials: 'include',
			headers: {
				'Content-Type': 'application/json'
			}
		}).then(res => res.json());
		console.log(response)
}

export async function VerifyWithCode(code: string){	
		await fetch(`http://localhost:7000/code/${code}`);
}

export async function Logout() {
		const response = await fetch("http://localhost:7000/logout", {
			method: "GET",
			credentials: "include",
			headers: {
				"Content-Type": "application/json"
			}
		}).then(res => res.json())
		console.log(response)
}

export async function Register(info: User) {
		const response = await fetch('http://localhost:7000/register', {
			method: 'POST',
			body: JSON.stringify(info),
			credentials: 'include',
			headers: {
				'Content-Type': 'application/json'
			}
		}).then(res => res.json());
		console.log(response)
}
