import Link from "next/link";
import { faPen } from "@fortawesome/free-solid-svg-icons";
import { FontAwesomeIcon } from "@fortawesome/react-fontawesome";
import { CommonButton } from "./CommonButton";

type PostType = {
  id: number;
  title: string;
  description: string;
  date: Date;
};

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
      className={`w-full max-w-7xl mx-auto px-4 sm:px-6 lg:px-8 grid ${
        !isDetailPage
          ? "grid-cols-1 sm:grid-cols-2 lg:grid-cols-3 gap-6"
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
            className="p-4 sm:p-6 bg-white dark:bg-gray-800 shadow-md rounded-lg"
          >
            <h3 className="text-xl sm:text-2xl font-bold text-gray-900 dark:text-gray-100">{`${year}/${month}/${day}(${dayOfWeek})`}</h3>
            <h3 className="text-lg sm:text-xl font-semibold mt-2 text-gray-900 dark:text-gray-100">
              {post.title}
            </h3>
            <p className="text-base sm:text-lg mt-2 text-gray-700 dark:text-gray-300">
              {isDetailPage
                ? post.description
                : truncateText(post.description, 40)}
            </p>
            <div className="flex justify-between items-center mt-4">
              {!isDetailPage && (
                <Link
                  href={`/dashboard/blog/edit/${post.id}`}
                  className="text-black dark:text-gray-300 text-base sm:text-lg transition-transform transform hover:scale-125"
                >
                  <FontAwesomeIcon icon={faPen} />
                </Link>
              )}
              {!isDetailPage && (
                <CommonButton 
                  href={`/dashboard/blog/${post.id}`}
                  className="p-2"
                >
                  Show more
                </CommonButton>
              )}
            </div>
          </div>
        );
      })}
    </div>
  );
};

export default BlogCard;
