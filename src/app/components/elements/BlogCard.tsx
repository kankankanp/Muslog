import { PostType } from "@/app/type/PostType";
import Link from "next/link";
import { faPen } from "@fortawesome/free-solid-svg-icons";
import { FontAwesomeIcon } from "@fortawesome/react-fontawesome";

type BlogCardProps = {
  isDetailPage?: boolean;
  posts: PostType[];
};

const BlogCard = ({ isDetailPage, posts }: BlogCardProps) => {
  const safePosts = Array.isArray(posts) ? posts : [];
  const truncateText = (text: string, length: number) => {
    if (text.length <= length) {
      return text;
    }
    return text.substring(0, length) + "...";
  };

  return (
    <div
      className={`max-w-md grid ${
        !isDetailPage
          ? "grid-cols-2 gap-4 justify-center md:grid-cols-1"
          : "grid-cols-1"
      }`}
    >
      {safePosts.map((post: PostType) => {
        const date = new Date(post.date);
        const year = date.getFullYear();
        const month = date.getMonth() + 1;
        const day = date.getDate();
        const daysOfWeek = ["Sun", "Mon", "Tue", "Wed", "Thu", "Fri", "Sat"];
        const dayOfWeek = daysOfWeek[date.getDay()];

        return (
          <div
            key={post.id}
            className="p-6 bg-white dark:bg-gray-800 shadow-md rounded-lg"
          >
            <h3 className="text-2xl font-bold text-gray-900 dark:text-gray-100">{`${year}/${month}/${day}(${dayOfWeek})`}</h3>
            <h3 className="text-xl font-semibold mt-2 text-gray-900 dark:text-gray-100">
              {post.title}
            </h3>
            <p className="text-lg mt-2 text-gray-700 dark:text-gray-300">
              {isDetailPage
                ? post.description
                : truncateText(post.description, 40)}
            </p>
            <div className="flex justify-between items-center mt-4">
              {!isDetailPage && (
                <Link
                  href={`/dashboard/blog/edit/${post.id}`}
                  className="text-black dark:text-gray-300 text-lg transition-transform transform hover:scale-125"
                >
                  <FontAwesomeIcon icon={faPen} />
                </Link>
              )}
              {!isDetailPage && (
                <Link
                  href={`/dashboard/blog/${post.id}`}
                  className="px-4 py-2 bg-blue-500 dark:bg-blue-600 text-white rounded-md hover:bg-blue-600 dark:hover:bg-blue-700 transition"
                >
                  Show more
                </Link>
              )}
            </div>
          </div>
        );
      })}
    </div>
  );
};

export default BlogCard;
