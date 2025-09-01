"use client";

import {
  faHome,
  faSignOutAlt,
  faUser,
  faPlusSquare,
} from "@fortawesome/free-solid-svg-icons";
import { FontAwesomeIcon } from "@fortawesome/react-fontawesome";
import Link from "next/link";
import { useRouter } from "next/navigation";
import toast from "react-hot-toast";
import { useDispatch } from "react-redux";
import { usePostLogout } from "@/libs/api/generated/orval/auth/auth";
import { logout } from "@/libs/store/authSlice";

export default function Sidebar() {
  const router = useRouter();
  const dispatch = useDispatch();
  const { mutate: logoutMutation, isPending } = usePostLogout();

  const handleLogout = () => {
    logoutMutation(undefined, {
      onSuccess: () => {
        dispatch(logout());
        toast.success("ログアウトしました");
        router.push("/login-or-signup");
      },
      onError: (error) => {
        console.error("Logout error:", error);
        toast.error("ログアウトに失敗しました。");
      },
    });
  };

  return (
    <aside className="fixed top-0 left-0 h-full w-64 flex-shrink-0 bg-gray-800 text-white flex flex-col z-10">
      <div className="p-4 text-2xl font-bold border-b border-gray-700">
        <Link href="/dashboard">Muslog</Link>
      </div>
      <nav className="flex-grow p-4 space-y-2">
        <Link
          href="/dashboard"
          className="flex items-center py-2.5 px-4 rounded transition duration-200 hover:bg-gray-700"
        >
          <FontAwesomeIcon icon={faHome} className="mr-3" />
          ホーム
        </Link>
        <Link
          href="/dashboard/me"
          className="flex items-center py-2.5 px-4 rounded transition duration-200 hover:bg-gray-700"
        >
          <FontAwesomeIcon icon={faUser} className="mr-3" />
          マイページ
        </Link>
        <Link
          href="/dashboard/post/add"
          className="flex items-center py-2.5 px-4 rounded transition duration-200 hover:bg-gray-700"
        >
          <FontAwesomeIcon icon={faPlusSquare} className="mr-3" />
          新規作成
        </Link>
      </nav>
      <div className="p-4 border-t border-gray-700">
        <button
          onClick={handleLogout}
          disabled={isPending}
          className="w-full flex items-center justify-center py-2.5 px-4 rounded transition duration-200 bg-red-600 hover:bg-red-700 disabled:bg-red-400"
        >
          <FontAwesomeIcon icon={faSignOutAlt} className="mr-3" />
          ログアウト
        </button>
      </div>
    </aside>
  );
}
