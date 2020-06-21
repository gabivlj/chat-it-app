import React, { useState } from 'react';
import { useMutation } from '@apollo/react-hooks';
import { SEND_MESSAGE } from '../../queries/post_queries';
import {
  sendMessage,
  sendMessageVariables
} from '../../queries/types/sendMessage';

type Props = {
  postId: string;
};

export default function InputChat({ postId }: Props) {
  const [input, setInput] = useState('');
  const [mutate] = useMutation<sendMessage, sendMessageVariables>(SEND_MESSAGE);

  function onSubmit(e: React.FormEvent<HTMLFormElement>) {
    const text = input.trim();
    if (!text.length) {
      setInput('');
      return;
    }
    e.preventDefault();
    mutate({ variables: { text, postId } });
    setInput('');
  }

  return (
    <div>
      <form onSubmit={onSubmit}>
        <div className="flex items-center border-b border-b-2 border-teal-500 py-2">
          <input
            className="appearance-none bg-transparent border-none w-full text-gray-700 mr-3 py-1 px-2 leading-tight focus:outline-none"
            type="text"
            value={input}
            onChange={e => {
              setInput(e.target.value);
            }}
          />

          <input
            className="flex-shrink-0 border-transparent border-4 text-teal-500 hover:text-teal-800 text-sm py-1 px-2 rounded"
            type="submit"
            value="Send"
          ></input>
        </div>
      </form>
    </div>
  );
}
