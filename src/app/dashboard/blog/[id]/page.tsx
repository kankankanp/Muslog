import BlogCard from "@/app/components/elements/BlogCard";
import Footer from "@/app/components/layouts/Footer";
import Header from "@/app/components/layouts/Header";
import { getAllBlogIds, getBlogById } from "@/app/lib/utils/blog";

export async function generateStaticParams() {
  const ids = await getAllBlogIds();

  return ids.map((id: number) => ({
    //メモ：動的ルーティングではURLパラメータを文字列として扱う必要あり
    id: id.toString(),
  }));
}

const ShowBlogDetails = async (props: { params: Promise<{ id: number }> }) => {
  const params = await props.params;
  const { id } = params;

  const post = await getBlogById(Number(id));
  const postarray: any = [post];

  return (
    <div className="dark:bg-gray-900 bg-gray-100">
      <Header />
      <main>
        <BlogCard isDetailPage={true} posts={postarray} />
      </main>
      <Footer />
    </div>
  );
};

export default ShowBlogDetails;
