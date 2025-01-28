"use client";

import { useRouter } from "next/navigation";
import { useRef } from "react";
import toast, { Toaster } from "react-hot-toast";
import "@/scss/modal.scss";
import { editBlog, deleteBlog, getBlogById } from "../../../lib/utils";

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
      duration: 1500,
    });

    router.push("/blog/page/1");
    router.refresh();
  };

  const handleDelete = async (e: React.FormEvent) => {
    await deleteBlog(params.id);
  };

  getBlogById(params.id)
    .then((data) => {
      titleRef.current!.value = data.title;
      descriptionRef.current!.value = data.description;
    })
    .catch((error) => {
      toast.error("Error", error);
    });

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
