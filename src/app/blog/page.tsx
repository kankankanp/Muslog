import Header from "../components/layouts/header/page";
import Footer from "../components/layouts/footer/page";
import BlogCard from "../components/elements/blog-card/page";
import AddBlogModal from "../components/elements/add-blog-modal/page";
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
