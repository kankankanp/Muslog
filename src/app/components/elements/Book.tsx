import { PostType } from "./BlogCard";
import "@/scss/book.scss";

type BookProps = {
  posts: PostType[];
};

export const Book = ({ posts }: BookProps) => {
  const pageCount = 6;
  const pages = Array.from({ length: pageCount }, (_, i) => 99 - i);

  return (
    <div className="book">
      {pages.map((zIndex, idx) => {
        const postIndex = idx - 1;
        const isCover = idx === 0 || idx === pages.length - 1;

        return (
          <label className="book-inner" key={idx}>
            <input className="book-inner__flip" type="checkbox" />
            <span
              className={`book-inner__page z-[${zIndex}] ${
                isCover ? "front-cover" : ""
              }`}
            >
              {isCover ? (
                "表紙"
              ) : posts[postIndex] ? (
                <div>
                  <h2>{posts[postIndex].title}</h2>
                  <p>{posts[postIndex].description}</p>
                </div>
              ) : (
                "None"
              )}
            </span>
            <span
              className={`book-inner__dummy dummy ${
                isCover ? "back-cover" : ""
              }`}
            >
              {isCover ? (
                "表紙"
              ) : posts[postIndex] ? (
                <div>
                  <h2>{posts[postIndex].title}</h2>
                  <p>{posts[postIndex].description}</p>
                </div>
              ) : (
                "None"
              )}
            </span>
          </label>
        );
      })}
    </div>
  );
};
