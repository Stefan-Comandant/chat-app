import type { User } from "$lib/interfaces.ts"
import { FetchConfig } from "$lib/interfaces.ts"

export async function Login(info: User) {
		const response = await fetch('/api/login', { ...FetchConfig, method: "POST", body: JSON.stringify(info)}).then(res => res.json());
		console.log(response)
}

export async function VerifyWithCode(code: string){	
		await fetch(`/api/code/${code}`);
}

export async function Logout() {
		const response = await fetch("/api/logout", FetchConfig).then(res => res.json())
		console.log(response)
}

export async function Register(info: User) {
		const response = await fetch('/api/register', { ...FetchConfig, method: "POST", body: JSON.stringify(info)}).then(res => res.json());
		console.log(response)
}
