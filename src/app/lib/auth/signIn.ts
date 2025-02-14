"use server";

import { signIn } from "./auth";


export async function loginAction(email: string, password: string) {
  try {
    const result = await signIn("credentials", {
      email,
      password,
      redirect: false,
    });

    console.log("Server Action Result:", result);

    if (result?.error) {
      return { error: "メールアドレスまたはパスワードが間違っています。" };
    }

    return { success: true };
  } catch (error) {
    console.error("ログインエラー:", error);
    return { error: "ログインに失敗しました。" };
  }
}
