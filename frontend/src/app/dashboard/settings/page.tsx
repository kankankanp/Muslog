'use client';

import { useState, useEffect } from 'react';
import toast from 'react-hot-toast';
import { useGetMe } from '@/libs/api/generated/orval/auth/auth';
import {
  useGetUsersMeSettings,
  usePutUsersMeSettings,
} from '@/libs/api/generated/orval/users/users';

export default function SettingsPage() {
  const [editorType, setEditorType] = useState<'markdown' | 'wysiwyg'>(
    'markdown'
  );

  const { data: userData } = useGetMe();
  const {
    data: settingsData,
    isLoading: isSettingsLoading,
    error: settingsError,
  } = useGetUsersMeSettings();
  const { mutate: updateSettings, isPending: isUpdating } =
    usePutUsersMeSettings();

  useEffect(() => {
    if (settingsError) {
      console.error('設定の取得に失敗:', settingsError);
    }
  }, [settingsError]);

  useEffect(() => {
    if (settingsData?.setting?.editorType) {
      setEditorType(settingsData.setting.editorType as 'markdown' | 'wysiwyg');
    }
  }, [settingsData]);

  const handleEditorTypeChange = (newEditorType: 'markdown' | 'wysiwyg') => {
    setEditorType(newEditorType);
    updateSettings(
      { data: { editorType: newEditorType } },
      {
        onSuccess: () => {
          toast.success('設定を更新しました');
        },
        onError: (error) => {
          console.error('設定の更新に失敗しました:', error);
          toast.error('設定の更新に失敗しました');
          // エラー時に元の値に戻す
          if (settingsData?.setting?.editorType) {
            setEditorType(
              settingsData.setting.editorType as 'markdown' | 'wysiwyg'
            );
          }
        },
      }
    );
  };

  if (isSettingsLoading) {
    return (
      <div className="flex items-center justify-center min-h-screen">
        <div className="text-lg">読み込み中...</div>
      </div>
    );
  }

  return (
    <>
      <div className="border-gray-100 border-b-2 bg-white px-8 py-6">
        <div className="flex items-center gap-3">
          <h1 className="text-3xl font-bold">設定</h1>
        </div>
      </div>

      <div className="p-8 max-w-4xl mx-auto">
        <div className="bg-white rounded-lg shadow-sm border border-gray-200">
          <div className="p-6">
            <h2 className="text-xl font-semibold mb-4">ユーザー情報</h2>
            <div className="space-y-3 text-gray-600">
              <div>
                <span className="font-medium">ユーザー名:</span>{' '}
                {userData?.name}
              </div>
              <div>
                <span className="font-medium">メールアドレス:</span>{' '}
                {userData?.email}
              </div>
            </div>
          </div>
        </div>

        <div className="bg-white rounded-lg shadow-sm border border-gray-200 mt-6">
          <div className="p-6">
            <h2 className="text-xl font-semibold mb-4">エディタ設定</h2>
            <div className="space-y-4">
              <div>
                <label className="block text-sm font-medium text-gray-700 mb-3">
                  記事作成時のエディタタイプ
                </label>
                <div className="flex gap-4">
                  <button
                    className={`px-6 py-3 rounded-lg border transition-colors ${
                      editorType === 'markdown'
                        ? 'bg-blue-50 border-blue-500 text-blue-700'
                        : 'bg-white border-gray-300 text-gray-700 hover:bg-gray-50'
                    } ${isUpdating ? 'opacity-50 cursor-not-allowed' : ''}`}
                    onClick={() => handleEditorTypeChange('markdown')}
                    disabled={isUpdating}
                  >
                    <div className="text-center">
                      <div className="font-medium">Markdownエディタ</div>
                      <div className="text-sm mt-1">
                        マークダウン記法でテキスト編集
                      </div>
                    </div>
                  </button>
                  <button
                    className={`px-6 py-3 rounded-lg border transition-colors ${
                      editorType === 'wysiwyg'
                        ? 'bg-blue-50 border-blue-500 text-blue-700'
                        : 'bg-white border-gray-300 text-gray-700 hover:bg-gray-50'
                    } ${isUpdating ? 'opacity-50 cursor-not-allowed' : ''}`}
                    onClick={() => handleEditorTypeChange('wysiwyg')}
                    disabled={isUpdating}
                  >
                    <div className="text-center">
                      <div className="font-medium">WYSIWYGエディタ</div>
                      <div className="text-sm mt-1">
                        ビジュアルエディタで直感的に編集
                      </div>
                    </div>
                  </button>
                </div>
              </div>
              {isUpdating && (
                <div className="text-sm text-gray-500">設定を更新中...</div>
              )}
              {settingsError ? (
                <div className="text-sm text-red-500">
                  設定の読み込みでエラーが発生しました。デフォルト設定で表示しています。
                </div>
              ) : null}
            </div>
          </div>
        </div>

        <div className="bg-white rounded-lg shadow-sm border border-gray-200 mt-6">
          <div className="p-6">
            <h2 className="text-xl font-semibold mb-4">アカウント設定</h2>
            <div className="text-gray-600">
              <p className="mb-4">
                メールアドレスやパスワードの変更機能は今後のアップデートで追加予定です。
              </p>
              <div className="text-sm text-gray-500">
                ※ 現在はエディタタイプの設定のみ利用可能です
              </div>
            </div>
          </div>
        </div>
      </div>
    </>
  );
}
