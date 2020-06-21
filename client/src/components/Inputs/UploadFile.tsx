import React from 'react';

type Props = {
  onChange: (e: any) => void;
};

export default function UploadFile({ onChange }: Props) {
  return (
    <div className="mt-3 flex flex-col items-center">
      <label className="w-32 h-1/4 flex mr-32 flex-col items-center bg-white text-blue rounded-lg shadow-lg tracking-wide border border-blue cursor-pointer hover:bg-green-200">
        <svg
          className="w-1/4 h-1/4"
          fill="currentColor"
          xmlns="http://www.w3.org/2000/svg"
          viewBox="0 0 20 20"
        >
          <path d="M16.88 9.1A4 4 0 0 1 16 17H5a5 5 0 0 1-1-9.9V7a3 3 0 0 1 4.52-2.59A4.98 4.98 0 0 1 17 8c0 .38-.04.74-.12 1.1zM11 11h3l-4-4-4 4h3v3h2v-3z" />
        </svg>
        <span className="mt-2 text-sm">Upload image</span>
        <input type="file" className="hidden" onChange={onChange} />
      </label>
    </div>
  );
}
