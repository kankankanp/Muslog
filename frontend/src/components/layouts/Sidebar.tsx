'use client';

import {
  User,
  Home,
  PlusCircle,
  Users,
  HelpCircle,
  LogOut,
  X,
  Megaphone,
  Settings,
} from 'lucide-react';
import Link from 'next/link';
import { useRouter } from 'next/navigation';
import toast from 'react-hot-toast';
import { useDispatch } from 'react-redux';
import { useSidebar } from '@/contexts/SidebarContext';
import { usePostLogout } from '@/libs/api/generated/orval/auth/auth';
import { logout } from '@/libs/store/authSlice';

const Sidebar = () => {
  const { isSidebarOpen, setIsSidebarOpen } = useSidebar();
  const router = useRouter();
  const dispatch = useDispatch();
  const { mutate: logoutMutation, isPending } = usePostLogout();

  const handleLogout = () => {
    logoutMutation(undefined, {
      onSuccess: () => {
        dispatch(logout());
        toast.success('ログアウトしました');
        router.push('/login-or-signup');
      },
      onError: (error) => {
        console.error('Logout error:', error);
        toast.error('ログアウトに失敗しました。');
      },
    });
  };

  return (
    <aside
      className={`fixed inset-y-0 left-0 z-50 w-64 bg-white border-r border-gray-200 p-4 transform transition-transform duration-300 ease-in-out md:relative md:translate-x-0 ${isSidebarOpen ? 'translate-x-0' : '-translate-x-full'}`}
    >
      <div className="flex justify-end md:hidden">
        <button
          onClick={() => setIsSidebarOpen(false)}
          className="text-gray-600 hover:text-gray-800"
        >
          <X size={24} />
        </button>
      </div>
      <nav>
        <ul className="flex flex-col gap-6 mt-8">
          <li className="mb-2">
            <Link
              href="/dashboard"
              className="flex items-center p-2 rounded-lg hover:bg-gray-100"
              onClick={() => setIsSidebarOpen(false)}
            >
              <Home className="mr-3 h-5 w-5" />
              <span>ホーム</span>
            </Link>
          </li>
          <li className="mb-2">
            <Link
              href="/dashboard/me"
              className="flex items-center p-2 rounded-lg hover:bg-gray-100"
              onClick={() => setIsSidebarOpen(false)}
            >
              <User className="mr-3 h-5 w-5" />
              <span>マイページ</span>
            </Link>
          </li>
          <li className="mb-2">
            <Link
              href="/dashboard/post/add"
              className="flex items-center p-2 rounded-lg hover:bg-gray-100"
              onClick={() => setIsSidebarOpen(false)}
            >
              <PlusCircle className="mr-3 h-5 w-5" />
              <span>記事作成</span>
            </Link>
          </li>
          <li className="mb-2">
            <Link
              href="/dashboard/community"
              className="flex items-center p-2 rounded-lg hover:bg-gray-100"
              onClick={() => setIsSidebarOpen(false)}
            >
              <Users className="mr-3 h-5 w-5" />
              <span>コミュニティ</span>
            </Link>
          </li>
          <li className="mb-2">
            <Link
              href="/dashboard/band-recruitments"
              className="flex items-center p-2 rounded-lg hover:bg-gray-100"
              onClick={() => setIsSidebarOpen(false)}
            >
              <Megaphone className="mr-3 h-5 w-5" />
              <span>バンド募集</span>
            </Link>
          </li>
          <li className="mb-2">
            <Link
              href="/dashboard/settings"
              className="flex items-center p-2 rounded-lg hover:bg-gray-100"
              onClick={() => setIsSidebarOpen(false)}
            >
              <Settings className="mr-3 h-5 w-5" />
              <span>設定</span>
            </Link>
          </li>
          <li className="mb-2">
            <Link
              href="/dashboard/help"
              className="flex items-center p-2 rounded-lg hover:bg-gray-100"
              onClick={() => setIsSidebarOpen(false)}
            >
              <HelpCircle className="mr-3 h-5 w-5" />
              <span>ヘルプ</span>
            </Link>
          </li>
          <li className="mb-2">
            <Link
              href="/login-or-signup"
              className="flex items-center p-2 rounded-lg hover:bg-gray-100"
            >
              <button
                onClick={() => {
                  handleLogout();
                  setIsSidebarOpen(false);
                }}
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
