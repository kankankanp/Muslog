'use client';

import Footer from "./components/layouts/Footer";
import { fetchAllBlogs } from "./lib/utils/blog";
import { Book } from "./components/elements/Book";
import { useState, useEffect } from "react";
import Header from "./components/layouts/Header";

export default function Home() {
  // const [posts, setPosts] = useState([]);

  // useEffect(() => {
  //   const getBlogs = async () => {
  //     const blogs = await fetchAllBlogs();
  //     setPosts(blogs);
  //   };
  //   getBlogs();
  // }, []);
  return (
    <>
      <Header />
      <main className="dark:bg-gray-900 bg-gray-100 min-h-screen flex items-center justify-center">
        <Book posts={[]} />
      </main>
      <Footer />
    </>
  );
}
