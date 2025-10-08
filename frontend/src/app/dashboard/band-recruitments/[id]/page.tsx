"use client";

import { Calendar, Edit, MapPin, Users } from "lucide-react";
import Link from "next/link";
import { useParams, useRouter } from "next/navigation";
import { useState } from "react";
import toast from "react-hot-toast";
import Spinner from "@/components/layouts/Spinner";
import { useGetMe } from "@/libs/api/generated/orval/auth/auth";
import {
  useGetBandRecruitmentsId,
  usePostBandRecruitmentsIdApply,
} from "@/libs/api/generated/orval/band-recruitments/band-recruitments";

const formatDateTime = (value?: string | null) => {
  if (!value) return undefined;
  const date = new Date(value);
  if (Number.isNaN(date.getTime())) {
    return value.split("T")[0];
  }
  return `${date.getFullYear()}-${String(date.getMonth() + 1).padStart(2, "0")}-${String(date.getDate()).padStart(2, "0")}`;
};

const BandRecruitmentDetailPage = () => {
  const params = useParams<{ id: string }>();
  const router = useRouter();
  const id = params?.id;
  const [message, setMessage] = useState("");

  const { data: me } = useGetMe();
  const { data, isLoading, isError } = useGetBandRecruitmentsId(id, {
    query: {
      enabled: !!id,
    },
  });
  const { mutateAsync: applyMutation, isPending: isApplying } =
    usePostBandRecruitmentsIdApply();

  if (isLoading) {
    return <Spinner />;
  }

  if (isError || !data?.recruitment) {
    return (
      <div className="p-8 text-center text-red-600">
        募集情報の取得に失敗しました。
      </div>
    );
  }

  const recruitment = data.recruitment;
  const isOwner = me?.id && recruitment.userId === me.id;
  const canApply =
    !isOwner &&
    recruitment.status !== "closed" &&
    !recruitment.hasApplied;

  const handleApply = async () => {
    if (!id) return;
    if (!message.trim()) {
      toast.error("応募メッセージを入力してください");
      return;
    }

    try {
      await applyMutation({
        id,
        data: {
          message,
        },
      });
      toast.success("応募が完了しました");
      router.push(`/dashboard/band-recruitments/${id}/apply/complete`);
    } catch (err) {
      console.error(err);
      toast.error("応募に失敗しました");
    }
  };

  return (
    <div className="mx-auto max-w-4xl space-y-8 p-6">
      <div className="flex flex-col gap-4 rounded-xl border border-gray-200 bg-white p-6 shadow-sm">
        <div className="flex flex-col gap-4 md:flex-row md:items-center md:justify-between">
          <div>
            <span
              className={`inline-flex items-center rounded-full px-3 py-1 text-xs font-semibold ${
                recruitment.status === "open"
                  ? "bg-green-100 text-green-700"
                  : "bg-gray-200 text-gray-600"
              }`}
            >
              {recruitment.status === "open" ? "募集中" : "募集終了"}
            </span>
            <h1 className="mt-2 text-3xl font-bold text-slate-900">
              {recruitment.title}
            </h1>
            <p className="mt-2 text-sm text-slate-500">
              投稿日: {formatDateTime(recruitment.createdAt)}
            </p>
          </div>
          <div className="flex gap-3">
            <Link
              href="/dashboard/band-recruitments"
              className="rounded-lg border border-gray-300 px-4 py-2 text-sm font-medium text-gray-600 hover:bg-gray-100"
            >
              一覧へ戻る
            </Link>
            {isOwner && (
              <Link
                href={`/dashboard/band-recruitments/${id}/edit`}
                className="inline-flex items-center gap-2 rounded-lg bg-indigo-600 px-4 py-2 text-sm font-semibold text-white hover:bg-indigo-700"
              >
                <Edit size={16} /> 編集
              </Link>
            )}
          </div>
        </div>
        <p className="text-base leading-relaxed text-slate-700 whitespace-pre-wrap">
          {recruitment.description}
        </p>
        <div className="grid grid-cols-1 gap-4 border-t border-gray-100 pt-4 md:grid-cols-2">
          <div className="flex items-center gap-2 text-sm text-slate-600">
            <Users size={18} className="text-slate-400" />
            募集パート: {recruitment.recruitingParts?.join(", ") || "未設定"}
          </div>
          <div className="flex items-center gap-2 text-sm text-slate-600">
            <MapPin size={18} className="text-slate-400" />
            活動地域: {recruitment.location || "未設定"}
          </div>
          <div className="flex items-center gap-2 text-sm text-slate-600">
            <Calendar size={18} className="text-slate-400" />
            募集締切: {formatDateTime(recruitment.deadline) ?? "未設定"}
          </div>
          <div className="text-sm text-slate-600">
            希望スキル / 経験: {recruitment.skillLevel || "特になし"}
          </div>
          <div className="text-sm text-slate-600">
            連絡方法: {recruitment.contact || "未設定"}
          </div>
          <div className="text-sm text-slate-600">
            応募数: {recruitment.applicationsCount ?? 0} 件
          </div>
        </div>
      </div>

      <div className="rounded-xl border border-gray-200 bg-white p-6 shadow-sm">
        <h2 className="text-xl font-semibold text-slate-900">応募する</h2>
        {isOwner && (
          <p className="mt-2 text-sm text-slate-500">
            自分が作成した募集には応募できません。
          </p>
        )}
        {recruitment.status === "closed" && (
          <p className="mt-2 text-sm text-red-500">この募集は締め切られています。</p>
        )}
        {recruitment.hasApplied && !isOwner && (
          <p className="mt-2 text-sm text-slate-500">すでに応募済みです。</p>
        )}
        {canApply && (
          <div className="mt-4 space-y-4">
            <textarea
              value={message}
              onChange={(e) => setMessage(e.target.value)}
              rows={5}
              placeholder="応募メッセージを入力してください (自己紹介や参加可能日など)"
              className="w-full rounded-md border border-gray-300 px-3 py-2 text-sm focus:border-indigo-500 focus:outline-none focus:ring-1 focus:ring-indigo-500"
            />
            <button
              type="button"
              onClick={handleApply}
              disabled={isApplying}
              className="rounded-lg bg-emerald-600 px-4 py-2 text-sm font-semibold text-white hover:bg-emerald-700 disabled:opacity-60"
            >
              {isApplying ? "送信中..." : "応募を送信"}
            </button>
          </div>
        )}
      </div>
    </div>
  );
};

export default BandRecruitmentDetailPage;
