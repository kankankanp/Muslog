import Link from "next/link";
import { signOut } from "@/auth";
import "@/scss/layout.scss";
import { FontAwesomeIcon } from "@fortawesome/react-fontawesome";
import { faHouse, faShareFromSquare } from "@fortawesome/free-solid-svg-icons";
import { faBookOpen } from "@fortawesome/free-solid-svg-icons";

const Header = () => {
  return (
    <header className="header">
      <h1 className="header__title">
        <a href="/">MyBlog</a>
      </h1>
      <nav className="header__nav">
        <ul className="header__items">
          <li className="header__item">
            <Link href="/">
              <FontAwesomeIcon icon={faHouse} />
            </Link>
          </li>
          <li className="header__item">
            <Link href="/dashboard/blog/page/1">
              <FontAwesomeIcon icon={faBookOpen} />
            </Link>
          </li>
          <li className="header__item">
            <form
              action={async () => {
                "use server";
                await signOut();
              }}
            >
              <button>
                <FontAwesomeIcon icon={faShareFromSquare} />
              </button>
            </form>
          </li>
        </ul>
      </nav>
    </header>
  );
};

export default Header;
