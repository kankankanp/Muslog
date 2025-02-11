import BlogCard from "@/app/components/elements/BlogCard";
import Pagination from "@/app/components/elements/Pagination";
import Footer from "@/app/components/layouts/Footer";
import Header from "@/app/components/layouts/Header";
import AddButton from "@/app/components/elements/AddButton";
import { fetchBlogsByPage } from "@/app/lib/utils";

export default async function Index(props: {
  params: Promise<{ page: number }>;
}) {
  const params = await props.params;
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
