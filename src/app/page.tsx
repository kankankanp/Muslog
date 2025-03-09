"use client"; // クライアントコンポーネントとして指定

import { useEffect, useState } from "react";
import Header from "./components/layouts/Header";
import Footer from "./components/layouts/Footer";
import { Book } from "./components/elements/Book";

export default function Home() {
  const [posts, setPosts] = useState([]);
  const [loading, setLoading] = useState(true);

  useEffect(() => {
    const fetchBlogs = async () => {
      try {
        const response = await fetch("/api/blog"); // APIルートからデータ取得
        if (!response.ok) {
          throw new Error("Failed to fetch blogs");
        }
        const data = await response.json();
        setPosts(data);
      } catch (error) {
        console.error("Error fetching blogs:", error);
      } finally {
        setLoading(false);
      }
    };

    fetchBlogs();
  }, []);

  return (
    <>
      <Header />
      <main className="dark:bg-gray-900 bg-gray-100 min-h-screen flex items-center justify-center">
        {loading ? (
          <p className="text-white">Loading...</p> // ローディング表示
        ) : (
          <Book posts={posts} />
        )}
      </main>
      <Footer />
    </>
  );
}
