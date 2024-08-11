import Link from "next/link";
import styles from "@/scss/layout.module.scss";

const AddButton = () => {
  return (
    <Link href="/blog/add" className={styles.blog__add}>
      <span></span>
    </Link>
  );
};

export default AddButton;
