"use client";

import PostBlog from "@/app/blog/add/page";
import "@/scss/blog-card.scss";
import Modal from "react-modal";
import { useRouter } from "next/navigation";
import { useEffect, useState } from "react";

const AddBlogModal = () => {
  const router = useRouter();
  const [isModalOpen, setIsModalOpen] = useState(false);

  useEffect(() => {
    const params = new URLSearchParams(window.location.search);
    if (params.get("modal") === "add-blog") {
      setIsModalOpen(true);
    } else {
      setIsModalOpen(false);
    }
  }, []);

  const openModal = () => {
    router.push("/blog?modal=add-blog");
  };

  const closeModal = () => {
    router.push("/blog");
  };

  return (
    <div>
      <button onClick={openModal} className="blog__add"></button>

      <Modal
        isOpen={isModalOpen}
        onRequestClose={closeModal}
        contentLabel="Add Blog Modal"
      >
        <PostBlog />
      </Modal>
    </div>
  );
};

export default AddBlogModal;
