"use client";

import { useParams } from "next/navigation";
import AddButton from "@/app/components/elements/buttons/AddButton";
import BlogCard from "@/app/components/elements/cards/BlogCard";
import Pagination from "@/app/components/elements/others/Pagination";
import { useGetBlogsPagePage } from "@/app/libs/api/generated/orval/blogs/blogs";

export default function Page() {
  const params = useParams();
  const { page } = params as { page: string };
  const pageIndex = Number(page);

  const { data: blogsData, isPending, error } = useGetBlogsPagePage(pageIndex);

  const posts = blogsData?.posts || [];
  const totalCount = blogsData?.totalCount || 0;

  if (isPending) return <div>Loading...</div>;
  if (error) return <div>Error: {(error as any).message}</div>;

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
