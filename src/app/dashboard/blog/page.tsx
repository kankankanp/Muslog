import AddButton from "@/app/components/elements/AddButton";
import BlogCard from "@/app/components/elements/BlogCard";
import Footer from "@/app/components/layouts/Footer";
import Header from "@/app/components/layouts/Header";
import { fetchAllBlogs } from "@/app/lib/utils";

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
