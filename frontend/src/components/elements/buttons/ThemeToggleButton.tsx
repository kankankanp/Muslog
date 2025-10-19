'use client';

import { faMoon, faSun } from '@fortawesome/free-solid-svg-icons';
import { FontAwesomeIcon } from '@fortawesome/react-fontawesome';
import { useTheme } from 'next-themes';

export default function ThemeToggleButton() {
  const { theme, setTheme } = useTheme();

  const isDark = theme === 'dark';

  const toggleTheme = () => {
    setTheme(isDark ? 'light' : 'dark');
  };

  return (
    <label className="inline-flex items-center cursor-pointer">
      <input
        type="checkbox"
        checked={isDark}
        onChange={toggleTheme}
        className="sr-only peer"
      />
      <div className="relative w-[76px] h-[32px] bg-gray-300 peer-focus:ring-4 peer-focus:ring-blue-300 dark:peer-focus:ring-blue-800 rounded-full peer dark:bg-gray-700 peer-checked:bg-blue-600 dark:peer-checked:bg-blue-600 transition-all">
        <div
          className={`absolute top-0.5 left-[4px] w-7 h-7 bg-white border border-gray-300 rounded-full dark:border-gray-600 transition-all flex items-center justify-center ${
            isDark ? 'translate-x-[40px]' : 'translate-x-0'
          }`}
        >
          <FontAwesomeIcon
            icon={isDark ? faMoon : faSun}
            className="w-4 h-4 text-gray-700 dark:text-gray-300"
          />
        </div>
      </div>
    </label>
  );
}
