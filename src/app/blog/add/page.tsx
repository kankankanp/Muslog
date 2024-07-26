'use client';

import { useRouter } from "next/navigation";
import { useRef } from "react";
import toast, { Toaster } from "react-hot-toast";

const postBlog = async ( 
  title: string | undefined,
  description: string | undefined,
) => {
  const res = await fetch('http://localhost:3000/api/blog', {
    method: 'POST',
    body: JSON.stringify({title, description}),
    headers: {
      'Content-Type': 'application/json'
    }
  })

  return res.json();
}

const PostBlog = () => {
  const router = useRouter();

  const titleRef = useRef<HTMLInputElement | null>(null);
  const descriptionRef = useRef<HTMLTextAreaElement | null>(null);

  const handleSubmit = async (e: React.FormEvent) => {
    e.preventDefault();

    toast.loading('Posting...')

    await postBlog(titleRef.current?.value, descriptionRef.current?.value)

    
    
    router.push('/blog');
    router.refresh();
    
    toast.success('Posted!', {
      duration: 2000,
    })
  }

  return (
    <>
      <Toaster />
      <form onSubmit={handleSubmit}>
        <input
          type="text"
          ref={titleRef}
          placeholder="タイトルを入力"
        />
        <textarea 
          ref={descriptionRef}
          name="" 
          placeholder="記事を入力"
          id=""
        ></textarea>
        <button>投稿</button>
      </form>
    </>
  )
}

export default PostBlog;