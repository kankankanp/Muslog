import Link from "next/link";
import Image from "next/image";

const Footer = () => {
  return (
    <footer className="bg-gray-800 text-gray-300 py-10">
      <div className="container mx-auto px-6 flex flex-col md:flex-row justify-between items-center">
        <div className="flex flex-col items-center md:items-start">
          <Link href="/" className="flex items-center gap-2 mb-3">
            <Image
              src="/logo.png"
              alt="BLOG"
              width={100}
              height={100}
              priority
            />
          </Link>
          <p className="text-sm text-gray-400">
            Simple Blog
          </p>
        </div>

        <nav className="mt-6 md:mt-0 flex flex-wrap justify-center gap-6 text-sm">
          {[
            { href: "/", label: "ホーム" },
            { href: "/dashboard", label: "管理" },
            { href: "/dashboard/blog/page/1", label: "記事" },
            { href: "/contact", label: "お問い合わせ" },
          ].map((item) => (
            <Link
              key={item.href}
              href={item.href}
              className="hover:text-blue-400 transition"
            >
              {item.label}
            </Link>
          ))}
        </nav>

        {/* <div className="mt-6 md:mt-0 flex gap-4">
          <Link href="https://twitter.com" target="_blank">
            <Image
              src="/twitter.png"
              alt="Twitter"
              width={30}
              height={30}
              className="hover:opacity-70 transition"
            />
          </Link>
          <Link href="https://facebook.com" target="_blank">
            <Image
              src="/facebook.png"
              alt="Facebook"
              width={30}
              height={30}
              className="hover:opacity-70 transition"
            />
          </Link>
          <Link href="https://instagram.com" target="_blank">
            <Image
              src="/instagram.png"
              alt="Instagram"
              width={30}
              height={30}
              className="hover:opacity-70 transition"
            />
          </Link>
        </div> */}
      </div>

      <div className="border-t border-gray-700 mt-6 pt-6 text-center text-gray-400 text-xs">
        © {new Date().getFullYear()} BLOG. All rights reserved.
      </div>
    </footer>
  );
};

export default Footer;
