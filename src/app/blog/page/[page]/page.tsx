import Pagination from "@/app/components/elements/Pagination";
import Header from "@/app/components/layouts/Header";
import AddBlogModal from "@/app/components/elements/AddBlogModal";
import BlogCard from "@/app/components/elements/BlogCard";
import Footer from "@/app/components/layouts/Footer";
import { fetchBlogsByPage } from "../../../lib/utils";

export default async function Index({ params }: { params: { page: number } }) {
  const pageIndex = params.page;
  const { posts, totalCount } = await fetchBlogsByPage(pageIndex);

  return (
    <>
      <Header />
      <main>
        <AddBlogModal />
        <BlogCard isDetailPage={false} posts={posts} />
      </main>
      <Pagination totalCount={totalCount} />
      <Footer />
    </>
  );
}
