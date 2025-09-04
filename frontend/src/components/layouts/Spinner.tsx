"use client";

export default function Spinner() {
  return (
    <div className="flex items-center justify-center w-full h-full">
      <div className="animate-spin rounded-full h-12 w-12 border-t-4 border-b-4 border-blue-500 border-opacity-40"></div>
    </div>
  );
}
