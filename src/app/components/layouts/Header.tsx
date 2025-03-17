"use client";

import Link from "next/link";
import { useSession, signOut } from "next-auth/react";
import ThemeToggleButton from "../elements/ThemeToggleButton";
import Image from "next/image";
import { useEffect, useState } from "react";

const Header = () => {
  const { data: session, status, update } = useSession();
  const [loading, setLoading] = useState(true);

  useEffect(() => {
    if (status === "loading") return;
    setLoading(false);
  }, [status]);

  useEffect(() => {
    update();
  }, []);

  return (
    <header className="flex flex-col md:flex-row md:justify-around items-center bg-white dark:bg-gray-800 px-4 md:px-8 border-b border-gray-300 dark:border-gray-700 py-4 md:py-0">
      <div className="flex items-center mb-4 md:mb-0">
        <Link href="/" className="flex items-center gap-2">
          <Image
            src="/logo.png"
            alt="BLOG"
            width={80}
            height={80}
            className="md:w-[100px] md:h-[100px]"
            priority
          />
        </Link>
      </div>
      <div className="flex flex-col md:flex-row gap-4 md:gap-[30px] w-full md:w-auto">
        <nav className="flex flex-wrap justify-center md:flex-nowrap items-center gap-4 md:gap-8">
          {[
            { href: "/", label: "ホーム", icon: "/home.png" },
            { href: "/dashboard", label: "管理", icon: "/description.png" },
            {
              href: "/dashboard/blog/page/1",
              label: "記事",
              icon: "/library.png",
            },
          ].map((item) => (
            <Link
              key={item.href}
              href={item.href}
              className="group relative text-gray-700 dark:text-gray-400 text-base md:text-lg font-medium hover:text-blue-600 dark:hover:text-blue-400 transition flex flex-col items-center"
            >
              <Image
                src={item.icon}
                alt=""
                width={30}
                height={30}
                className="md:w-[40px] md:h-[40px] group-hover:-translate-y-1 transition-transform duration-200"
                priority
              />
              <span>{item.label}</span>
              <span className="absolute left-0 bottom-[-4px] w-full h-[2px] bg-blue-600 dark:bg-blue-400 scale-x-0 group-hover:scale-x-100 transition-transform"></span>
            </Link>
          ))}
        </nav>
        <div className="flex flex-wrap justify-center md:flex-nowrap items-center gap-4 md:gap-6">
          {loading ? (
            <p className="text-gray-700 dark:text-gray-200">Loading...</p>
          ) : session ? (
            <button
              onClick={() => signOut()}
              className="group relative text-gray-700 dark:text-gray-400 text-base md:text-lg font-medium hover:text-blue-600 dark:hover:text-blue-400 transition flex flex-col items-center"
            >
              <Image
                src="/logout.png"
                alt=""
                width={30}
                height={30}
                priority
                className="md:w-[40px] md:h-[40px] group-hover:-translate-y-1 transition-transform duration-200 text-center"
              />
              <span>ログアウト</span>
              <span className="absolute left-0 bottom-[-4px] w-full h-[2px] bg-blue-600 dark:bg-blue-400 scale-x-0 group-hover:scale-x-100 transition-transform"></span>
            </button>
          ) : (
            <div className="flex flex-wrap md:flex-nowrap gap-4 items-center">
              <Link
                href="/registration/login"
                className="group relative text-gray-700 dark:text-gray-400 text-base md:text-lg font-medium hover:text-blue-700 dark:hover:text-blue-400 transition flex flex-col items-center"
              >
                <Image
                  src="/login.png"
                  alt=""
                  width={30}
                  height={30}
                  priority
                  className="md:w-[40px] md:h-[40px] group-hover:-translate-y-1 transition-transform duration-200"
                />
                <span>ログイン</span>
                <span className="absolute left-0 bottom-[-4px] w-full h-[2px] bg-blue-600 dark:bg-blue-400 scale-x-0 group-hover:scale-x-100 transition-transform"></span>
              </Link>
              <Link
                href="/registration/signup"
                className="px-3 md:px-4 bg-red-500 text-white font-medium rounded-full hover:bg-red-600 dark:hover:bg-red-500 transition h-10 md:h-12 w-[120px] md:w-[150px] flex items-center justify-center text-sm md:text-base"
              >
                <span>新規登録</span>
              </Link>
            </div>
          )}
          <ThemeToggleButton />
        </div>
      </div>
    </header>
  );
};

export default Header;
