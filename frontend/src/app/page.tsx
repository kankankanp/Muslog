'use client';

import { motion } from 'framer-motion';
import {
  ArrowRight,
  Music,
  Users,
  BookOpen,
  Star,
  Heart,
  UserPlus,
  Music2,
  PenTool,
} from 'lucide-react';
import Image from 'next/image';
import Link from 'next/link';

const fadeInUp = {
  initial: { opacity: 0, y: 60 },
  animate: { opacity: 1, y: 0 },
  transition: { duration: 0.6 },
};

const staggerContainer = {
  animate: {
    transition: {
      staggerChildren: 0.2,
    },
  },
};

const WaveDivider = ({
  upperSectionBg,
  lowerSectionBg,
}: {
  upperSectionBg: string;
  lowerSectionBg: string;
}) => (
  <div className={`relative h-24 overflow-hidden`}>
    <svg
      className={`absolute bottom-0 w-full h-24 ${lowerSectionBg}`}
      viewBox="0 0 1440 100"
      preserveAspectRatio="none"
      fill="currentColor"
    >
      <path
        d="M0,50 C150,100 350,0 500,50 C650,100 850,0 1000,50 C1150,100 1350,0 1440,50 L1440,100 L0,100 Z"
        className={upperSectionBg}
      />
    </svg>
  </div>
);

const WaveDividerInverted = ({
  upperSectionBg,
  lowerSectionBg,
}: {
  upperSectionBg: string;
  lowerSectionBg: string;
}) => (
  <div className={`relative h-24 overflow-hidden`}>
    <svg
      className={`absolute bottom-0 w-full h-24 ${lowerSectionBg}`}
      viewBox="0 0 1440 100"
      preserveAspectRatio="none"
      fill="currentColor"
    >
      <path
        d="M0,50 C150,0 350,100 500,50 C650,0 850,100 1000,50 C1150,0 1350,100 1440,50 L1440,100 L0,100 Z"
        className={upperSectionBg}
      />
    </svg>
  </div>
);

