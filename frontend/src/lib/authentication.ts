import type { User } from "$lib/interfaces.ts"
import { FetchConfig } from "$lib/interfaces.ts"

export async function Login(info: User) {
		let response = await fetch('/api/login', { ...FetchConfig, method: "POST", body: JSON.stringify(info)});
    /*
     TODO: fix the wating for code verification timeout
     TODO: add a better way the verify emails
     TODO:
    */
    if (response.ok) response = await response.json();
    else response = JSON.parse(await response.text());
		console.log(response)
    return response
}

export async function VerifyWithCode(code: string){	
		await fetch(`/api/code/${code}`);
}

export async function Logout() {
		let response = await fetch("/api/logout", FetchConfig)
    if (response.ok) response = await response.json();
    else response = JSON.parse(await response.text());
		console.log(response)
}

export async function Register(info: User) {
		let response = await fetch('/api/register', { ...FetchConfig, method: "POST", body: JSON.stringify(info)});
		if (response.ok) response = await response.json();
    else response = JSON.parse(await response.text());
    console.log(response)
}
