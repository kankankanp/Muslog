import Header from "../components/layouts/header/page";
import Footer from "../components/layouts/footer/page";
import BlogCard from "../components/elements/blog-card/page";
import AddBlogModal from "../components/elements/add-blog-modal/page";
import Pagination from "../components/elements/pagination/page";
import { countData } from "../api/blog/route";

export const fetchAllBlogs = async () => {
  const res = await fetch("http://localhost:3000/api/blog", {
    cache: "no-store",
  });

  const data = await res.json();

  return data.posts;
};

const Blog = async () => {
  const posts = await fetchAllBlogs();
  const [count] = await Promise.all([countData]);
  const PER_PAGE = 4;
  const postPerPage = posts.slice(0, PER_PAGE);

  return (
    <>
      <Header />
      <main>
        <AddBlogModal />
        <BlogCard isDetailPage={false} posts={postPerPage} />
      </main>
      <Pagination totalCount={count} />
      <Footer />
    </>
  );
};

export default Blog;
