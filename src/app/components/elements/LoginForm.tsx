import { redirect } from "next/navigation";
import { signIn } from "@/app/lib/auth/auth";
import { z } from "zod";

const loginSchema = z.object({
  email: z.string().email("有効なメールアドレスを入力してください"),
  password: z.string().min(6, "パスワードは6文字以上で入力してください"),
});

// MEMO: ServerActionで実装
export default function LoginForm() {
  async function handleSubmit(formData: FormData): Promise<void> {
    "use server";

    const data = {
      email: formData.get("email")?.toString().trim() ?? "",
      password: formData.get("password")?.toString().trim() ?? "",
    };

    const validation = loginSchema.safeParse(data);
    if (!validation.success) {
      throw new Error(validation.error.errors[0].message);
    }

    try {
      const result = await signIn("credentials", {
        email: data.email,
        password: data.password,
        redirect: false,
      });

      if (result?.error) {
        throw new Error("メールアドレスまたはパスワードが間違っています。");
      }
    } catch (error) {
      throw new Error("ログインに失敗しました。");
    }

    redirect("/dashboard");
  }

  return (
    <div className="flex items-center justify-center min-h-screen bg-gray-100">
      <form
        action={handleSubmit}
        className="bg-white p-6 rounded-lg shadow-lg w-96 space-y-4"
      >
        <h2 className="text-2xl font-bold text-center text-gray-700">
          ログイン
        </h2>
        <div>
          <label className="block text-gray-700">メールアドレス:</label>
          <input
            className="w-full mt-1 p-2 border border-gray-300 rounded-md focus:ring-2 focus:ring-blue-500"
            type="email"
            name="email"
            required
          />
        </div>
        <div>
          <label className="block text-gray-700">パスワード:</label>
          <input
            className="w-full mt-1 p-2 border border-gray-300 rounded-md focus:ring-2 focus:ring-blue-500"
            type="password"
            name="password"
            required
          />
        </div>
        <button className="w-full bg-blue-500 text-white p-2 rounded-md hover:bg-blue-600 transition duration-300">
          ログインする
        </button>
      </form>
    </div>
  );
}
