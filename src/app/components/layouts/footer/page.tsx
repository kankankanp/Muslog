import Image from "next/image";
import "@/scss/layout.scss";

const Footer = () => {
  return (
    <footer className="footer">
      <p>Provided by</p>
      <Image
        src="/vercel.svg"
        alt="Vercel Logo"
        className="dark:invert"
        width={150}
        height={36}
        priority
      />
    </footer>
  );
};

export default Footer;
