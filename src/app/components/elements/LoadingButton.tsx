"use client";

import { useFormStatus } from "react-dom";

type LoadingButtonProps = {
  label: string;
  color?: "blue" | "green";
};

export default function LoadingButton({
  label,
  color = "blue",
}: LoadingButtonProps) {
  const { pending } = useFormStatus();
  const buttonColor = `bg-${color}-500 hover:bg-${color}-600 dark:hover:bg-${color}-400`;

  return (
    <button
      type="submit"
      disabled={pending}
      className={`w-full p-2 rounded-md transition duration-300 flex items-center justify-center text-white disabled:opacity-50 ${buttonColor}`}
    >
      {pending ? (
        <div className="w-5 h-5 border-2 border-white border-t-transparent rounded-full animate-spin" />
      ) : (
        label
      )}
    </button>
  );
}
