import BlogCard from "@/app/components/elements/BlogCard";
import Pagination from "@/app/components/elements/Pagination";
import Footer from "@/app/components/layouts/Footer";
import Header from "@/app/components/layouts/Header";
import AddButton from "@/app/components/elements/AddButton";
import { fetchBlogsByPage } from "@/app/lib/utils/blog";

export default async function Index(props: {
  params: Promise<{ page: number }>;
}) {
  const params = await props.params;
  const pageIndex = params.page;
  const { posts, totalCount } = await fetchBlogsByPage(pageIndex);
  console.log(posts);

  return (
    <>
      <Header />
      <main className="dark:bg-gray-900 bg-gray-100">
        <AddButton />
        <BlogCard isDetailPage={false} posts={posts} />
        <Pagination totalCount={totalCount} />
      </main>
      <Footer />
    </>
  );
}
