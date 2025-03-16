import BlogCard from "@/app/components/elements/BlogCard";
import Pagination from "@/app/components/elements/Pagination";
import Footer from "@/app/components/layouts/Footer";
import Header from "@/app/components/layouts/Header";
import AddButton from "@/app/components/elements/AddButton";
import { fetchBlogsByPage } from "@/app/lib/utils/blog";
import { auth } from "@/app/lib/auth/auth";
import { redirect } from "next/navigation";

export default async function Index(props: {
  params: Promise<{ page: number }>;
}) {
  const session = await auth();
  if (!session?.user?.email) {
    redirect("/registration/login");
  }

  const params = await props.params;
  const pageIndex = params.page;
  const { posts, totalCount } = await fetchBlogsByPage(pageIndex);

  return (
    <>
      <main className="dark:bg-gray-900 bg-gray-100 pt-8">
        <AddButton />
        <BlogCard isDetailPage={false} posts={posts} />
        <Pagination totalCount={totalCount} />
      </main>
    </>
  );
}
