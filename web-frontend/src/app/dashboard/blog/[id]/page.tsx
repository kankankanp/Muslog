import BlogCard from "@/app/components/elements/cards/BlogCard";
import { BlogsService } from "@/app/libs/api/generated";

export async function generateStaticParams() {
  const response = await BlogsService.getBlogs();
  const ids = response.posts?.map(post => post.id) || [];

  return ids.map((id: number) => ({
    id: id.toString(),
  }));
}

export default async function Page(props: { params: { id: string } }) {
  const { id } = props.params;
  const response = await BlogsService.getBlogs1(Number(id));
  const post = response.post;

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
}
