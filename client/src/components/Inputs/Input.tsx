import React from 'react';

type Props = {
  className: string;
  onChange: (e: React.FormEvent<HTMLInputElement>) => void;
  value: string;
  name?: string;
  id?: string | undefined;
  placeholder?: string | undefined;
  type: string;
  label: string;
};

export default function Input({
  className,
  onChange,
  value,
  placeholder = '',
  id = '',
  type,
  name,
  label
}: Props) {
  return (
    <>
      <label className="block text-gray-700 text-sm font-bold mb-2">
        {label}
      </label>
      <input
        className={`shadow appearance-none border ${className} rounded w-full py-2 px-3 text-gray-700 leading-tight focus:outline-none focus:shadow-outline`}
        id={id}
        type={type}
        placeholder={placeholder}
        name={name || ''}
        onChange={onChange}
        value={value}
      />
    </>
  );
}
