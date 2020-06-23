import React from 'react';

type Props = {
  onClick: (obj: any) => void;
  text: String;
  icon?: React.ReactFragment;
};

export default function Button({ onClick, text, icon }: Props) {
  const iconRender = icon ? icon : null;
  return (
    <button
      className="bg-gray-300 hover:bg-gray-400 text-gray-800 font-bold py-2 px-4 rounded inline-flex items-center"
      onClick={onClick}
    >
      {iconRender}
      <span>{text}</span>
    </button>
  );
}
