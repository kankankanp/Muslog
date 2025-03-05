import "./OtherBook.css";

export const OtherBook = () => {
  const pageCount = 6; // 繰り返す回数
  const pages = Array.from({ length: pageCount }, (_, i) => 99 - i);

  return (
    <div className="cssbk">
      {pages.map((zIndex, idx) => (
        <label className="cssbk-inner" key={idx}>
          <input className="cssbk-inner__flip" type="checkbox" />
          <span
            className={`cssbk-inner__page z-[${zIndex}] ${
              zIndex <= 98 ? "content-page" : ""
            }`}
          >
            中の文章
          </span>
          <span
            className={`cssbk-inner__dummy dummy ${
              zIndex <= 98 ? "content-page" : ""
            }`}
          >
            中の文章
          </span>
        </label>
      ))}
    </div>
  );
};
