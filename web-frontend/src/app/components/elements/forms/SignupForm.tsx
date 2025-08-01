"use client";

import { zodResolver } from "@hookform/resolvers/zod";
import { useRouter } from "next/navigation";
import { useForm } from "react-hook-form";
import toast from "react-hot-toast";
import { z } from "zod";
import LoadingButton from "../buttons/LoadingButton";
import { usePostAuthRegister } from "@/app/libs/api/generated/orval/auth/auth";

const signupSchema = z.object({
  name: z.string().min(1, "名前を入力してください"),
  email: z
    .string()
    .email({ message: "有効なメールアドレスを入力してください。" }),
  password: z
    .string()
    .min(6, { message: "パスワードは6文字以上で入力してください。" }),
});

type SignupFormInputs = z.infer<typeof signupSchema>;

export default function SignupForm() {
  const router = useRouter();
  const { mutate: signupMutation, isPending } = usePostAuthRegister();

  const {
    register,
    handleSubmit,
    formState: { errors },
  } = useForm<SignupFormInputs>({
    resolver: zodResolver(signupSchema),
  });

  const onSubmit = (data: SignupFormInputs) => {
    signupMutation({ data }, {
      onSuccess: () => {
        toast.success("登録しました。");
        router.push("/dashboard");
      },
      onError: (error: any) => {
        toast.error("登録に失敗しました。");
        console.error("Signup failed:", error);
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
          新規登録
        </h2>

        {errors.name && (
          <div className="text-red-500 text-sm text-center">
            {errors.name.message}
          </div>
        )}
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
            名前:
          </label>
          <input
            className="w-full mt-1 p-2 border border-gray-300 dark:border-gray-600 rounded-md focus:ring-2 focus:ring-green-500 dark:focus:ring-green-400 bg-white dark:bg-gray-700 text-gray-900 dark:text-gray-200"
            type="text"
            {...register("name")}
          />
        </div>
        <div>
          <label className="block text-gray-700 dark:text-gray-300">
            メールアドレス:
          </label>
          <input
            className="w-full mt-1 p-2 border border-gray-300 dark:border-gray-600 rounded-md focus:ring-2 focus:ring-green-500 dark:focus:ring-green-400 bg-white dark:bg-gray-700 text-gray-900 dark:text-gray-200"
            type="email"
            {...register("email")}
          />
        </div>
        <div>
          <label className="block text-gray-700 dark:text-gray-300">
            パスワード:
          </label>
          <input
            className="w-full mt-1 p-2 border border-gray-300 dark:border-gray-600 rounded-md focus:ring-2 focus:ring-green-500 dark:focus:ring-green-400 bg-white dark:bg-gray-700 text-gray-900 dark:text-gray-200"
            type="password"
            {...register("password")}
          />
        </div>
        <LoadingButton
          label={"登録する"}
          color={"green"}
          isPending={isPending}
        />
      </form>
    </div>
  );
}
