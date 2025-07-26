"use client";

import { useRouter } from "next/navigation";
import AddButton from "@/app/components/elements/buttons/AddButton";
import BlogCard from "@/app/components/elements/cards/BlogCard";
import Pagination from "@/app/components/elements/others/Pagination";
import { BlogsService } from "@/app/libs/api/generated";
import { useSelector } from "react-redux";
import { RootState } from "@/app/libs/store/store";
import { useEffect } from "react";
import { useGetBlogsByPage } from "@/app/libs/hooks/api/useBlogs";

export default function Page({ params }: { params: { page: number } }) {
  const session = useSelector((state: RootState) => state.auth);
  const router = useRouter();

  useEffect(() => {
    if (!session?.accessToken) {
      router.push("/registration/login");
    }
  }, [session, router]);

  const pageIndex = params.page;

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
