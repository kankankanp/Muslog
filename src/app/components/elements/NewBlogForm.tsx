"use client";

import { useForm } from "react-hook-form";
import { z } from "zod";
import { zodResolver } from "@hookform/resolvers/zod";
import toast from "react-hot-toast";
import { useRouter } from "next/navigation";
import { Track } from "@/app/components/elements/SelectMusciArea";
import "@/scss/new-blog-form.scss";

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
      track, // 曲情報も一緒に送信
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
    <form onSubmit={handleSubmit(onSubmit)} className="form">
      <div className="form__title">
        <label htmlFor="title">タイトル</label>
        <input type="text" {...register("title")} />
        {errors.title && <p className="error">{errors.title.message}</p>}
      </div>
      <div className="form__description">
        <label htmlFor="description">内容</label>
        <textarea {...register("description")} name="description"></textarea>
        {errors.description && (
          <p className="error">{errors.description.message}</p>
        )}
      </div>
      <button className="form__btn" type="submit" disabled={isSubmitting}>
        <span></span>Post
      </button>
    </form>
  );
};

export default NewBlogForm;
