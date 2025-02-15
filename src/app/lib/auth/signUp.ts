"use server";

import bcrypt from "bcryptjs";
import prisma from "../db/prisma";

export async function signupAction(
  name: string,
  email: string,
  password: string
) {
  try {
    const existingUser = await prisma.user.findUnique({
      where: { email },
    });

    if (existingUser) {
      console.log("⚠️ ユーザーはすでに登録されています:", email);
      return { error: "このメールアドレスは既に登録されています。" };
    }

    const hashedPassword = await bcrypt.hash(password, 10);

    const user = await prisma.user.create({
      data: {
        name,
        email,
        password: hashedPassword,
        emailVerified: false,
      },
    });

    console.log("新規ユーザー作成成功:", user);

    return { success: true, user };
  } catch (error) {
    return { error: "ユーザー登録に失敗しました。" };
  }
}
