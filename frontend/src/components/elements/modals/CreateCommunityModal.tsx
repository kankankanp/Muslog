import { zodResolver } from "@hookform/resolvers/zod";
import React from "react";
import { useForm } from "react-hook-form";
import toast from "react-hot-toast";
import Modal from "react-modal";
import * as z from "zod";
import { usePostCommunities } from "@/libs/api/generated/orval/communities/communities";

const createCommunitySchema = z.object({
  name: z.string().min(1, "Community name is required"),
  description: z.string().min(1, "Description is required"),
});

type CreateCommunityFormData = z.infer<typeof createCommunitySchema>;

interface CreateCommunityModalProps {
  isOpen: boolean;
  onClose: () => void;
  onCommunityCreated?: () => void;
}

const CreateCommunityModal: React.FC<CreateCommunityModalProps> = ({
  isOpen,
  onClose,
  onCommunityCreated,
}) => {
  const {
    register,
    handleSubmit,
    reset,
    formState: { errors },
  } = useForm<CreateCommunityFormData>({
    resolver: zodResolver(createCommunitySchema),
  });

  const createCommunityMutation = usePostCommunities();

  const onSubmit = async (data: CreateCommunityFormData) => {
    try {
      await createCommunityMutation.mutateAsync({ data });
      toast.success("Community created successfully!");
      reset();
      onCommunityCreated?.();
      onClose(); // Close modal on success
    } catch (error) {
      toast.error("Failed to create community.");
      console.error("Error creating community:", error);
    }
  };

  return (
    <Modal
      isOpen={isOpen}
      onRequestClose={onClose}
      contentLabel="Create Community"
      className="absolute top-1/2 left-1/2 -translate-x-1/2 -translate-y-1/2 bg-white p-6 rounded-lg shadow-lg max-w-md w-full outline-none overflow-auto" // Updated
      overlayClassName="fixed inset-0 bg-black bg-opacity-75 z-50" // Updated
      ariaHideApp={false} // ← 追加
    >
      <div className="p-6 border rounded-lg shadow-md bg-white dark:bg-gray-800">
        <h2 className="text-2xl font-bold text-gray-900 dark:text-white mb-4">
          コミュニティを作成する
        </h2>
        <form onSubmit={handleSubmit(onSubmit)} className="space-y-4">
          <div>
            <label
              htmlFor="name"
              className="block text-sm font-medium text-gray-700 dark:text-gray-300"
            >
              名前
            </label>
            <input
              type="text"
              id="name"
              {...register("name")}
              className="mt-1 block w-full px-3 py-2 border border-gray-300 rounded-md shadow-sm focus:outline-none focus:ring-indigo-500 focus:border-indigo-500 sm:text-sm dark:bg-gray-700 dark:border-gray-600 dark:text-white"
            />
            {errors.name && (
              <p className="mt-1 text-sm text-red-600">{errors.name.message}</p>
            )}
          </div>
          <div>
            <label
              htmlFor="description"
              className="block text-sm font-medium text-gray-700 dark:text-gray-300"
            >
              説明
            </label>
            <textarea
              id="description"
              {...register("description")}
              rows={3}
              className="mt-1 block w-full px-3 py-2 border border-gray-300 rounded-md shadow-sm focus:outline-none focus:ring-indigo-500 focus:border-indigo-500 sm:text-sm dark:bg-gray-700 dark:border-gray-600 dark:text-white"
            ></textarea>
            {errors.description && (
              <p className="mt-1 text-sm text-red-600">
                {errors.description.message}
              </p>
            )}
          </div>
          <button
            type="submit"
            className="mx-auto flex justify-center py-2 px-4 border border-transparent rounded-md shadow-sm text-sm font-medium text-white bg-indigo-600 hover:bg-indigo-700 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-indigo-500 dark:bg-indigo-500 dark:hover:bg-indigo-600"
          >
            作成する
          </button>
          <button
            type="button"
            onClick={onClose}
            className="ml-2 py-2 px-4 border border-gray-300 rounded-md shadow-sm text-sm font-medium text-gray-700 bg-white hover:bg-gray-50 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-indigo-500 dark:bg-gray-700 dark:border-gray-600 dark:text-white dark:hover:bg-gray-600"
          >
            キャンセル
          </button>
        </form>
      </div>
    </Modal>
  );
};

export default CreateCommunityModal;
