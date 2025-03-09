"use client";

import { useRouter } from "next/navigation";
import { useEffect } from "react";
import { useForm } from "react-hook-form";
import { z } from "zod";
import { zodResolver } from "@hookform/resolvers/zod";
import toast, { Toaster } from "react-hot-toast";
import { getBlogById, editBlog, deleteBlog } from "@/app/lib/utils/blog";

const schema = z.object({
  title: z.string().min(1, "タイトルを入力してください"),
  description: z.string().min(1, "内容を入力してください"),
});

const EditPost = ({ params }: { params: { id: number } }) => {
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

  useEffect(() => {
    getBlogById(params.id)
      .then((data) => {
        reset(data);
      })
      .catch(() => {
        toast.error("Error fetching data");
      });
  }, [params.id, reset]);

  const onSubmit = async (data: { title: string; description: string }) => {
    try {
      await editBlog(data.title, data.description, params.id);
      toast.success("Updated!", { duration: 1500 });
      setTimeout(() => {
        router.push("/dashboard/blog/page/1");
        router.refresh();
      }, 2000);
    } catch (error) {
      toast.error("Failed to update post.");
    }
  };

  const handleDelete = async () => {
    try {
      await deleteBlog(params.id);
      toast.success("Deleted!");
      setTimeout(() => {
        router.push("/dashboard/blog/page/1");
        router.refresh();
      }, 2000);
    } catch (error) {
      toast.error("Failed to delete post.");
    }
  };

  return (
    <>
      <Toaster />
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
        <div className="form__btn-area">
          <button
            className="form__btn form__btn--update"
            type="submit"
            disabled={isSubmitting}
          >
            <span></span>Update
          </button>
          <button
            className="form__btn form__btn--delete"
            type="button"
            onClick={handleDelete}
          >
            <span></span>Delete
          </button>
        </div>
      </form>
    </>
  );
};

export default EditPost;
