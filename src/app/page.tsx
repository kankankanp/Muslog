import Header from "./components/layouts/Header";
import Footer from "./components/layouts/Footer";
import DemoBook from "./components/elements/Book";
import { fetchAllBlogs } from "./lib/utils/blog";
import Book from "./components/elements/Book";

export default async function Home() {
  const posts = await fetchAllBlogs();

  return (
    <>
      <Header />
      <main className="dark:bg-gray-900 bg-gray-100 min-h-screen flex items-center justify-center">
        <Book posts={posts} />
      </main>
      <Footer />
    </>
  );
}
