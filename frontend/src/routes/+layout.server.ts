import { FetchConfig } from "$lib/interfaces.ts";
import { redirect } from "@sveltejs/kit";

export const load = async ({ cookies, route, fetch }: any) => {
  const cookie = cookies.get("session_cookie")

  if (!cookie && !["/login", "/register"].includes(route.id)){
    redirect(303, "/login")
  }

  let userData = await fetch("http://localhost:9000/api/getUserData", FetchConfig).then((res: Response) => res.json())
  if (!userData) userData = { username: "Guest", email: "example@gmail.com", password: "guest"}

  return{
    USER: userData.response,
  }
};