import type { User } from "$lib/interfaces.ts"
import { FetchConfig } from "$lib/interfaces.ts"

export async function Login(info: User) {
		const response = await fetch('http://localhost:7000/login', { ...FetchConfig, method: "POST", body: JSON.stringify(info)}).then(res => res.json());
		console.log(response)
}

export async function VerifyWithCode(code: string){	
		await fetch(`http://localhost:7000/code/${code}`);
}

export async function Logout() {
		const response = await fetch("http://localhost:7000/logout", FetchConfig).then(res => res.json())
		console.log(response)
}

export async function Register(info: User) {
		const response = await fetch('http://localhost:7000/register', { ...FetchConfig, method: "POST", body: JSON.stringify(info)}).then(res => res.json());
		console.log(response)
}
