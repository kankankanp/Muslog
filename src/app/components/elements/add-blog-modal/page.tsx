"use client";

import PostBlog from "@/app/blog/add/page";
import styles from "@/scss/blog.module.scss";

import { useRouter, useSearchParams } from "next/navigation";
import { useEffect, useState } from "react";
import Modal from "react-modal";

const AddBlogModal = () => {
  const router = useRouter();
  const searchParams = useSearchParams();
  const [isModalOpen, setIsModalOpen] = useState(false);

  useEffect(() => {
    if (searchParams.get("modal") === "add-blog") {
      setIsModalOpen(true);
    } else {
      setIsModalOpen(false);
    }
  }, [searchParams]);

  const openModal = () => {
    router.push("/blog?modal=add-blog");
  };

  const closeModal = () => {
    router.push("/blog");
  };

  return (
    <div>
      <button onClick={openModal} className={styles.blog__add}></button>

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
