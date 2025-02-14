"use client";

import { useState } from "react";
import { useRouter } from "next/navigation";
import { signUpAction } from "@/app/lib/auth/signUp";

export default function SignUpForm() {
  const router = useRouter();
  const [error, setError] = useState<string | null>(null);
  const [loading, setLoading] = useState(false);

  async function handleSubmit(event: React.FormEvent<HTMLFormElement>) {
    event.preventDefault();
    setError(null);
    setLoading(true);

    const formData = new FormData(event.currentTarget);
    const name = formData.get("name")?.toString().trim();
    const email = formData.get("email")?.toString().trim();
    const password = formData.get("password")?.toString().trim();

    if (!name || !email || !password) {
      setError("名前、メールアドレス、パスワードを入力してください。");
      setLoading(false);
      return;
    }

    try {
      // Server Action を実行
      const result = await signUpAction(name, email, password);

      console.log("SignUp Result:", result);

      if (result?.error) {
        setError(result.error);
      } else {
        router.push("/dashboard"); // 登録成功時にリダイレクト
      }
    } catch (error) {
      console.error("サインアップエラー:", error);
      setError("サインアップに失敗しました。");
    } finally {
      setLoading(false);
    }
  }

  return (
    <div className="flex items-center justify-center min-h-screen bg-gray-100">
      <form
        onSubmit={handleSubmit}
        className="bg-white p-6 rounded-lg shadow-lg w-96 space-y-4"
      >
        <h2 className="text-2xl font-bold text-center text-gray-700">
          新規登録
        </h2>
        <div>
          <label className="block text-gray-700">名前:</label>
          <input
            className="w-full mt-1 p-2 border border-gray-300 rounded-md focus:ring-2 focus:ring-blue-500"
            type="text"
            name="name"
            required
          />
        </div>
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
        {error && (
          <div className="text-red-500 text-sm text-center">{error}</div>
        )}
        <button
          className="w-full bg-green-500 text-white p-2 rounded-md hover:bg-green-600 transition duration-300 disabled:opacity-50"
          disabled={loading}
          aria-disabled={loading}
        >
          {loading ? "登録中..." : "登録する"}
        </button>
      </form>
    </div>
  );
}
