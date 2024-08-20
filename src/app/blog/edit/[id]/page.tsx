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
    `https://my-next-blog-iota-six.vercel.app/api/blog/${id}`,
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
    `https://my-next-blog-iota-six.vercel.app/api/blog/${id}`,
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
    `https://my-next-blog-iota-six.vercel.app/api/blog/${id}`
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

    await editBlog(
      titleRef.current?.value,
      descriptionRef.current?.value,
      params.id
    );

    toast.success("Updated!", {
      duration: 2000,
    });

    setTimeout(() => {
      router.push("/blog/page/1");
      router.refresh();
    }, 1500);
  };

  const handleDelete = async (e: React.FormEvent) => {
    await deleteBlog(params.id);

    // toast.error("Deleted!", {
    //   duration: 2000,
    // });

    // setTimeout(() => {
    //   router.push("/blog/page/1");
    //   router.refresh();
    // }, 1500);
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
      <Toaster
        toastOptions={{
          success: {
            iconTheme: {
              primary: "#4bbeee",
              secondary: "#fff",
            },
          },
          // error: {
          //   iconTheme: {
          //     primary: "red",
          //     secondary: "#fff",
          //   },
          // },
        }}
      />
      <form onSubmit={handleSubmit} className="form">
        <div className="form__title">
          <label htmlFor="title">タイトル</label>
          <input type="text" ref={titleRef} />
        </div>
        <div className="form__description">
          <label htmlFor="description">内容</label>
          <textarea ref={descriptionRef} name="description"></textarea>
        </div>
        <div className="form__btn-area">
          <button className="form__btn form__btn--update">
            <span></span>Update
          </button>
          <button
            className="form__btn form__btn--delete"
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
