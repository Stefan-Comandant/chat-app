import type { User, VerificationSession } from "$lib/interfaces.ts"
import { FetchConfig } from "$lib/interfaces.ts"

export async function Login(info: User) {
		const response = await fetch('/api/login', { ...FetchConfig, method: "POST", body: JSON.stringify(info)});
    const data = await response.json()
		console.log(data)
    return data
}

export async function VerifyWithCode(info: VerificationSession){	
		const response = await fetch("/api/code/", { ...FetchConfig, method: "POST", body: JSON.stringify(info)} );
    const data = await response.json()
    console.log(data)

    return data
}

export async function Logout() {
		const response = await fetch("/api/logout", FetchConfig)
    const data = await response.json()
		console.log(data)
    
    return data
}

export async function Register(info: User) {
		const response = await fetch('/api/register', { ...FetchConfig, method: "POST", body: JSON.stringify(info)});
    const data = await response.json();
    console.log(data)

    return data
}
