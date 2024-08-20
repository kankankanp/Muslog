"use client";

import { useRouter } from "next/navigation";
import { useRef } from "react";
import toast, { Toaster } from "react-hot-toast";
import "@/scss/modal.scss";

const postBlog = async (
  title: string | undefined,
  description: string | undefined
) => {
  const res = await fetch(
    "https://my-next-blog-m1sli2z91-southvillages-projects.vercel.app/api/blog",
    {
      method: "POST",
      body: JSON.stringify({ title, description }),
      headers: {
        "Content-Type": "application/json",
      },
    }
  );

  return res.json();
};

const PostBlog = () => {
  const router = useRouter();

  const titleRef = useRef<HTMLInputElement | null>(null);
  const descriptionRef = useRef<HTMLTextAreaElement | null>(null);

  const handleSubmit = async (e: React.FormEvent) => {
    e.preventDefault();

    await postBlog(titleRef.current?.value, descriptionRef.current?.value);

    toast.success("Posted!", {
      duration: 2000,
    });

    setTimeout(() => {
      router.push("/blog/page/1");
      router.refresh();
    }, 1500);
  };

  return (
    <>
      <Toaster />
      <form onSubmit={handleSubmit} className="form">
        <div className="form__title">
          <label htmlFor="title">タイトル</label>
          <input type="text" ref={titleRef} />
        </div>
        <div className="form__description">
          <label htmlFor="description">内容</label>
          <textarea ref={descriptionRef} name="description"></textarea>
        </div>
        <button className="form__btn">
          <span></span>Post
        </button>
      </form>
    </>
  );
};

export default PostBlog;
