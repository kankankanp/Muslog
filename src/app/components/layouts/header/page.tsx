import Link from "next/link";
import "@/scss/layout.scss";

const Header = () => {
  return (
    <header className="header">
      <h1 className="header__title">
        <a href="#">MyBlog</a>
      </h1>
      <nav className="header__nav">
        <ul className="header__items">
          <li className="header__item">
            <Link href="/about">About</Link>
          </li>
          <li className="header__item">
            <Link href="/blog/page/1">Blog</Link>
          </li>
        </ul>
      </nav>
    </header>
  );
};

export default Header;
