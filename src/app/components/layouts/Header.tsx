import Link from "next/link";
import "@/scss/layout.scss";
import { FontAwesomeIcon } from "@fortawesome/react-fontawesome";
import { faHouse } from "@fortawesome/free-solid-svg-icons";
import { faBookOpen } from "@fortawesome/free-solid-svg-icons";

const Header = () => {
  return (
    <header className="header">
      <h1 className="header__title">
        <a href="#">MyBlog</a>
      </h1>
      <nav className="header__nav">
        <ul className="header__items">
          <li className="header__item">
            <Link href="/">
              <FontAwesomeIcon icon={faHouse} />
            </Link>
          </li>
          <li className="header__item">
            <Link href="/blog">
              <FontAwesomeIcon icon={faBookOpen} />
            </Link>
          </li>
        </ul>
      </nav>
    </header>
  );
};

export default Header;
