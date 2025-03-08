"use client";

import { useForm } from "react-hook-form";
import { z } from "zod";
import { zodResolver } from "@hookform/resolvers/zod";
import toast from "react-hot-toast";
import { useRouter } from "next/navigation";
import { Track } from "@/app/components/elements/SelectMusciArea";
import Image from "next/image";

const ENDPOINT = process.env.NEXT_PUBLIC_API_URL;

const postBlog = async (
  title: string,
  description: string,
  track: Track | null
) => {
  const res = await fetch(`${ENDPOINT}/api/blog`, {
    method: "POST",
    body: JSON.stringify({
      title,
      description,
      track,
    }),
    headers: {
      "Content-Type": "application/json",
    },
  });

  return res.json();
};

const schema = z.object({
  title: z.string().min(1, "タイトルを入力してください"),
  description: z.string().min(1, "内容を入力してください"),
});

type NewBlogFormProps = {
  selectedTrack: Track | null;
};

const NewBlogForm = ({ selectedTrack }: NewBlogFormProps) => {
  console.log(selectedTrack);
  const router = useRouter();
  const {
    register,
    handleSubmit,
    formState: { errors, isSubmitting },
    reset,
  } = useForm<{ title: string; description: string }>({
    resolver: zodResolver(schema),
  });

  const onSubmit = async (data: { title: string; description: string }) => {
    try {
      await postBlog(data.title, data.description, selectedTrack);

      toast.success("Posted!", { duration: 1500 });
      setTimeout(() => {
        reset();
        router.push("/dashboard/blog/page/1");
        router.refresh();
      }, 2000);
    } catch {
      toast.error("Failed to post.");
    }
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
        <textarea
          {...register("description")}
          name="description"
          rows={5}
          className="w-full px-4 py-2 border rounded-md focus:outline-none focus:ring-2 focus:ring-green-500"
        ></textarea>
        {errors.description && (
          <p className="mt-1 text-sm text-red-500">
            {errors.description.message}
          </p>
        )}
      </div>

      {selectedTrack ? (
        <div className="flex items-center gap-4 p-4 mt-6 border rounded-md bg-gray-50">
          <Image
            width={50}
            height={50}
            src={selectedTrack.albumImageUrl}
            alt={selectedTrack.name}
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

      <button
        type="submit"
        disabled={isSubmitting}
        className="mt-6 w-full px-4 py-2 bg-green-500 text-white rounded-md hover:bg-green-600 disabled:opacity-50"
      >
        Post
      </button>
    </form>
  );
};

export default NewBlogForm;
