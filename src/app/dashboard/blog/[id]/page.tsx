import BlogCard from "@/app/components/elements/BlogCard";
import { getAllBlogIds, getBlogById } from "@/app/lib/utils/blog";

export async function generateStaticParams() {
  const ids = await getAllBlogIds();

  return ids.map((id: number) => ({
    //メモ：動的ルーティングではURLパラメータを文字列として扱う必要あり
    id: id.toString(),
  }));
}

export default async function Page(props: { params: Promise<{ id: number }> }) {
  const params = await props.params;
  const { id } = params;
  const post = await getBlogById(Number(id));

  if (!post) {
    return (
      <div className="dark:bg-gray-900 bg-gray-100 h-screen flex items-center justify-center">
        <h1 className="text-2xl font-bold text-red-500">
          記事が見つかりません
        </h1>
      </div>
    );
  }

  return (
    <div className="dark:bg-gray-900 bg-gray-100">
      <main>
        <BlogCard isDetailPage={true} posts={[post]} />
      </main>
    </div>
  );
};
