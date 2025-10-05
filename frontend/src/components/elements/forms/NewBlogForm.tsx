"use client";

import { zodResolver } from "@hookform/resolvers/zod";
import "easymde/dist/easymde.min.css";
import dynamic from "next/dynamic";
import Image from "next/image";
import { useRouter } from "next/navigation";
import { useEffect } from "react";
import { Controller, useForm } from "react-hook-form";
import toast from "react-hot-toast";
const SimpleMDEEditor = dynamic(() => import("react-simplemde-editor"), {
  ssr: false,
});
import Select from "react-select";
import { z } from "zod";
import { CommonButton } from "../buttons/CommonButton";
import { Track } from "@/libs/api/generated/orval/model/track";
import { usePostPosts } from "@/libs/api/generated/orval/posts/posts";
import {
  usePostTagsPostsPostID,
  useGetTags,
} from "@/libs/api/generated/orval/tags/tags";

const schema = z.object({
  title: z.string().min(1, "タイトルを入力してください"),
  description: z.string().min(1, "内容を入力してください"),
  tags: z.array(z.string()).optional(), // tagsを文字列の配列に変更
  track: z
    .object({
      spotifyId: z.string().optional(),
      name: z.string().optional(),
      artistName: z.string().optional(),
      albumImageUrl: z.string().optional(),
    })
    .nullable(),
});

type FormData = {
  title: string;
  description: string;
  track: Track | null;
  tags?: string[]; // tagsの型を文字列の配列に変更
};

type NewBlogFormProps = {
  selectedTrack: Track | null;
};

const NewBlogForm = ({ selectedTrack }: NewBlogFormProps) => {
  const router = useRouter();
  const {
    register,
    handleSubmit,
    setValue,
    formState: { errors, isSubmitting },
    reset,
    control,
  } = useForm<FormData>({
    resolver: zodResolver(schema),
    defaultValues: {
      title: "",
      description: "",
      track: null,
      tags: [], // デフォルト値を空の配列に
    },
  });

  useEffect(() => {
    setValue("track", selectedTrack);
  }, [selectedTrack, setValue]);

  const createPostMutation = usePostPosts();
  const addTagsToPostMutation = usePostTagsPostsPostID();
  const { data: allTagsData } = useGetTags(); // 全てのタグを取得

  const tagOptions =
    allTagsData?.tags?.map((tag) => ({
      value: tag.name || "",
      label: tag.name || "",
    })) || [];

  const onSubmit = async (data: FormData) => {
    createPostMutation.mutate(
      { data: { ...data, tags: undefined } }, // tagsは別途処理するためundefinedにする
      {
        onSuccess: (response) => {
          toast.success("ブログが作成されました");
          reset();
          router.push("/dashboard");

          if (data.tags && data.tags.length > 0 && response.post?.id) {
            addTagsToPostMutation.mutate({
              postID: response.post.id,
              data: { tag_names: data.tags },
            });
          }
        },
        onError: (error: any) => {
          toast.error(error.message || "ブログの作成に失敗しました");
        },
      },
    );
  };

  return (
    <form
      onSubmit={handleSubmit(onSubmit)}
      className="w-full max-w-2xl mx-auto bg-white p-6 rounded-lg shadow-md"
    >
      <div className="mb-4">
        <label
          htmlFor="title"
          className="block text-sm font-medium text-gray-700 mb-1"
        >
          タイトル
        </label>
        <input
          type="text"
          {...register("title")}
          className="w-full px-4 py-2 border rounded-md focus:outline-none focus:ring-2 focus:ring-green-500"
        />
        {errors.title && (
          <p className="mt-1 text-sm text-red-500">{errors.title.message}</p>
        )}
      </div>

      <div className="mb-4">
        <label
          htmlFor="description"
          className="block text-sm font-medium text-gray-700 mb-1"
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
          <p className="mt-1 text-sm text-red-500">
            {errors.description.message}
          </p>
        )}
      </div>

      <div className="mb-4">
        <label
          htmlFor="tags"
          className="block text-sm font-medium text-gray-700 mb-1"
        >
          タグ
        </label>
        <Controller
          name="tags"
          control={control}
          render={({ field }) => (
            <Select
              {...field}
              options={tagOptions}
              isMulti
              isClearable
              placeholder="タグを選択または検索..."
              classNamePrefix="react-select"
              onChange={(selectedOptions) =>
                field.onChange(
                  selectedOptions
                    ? selectedOptions.map((option) => option.value)
                    : [],
                )
              }
              value={tagOptions.filter((option) =>
                (field.value || []).includes(option.value),
              )}
            />
          )}
        />
        {errors.tags && (
          <p className="mt-1 text-sm text-red-500">{errors.tags.message}</p>
        )}
      </div>

      {selectedTrack ? (
        <div className="flex items-center gap-4 p-4 mt-6 border rounded-md bg-gray-50">
          <Image
            width={50}
            height={50}
            src={selectedTrack.albumImageUrl || "/default-image.jpg"}
            alt={selectedTrack.name || ""}
            className="w-16 h-16 object-cover rounded"
          />
          <div>
            <p className="text-lg font-semibold">{selectedTrack.name}</p>
            <p className="text-sm text-gray-500">{selectedTrack.artistName}</p>
          </div>
        </div>
      ) : (
        <p className="mt-4 text-sm text-gray-500">※ 曲が選択されていません。</p>
      )}
      <div className="mt-4 grid grid-cols-1 md:grid-cols-2 gap-8">
        <button
          type="submit"
          disabled={isSubmitting}
          className="px-4 py-2 bg-green-500 text-white rounded-md hover:bg-green-600 dark:hover:bg-green-700 disabled:opacity-50"
        >
          Post
        </button>
        <CommonButton href={`/dashboard/blog/page/1`}>Back</CommonButton>
      </div>
    </form>
  );
};

export default NewBlogForm;
