"use server";

import { redirect } from "next/navigation";
import { z } from "zod";

const loginSchema = z.object({
  email: z.string().email("有効なメールアドレスを入力してください"),
  password: z.string().min(6, "パスワードは6文字以上で入力してください"),
});

export type LoginState = {
  message: string | null;
};

export async function loginAction(
  _: LoginState,
  formData: FormData
): Promise<LoginState> {
  const data = {
    email: formData.get("email")?.toString().trim() ?? "",
    password: formData.get("password")?.toString().trim() ?? "",
  };

  const validation = loginSchema.safeParse(data);
  if (!validation.success) {
    return {
      message: validation.error.errors[0].message,
    };
  }

  try {
    const response = await fetch(`${process.env.NEXT_PUBLIC_API_URL}/auth/login`, {
      method: "POST",
      headers: {
        "Content-Type": "application/json",
      },
      body: JSON.stringify(data),
    });

    if (!response.ok) {
      return {
        message: "メールアドレスまたはパスワードが間違っています。",
      };
    }
  } catch {
    return {
      message: "ログインに失敗しました。",
    };
  }

  redirect("/");
}
