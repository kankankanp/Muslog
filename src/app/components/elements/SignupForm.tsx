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

// MEMO: ServerActionで実装
// TODO: RCCにしてバリデーションメッセージを実装する
export default function Fff() {
  async function handleSubmit(formData: FormData): Promise<void> {
    "use server";

    const data = {
      name: formData.get("name")?.toString().trim() ?? "",
      email: formData.get("email")?.toString().trim() ?? "",
      password: formData.get("password")?.toString().trim() ?? "",
    };

    const validation = signupSchema.safeParse(data);
    if (!validation.success) {
      throw new Error(validation.error.errors[0].message);
    }

    try {
      const existingUser = await prisma.user.findUnique({
        where: { email: data.email },
      });

      if (existingUser) {
        throw new Error("このメールアドレスは既に登録されています。");
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
    } catch (error) {
      throw new Error("ユーザー登録に失敗しました。");
    }
    redirect("/dashboard");
  }

  return (
    <div className="flex items-center justify-center min-h-screen bg-gray-100 dark:bg-gray-900">
      <form
        action={handleSubmit}
        className="bg-white dark:bg-gray-800 p-6 rounded-lg shadow-lg w-96 space-y-4"
      >
        <h2 className="text-2xl font-bold text-center text-gray-700 dark:text-gray-200">
          新規登録
        </h2>

        <div>
          <label className="block text-gray-700 dark:text-gray-300">
            名前:
          </label>
          <input
            className="w-full mt-1 p-2 border border-gray-300 dark:border-gray-600 rounded-md focus:ring-2 focus:ring-blue-500 dark:focus:ring-blue-400 bg-white dark:bg-gray-700 text-gray-900 dark:text-gray-200"
            type="text"
            name="name"
            required
          />
        </div>
        <div>
          <label className="block text-gray-700 dark:text-gray-300">
            メールアドレス:
          </label>
          <input
            className="w-full mt-1 p-2 border border-gray-300 dark:border-gray-600 rounded-md focus:ring-2 focus:ring-blue-500 dark:focus:ring-blue-400 bg-white dark:bg-gray-700 text-gray-900 dark:text-gray-200"
            type="email"
            name="email"
            required
          />
        </div>
        <div>
          <label className="block text-gray-700 dark:text-gray-300">
            パスワード:
          </label>
          <input
            className="w-full mt-1 p-2 border border-gray-300 dark:border-gray-600 rounded-md focus:ring-2 focus:ring-blue-500 dark:focus:ring-blue-400 bg-white dark:bg-gray-700 text-gray-900 dark:text-gray-200"
            type="password"
            name="password"
            required
          />
        </div>
        <button className="w-full bg-green-500 text-white p-2 rounded-md hover:bg-green-600 dark:hover:bg-green-400 transition duration-300">
          登録する
        </button>
      </form>
    </div>
  );
}
