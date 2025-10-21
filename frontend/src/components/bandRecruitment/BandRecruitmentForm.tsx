'use client';

import { useEffect, useState } from 'react';

export type BandRecruitmentFormValues = {
  title: string;
  description: string;
  genre: string;
  location: string;
  recruitingParts: string;
  skillLevel: string;
  contact: string;
  deadline: string;
  status: string;
};

type BandRecruitmentFormProps = {
  initialValues?: Partial<BandRecruitmentFormValues>;
  submitLabel: string;
  isSubmitting?: boolean;
  onSubmit: (values: BandRecruitmentFormValues) => void;
  onCancel?: () => void;
};

const defaultValues: BandRecruitmentFormValues = {
  title: '',
  description: '',
  genre: '',
  location: '',
  recruitingParts: '',
  skillLevel: '',
  contact: '',
  deadline: '',
  status: 'open',
};

const BandRecruitmentForm: React.FC<BandRecruitmentFormProps> = ({
  initialValues,
  submitLabel,
  isSubmitting,
  onSubmit,
  onCancel,
}) => {
  const [values, setValues] = useState<BandRecruitmentFormValues>({
    ...defaultValues,
    ...initialValues,
  });
  const [error, setError] = useState<string | null>(null);

  useEffect(() => {
    setValues((prev) => ({ ...prev, ...initialValues }));
  }, [initialValues]);

  const handleChange = (
    field: keyof BandRecruitmentFormValues,
    value: string
  ) => {
    setValues((prev) => ({ ...prev, [field]: value }));
  };

  const handleSubmit = (e: React.FormEvent) => {
    e.preventDefault();
    if (!values.title.trim()) {
      setError('タイトルを入力してください。');
      return;
    }
    if (!values.description.trim()) {
      setError('募集内容を入力してください。');
      return;
    }
    if (!values.contact.trim()) {
      setError('連絡方法を入力してください。');
      return;
    }

    setError(null);
    onSubmit({ ...values });
  };

  return (
    <form onSubmit={handleSubmit} className="space-y-6">
      {error && <p className="text-red-600 text-sm">{error}</p>}
      <div>
        <label className="block text-sm font-medium text-slate-700">
          タイトル
        </label>
        <input
          type="text"
          value={values.title}
          onChange={(e) => handleChange('title', e.target.value)}
          className="mt-1 w-full rounded-md border border-gray-300 px-3 py-2 focus:border-indigo-500 focus:outline-none focus:ring-1 focus:ring-indigo-500"
          placeholder="例: ロックバンドのギタリスト募集"
        />
      </div>
      <div>
        <label className="block text-sm font-medium text-slate-700">
          募集内容
        </label>
        <textarea
          value={values.description}
          onChange={(e) => handleChange('description', e.target.value)}
          className="mt-1 w-full rounded-md border border-gray-300 px-3 py-2 focus:border-indigo-500 focus:outline-none focus:ring-1 focus:ring-indigo-500"
          rows={6}
          placeholder="募集の背景や求める人物像、活動頻度などを記入してください。"
        />
      </div>
      <div className="grid grid-cols-1 gap-4 md:grid-cols-2">
        <div>
          <label className="block text-sm font-medium text-slate-700">
            ジャンル
          </label>
          <input
            type="text"
            value={values.genre}
            onChange={(e) => handleChange('genre', e.target.value)}
            className="mt-1 w-full rounded-md border border-gray-300 px-3 py-2 focus:border-indigo-500 focus:outline-none focus:ring-1 focus:ring-indigo-500"
            placeholder="例: ロック、J-POP"
          />
        </div>
        <div>
          <label className="block text-sm font-medium text-slate-700">
            活動地域
          </label>
          <input
            type="text"
            value={values.location}
            onChange={(e) => handleChange('location', e.target.value)}
            className="mt-1 w-full rounded-md border border-gray-300 px-3 py-2 focus:border-indigo-500 focus:outline-none focus:ring-1 focus:ring-indigo-500"
            placeholder="例: 東京 / オンライン"
          />
        </div>
      </div>
      <div className="grid grid-cols-1 gap-4 md:grid-cols-2">
        <div>
          <label className="block text-sm font-medium text-slate-700">
            募集パート (カンマ区切り)
          </label>
          <input
            type="text"
            value={values.recruitingParts}
            onChange={(e) => handleChange('recruitingParts', e.target.value)}
            className="mt-1 w-full rounded-md border border-gray-300 px-3 py-2 focus:border-indigo-500 focus:outline-none focus:ring-1 focus:ring-indigo-500"
            placeholder="例: Vo, Gt, Dr"
          />
        </div>
        <div>
          <label className="block text-sm font-medium text-slate-700">
            希望スキル / 経験
          </label>
          <input
            type="text"
            value={values.skillLevel}
            onChange={(e) => handleChange('skillLevel', e.target.value)}
            className="mt-1 w-full rounded-md border border-gray-300 px-3 py-2 focus:border-indigo-500 focus:outline-none focus:ring-1 focus:ring-indigo-500"
            placeholder="例: 初心者歓迎 / 中級以上"
          />
        </div>
      </div>
      <div className="grid grid-cols-1 gap-4 md:grid-cols-2">
        <div>
          <label className="block text-sm font-medium text-slate-700">
            連絡方法
          </label>
          <input
            type="text"
            value={values.contact}
            onChange={(e) => handleChange('contact', e.target.value)}
            className="mt-1 w-full rounded-md border border-gray-300 px-3 py-2 focus:border-indigo-500 focus:outline-none focus:ring-1 focus:ring-indigo-500"
            placeholder="例: sample@example.com / Discord ID"
          />
        </div>
        <div>
          <label className="block text-sm font-medium text-slate-700">
            募集締切
          </label>
          <input
            type="date"
            value={values.deadline}
            onChange={(e) => handleChange('deadline', e.target.value)}
            className="mt-1 w-full rounded-md border border-gray-300 px-3 py-2 focus:border-indigo-500 focus:outline-none focus:ring-1 focus:ring-indigo-500"
          />
        </div>
      </div>
      <div>
        <label className="block text-sm font-medium text-slate-700">
          ステータス
        </label>
        <select
          value={values.status}
          onChange={(e) => handleChange('status', e.target.value)}
          className="mt-1 w-full rounded-md border border-gray-300 px-3 py-2 focus:border-indigo-500 focus:outline-none focus:ring-1 focus:ring-indigo-500"
        >
          <option value="open">募集中</option>
          <option value="closed">募集終了</option>
        </select>
      </div>
      <div className="flex justify-end gap-3 pt-4">
        {onCancel && (
          <button
            type="button"
            onClick={onCancel}
            className="rounded-lg border border-gray-300 px-4 py-2 text-sm font-medium text-gray-600 hover:bg-gray-100"
          >
            キャンセル
          </button>
        )}
        <button
          type="submit"
          disabled={isSubmitting}
          className="rounded-lg bg-indigo-600 px-4 py-2 text-sm font-semibold text-white hover:bg-indigo-700 disabled:opacity-60"
        >
          {isSubmitting ? '送信中...' : submitLabel}
        </button>
      </div>
    </form>
  );
};

export default BandRecruitmentForm;
