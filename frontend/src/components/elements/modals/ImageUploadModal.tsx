import React, { useState, useCallback, useRef } from "react";
import { X, UploadCloud } from "lucide-react";
import Image from "next/image";

interface ImageUploadModalProps {
  isOpen: boolean;
  onClose: () => void;
  onImageUpload: (file: File) => void;
  currentImageUrl?: string; // Optional: for displaying current image in edit mode
}

const ImageUploadModal: React.FC<ImageUploadModalProps> = ({
  isOpen,
  onClose,
  onImageUpload,
  currentImageUrl,
}) => {
  const [dragActive, setDragActive] = useState(false);
  const [selectedFile, setSelectedFile] = useState<File | null>(null);
  const inputRef = useRef<HTMLInputElement>(null);

  const handleDrag = useCallback((e: React.DragEvent<HTMLDivElement>) => {
    e.preventDefault();
    e.stopPropagation();
    if (e.type === "dragenter" || e.type === "dragover") {
      setDragActive(true);
    } else if (e.type === "dragleave") {
      setDragActive(false);
    }
  }, []);

  const handleDrop = useCallback((e: React.DragEvent<HTMLDivElement>) => {
    e.preventDefault();
    e.stopPropagation();
    setDragActive(false);
    if (e.dataTransfer.files && e.dataTransfer.files[0]) {
      setSelectedFile(e.dataTransfer.files[0]);
    }
  }, []);

  const handleChange = useCallback((e: React.ChangeEvent<HTMLInputElement>) => {
    e.preventDefault();
    if (e.target.files && e.target.files[0]) {
      setSelectedFile(e.target.files[0]);
    }
  }, []);

  const handleUpload = () => {
    if (selectedFile) {
      onImageUpload(selectedFile);
      setSelectedFile(null); // Clear selected file after upload
      onClose();
    } else {
      alert("ファイルを選択してください。");
    }
  };

  if (!isOpen) return null;

  const imageUrl = selectedFile ? URL.createObjectURL(selectedFile) : currentImageUrl;

  return (
    <div className="fixed inset-0 bg-black bg-opacity-50 flex items-center justify-center z-50">
      <div className="bg-white dark:bg-gray-800 p-6 rounded-lg shadow-lg w-full max-w-md relative">
        <button
          onClick={onClose}
          className="absolute top-3 right-3 text-gray-500 hover:text-gray-700 dark:text-gray-400 dark:hover:text-gray-200"
        >
          <X size={24} />
        </button>
        <h2 className="text-2xl font-bold mb-4 text-gray-900 dark:text-gray-100">画像を追加・変更</h2>

        <div
          className={`border-2 border-dashed rounded-lg p-6 text-center transition-colors ${dragActive ? "border-blue-500 bg-blue-50 dark:bg-blue-900" : "border-gray-300 dark:border-gray-600 bg-gray-50 dark:bg-gray-700"}`}
          onDragEnter={handleDrag}
          onDragLeave={handleDrag}
          onDragOver={handleDrag}
          onDrop={handleDrop}
          onClick={() => inputRef.current?.click()}
        >
          <input
            type="file"
            ref={inputRef}
            className="hidden"
            onChange={handleChange}
            accept="image/*"
          />
          {imageUrl ? (
            <div className="relative w-full h-48 mb-4 rounded-md overflow-hidden">
              <Image src={imageUrl} alt="Preview" layout="fill" objectFit="contain" />
            </div>
          ) : (
            <div className="flex flex-col items-center justify-center text-gray-500 dark:text-gray-400">
              <UploadCloud size={48} className="mb-2" />
              <p className="mb-1">画像をドラッグ&ドロップ</p>
              <p className="text-sm">またはクリックしてファイルを選択</p>
            </div>
          )}
        </div>

        <button
          onClick={handleUpload}
          className="mt-4 w-full px-4 py-2 bg-blue-600 text-white rounded-md hover:bg-blue-700 transition-colors"
        >
          アップロード
        </button>
      </div>
    </div>
  );
};

export default ImageUploadModal;