export default function Page() {
  return (
    <div className="min-h-screen bg-gradient-to-b from-emerald-50 to-white dark:from-gray-950 dark:to-gray-900">
      <section className="relative h-screen flex items-center justify-center text-center px-4 md:text-left bg-emerald-50 dark:bg-gray-900">
        <div className="flex flex-col md:flex-row max-w-6xl mx-auto items-center justify-between gap-2">
          <motion.div
            className="md:w-1/2"
            initial={{ opacity: 0, x: -60 }}
            animate={{ opacity: 1, x: 0 }}
            transition={{ duration: 0.8 }}
          >
            <motion.h1
              className="text-2xl md:text-6xl font-bold text-emerald-800 dark:text-emerald-100 mb-2"
              initial={{ opacity: 0, y: 20 }}
              animate={{ opacity: 1, y: 0 }}
              transition={{ duration: 0.8, delay: 0.2 }}
            >
              あなたの音楽の物語を
              <span className="inline-block">シェアしよう。</span>
            </motion.h1>
            <motion.p
              className="text-xl text-emerald-600 dark:text-emerald-300 mb-8"
              initial={{ opacity: 0, y: 20 }}
              animate={{ opacity: 1, y: 0 }}
              transition={{ duration: 0.8, delay: 0.4 }}
            >
              好きな曲と共に、
              <span className="inline-block">あなたの想いを綴る。</span>
              <br />
              新しい音楽との出会いの場へ。
            </motion.p>
            <motion.div
              className="flex flex-wrap gap-4"
              initial={{ opacity: 0, y: 20 }}
              animate={{ opacity: 1, y: 0 }}
              transition={{ duration: 0.8, delay: 0.6 }}
            >
              <Link
                href="/dashboard"
                className="inline-flex items-center px-8 py-4 bg-emerald-600 dark:bg-emerald-700 text-white rounded-full text-lg font-semibold hover:bg-emerald-700 dark:hover:bg-emerald-800 transition-colors"
              >
                はじめる
                <ArrowRight className="ml-2 w-5 h-5" />
              </Link>
              <Link
                href="/login-or-signup"
                className="inline-flex items-center px-8 py-4 bg-white dark:bg-gray-700 text-emerald-600 dark:text-emerald-300 rounded-full text-lg font-semibold hover:bg-gray-100 dark:hover:bg-gray-600 transition-colors border border-emerald-600 dark:border-emerald-300"
              >
                ログイン
              </Link>
            </motion.div>
          </motion.div>
          <motion.div
            className="md:w-1/2 mt-8 md:mt-0"
            initial={{ opacity: 0, x: 60 }}
            animate={{ opacity: 1, x: 0 }}
            transition={{ duration: 0.8, delay: 0.8 }}
          >
            <Image
              src="/mv.png"
              width={200}
              height={200}
              alt="音楽コミュニティのイラスト"
              className="w-full h-auto object-contain md:w-[560px]"
            />
          </motion.div>
        </div>
      </section>

      <WaveDivider
        upperSectionBg="text-white dark:text-gray-800"
        lowerSectionBg="bg-emerald-50 dark:bg-gray-900"
      />

      <section className="py-20 px-4 bg-white dark:bg-gray-800">
        <motion.div
          className="max-w-6xl mx-auto"
          initial="initial"
          whileInView="animate"
          viewport={{ once: true }}
          variants={staggerContainer}
        >
          <motion.h2
            className="text-3xl font-bold text-emerald-800 dark:text-emerald-100 text-center mb-16"
            variants={fadeInUp}
          >
            3つの特徴
          </motion.h2>
          <div className="grid md:grid-cols-3 gap-12">
            {[
              {
                icon: Music,
                title: '音楽との出会い',
                description:
                  'あなたの好きな曲を共有し、新しい音楽との出会いを創出します。',
              },
              {
                icon: BookOpen,
                title: 'ブログ機能',
                description:
                  '音楽にまつわる思い出や想いを、ブログとして残すことができます。',
              },
              {
                icon: Users,
                title: 'コミュニティ',
                description: '同じ音楽を愛する仲間と、想いを共有できます。',
              },
            ].map((feature, index) => (
              <motion.div
                key={index}
                className="text-center"
                variants={fadeInUp}
              >
                <motion.div
                  className="w-16 h-16 bg-emerald-100 dark:bg-gray-700 rounded-full flex items-center justify-center mx-auto mb-6"
                  whileHover={{ scale: 1.1 }}
                  transition={{ type: 'spring', stiffness: 400, damping: 10 }}
                >
                  <feature.icon className="w-8 h-8 text-emerald-600 dark:text-emerald-300" />
                </motion.div>
                <h3 className="text-xl font-semibold text-emerald-800 dark:text-emerald-100 mb-4">
                  {feature.title}
                </h3>
                <p className="text-emerald-600 dark:text-emerald-300">
                  {feature.description}
                </p>
              </motion.div>
            ))}
          </div>
        </motion.div>
      </section>

      <WaveDividerInverted
        upperSectionBg="text-emerald-50 dark:text-gray-900"
        lowerSectionBg="bg-white dark:bg-gray-800"
      />

      <section className="py-20 px-4 bg-emerald-50 dark:bg-gray-900">
        <motion.div
          className="max-w-6xl mx-auto"
          initial="initial"
          whileInView="animate"
          viewport={{ once: true }}
          variants={staggerContainer}
        >
          <motion.h2
            className="text-3xl font-bold text-emerald-800 dark:text-emerald-100 text-center mb-16"
            variants={fadeInUp}
          >
            簡単3ステップ
          </motion.h2>
          <div className="grid md:grid-cols-3 gap-12">
            {[
              {
                number: '01',
                title: 'アカウント作成',
                description: '簡単な登録で、すぐに始められます。',
                icon: UserPlus,
              },
              {
                number: '02',
                title: '音楽を選択',
                description: 'あなたの好きな曲を選んでください。',
                icon: Music2,
              },
              {
                number: '03',
                title: 'ブログを投稿',
                description: '音楽にまつわる想いを綴りましょう。',
                icon: PenTool,
              },
            ].map((step, index) => (
              <motion.div
                key={index}
                className="bg-white dark:bg-gray-800 p-8 rounded-lg shadow-sm"
                variants={fadeInUp}
                whileHover={{ y: -10 }}
                transition={{ type: 'spring', stiffness: 300 }}
              >
                <div className="flex items-center justify-between mb-4">
                  <div className="text-4xl font-bold text-emerald-600 dark:text-emerald-300">
                    {step.number}
                  </div>
                  <motion.div
                    className="w-12 h-12 bg-emerald-100 dark:bg-gray-700 rounded-full flex items-center justify-center"
                    whileHover={{ scale: 1.1 }}
                    transition={{ type: 'spring', stiffness: 400, damping: 10 }}
                  >
                    <step.icon className="w-6 h-6 text-emerald-600 dark:text-emerald-300" />
                  </motion.div>
                </div>
                <h3 className="text-xl font-semibold text-emerald-800 dark:text-emerald-100 mb-4">
                  {step.title}
                </h3>
                <p className="text-emerald-600 dark:text-emerald-300">
                  {step.description}
                </p>
              </motion.div>
            ))}
          </div>
        </motion.div>
      </section>

      <WaveDivider
        upperSectionBg="text-white dark:text-gray-800"
        lowerSectionBg="bg-emerald-50 dark:bg-gray-900"
      />

      <section className="py-20 px-4 bg-white dark:bg-gray-800">
        <motion.div
          className="max-w-6xl mx-auto"
          initial="initial"
          whileInView="animate"
          viewport={{ once: true }}
          variants={staggerContainer}
        >
          <motion.h2
            className="text-3xl font-bold text-emerald-800 dark:text-emerald-100 text-center mb-16"
            variants={fadeInUp}
          >
            こんな方におすすめ
          </motion.h2>
          <div className="grid md:grid-cols-2 gap-12">
            {[
              {
                icon: Star,
                title: '音楽が好きな方',
                description:
                  '好きな曲を共有し、新しい音楽との出会いを楽しめます。',
              },
              {
                icon: Heart,
                title: '思い出を残したい方',
                description:
                  '音楽にまつわる大切な思い出を、ブログとして残せます。',
              },
            ].map((benefit, index) => (
              <motion.div
                key={index}
                className="flex items-start space-x-4"
                variants={fadeInUp}
                whileHover={{ x: 10 }}
                transition={{ type: 'spring', stiffness: 300 }}
              >
                <benefit.icon className="w-6 h-6 text-emerald-600 dark:text-emerald-300 mt-1" />
                <div>
                  <h3 className="text-xl font-semibold text-emerald-800 dark:text-emerald-100 mb-2">
                    {benefit.title}
                  </h3>
                  <p className="text-emerald-600 dark:text-emerald-300">
                    {benefit.description}
                  </p>
                </div>
              </motion.div>
            ))}
          </div>
        </motion.div>
      </section>

      <WaveDividerInverted
        upperSectionBg="text-emerald-50 dark:text-gray-900"
        lowerSectionBg="bg-white dark:bg-gray-800"
      />
    </div>
  );
}
