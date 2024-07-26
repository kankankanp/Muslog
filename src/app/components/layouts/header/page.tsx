import Link from "next/link";
import styles from "@/scss/layout.module.scss";

const Header = () => {
  return (
    <header className={styles.header}>
      <h1 className={styles.header__title}>
        <a href="#">MyBlog</a>
      </h1>
      <nav className={styles.header__nav}>
        <ul className={styles.header__items}>
          <li className={styles.header__item}>
            <Link href="/">Home</Link>
          </li>
          <li className={styles.header__item}>
            <Link href="/about">About</Link>
          </li>
          <li className={styles.header__item}>
            <Link href="/blog">Blog</Link>
          </li>
        </ul>
      </nav>
    </header>
  );
};

export default Header;
