"use client";

import { useParams } from "next/navigation";
import AddButton from "@/components/elements/buttons/AddButton";
import BlogCard from "@/components/elements/cards/BlogCard";
import Pagination from "@/components/elements/others/Pagination";
import { useGetPostsPagePage } from "@/libs/api/generated/orval/posts/posts";

export default function Page() {
  const params = useParams();
  const { page } = params as { page: string };
  const pageIndex = Number(page);

  const { data: postsData, isPending, error } = useGetPostsPagePage(pageIndex);
  console.log(postsData);

  const posts = postsData?.posts || [];
  const totalCount = postsData?.totalCount || 0;

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
