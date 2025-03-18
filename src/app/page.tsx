"use client";

import { useState, useEffect } from "react";
import { Book } from "./components/elements/Book";
import { fetchAllBlogs } from "./lib/utils/blog";

export default function Home() {
  const [posts, setPosts] = useState([]);

  useEffect(() => {
    const getBlogs = async () => {
      const blogs = await fetchAllBlogs();
      setPosts(blogs);
    };
    getBlogs();
  }, []);
  return (
    <>
      <main className="dark:bg-gray-900 bg-gray-100 min-h-screen flex items-center justify-center">
        <Book posts={posts} />
      </main>
    </>
  );
}
