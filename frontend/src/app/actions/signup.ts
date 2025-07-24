"use server";

import bcrypt from "bcryptjs";
import { redirect } from "next/navigation";
import { z } from "zod";
import { signIn } from "@/app/lib/auth/auth";
import prisma from "@/app/lib/db/prisma";

const signupSchema = z.object({
  name: z.string().min(2, "名前は2文字以上で入力してください"),
  email: z.string().email("有効なメールアドレスを入力してください"),
  password: z.string().min(6, "パスワードは6文字以上で入力してください"),
});

export type SignupState = {
  message: string | null;
};

export async function signupAction(
  _: SignupState,
  formData: FormData
): Promise<SignupState> {
  const data = {
    name: formData.get("name")?.toString().trim() ?? "",
    email: formData.get("email")?.toString().trim() ?? "",
    password: formData.get("password")?.toString().trim() ?? "",
  };

  const validation = signupSchema.safeParse(data);
  if (!validation.success) {
    return { message: validation.error.errors[0].message };
  }

  try {
    const existingUser = await prisma.user.findUnique({
      where: { email: data.email },
    });

    if (existingUser) {
      return { message: "このメールアドレスは既に登録されています。" };
    }

    const hashedPassword = await bcrypt.hash(data.password, 10);

    await prisma.user.create({
      data: {
        name: data.name,
        email: data.email,
        password: hashedPassword,
      },
    });

    await signIn("credentials", {
      email: data.email,
      password: data.password,
      redirect: false,
    });
  } catch {
    return { message: "ユーザー登録に失敗しました。" };
  }

  redirect("/dashboard");
}
