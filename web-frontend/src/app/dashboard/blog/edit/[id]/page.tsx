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
import {
  useGetBlogsId,
  usePutBlogsId,
} from "@/libs/api/generated/orval/blogs/blogs";
import { Tag } from "@/libs/api/generated/orval/model/tag";
import { useDeletePostsId } from "@/libs/api/generated/orval/posts/posts";
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

export default function Page() {
  const router = useRouter();
  const params = useParams();
  const { id } = params as { id: string };
  const { data: post, isPending, error } = useGetBlogsId(Number(id));
  const { data: tags } = useGetTagsPostsPostID(Number(id));
  const { mutate: updatePost } = usePutBlogsId();
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
    if (post) {
      reset({ title: post.post?.title, description: post.post?.description });
    }
    if (tags && tags.tags) {
      setValue("tags", tags.tags.map((tag: Tag) => tag.name).join(", "));
    }
  }, [post, tags, reset, setValue]);

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
          toast.success("更新しました！");
          reset();
        },
        onError: (error: any) => {
          console.error("Update error:", error);
          toast.error("更新に失敗しました。");
        },
      }
    );

    const currentTagNames = tags?.tags?.map((tag: Tag) => tag.name) || [];
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
    deletePost(
      { id: Number(id) },
      {
        onSuccess: () => {
          toast.success("削除しました。");
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

  if (isPending) return <div>Loading...</div>;
  if (error || !post) return <div>Error loading post.</div>;

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
          <Controller
            name="description"
            control={control}
            render={({ field }) => {
              const memoizedOptions = () => ({
                spellChecker: false,
                hideIcons: ["side-by-side", "fullscreen"] as const,
              });
              return (
                <SimpleMDEEditor
                  key="description-editor"
                  value={field.value}
                  onChange={field.onChange}
                  // options={memoizedOptions}
                />
              );
            }}
          />
          {errors.description && (
            <p className="text-red-500 text-sm mt-1">
              {errors.description.message}
            </p>
          )}
        </div>
        <div className="mb-4">
          <label
            htmlFor="tags"
            className="block text-gray-700 dark:text-gray-300 font-medium mb-1"
          >
            タグ (カンマ区切り)
          </label>
          <input
            type="text"
            {...register("tags")}
            className="w-full p-2 border rounded-md dark:bg-gray-700 dark:text-white focus:ring focus:ring-indigo-300"
            placeholder="例: プログラミング, 日常, 音楽"
          />
          {errors.tags && (
            <p className="text-red-500 text-sm mt-1">{errors.tags.message}</p>
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
