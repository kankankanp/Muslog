import Header from "../components/layouts/header/page";
import Footer from "../components/layouts/footer/page";
import BlogCard from "../components/elements/blog-card/page";
import AddBlogModal from "../components/elements/add-blog-modal/page";
import Pagination from "../components/elements/pagination/page";
import { countAllBlogs } from "./utils";
import { fetchAllBlogs } from "./utils";

const Blog = async () => {
  const posts = await fetchAllBlogs();
  const count = await countAllBlogs();
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
