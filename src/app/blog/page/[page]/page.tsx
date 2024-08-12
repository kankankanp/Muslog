import Pagination from "@/app/components/elements/pagination/page";
import { fetchAllBlogs } from "../../utils";
import { countAllBlogs } from "../../utils";
import Header from "@/app/components/layouts/header/page";
import AddBlogModal from "@/app/components/elements/add-blog-modal/page";
import BlogCard from "@/app/components/elements/blog-card/page";
import Footer from "@/app/components/layouts/footer/page";

export const generateStaticParams = async () => {
  const count = await countAllBlogs();
  const range = (start: number, end: number) =>
    [...Array(end - start + 1)].map((_, i) => start + i);

  const paths = range(2, Math.ceil(count / 2)).map((num) => ({
    page: `${num}`,
  }));
  return paths;
};

export default async function Index({ params }: { params: { page: number } }) {
  const count = await countAllBlogs();
  const posts = await fetchAllBlogs();
  const PER_PAGE = 4;
  const pageIndex = params.page;
  const postPerPage = posts.slice(
    PER_PAGE * pageIndex - PER_PAGE,
    PER_PAGE * pageIndex
  );

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
}
