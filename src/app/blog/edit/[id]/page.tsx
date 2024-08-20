"use client";

import { useRouter } from "next/navigation";
import { useEffect, useRef } from "react";
import toast, { Toaster } from "react-hot-toast";
import "@/scss/modal.scss";

const editBlog = async (
  title: string | undefined,
  description: string | undefined,
  id: number
) => {
  const res = await fetch(
    `https://my-next-blog-m1sli2z91-southvillages-projects.vercel.app/api/blog/${id}`,
    {
      method: "PUT",
      body: JSON.stringify({ title, description, id }),
      headers: {
        "Content-Type": "application/json",
      },
    }
  );

  return res.json();
};

const deleteBlog = async (id: number) => {
  const res = await fetch(
    `https://my-next-blog-m1sli2z91-southvillages-projects.vercel.app/api/blog/${id}`,
    {
      method: "DELETE",
      headers: {
        "Content-Type": "application/json",
      },
    }
  );

  return res.json();
};

const getBlogById = async (id: number) => {
  const res = await fetch(
    `https://my-next-blog-m1sli2z91-southvillages-projects.vercel.app/api/blog/${id}`
  );
  const data = await res.json();
  return data.post;
};

const EditPost = ({ params }: { params: { id: number } }) => {
  const router = useRouter();
  const titleRef = useRef<HTMLInputElement | null>(null);
  const descriptionRef = useRef<HTMLTextAreaElement | null>(null);

  const handleSubmit = async (e: React.FormEvent) => {
    e.preventDefault();

    toast.loading("Editing...");

    await editBlog(
      titleRef.current?.value,
      descriptionRef.current?.value,
      params.id
    );

    router.push("/blog");
    router.refresh();

    toast.success("Updated!", {
      duration: 2000,
    });
  };

  const handleDelete = async (e: React.FormEvent) => {
    toast.loading("deleting");
    await deleteBlog(params.id);

    router.push("/blog");
    router.refresh();
  };

  useEffect(() => {
    getBlogById(params.id)
      .then((data) => {
        titleRef.current!.value = data.title;
        descriptionRef.current!.value = data.description;
      })
      .catch((error) => {
        toast.error("Error");
      });
  }, []);

  return (
    <>
      <Toaster />
      <form onSubmit={handleSubmit}>
        <input type="text" ref={titleRef} placeholder="タイトルを入力" />
        <textarea
          ref={descriptionRef}
          name=""
          placeholder="記事を入力"
          id=""
        ></textarea>
        <button>更新</button>
        <button onClick={handleDelete}>削除</button>
      </form>
    </>
  );
};

export default EditPost;
