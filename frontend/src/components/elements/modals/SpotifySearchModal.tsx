'use client';

import React from 'react';
import Modal from 'react-modal';
import SelectMusicArea from '@/components/elements/others/SelectMusicArea';
import { Track } from '@/libs/api/generated/orval/model/track';

type SpotifySearchModalProps = {
  isOpen: boolean;
  onClose: () => void;
  onSelectTracks: (tracks: Track[]) => void;
  initialSelectedTracks: Track[];
};

const SpotifySearchModal = ({
  isOpen,
  onClose,
  onSelectTracks,
  initialSelectedTracks,
}: SpotifySearchModalProps): JSX.Element => {
  return (
    <Modal
      isOpen={isOpen}
      onRequestClose={onClose}
      contentLabel="Spotify曲選択"
      className="absolute top-1/2 left-1/2 -translate-x-1/2 -translate-y-1/2 bg-white p-6 rounded-lg shadow-lg max-w-md w-full outline-none overflow-auto"
      overlayClassName="fixed inset-0 bg-black bg-opacity-75 z-50"
    >
      <div className="p-6 bg-white rounded-lg shadow-lg max-w-md mx-auto my-20">
        <h2 className="text-2xl font-bold mb-4">Spotifyから曲を選択</h2>
        <SelectMusicArea
          onSelect={onSelectTracks}
          initialSelectedTracks={initialSelectedTracks}
        />
        <div className="flex justify-end mt-4">
          <button className="px-4 py-2 bg-gray-300 rounded" onClick={onClose}>
            閉じる
          </button>
        </div>
      </div>
    </Modal>
  );
};

export default SpotifySearchModal;
