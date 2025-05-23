"use client";

import { useFormState } from "react-dom";
import LoadingButton from "../buttons/LoadingButton";
import { loginAction, type LoginState } from "@/app/actions/login";

const initialState: LoginState = {
  message: null,
};

export default function LoginForm() {
  const [state, formAction] = useFormState(loginAction, initialState);

  return (
    <div className="flex items-center justify-center min-h-screen bg-gray-100 dark:bg-gray-900">
      <form
        action={formAction}
        noValidate
        className="bg-white dark:bg-gray-800 p-6 rounded-lg shadow-lg w-96 space-y-4"
      >
        <h2 className="text-2xl font-bold text-center text-gray-700 dark:text-gray-200">
          ログイン
        </h2>

        {state.message && (
          <div className="text-red-500 text-sm text-center">
            {state.message}
          </div>
        )}

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

        <div className="bg-blue-200 rounded-md p-4 text-gray-700 dark:text-gray-300 flex flex-col gap-2">
          <p className="text-center text-md font-medium">
            ポートフォリオを閲覧の方は
            <span className="inline-block">
              下記を入力してログインしてください。
            </span>
          </p>
          <div className="flex justify-between items-center bg-white rounded-md p-2">
            <div className="flex flex-col">
              <p className="text-sm">メールアドレス:</p>
              <p className="font-mono">user1@example.com</p>
            </div>
            <button
              type="button"
              className="ml-2 px-2 py-1 text-sm border-gray-400 border rounded hover:bg-blue-500 hover:text-white transition"
              onClick={() => navigator.clipboard.writeText("user1@example.com")}
            >
              コピー
            </button>
          </div>

          <div className="flex justify-between items-center bg-white rounded-md p-2">
            <div className="flex flex-col">
              <p className="text-sm">パスワード:</p>
              <p className="font-mono">password</p>
            </div>
            <button
              type="button"
              className="ml-2 px-2 py-1 text-sm border-gray-400 border rounded hover:bg-blue-500 hover:text-white transition"
              onClick={() => navigator.clipboard.writeText("password")}
            >
              コピー
            </button>
          </div>
        </div>

        <LoadingButton label={"ログイン"} color={"blue"} />
      </form>
    </div>
  );
}
