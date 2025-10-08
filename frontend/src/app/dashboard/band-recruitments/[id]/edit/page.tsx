"use client";

import { useParams, useRouter } from "next/navigation";
import { useMemo } from "react";
import toast from "react-hot-toast";
import BandRecruitmentForm, {
  type BandRecruitmentFormValues,
} from "@/components/bandRecruitment/BandRecruitmentForm";
import Spinner from "@/components/layouts/Spinner";
import {
  useGetBandRecruitmentsId,
  usePutBandRecruitmentsId,
} from "@/libs/api/generated/orval/band-recruitments/band-recruitments";

const toDateInputValue = (value?: string | null) => {
  if (!value) return "";
  const date = new Date(value);
  if (Number.isNaN(date.getTime())) {
    // value might already be yyyy-mm-dd
    return value.split("T")[0];
  }
  const year = date.getFullYear();
  const month = String(date.getMonth() + 1).padStart(2, "0");
  const day = String(date.getDate()).padStart(2, "0");
  return `${year}-${month}-${day}`;
};

const BandRecruitmentEditPage = () => {
  const params = useParams<{ id: string }>();
  const router = useRouter();
  const id = params?.id;

  const { data, isLoading, isError } = useGetBandRecruitmentsId(id, {
    query: {
      enabled: !!id,
    },
  });

  const { mutateAsync, isPending } = usePutBandRecruitmentsId();

  const initialValues: BandRecruitmentFormValues | undefined = useMemo(() => {
    const recruitment = data?.recruitment;
    if (!recruitment) return undefined;
    return {
      title: recruitment.title ?? "",
      description: recruitment.description ?? "",
      genre: recruitment.genre ?? "",
      location: recruitment.location ?? "",
      recruitingParts: recruitment.recruitingParts?.join(", ") ?? "",
      skillLevel: recruitment.skillLevel ?? "",
      contact: recruitment.contact ?? "",
      deadline: toDateInputValue(recruitment.deadline),
      status: recruitment.status ?? "open",
    };
  }, [data]);

  const handleSubmit = async (values: BandRecruitmentFormValues) => {
    if (!id) return;
    const recruitingParts = values.recruitingParts
      .split(",")
      .map((part) => part.trim())
      .filter(Boolean);

    try {
      const response = await mutateAsync({
        id,
        data: {
          title: values.title,
          description: values.description,
          genre: values.genre || undefined,
          location: values.location || undefined,
          recruitingParts,
          skillLevel: values.skillLevel || undefined,
          contact: values.contact,
          deadline: values.deadline ? `${values.deadline}T00:00:00Z` : undefined,
          status: values.status,
        },
      });

      toast.success("募集内容を更新しました");
      router.push(`/dashboard/band-recruitments/${response.recruitment?.id ?? id}`);
    } catch (err) {
      console.error(err);
      toast.error("募集内容の更新に失敗しました");
    }
  };

  if (isLoading || !initialValues) {
    return <Spinner />;
  }

  if (isError) {
    return (
      <div className="p-8 text-center text-red-600">
        募集情報の取得に失敗しました。
      </div>
    );
  }

  return (
    <div className="mx-auto max-w-3xl space-y-6 p-6">
      <div className="border-b border-gray-200 pb-4">
        <h1 className="text-3xl font-bold text-slate-900">募集内容を編集</h1>
        <p className="mt-2 text-sm text-slate-600">
          情報を更新し保存すると、即座に反映されます。
        </p>
      </div>
      <BandRecruitmentForm
        initialValues={initialValues}
        submitLabel="変更を保存"
        onSubmit={handleSubmit}
        isSubmitting={isPending}
        onCancel={() => router.push(`/dashboard/band-recruitments/${id}`)}
      />
    </div>
  );
};

export default BandRecruitmentEditPage;
