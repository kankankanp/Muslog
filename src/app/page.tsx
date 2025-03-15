'use client';

import Footer from "./components/layouts/Footer";
import { Book } from "./components/elements/Book";

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
      <main className="dark:bg-gray-900 bg-gray-100 min-h-screen flex items-center justify-center">
        <Book posts={[]} />
      </main>
      <Footer />
    </>
  );
}
