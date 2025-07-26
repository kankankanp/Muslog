import { BookOpen } from "lucide-react";
import Link from "next/link";

export default async function Page() {
  return (
    <div className="dark:bg-gray-900 bg-gray-100 min-h-screen">
      <div className="py-8 px-4 max-w-6xl mx-auto">
        {/* <ProfileCard name={user.name} email={user.email} /> */}
        <div className="mt-8">
          <Link href="/library" className="block">
            <div className="bg-white dark:bg-gray-800 rounded-lg shadow-md hover:shadow-lg transition-shadow duration-300 p-6 flex items-center space-x-4">
              <div className="flex-shrink-0 w-16 h-16 bg-emerald-100 dark:bg-gray-700 rounded-full flex items-center justify-center">
                <BookOpen className="w-8 h-8 text-emerald-600 dark:text-emerald-300" />
              </div>
              <div>
                <h2 className="text-2xl font-semibold text-emerald-800 dark:text-emerald-100 mb-2">
                  あなたのライブラリ
                </h2>
                <p className="text-gray-600 dark:text-gray-400">
                  これまでに投稿したブログ記事をすべて見る。
                </p>
              </div>
            </div>
          </Link>
        </div>
      </div>
    </div>
  );
}
