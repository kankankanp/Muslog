"use client";

import {
  User,
  Home,
  PlusCircle,
  Users,
  HelpCircle,
  LogOut,
} from "lucide-react";
import Link from "next/link";
import { useRouter } from "next/navigation";
import toast from "react-hot-toast";
import { useDispatch } from "react-redux";
import { usePostLogout } from "@/libs/api/generated/orval/auth/auth";
import { logout } from "@/libs/store/authSlice";

const Sidebar = () => {
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
    <aside className="h-screen w-64 bg-white border-r border-gray-200 p-4">
      <nav>
        <ul className="flex flex-col gap-6 mt-8">
          <li className="mb-2">
            <Link
              href="/dashboard"
              className="flex items-center p-2 rounded-lg hover:bg-gray-100"
            >
              <Home className="mr-3 h-5 w-5" />
              <span>ホーム</span>
            </Link>
          </li>
          <li className="mb-2">
            <Link
              href="/dashboard/me"
              className="flex items-center p-2 rounded-lg hover:bg-gray-100"
            >
              <User className="mr-3 h-5 w-5" />
              <span>マイページ</span>
            </Link>
          </li>
          <li className="mb-2">
            <Link
              href="/dashboard/post/add"
              className="flex items-center p-2 rounded-lg hover:bg-gray-100"
            >
              <PlusCircle className="mr-3 h-5 w-5" />
              <span>記事作成</span>
            </Link>
          </li>
          <li className="mb-2">
            <Link
              href="/dashboard/community"
              className="flex items-center p-2 rounded-lg hover:bg-gray-100"
            >
              <Users className="mr-3 h-5 w-5" />
              <span>コミュニティ</span>
            </Link>
          </li>
          <li className="mb-2">
            <Link
              href="/dashboard/help"
              className="flex items-center p-2 rounded-lg hover:bg-gray-100"
            >
              <HelpCircle className="mr-3 h-5 w-5" />
              <span>ヘルプ</span>
            </Link>
          </li>
          <li className="mb-2">
            <Link
              href="/logout"
              className="flex items-center p-2 rounded-lg hover:bg-gray-100"
            >
              <button
                onClick={handleLogout}
                disabled={isPending}
                className="flex"
              >
                <LogOut className="mr-3 h-5 w-5" />
                ログアウト
              </button>
            </Link>
          </li>
        </ul>
      </nav>
    </aside>
  );
};

export default Sidebar;
