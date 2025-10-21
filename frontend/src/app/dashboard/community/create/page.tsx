'use client';

import { zodResolver } from '@hookform/resolvers/zod';
import { useRouter } from 'next/navigation';
import { useForm } from 'react-hook-form';
import toast from 'react-hot-toast';
import * as z from 'zod';
import { usePostCommunities } from '@/libs/api/generated/orval/communities/communities';

const createCommunitySchema = z.object({
  name: z.string().min(1, 'コミュニティ名は必須です'),
  description: z.string().min(1, '説明は必須です'),
});

type CreateCommunityFormData = z.infer<typeof createCommunitySchema>;

const CommunityCreatePage = () => {
  const router = useRouter();
  const {
    register,
    handleSubmit,
    formState: { errors },
    reset,
  } = useForm<CreateCommunityFormData>({
    resolver: zodResolver(createCommunitySchema),
  });

  const { mutateAsync, isPending } = usePostCommunities();

  const onSubmit = async (data: CreateCommunityFormData) => {
    try {
      const response = await mutateAsync({ data });
      toast.success('コミュニティを作成しました');
      reset();
      const newId = response?.community?.id;
      if (newId) {
        router.push(`/dashboard/community/${newId}`);
      } else {
        router.push('/dashboard/community');
      }
    } catch (error) {
      console.error('Error creating community:', error);
      toast.error('コミュニティの作成に失敗しました');
    }
  };

  return (
    <div className="mx-auto max-w-3xl space-y-6 p-6">
      <div className="border-b border-gray-200 pb-4">
        <h1 className="text-3xl font-bold text-slate-900">
          コミュニティを作成
        </h1>
        <p className="mt-2 text-sm text-slate-600">
          必要事項を入力して新しいコミュニティを作成しましょう。
        </p>
      </div>
      <form onSubmit={handleSubmit(onSubmit)} className="space-y-6">
        <div>
          <label
            htmlFor="name"
            className="block text-sm font-medium text-slate-700"
          >
            コミュニティ名
          </label>
          <input
            id="name"
            type="text"
            {...register('name')}
            disabled={isPending}
            className="mt-2 block w-full rounded-md border border-slate-300 px-3 py-2 shadow-sm focus:border-indigo-500 focus:outline-none focus:ring-1 focus:ring-indigo-500"
            placeholder="例: 東京インディーロック好き"
          />
          {errors.name && (
            <p className="mt-1 text-sm text-red-600">{errors.name.message}</p>
          )}
        </div>
        <div>
          <label
            htmlFor="description"
            className="block text-sm font-medium text-slate-700"
          >
            説明
          </label>
          <textarea
            id="description"
            rows={4}
            {...register('description')}
            disabled={isPending}
            className="mt-2 block w-full rounded-md border border-slate-300 px-3 py-2 shadow-sm focus:border-indigo-500 focus:outline-none focus:ring-1 focus:ring-indigo-500"
            placeholder="コミュニティの目的や参加条件などを記載してください"
          />
          {errors.description && (
            <p className="mt-1 text-sm text-red-600">
              {errors.description.message}
            </p>
          )}
        </div>
        <div className="flex justify-end gap-3">
          <button
            type="button"
            onClick={() => router.push('/dashboard/community')}
            className="rounded-md border border-slate-300 px-4 py-2 text-sm font-medium text-slate-700 transition hover:bg-slate-50"
            disabled={isPending}
          >
            キャンセル
          </button>
          <button
            type="submit"
            disabled={isPending}
            className="rounded-md bg-indigo-600 px-4 py-2 text-sm font-medium text-white shadow-sm transition hover:bg-indigo-700 focus:outline-none focus:ring-2 focus:ring-indigo-500 focus:ring-offset-1 disabled:cursor-not-allowed disabled:bg-indigo-300"
          >
            {isPending ? '作成中...' : 'コミュニティを作成'}
          </button>
        </div>
      </form>
    </div>
  );
};

export default CommunityCreatePage;
