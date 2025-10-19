'use client';

import { Calendar, MapPin, Users } from 'lucide-react';
import Link from 'next/link';
import type { BandRecruitment } from '@/libs/api/generated/orval/model/bandRecruitment';

const statusStyles: Record<string, string> = {
  open: 'bg-green-100 text-green-700 border border-green-200',
  closed: 'bg-gray-200 text-gray-600 border border-gray-300',
};

type BandRecruitmentCardProps = {
  recruitment: BandRecruitment;
};

const formatDate = (value?: string | null) => {
  if (!value) return '';
  const date = new Date(value);
  if (Number.isNaN(date.getTime())) return value;
  return date.toLocaleDateString();
};

const BandRecruitmentCard: React.FC<BandRecruitmentCardProps> = ({
  recruitment,
}) => {
  const status = recruitment.status ?? 'open';
  const badgeClass = statusStyles[status] ?? statusStyles.open;

  return (
    <div className="flex flex-col justify-between rounded-xl border border-gray-200 bg-white p-6 shadow-sm hover:shadow-md transition-shadow duration-200">
      <div>
        <div className="flex items-start justify-between gap-2">
          <h3 className="text-xl font-semibold text-slate-900">
            {recruitment.title ?? 'タイトル未設定'}
          </h3>
          <span
            className={`px-3 py-1 rounded-full text-sm font-medium ${badgeClass}`}
          >
            {status === 'open' ? '募集中' : '募集終了'}
          </span>
        </div>
        <p className="mt-3 text-sm text-slate-600 line-clamp-3">
          {recruitment.description ?? '募集内容が入力されていません。'}
        </p>
        <div className="mt-4 space-y-2 text-sm text-slate-600">
          {recruitment.genre && (
            <div>
              <span className="font-semibold">ジャンル:</span>{' '}
              {recruitment.genre}
            </div>
          )}
          <div className="flex items-center gap-2">
            <MapPin size={16} className="text-slate-400" />
            <span>{recruitment.location ?? '場所未設定'}</span>
          </div>
          <div className="flex items-center gap-2">
            <Users size={16} className="text-slate-400" />
            <span>
              募集パート: {recruitment.recruitingParts?.join(', ') || '未設定'}
            </span>
          </div>
          {recruitment.deadline && (
            <div className="flex items-center gap-2">
              <Calendar size={16} className="text-slate-400" />
              <span>締切: {formatDate(recruitment.deadline)}</span>
            </div>
          )}
        </div>
      </div>
      <Link
        href={`/dashboard/band-recruitments/${recruitment.id}`}
        className="mt-6 inline-flex items-center justify-center rounded-lg bg-indigo-600 px-4 py-2 text-sm font-semibold text-white hover:bg-indigo-700"
      >
        詳細を見る
      </Link>
    </div>
  );
};

export default BandRecruitmentCard;
