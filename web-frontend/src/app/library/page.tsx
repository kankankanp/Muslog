import "@/scss/book.scss";
import { fetchAllBlogs } from "@/app/lib/utils/blog";
import { Book } from "../components/elements/others/Book";

export default async function Page() {
  const posts = await fetchAllBlogs();

  return (
    <div className="dark:bg-gray-900 bg-gray-100 min-h-screen py-8">
      {posts.length > 0 ? (
        <Book posts={posts} />
      ) : (
        <p className="text-center text-gray-600 dark:text-gray-400">
          まだブログ記事がありません。
        </p>
      )}
    </div>
  );
}
