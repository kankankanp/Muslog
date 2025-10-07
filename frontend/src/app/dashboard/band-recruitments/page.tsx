"use client";

import { Filter, RefreshCw } from "lucide-react";
import Link from "next/link";
import { useMemo, useState } from "react";
import BandRecruitmentCard from "@/components/bandRecruitment/BandRecruitmentCard";
import Spinner from "@/components/layouts/Spinner";
import { useGetBandRecruitments } from "@/libs/api/generated/orval/band-recruitments/band-recruitments";

const PER_PAGE = 9;

const BandRecruitmentsPage = () => {
  const [keyword, setKeyword] = useState("");
  const [genre, setGenre] = useState("");
  const [location, setLocation] = useState("");
  const [status, setStatus] = useState("all");
  const [page, setPage] = useState(1);

  const queryParams = useMemo(
    () => ({
      keyword: keyword || undefined,
      genre: genre || undefined,
      location: location || undefined,
      status: status === "all" ? undefined : status,
      page,
      perPage: PER_PAGE,
    }),
    [keyword, genre, location, status, page],
  );

  const { data, isLoading, isError, error } = useGetBandRecruitments(queryParams);

  const recruitments = data?.recruitments ?? [];
  const totalCount = data?.totalCount ?? 0;
  const totalPages = Math.max(1, Math.ceil(totalCount / PER_PAGE));

  const resetFilters = () => {
    setKeyword("");
    setGenre("");
    setLocation("");
    setStatus("all");
    setPage(1);
  };

  if (isLoading) {
    return <Spinner />;
  }

  if (isError) {
    return (
      <div className="p-8 text-center text-red-600">
        データの取得に失敗しました。
        <div className="mt-2 text-sm text-gray-500">
          {error instanceof Error ? error.message : "Unknown error"}
        </div>
      </div>
    );
  }

  return (
    <div className="space-y-6 p-6">
      <div className="flex flex-col gap-4 border-b border-gray-200 pb-4 md:flex-row md:items-center md:justify-between">
        <h1 className="text-3xl font-bold text-slate-900">バンド募集</h1>
        <Link
          href="/dashboard/band-recruitments/create"
          className="inline-flex items-center justify-center rounded-lg bg-indigo-600 px-4 py-2 text-sm font-semibold text-white hover:bg-indigo-700"
        >
          新規募集を作成
        </Link>
      </div>

      <div className="rounded-xl border border-gray-200 bg-white p-5 shadow-sm">
        <h2 className="flex items-center gap-2 text-lg font-semibold text-slate-800">
          <Filter size={18} />
          条件で絞り込む
        </h2>
        <div className="mt-4 grid grid-cols-1 gap-4 md:grid-cols-2 lg:grid-cols-4">
          <div>
            <label className="block text-sm font-medium text-slate-600">キーワード</label>
            <input
              type="text"
              value={keyword}
              onChange={(e) => {
                setKeyword(e.target.value);
                setPage(1);
              }}
              className="mt-1 w-full rounded-md border border-gray-300 px-3 py-2 text-sm focus:border-indigo-500 focus:outline-none focus:ring-1 focus:ring-indigo-500"
              placeholder="タイトル・本文から検索"
            />
          </div>
          <div>
            <label className="block text-sm font-medium text-slate-600">ジャンル</label>
            <input
              type="text"
              value={genre}
              onChange={(e) => {
                setGenre(e.target.value);
                setPage(1);
              }}
              className="mt-1 w-full rounded-md border border-gray-300 px-3 py-2 text-sm focus:border-indigo-500 focus:outline-none focus:ring-1 focus:ring-indigo-500"
              placeholder="例: ロック"
            />
          </div>
          <div>
            <label className="block text-sm font-medium text-slate-600">活動地域</label>
            <input
              type="text"
              value={location}
              onChange={(e) => {
                setLocation(e.target.value);
                setPage(1);
              }}
              className="mt-1 w-full rounded-md border border-gray-300 px-3 py-2 text-sm focus:border-indigo-500 focus:outline-none focus:ring-1 focus:ring-indigo-500"
              placeholder="例: 東京 / オンライン"
            />
          </div>
          <div>
            <label className="block text-sm font-medium text-slate-600">ステータス</label>
            <select
              value={status}
              onChange={(e) => {
                setStatus(e.target.value);
                setPage(1);
              }}
              className="mt-1 w-full rounded-md border border-gray-300 px-3 py-2 text-sm focus:border-indigo-500 focus:outline-none focus:ring-1 focus:ring-indigo-500"
            >
              <option value="all">すべて</option>
              <option value="open">募集中</option>
              <option value="closed">募集終了</option>
            </select>
          </div>
        </div>
        <div className="mt-4 flex items-center justify-end gap-3">
          <button
            type="button"
            onClick={resetFilters}
            className="inline-flex items-center gap-2 rounded-md border border-gray-300 px-3 py-2 text-sm text-gray-600 hover:bg-gray-100"
          >
            <RefreshCw size={14} />
            リセット
          </button>
        </div>
      </div>

      <div>
        <div className="flex items-center justify-between text-sm text-slate-600">
          <span>{totalCount} 件の募集が見つかりました</span>
        </div>
        {recruitments.length === 0 ? (
          <div className="mt-6 rounded-lg border border-dashed border-gray-300 p-12 text-center text-slate-500">
            条件に合う募集は見つかりませんでした。
          </div>
        ) : (
          <div className="mt-6 grid grid-cols-1 gap-6 md:grid-cols-2 xl:grid-cols-3">
            {recruitments.map((recruitment) => (
              <BandRecruitmentCard
                key={recruitment.id}
                recruitment={recruitment}
              />
            ))}
          </div>
        )}
      </div>

      {totalPages > 1 && (
        <div className="flex justify-center gap-2 pt-4">
          {Array.from({ length: totalPages }, (_, index) => index + 1).map(
            (pageNumber) => (
              <button
                key={pageNumber}
                onClick={() => setPage(pageNumber)}
                className={`rounded-md px-3 py-1 text-sm font-medium ${
                  pageNumber === page
                    ? "bg-indigo-600 text-white"
                    : "border border-gray-300 text-slate-700 hover:bg-gray-100"
                }`}
              >
                {pageNumber}
              </button>
            ),
          )}
        </div>
      )}
    </div>
  );
};

export default BandRecruitmentsPage;
