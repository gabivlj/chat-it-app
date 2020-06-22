import React from 'react';

type Props = {
  onChange: (e: any) => void;
  uploadImageText?: String;
};

export default function UploadFile({ onChange, uploadImageText }: Props) {
  if (uploadImageText === undefined || uploadImageText === null) {
    uploadImageText = 'Upload image';
  }
  return (
    <div className="mt-3 flex flex-col items-center">
      <label className="w-32 h-1/4 flex mr-32 flex-col items-center bg-gray-200 text-blue rounded-lg shadow-lg tracking-wide border border-blue cursor-pointer hover:bg-white">
        <svg
          className="w-1/6 h-1/6"
          fill="currentColor"
          xmlns="http://www.w3.org/2000/svg"
          viewBox="0 0 20 20"
        >
          <path
            fill="currentColor"
            d="M9,16V10H5L12,3L19,10H15V16H9M5,20V18H19V20H5Z"
          />
        </svg>
        <span className="mt-2 text-sm">{uploadImageText}</span>
        <input type="file" className="hidden" onChange={onChange} />
      </label>
    </div>
  );
}
