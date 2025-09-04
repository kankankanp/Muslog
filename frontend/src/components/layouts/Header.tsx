"use client";

import { Menu } from "lucide-react";
import Image from "next/image";
import Link from "next/link";
import { useSidebar } from "@/contexts/SidebarContext";

const Header = () => {
  const { setIsSidebarOpen } = useSidebar();

  return (
    <header className="bg-gray-800 text-white px-4 flex items-center justify-between">
      <div className="flex items-center gap-4">
        <button
          className="md:hidden text-white mr-4"
          onClick={() => setIsSidebarOpen(true)}
        >
          <Menu size={24} />
        </button>
        <Link href="/dashboard" className="flex items-center gap-2">
          <Image src="/logo.png" width={60} height={60} alt="Muslog" />
          <h1 className="text-2xl">Muslog</h1>
        </Link>
      </div>
      {/* <div className="relative flex-grow mx-4 max-w-lg">
        <input
          type="text"
          placeholder="検索"
          className="w-full pl-4 pr-10 py-2 rounded-full bg-gray-700 text-white placeholder-gray-400 focus:outline-none focus:ring-2 focus:ring-purple-500"
        />
        <Search className="absolute right-3 top-1/2 -translate-y-1/2 h-5 w-5 text-gray-400" />
      </div> */}
    </header>
  );
};

export default Header;
