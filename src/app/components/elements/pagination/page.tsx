"use client";
import Link from "next/link";
import { useParams } from "next/navigation";
import "@/scss/pagination.scss";

// type Props = {
//   totalCount: number;
// };

const range = (start: number, end: number) =>
  Array.from({ length: end - start + 1 }, (_, i) => start + i);

const PER_PAGE = 4;

const Pagination = ({ totalCount }: any) => {
  const params = useParams();
  const currentPage = params?.page ? Number(params.page) : 1;
  console.log(Number(totalCount));
  return (
    <ul className="pagination">
      {range(1, Math.ceil(Number(totalCount) / PER_PAGE)).map(
        (number, index) => (
          <li
            className={`pagination__item ${
              currentPage === number ? "is-active" : ""
            }`}
            key={index}
          >
            <Link
              className="pagination__item-link"
              href={`/blog/page/${number}`}
              aria-current={currentPage === number ? "page" : undefined}
            >
              {number}
            </Link>
          </li>
        )
      )}
    </ul>
  );
};

export default Pagination;
