"use client";

import { CheckCircle } from "lucide-react";
import Link from "next/link";
import { useParams } from "next/navigation";

const BandRecruitmentApplyCompletePage = () => {
  const params = useParams<{ id: string }>();
  const id = params?.id;

  return (
    <div className="mx-auto flex max-w-xl flex-col items-center gap-6 p-10 text-center">
      <CheckCircle className="h-16 w-16 text-emerald-500" />
      <div>
        <h1 className="text-3xl font-bold text-slate-900">応募が完了しました！</h1>
        <p className="mt-3 text-sm text-slate-600">
          応募内容は募集主に送信されました。返信をお待ちください。
        </p>
      </div>
      <div className="flex flex-col gap-3 sm:flex-row">
        {id && (
          <Link
            href={`/dashboard/band-recruitments/${id}`}
            className="rounded-lg border border-gray-300 px-4 py-2 text-sm font-medium text-gray-600 hover:bg-gray-100"
          >
            募集詳細に戻る
          </Link>
        )}
        <Link
          href="/dashboard/band-recruitments"
          className="rounded-lg bg-indigo-600 px-4 py-2 text-sm font-semibold text-white hover:bg-indigo-700"
        >
          他の募集を探す
        </Link>
      </div>
    </div>
  );
};

export default BandRecruitmentApplyCompletePage;
