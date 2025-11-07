import { useEffect } from 'react';
import { useGetUsersMeSettings } from '@/libs/api/generated/orval/users/users';

export const useUserSettings = () => {
  const { data: settingsData, isLoading, error } = useGetUsersMeSettings();

  const editorType = settingsData?.setting?.editorType || 'markdown';

  useEffect(() => {
    if (error) {
      console.error('設定の取得に失敗:', error);
    }
  }, [error]);

  return {
    editorType: editorType as 'markdown' | 'wysiwyg',
    isLoading,
    error,
    settings: settingsData?.setting,
  };
};
