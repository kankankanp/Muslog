'use client';

import { useRouter } from 'next/navigation';
import toast from 'react-hot-toast';
import BandRecruitmentForm, {
  type BandRecruitmentFormValues,
} from '@/components/bandRecruitment/BandRecruitmentForm';
import { usePostBandRecruitments } from '@/libs/api/generated/orval/band-recruitments/band-recruitments';

const BandRecruitmentCreatePage = () => {
  const router = useRouter();
  const { mutateAsync, isPending } = usePostBandRecruitments();

  const handleSubmit = async (values: BandRecruitmentFormValues) => {
    const recruitingParts = values.recruitingParts
      .split(',')
      .map((part) => part.trim())
      .filter(Boolean);

    try {
      const response = await mutateAsync({
        data: {
          title: values.title,
          description: values.description,
          genre: values.genre || undefined,
          location: values.location || undefined,
          recruitingParts,
          skillLevel: values.skillLevel || undefined,
          contact: values.contact,
          deadline: values.deadline
            ? `${values.deadline}T00:00:00Z`
            : undefined,
          status: values.status,
        },
      });

      const newId = response.recruitment?.id;
      toast.success('募集を作成しました');
      if (newId) {
        router.push(`/dashboard/band-recruitments/${newId}`);
      } else {
        router.push('/dashboard/band-recruitments');
      }
    } catch (err) {
      console.error(err);
      toast.error('募集の作成に失敗しました');
    }
  };

  return (
    <div className="mx-auto max-w-3xl space-y-6 p-6">
      <div className="border-b border-gray-200 pb-4">
        <h1 className="text-3xl font-bold text-slate-900">バンド募集を作成</h1>
        <p className="mt-2 text-sm text-slate-600">
          必要事項を入力して新しいバンド募集を投稿しましょう。
        </p>
      </div>
      <BandRecruitmentForm
        submitLabel="募集を作成"
        onSubmit={handleSubmit}
        isSubmitting={isPending}
        onCancel={() => router.push('/dashboard/band-recruitments')}
      />
    </div>
  );
};

export default BandRecruitmentCreatePage;
