"use client";

import { authenticate } from "../../lib/actions";
import { useFormState, useFormStatus } from "react-dom";

export default function LoginForm() {
  const [state, formAction] = useFormState(authenticate, true);

  return (
    <div className="flex items-center justify-center min-h-screen bg-gray-100">
      <form
        action={formAction}
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
        {!state && (
          <div className="text-red-500 text-sm text-center">
            メールアドレスかパスワードが違います。
          </div>
        )}
        <SubmitButton />
      </form>
    </div>
  );
}

function SubmitButton() {
  const { pending } = useFormStatus();
  return (
    <button
      className="w-full bg-blue-500 text-white p-2 rounded-md hover:bg-blue-600 transition duration-300 disabled:opacity-50"
      aria-disabled={pending}
    >
      {pending ? "ログイン中..." : "ログインする"}
    </button>
  );
}
