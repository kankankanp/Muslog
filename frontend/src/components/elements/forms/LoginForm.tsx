"use client";

import { zodResolver } from "@hookform/resolvers/zod";
import { useRouter, useSearchParams } from "next/navigation";
import { Suspense } from "react";
import { useForm } from "react-hook-form";
import toast from "react-hot-toast";
import { useDispatch } from "react-redux";
import { z } from "zod";
import LoadingButton from "../buttons/LoadingButton";
import Spinner from "@/components/layouts/Spinner";
import { GUEST_EMAIL, GUEST_PASSWORD } from "@/constants/guestUser";
import {
  useGetAuthGoogle,
  usePostAuthLogin,
} from "@/libs/api/generated/orval/auth/auth";
import { login } from "@/libs/store/authSlice";

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
  const searchParams = useSearchParams();
  const returnUrl = searchParams.get("returnUrl");

  const {
    register,
    handleSubmit,
    formState: { errors },
    setValue,
  } = useForm<LoginFormInputs>({ resolver: zodResolver(loginSchema) });

  const { mutate: loginMutation, isPending } = usePostAuthLogin();
  const { refetch: getAuthGoogle, isFetching: isGoogleLoginPending } =
    useGetAuthGoogle({
      query: {
        enabled: false, // Initially disable the query
      },
    });

  const onSubmit = (data: LoginFormInputs) => {
    loginMutation(
      { data },
      {
        onSuccess: (user: any) => {
          if (user?.user) dispatch(login(user.user));
          toast.success("ログインに成功しました");
          // returnUrlがある場合はそこに、なければ/dashboardに遷移
          const redirectTo = returnUrl
            ? decodeURIComponent(returnUrl)
            : "/dashboard";
          router.push(redirectTo);
        },
        onError: (error: any) => {
          console.error("login error:", error);
          toast.error("ログインに失敗しました");
        },
      },
    );
  };

  const handleGuestLogin = () => {
    setValue("email", GUEST_EMAIL, { shouldValidate: true });
    setValue("password", GUEST_PASSWORD, { shouldValidate: true });
    handleSubmit(onSubmit)();
  };

  const handleGoogleLogin = async () => {
    try {
      const { data } = await getAuthGoogle();
      if (data?.authURL) {
        window.location.href = data.authURL;
      }
    } catch (error) {
      console.error("Google login error:", error);
      toast.error("Googleログインに失敗しました。");
    }
  };

  return (
    <Suspense fallback={<Spinner />}>
      <div className="flex items-center justify-center dark:bg-gray-900 h-[360px]">
        <form
          onSubmit={handleSubmit(onSubmit)}
          noValidate
          className="dark:bg-gray-800 md:w-96 space-y-4"
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
              メールアドレス
            </label>
            <input
              className="w-full mt-1 p-2 max-md:px-[52px] border border-gray-300 dark:border-gray-600 rounded-md focus:ring-2 focus:ring-blue-500 dark:focus:ring-blue-400 bg-white dark:bg-gray-700 text-gray-900 dark:text-gray-200"
              type="email"
              {...register("email")}
            />
          </div>

          <div>
            <label className="block text-gray-700 dark:text-gray-300">
              パスワード
            </label>
            <input
              className="w-full mt-1 p-2 border border-gray-300 dark:border-gray-600 rounded-md focus:ring-2 focus:ring-blue-500 dark:focus:ring-400 bg-white dark:bg-gray-700 text-gray-900 dark:text-gray-200"
              type="password"
              {...register("password")}
            />
          </div>

          <div className="flex flex-col gap-4 w-[60%] mx-auto font-bold max-md:text-sm">
            <LoadingButton
              label={"ログイン"}
              color={"blue"}
              isPending={isPending}
            />
            <button
              type="button"
              onClick={handleGuestLogin}
              disabled={isPending}
              className="px-4 py-3 rounded-full border border-gray-300 bg-white hover:bg-blue-50 dark:bg-gray-700 dark:hover:bg-gray-600 dark:border-gray-600 text-center"
              aria-label="ゲストログイン"
            >
              ゲストログイン
            </button>
            <button
              type="button"
              onClick={handleGoogleLogin}
              disabled={isPending || isGoogleLoginPending}
              className="px-4 py-3 rounded-full border border-red-500 bg-red-500 text-white hover:bg-red-600 dark:bg-red-600 dark:hover:bg-red-700 dark:border-red-700 text-center"
              aria-label="Googleでログイン"
            >
              Googleでログイン
            </button>
          </div>
        </form>
      </div>
    </Suspense>
  );
}
