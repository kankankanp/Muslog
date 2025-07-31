"use client";

import { zodResolver } from "@hookform/resolvers/zod";
import { useRouter } from "next/navigation";
import { useForm } from "react-hook-form";
import toast from "react-hot-toast";
import { useDispatch } from "react-redux";
import { z } from "zod";
import LoadingButton from "../buttons/LoadingButton";
import { usePostAuthLogin } from "@/app/libs/api/generated/orval/auth/auth";
import { login } from "@/app/libs/store/authSlice";

const loginSchema = z.object({
  email: z
    .string()
    .email({ message: "有効なメールアドレスを入力してください。" }),
  password: z
    .string()
    .min(6, { message: "パスワードは6文字以上で入力してください。" }),
});

type LoginFormInputs = z.infer<typeof loginSchema>;

export default function LoginForm() {
  const dispatch = useDispatch();
  const router = useRouter();

  const {
    register,
    handleSubmit,
    formState: { errors },
  } = useForm<LoginFormInputs>({
    resolver: zodResolver(loginSchema),
  });

  const { mutate: loginMutation, isPending } = usePostAuthLogin();

  const onSubmit = (data: LoginFormInputs) => {
    loginMutation({ data }, {
      onSuccess: (user) => {
        if (user.user) {
          dispatch(login(user.user));
        }
        toast.success("ログインに成功しました");
        router.push("/dashboard");
      },
      onError: (error: any) => {
        toast.error(error.message || "ログインに失敗しました");
      },
    });
  };

  return (
    <div className="flex items-center justify-center min-h-screen bg-gray-100 dark:bg-gray-900">
      <form
        onSubmit={handleSubmit(onSubmit)}
        noValidate
        className="bg-white dark:bg-gray-800 p-6 rounded-lg shadow-lg w-96 space-y-4"
      >
        <h2 className="text-2xl font-bold text-center text-gray-700 dark:text-gray-200">
          ログイン
        </h2>

        {errors.email && (
          <div className="text-red-500 text-sm text-center">
            {errors.email.message}
          </div>
        )}
        {errors.password && (
          <div className="text-red-500 text-sm text-center">
            {errors.password.message}
          </div>
        )}

        <div>
          <label className="block text-gray-700 dark:text-gray-300">
            メールアドレス:
          </label>
          <input
            className="w-full mt-1 p-2 border border-gray-300 dark:border-gray-600 rounded-md focus:ring-2 focus:ring-blue-500 dark:focus:ring-blue-400 bg-white dark:bg-gray-700 text-gray-900 dark:text-gray-200"
            type="email"
            {...register("email")}
          />
        </div>

        <div>
          <label className="block text-gray-700 dark:text-gray-300">
            パスワード:
          </label>
          <input
            className="w-full mt-1 p-2 border border-gray-300 dark:border-gray-600 rounded-md focus:ring-2 focus:ring-blue-500 dark:focus:ring-400 bg-white dark:bg-gray-700 text-gray-900 dark:text-gray-200"
            type="password"
            {...register("password")}
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
              <p className="font-mono">EygQJpu@NillQOs.net</p>
            </div>
            <button
              type="button"
              className="ml-2 px-2 py-1 text-sm border-gray-400 border rounded hover:bg-blue-500 hover:text-white transition"
              onClick={() => navigator.clipboard.writeText("EygQJpu@NillQOs.net")}
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

        <LoadingButton
          label={"ログイン"}
          color={"blue"}
          isPending={isPending}
        />
      </form>
    </div>
  );
}
