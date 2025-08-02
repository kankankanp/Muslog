"use client";

import { zodResolver } from "@hookform/resolvers/zod";
import "easymde/dist/easymde.min.css";
import { useParams, useRouter } from "next/navigation";
import { useEffect } from "react";
import { Controller, useForm } from "react-hook-form";
import toast, { Toaster } from "react-hot-toast";
import SimpleMDEEditor from "react-simplemde-editor";
import { z } from "zod";
import { CommonButton } from "@/components/elements/buttons/CommonButton";
import { Tag } from "@/libs/api/generated/orval/model/tag";
import {
  useGetPostsId,
  usePutPostsId,
  useDeletePostsId,
} from "@/libs/api/generated/orval/posts/posts";
import {
  useGetTagsPostsPostID,
  usePostTagsPostsPostID,
  useDeleteTagsPostsPostID,
} from "@/libs/api/generated/orval/tags/tags";

const schema = z.object({
  title: z.string().min(1, "タイトルを入力してください"),
  description: z.string().min(1, "内容を入力してください"),
  tags: z.string().optional(),
});

export default function PostEditPage() {
  const router = useRouter();
  const params = useParams();
  const { id } = params as { id: string };
  const { data: postData, isPending, error } = useGetPostsId(Number(id));
  const { data: tagsData } = useGetTagsPostsPostID(Number(id));
  const { mutate: updatePost } = usePutPostsId();
  const { mutate: deletePost } = useDeletePostsId();
  const addTagsToPostMutation = usePostTagsPostsPostID();
  const removeTagsFromPostMutation = useDeleteTagsPostsPostID();

  const {
    register,
    handleSubmit,
    setValue,
    formState: { errors, isSubmitting },
    reset,
    control,
  } = useForm<{ title: string; description: string; tags?: string }>({
    resolver: zodResolver(schema),
  });

  useEffect(() => {
    if (postData?.post) {
      reset({
        title: postData.post.title,
        description: postData.post.description,
      });
    }
    if (tagsData?.tags) {
      setValue("tags", tagsData.tags.map((tag: Tag) => tag.name).join(", "));
    }
  }, [postData, tagsData, reset, setValue]);

  const onSubmit = async (data: {
    title: string;
    description: string;
    tags?: string;
  }) => {
    updatePost(
      {
        id: Number(id),
        data: {
          title: data.title,
          description: data.description,
        },
      },
      {
        onSuccess: () => {
          toast.success("投稿を更新しました！");
          reset();
        },
        onError: (error: any) => {
          console.error("Update error:", error);
          toast.error("更新に失敗しました。");
        },
      }
    );

    // タグの更新処理
    const currentTagNames = tagsData?.tags?.map((tag: Tag) => tag.name) || [];
    const newTagNames = data.tags
      ? data.tags
          .split(",")
          .map((tag) => tag.trim())
          .filter((tag) => tag.length > 0)
      : [];

    const tagsToAdd = newTagNames.filter(
      (tagName) => !currentTagNames.includes(tagName)
    );
    const tagsToRemove = currentTagNames.filter(
      (tagName): tagName is string =>
        typeof tagName === "string" && !newTagNames.includes(tagName)
    );

    if (tagsToAdd.length > 0) {
      addTagsToPostMutation.mutate({
        postID: Number(id),
        data: { tag_names: tagsToAdd },
      });
    }
    if (tagsToRemove.length > 0) {
      removeTagsFromPostMutation.mutate({
        postID: Number(id),
        data: { tag_names: tagsToRemove },
      });
    }
  };

  const handleDelete = async () => {
    if (!confirm("本当にこの投稿を削除しますか？")) return;

    deletePost(
      { id: Number(id) },
      {
        onSuccess: () => {
          toast.success("投稿を削除しました。");
          setTimeout(() => {
            router.push("/dashboard/blog/page/1");
            router.refresh();
          }, 2000);
        },
        onError: (error: any) => {
          console.error("Delete error:", error);
          toast.error("削除に失敗しました。");
        },
      }
    );
  };

  if (isPending) {
    return (
      <div className="flex items-center justify-center min-h-screen bg-gray-100 dark:bg-gray-900">
        <div className="text-center">
          <div className="animate-spin rounded-full h-8 w-8 border-b-2 border-blue-500 mx-auto mb-4"></div>
          <p className="text-gray-600 dark:text-gray-300">
            投稿を読み込み中...
          </p>
        </div>
      </div>
    );
  }

  if (error || !postData?.post) {
    return (
      <div className="flex items-center justify-center min-h-screen bg-gray-100 dark:bg-gray-900">
        <div className="text-center">
          <p className="text-red-600 dark:text-red-400">
            投稿の読み込みに失敗しました。
          </p>
          <CommonButton href="/dashboard/blog/page/1">戻る</CommonButton>
        </div>
      </div>
    );
  }

  return (
    <div className="min-h-screen flex items-center justify-center bg-gray-100 dark:bg-gray-900 p-4">
      <Toaster />
      <form
        onSubmit={handleSubmit(onSubmit)}
        className="bg-white dark:bg-gray-800 shadow-lg rounded-lg p-6 w-full max-w-2xl"
      >
        <h2 className="text-2xl font-semibold text-center text-gray-900 dark:text-gray-100 mb-6">
          投稿の編集
        </h2>

        <div className="mb-4">
          <label
            htmlFor="title"
            className="block text-gray-700 dark:text-gray-300 font-medium mb-2"
          >
            タイトル
          </label>
          <input
            type="text"
            {...register("title")}
            className="w-full p-3 border border-gray-300 dark:border-gray-600 rounded-md dark:bg-gray-700 dark:text-white focus:ring-2 focus:ring-blue-500 focus:border-transparent"
            placeholder="投稿のタイトルを入力してください"
          />
          {errors.title && (
            <p className="text-red-500 text-sm mt-1">{errors.title.message}</p>
          )}
        </div>

        <div className="mb-4">
          <label
            htmlFor="description"
            className="block text-gray-700 dark:text-gray-300 font-medium mb-2"
          >
            内容
          </label>
          <Controller
            name="description"
            control={control}
            render={({ field }) => (
              <SimpleMDEEditor
                key="description-editor"
                value={field.value}
                onChange={field.onChange}
                options={{
                  spellChecker: false,
                  hideIcons: ["side-by-side", "fullscreen"],
                  placeholder: "投稿の内容をMarkdownで入力してください",
                }}
              />
            )}
          />
          {errors.description && (
            <p className="text-red-500 text-sm mt-1">
              {errors.description.message}
            </p>
          )}
        </div>

        <div className="mb-6">
          <label
            htmlFor="tags"
            className="block text-gray-700 dark:text-gray-300 font-medium mb-2"
          >
            タグ (カンマ区切り)
          </label>
          <input
            type="text"
            {...register("tags")}
            className="w-full p-3 border border-gray-300 dark:border-gray-600 rounded-md dark:bg-gray-700 dark:text-white focus:ring-2 focus:ring-blue-500 focus:border-transparent"
            placeholder="例: プログラミング, 日常, 音楽"
          />
          {errors.tags && (
            <p className="text-red-500 text-sm mt-1">{errors.tags.message}</p>
          )}
        </div>

        <div className="flex flex-col sm:flex-row gap-3">
          <button
            className="flex-1 bg-blue-500 hover:bg-blue-600 text-white font-semibold py-3 px-4 rounded-md transition duration-200 disabled:opacity-50 disabled:cursor-not-allowed"
            type="submit"
            disabled={isSubmitting}
          >
            {isSubmitting ? "更新中..." : "更新"}
          </button>
          <button
            className="flex-1 bg-red-500 hover:bg-red-600 text-white font-semibold py-3 px-4 rounded-md transition duration-200"
            type="button"
            onClick={handleDelete}
          >
            削除
          </button>
          <CommonButton href="/dashboard">キャンセル</CommonButton>
        </div>
      </form>
    </div>
  );
}
