import AddButton from "../components/elements/AddButton";
import BlogCard from "../components/elements/BlogCard";
import Footer from "../components/layouts/Footer";
import Header from "../components/layouts/Header";
import { fetchAllBlogs } from "../lib/utils";

const Blog = async () => {
  const posts = await fetchAllBlogs();

  return (
    <>
      <Header />
      <main>
        <AddButton />
        <BlogCard isDetailPage={false} posts={posts} />
      </main>
      <Footer />
    </>
  );
};

export default Blog;
