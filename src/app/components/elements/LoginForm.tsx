"use client";

import { useState } from "react";
import { useRouter } from "next/navigation";
import { signIn } from "@/app/lib/auth";

export default function LoginForm() {
  const router = useRouter();
  const [error, setError] = useState<string | null>(null);
  const [loading, setLoading] = useState(false);

  async function handleSubmit(event: React.FormEvent<HTMLFormElement>) {
    event.preventDefault();
    setError(null);
    setLoading(true);

    const formData = new FormData(event.currentTarget);
    const email = formData.get("email")?.toString().trim();
    const password = formData.get("password")?.toString().trim();

    if (!email || !password) {
      setError("メールアドレスとパスワードを入力してください。");
      setLoading(false);
      return;
    }

    try {
      const result = await signIn("credentials", {
        email,
        password,
        redirect: false, // 手動でリダイレクトを管理
      });

      if (result?.error) {
        setError("メールアドレスまたはパスワードが間違っています。");
      } else {
        router.push("/dashboard"); // ログイン成功時にリダイレクト
      }
    } catch (error) {
      console.error("ログインエラー:", error);
      setError("ログインに失敗しました。");
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
        {error && (
          <div className="text-red-500 text-sm text-center">{error}</div>
        )}
        <button
          className="w-full bg-blue-500 text-white p-2 rounded-md hover:bg-blue-600 transition duration-300 disabled:opacity-50"
          disabled={loading}
          aria-disabled={loading}
        >
          {loading ? "ログイン中..." : "ログインする"}
        </button>
      </form>
    </div>
  );
}
