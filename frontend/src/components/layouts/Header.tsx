// components/Header.tsx
import { Search } from "lucide-react";
import Image from "next/image";

const Header = () => {
  return (
    <header className="bg-gray-800 text-white py-2 px-4 flex items-center justify-between">
      <div className="flex items-center">
        <Image src="/logo.png" width={50} height={50} alt="Muslog" />
      </div>
      <div className="relative flex-grow mx-4 max-w-lg">
        <input
          type="text"
          placeholder="検索"
          className="w-full pl-4 pr-10 py-2 rounded-full bg-gray-700 text-white placeholder-gray-400 focus:outline-none focus:ring-2 focus:ring-purple-500"
        />
        <Search className="absolute right-3 top-1/2 -translate-y-1/2 h-5 w-5 text-gray-400" />
      </div>
    </header>
  );
};

export default Header;
