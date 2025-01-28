import Link from "next/link";

const AddButton = () => {
  return (
    <Link href="/blog/add" className="blog__add">
      <span></span>
    </Link>
  );
};

export default AddButton;
