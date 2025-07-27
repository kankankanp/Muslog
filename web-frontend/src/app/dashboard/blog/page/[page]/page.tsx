"use client";

import { useRouter } from "next/navigation";
import { useParams } from "next/navigation";
import { useEffect } from "react";
import { useSelector } from "react-redux";
import AddButton from "@/app/components/elements/buttons/AddButton";
import BlogCard from "@/app/components/elements/cards/BlogCard";
import Pagination from "@/app/components/elements/others/Pagination";
import { useGetBlogsByPage } from "@/app/libs/hooks/api/useBlogs";
import { RootState } from "@/app/libs/store/store";

export default function Page() {
  const session = useSelector((state: RootState) => state.auth);
  const router = useRouter();
  const params = useParams();
  const { page } = params as { page: string };
  const pageIndex = Number(page);

  useEffect(() => {
    if (!session?.user) {
      router.push("/registration/login");
    }
  }, [session, router]);

  const { data: blogsData, isPending, error } = useGetBlogsByPage(pageIndex);

  const posts = blogsData?.posts || [];
  const totalCount = blogsData?.totalCount || 0;

  if (isPending) return <div>Loading...</div>;
  if (error) return <div>Error: {error.message}</div>;

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
