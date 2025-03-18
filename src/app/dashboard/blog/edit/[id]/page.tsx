"use client";

import { zodResolver } from "@hookform/resolvers/zod";
import { useRouter } from "next/navigation";
import { useEffect, useCallback } from "react";
import { useForm } from "react-hook-form";
import toast, { Toaster } from "react-hot-toast";
import { z } from "zod";
import { CommonButton } from "@/app/components/elements/CommonButton";
import { getBlogById, editBlog, deleteBlog } from "@/app/lib/utils/blog";

const schema = z.object({
  title: z.string().min(1, "タイトルを入力してください"),
  description: z.string().min(1, "内容を入力してください"),
});

export default function Page({ params }: { params: { id: number } }) {
  const router = useRouter();
  const {
    register,
    handleSubmit,
    formState: { errors, isSubmitting },
    reset,
  } = useForm<{ title: string; description: string }>({
    resolver: zodResolver(schema),
    defaultValues: { title: "", description: "" },
  });

  const fetchPost = useCallback(async () => {
    try {
      const data = await getBlogById(params.id);
      reset(data);
    } catch {
      toast.error("データの取得に失敗しました");
    }
  }, [params.id, reset]);

  useEffect(() => {
    fetchPost();
  }, [fetchPost]);

  const onSubmit = async (data: { title: string; description: string }) => {
    try {
      await editBlog(data.title, data.description, params.id);
      toast.success("更新しました！", { duration: 1500 });
      setTimeout(() => {
        router.push("/dashboard/blog/page/1");
        router.refresh();
      }, 2000);
    } catch (error) {
      toast.error("更新に失敗しました。");
    }
  };

  const handleDelete = async () => {
    try {
      await deleteBlog(params.id);
      toast.success("削除しました！");
      setTimeout(() => {
        router.push("/dashboard/blog/page/1");
        router.refresh();
      }, 2000);
    } catch (error) {
      toast.error("削除に失敗しました。");
    }
  };

  return (
    <div className="min-h-screen flex items-center justify-center bg-gray-100 dark:bg-gray-900">
      <Toaster />
      <form
        onSubmit={handleSubmit(onSubmit)}
        className="bg-white dark:bg-gray-800 shadow-lg rounded-lg p-6 w-full max-w-lg"
      >
        <h2 className="text-2xl font-semibold text-center text-gray-900 dark:text-gray-100 mb-4">
          記事の編集
        </h2>
        <div className="mb-4">
          <label
            htmlFor="title"
            className="block text-gray-700 dark:text-gray-300 font-medium mb-1"
          >
            タイトル
          </label>
          <input
            type="text"
            {...register("title")}
            className="w-full p-2 border rounded-md dark:bg-gray-700 dark:text-white focus:ring focus:ring-indigo-300"
          />
          {errors.title && (
            <p className="text-red-500 text-sm mt-1">{errors.title.message}</p>
          )}
        </div>
        <div className="mb-4">
          <label
            htmlFor="description"
            className="block text-gray-700 dark:text-gray-300 font-medium mb-1"
          >
            内容
          </label>
          <textarea
            {...register("description")}
            className="w-full p-2 border rounded-md dark:bg-gray-700 dark:text-white focus:ring focus:ring-indigo-300"
            rows={5}
          ></textarea>
          {errors.description && (
            <p className="text-red-500 text-sm mt-1">
              {errors.description.message}
            </p>
          )}
        </div>
        <div className="flex space-x-4">
          <button
            className="flex-1 bg-blue-500 hover:bg-blue-600 text-white font-semibold py-2 px-4 rounded-md transition disabled:opacity-50"
            type="submit"
            disabled={isSubmitting}
          >
            更新
          </button>
          <button
            className="flex-1 bg-red-500 hover:bg-red-600 text-white font-semibold py-2 px-4 rounded-md transition"
            type="button"
            onClick={handleDelete}
          >
            削除
          </button>
          <CommonButton href={`/dashboard/blog/page/1`}>Back</CommonButton>
        </div>
      </form>
    </div>
  );
}
