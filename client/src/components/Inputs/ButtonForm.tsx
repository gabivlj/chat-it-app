import React from 'react';

type Props = {
  onClick: (e: React.MouseEvent<HTMLButtonElement, MouseEvent>) => void;

  children: React.ReactNode;
};

export default function ButtonForm({ onClick, children }: Props) {
  return (
    <button
      className="bg-blue-500 hover:bg-blue-700 text-white font-bold py-2 px-4 rounded focus:outline-none focus:shadow-outline"
      type="button"
      onClick={onClick}
    >
      {children}
    </button>
  );
}
