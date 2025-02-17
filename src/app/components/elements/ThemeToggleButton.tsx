"use client";

import { useEffect } from "react";
import { useTheme } from "next-themes";
import { FontAwesomeIcon } from "@fortawesome/react-fontawesome";
import { faMoon, faSun } from "@fortawesome/free-solid-svg-icons";
import { RootState } from "@/app/lib/store/store";
import { setTheme } from "@/app/lib/store/themeSlice";
import { useDispatch, useSelector } from "react-redux";

export default function ThemeToggleButton() {
  const dispatch = useDispatch();
  const theme = useSelector((state: RootState) => state.theme.theme);
  const { setTheme: setNextTheme } = useTheme();

  useEffect(() => {
    setNextTheme(theme);
    document.documentElement.classList.toggle("dark", theme === "dark");
  }, [theme, setNextTheme]);

  const toggleTheme = () => {
    const newTheme = theme === "light" ? "dark" : "light";
    dispatch(setTheme(newTheme));
  };

  return (
    <label className="inline-flex items-center cursor-pointer">
      <input
        type="checkbox"
        checked={theme === "dark"}
        onChange={toggleTheme}
        className="sr-only peer"
      />
      <div className="relative w-14 h-7 bg-gray-300 peer-focus:ring-4 peer-focus:ring-blue-300 dark:peer-focus:ring-blue-800 rounded-full peer dark:bg-gray-700 peer-checked:bg-blue-600 dark:peer-checked:bg-blue-600 transition-all">
        <div className="absolute top-0.5 left-[4px] w-6 h-6 bg-white border border-gray-300 rounded-full dark:border-gray-600 transition-all flex items-center justify-center peer-checked:translate-x-7">
          <FontAwesomeIcon
            icon={theme === "dark" ? faMoon : faSun}
            className="w-4 h-4 text-gray-700 dark:text-gray-300"
          />
        </div>
      </div>
    </label>
  );
}
