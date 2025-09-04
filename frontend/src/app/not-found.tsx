"use client";

import { motion } from "framer-motion";
import { Home, Search, AlertCircle } from "lucide-react";
import Link from "next/link";

const fadeInUp = {
  initial: { opacity: 0, y: 60 },
  animate: { opacity: 1, y: 0 },
  transition: { duration: 0.6 },
};

export default function NotFound() {
  return (
    <div className="min-h-screen bg-gradient-to-b from-emerald-50 to-white dark:from-gray-950 dark:to-gray-900 flex items-center justify-center px-4">
      <motion.div
        className="text-center max-w-md w-full"
        initial="initial"
        animate="animate"
        variants={{
          animate: {
            transition: {
              staggerChildren: 0.2,
            },
          },
        }}
      >
        <motion.div
          className="w-24 h-24 bg-emerald-100 dark:bg-gray-700 rounded-full flex items-center justify-center mx-auto mb-8"
          variants={fadeInUp}
          whileHover={{ scale: 1.1 }}
          transition={{ type: "spring", stiffness: 400, damping: 10 }}
        >
          <AlertCircle className="w-12 h-12 text-emerald-600 dark:text-emerald-300" />
        </motion.div>

        <motion.h1
          className="text-6xl font-bold text-emerald-800 dark:text-emerald-100 mb-4"
          variants={fadeInUp}
        >
          404
        </motion.h1>

        <motion.h2
          className="text-2xl font-semibold text-emerald-700 dark:text-emerald-200 mb-4"
          variants={fadeInUp}
        >
          ページが見つかりません。
        </motion.h2>

        <motion.p
          className="text-emerald-600 dark:text-emerald-300 mb-8"
          variants={fadeInUp}
        >
          お探しのページは存在しないか、移動された可能性があります。
        </motion.p>

        <motion.div
          className="flex flex-col sm:flex-row gap-4 justify-center"
          variants={fadeInUp}
        >
          <Link
            href="/dashboard"
            className="inline-flex items-center px-6 py-3 bg-emerald-600 dark:bg-emerald-700 text-white rounded-full font-semibold hover:bg-emerald-700 dark:hover:bg-emerald-800 transition-colors"
          >
            <Home className="mr-2 w-5 h-5" />
            ホームに戻る
          </Link>
          <Link
            href="/login-or-signup"
            className="inline-flex items-center px-6 py-3 bg-white dark:bg-gray-800 text-emerald-600 dark:text-emerald-300 border-2 border-emerald-600 dark:border-emerald-300 rounded-full font-semibold hover:bg-emerald-50 dark:hover:bg-gray-700 transition-colors"
          >
            <Search className="mr-2 w-5 h-5" />
            ログインページへ
          </Link>
        </motion.div>
      </motion.div>
    </div>
  );
}
