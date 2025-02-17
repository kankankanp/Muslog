"use client";

import { useEffect } from "react";
import { useSelector, useDispatch } from "react-redux";
import { useTheme } from "next-themes";
import { RootState } from "@/app/lib/store/store";
import { setTheme } from "@/app/lib/store/themeSlice";
import { faMoon, faSun } from "@fortawesome/free-solid-svg-icons";
import { FontAwesomeIcon } from "@fortawesome/react-fontawesome";

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
    <button
      onClick={toggleTheme}
      className="p-2 bg-gray-200 dark:bg-gray-800 text-gray-900 dark:text-gray-100 rounded-md"
    >
      <FontAwesomeIcon icon={theme === 'light' ? faMoon : faSun} className="w-5 h-5" />
    </button>
  );
}
