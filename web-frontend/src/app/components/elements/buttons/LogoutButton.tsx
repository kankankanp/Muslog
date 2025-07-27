"use client";

import Image from "next/image";
import { useRouter } from "next/navigation";
import toast from "react-hot-toast";
import { useDispatch } from "react-redux";
import { useLogout } from "@/app/libs/hooks/api/useAuth";
import { logout } from "@/app/libs/store/authSlice";
import { AppDispatch } from "@/app/libs/store/store";

const LogoutButton = () => {
  const dispatch: AppDispatch = useDispatch();
  const router = useRouter();
  const { mutate: logoutMutation } = useLogout();

  const handleLogout = () => {
    logoutMutation(undefined, {
      onSuccess: () => {
        dispatch(logout());
        toast.success("ログアウトしました。");
        router.push("/login");
      },
      onError: (error) => {
        toast.error("ログアウトに失敗しました。");
        console.error("Logout failed:", error);
      },
    });
  };

  return (
    <button
      onClick={handleLogout}
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
  );
};

export default LogoutButton;
