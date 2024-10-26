import Header from "../components/layouts/Header";
import Footer from "../components/layouts/Footer";
import BlogCard from "../components/elements/BlogCard";
import AddBlogModal from "../components/elements/AddBlogModal";
import { fetchAllBlogs } from "../lib/utils";

const Blog = async () => {
  const posts = await fetchAllBlogs();

  return (
    <>
      <Header />
      <main>
        <AddBlogModal />
        <BlogCard isDetailPage={false} posts={posts} />
      </main>
      <Footer />
    </>
  );
};

export default Blog;
