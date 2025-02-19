import Link from "next/link";
import { auth, signOut } from "@/app/lib/auth/auth";
import ThemeToggleButton from "../elements/ThemeToggleButton";
import Image from "next/image";

const Header = async () => {
  const session = await auth();

  return (
    <header className="flex justify-around items-center bg-white px-8 border-b border-gray-300">
      <div className="flex items-center">
        <Link href="/" className="flex items-center gap-2">
          <Image src="/logo.png" alt="BLOG" width={100} height={100} priority />
          {/* <span className="absolute bottom-[-15px] left-[15px] text-gray-800 font-bold text-2xl">BLOG</span> */}
        </Link>
      </div>

      <div className="flex gap-[30px]">
        <nav className="flex items-center gap-8">
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
              className="group relative text-gray-700 text-lg font-medium hover:text-blue-600 transition flex flex-col items-center"
            >
              <Image
                src={item.icon}
                alt=""
                width={40}
                height={40}
                className="group-hover:-translate-y-1 transition-transform duration-200"
                priority
              />
              <span>{item.label}</span>
              <span className="absolute left-0 bottom-[-4px] w-full h-[2px] bg-blue-600 scale-x-0 group-hover:scale-x-100 transition-transform"></span>
            </Link>
          ))}
        </nav>
        <div className="flex items-center gap-6">
          {session?.user ? (
            <form
              action={async () => {
                "use server";
                await signOut();
              }}
            >
              <button className="group relative text-gray-700 text-lg font-medium hover:text-blue-600 transition flex flex-col items-center">
                <Image
                  src="/logout.png"
                  alt=""
                  width={40}
                  height={40}
                  priority
                  className="group-hover:-translate-y-1 transition-transform duration-200 text-center"
                />
                <span>ログアウト</span>
                <span className="absolute left-0 bottom-[-4px] w-full h-[2px] bg-blue-600 scale-x-0 group-hover:scale-x-100 transition-transform"></span>
              </button>
            </form>
          ) : (
            <div className="flex gap-4 items-center">
              <Link
                href="/registration/login"
                className="group relative text-blue-600 text-lg font-medium hover:text-blue-700 transition flex flex-col items-center"
              >
                <Image
                  src="/login.png"
                  alt=""
                  width={40}
                  height={40}
                  priority
                  className="group-hover:-translate-y-1 transition-transform duration-200"
                />
                <span>ログイン</span>
                <span className="absolute left-0 bottom-[-4px] w-full h-[2px] bg-blue-600 scale-x-0 group-hover:scale-x-100 transition-transform"></span>
              </Link>
              <Link
                href="/registration/signup"
                className="px-4 bg-red-500 text-white font-medium rounded-full hover:bg-black-600 transition h-12 w-[150px] flex items-center justify-center"
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
