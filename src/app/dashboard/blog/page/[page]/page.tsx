import { redirect } from "next/navigation";
import AddButton from "@/app/components/elements/AddButton";
import BlogCard from "@/app/components/elements/BlogCard";
import Pagination from "@/app/components/elements/Pagination";
import { auth } from "@/app/lib/auth/auth";
import { fetchBlogsByPage } from "@/app/lib/utils/blog";

export default async function Page(props: {
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
