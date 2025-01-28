import BlogCard from "@/app/components/elements/BlogCard";
import Pagination from "@/app/components/elements/Pagination";
import Footer from "@/app/components/layouts/Footer";
import Header from "@/app/components/layouts/Header";
import { fetchBlogsByPage } from "../../../lib/utils";
import AddButton from "@/app/components/elements/AddButton";

export default async function Index({ params }: { params: { page: number } }) {
  const pageIndex = params.page;
  const { posts, totalCount } = await fetchBlogsByPage(pageIndex);

  return (
    <>
      <Header />
      <main>
        <AddButton />
        <BlogCard isDetailPage={false} posts={posts} />
      </main>
      <Pagination totalCount={totalCount} />
      <Footer />
    </>
  );
}
