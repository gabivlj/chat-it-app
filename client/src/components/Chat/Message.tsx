import React from 'react';
import { getPostAndMessages_messagesPost } from '../../queries/types/getPostAndMessages';
import { GetUser_user } from '../../queries/types/GetUser';
import { isLogged_loged_user } from '../../queries/types/isLogged';

type Props = {
  message: getPostAndMessages_messagesPost;
  currentUser?: isLogged_loged_user | null;
};

export default function Message({ message, currentUser }: Props) {
  const float =
    currentUser && message.user.username === currentUser.username
      ? 'float-right'
      : 'float-left';
  return (
    <>
      <div
        className={`rounded-md p-2 bg-indigo-800 items-center text-indigo-100 leading-none  m-3 max-w-lg h-auto ${float} clear-both w-lg sm:w-2/3 xl:w-1/3 md:w-1/3 lg:w-2/3 w-2/3`}
        role="alert"
      >
        <span className="rounded-full bg-indigo-500 uppercase px-2 py-1 text-xs font-bold mr-3 w-auto inline-block">
          {message.user.username}
        </span>
        <span
          style={{ wordWrap: 'break-word' }}
          className="font-semibold mt-3 mr-2 text-left max-w-lg"
        >
          {message.text}
        </span>
        <svg
          className="fill-current opacity-75 h-4 w-4"
          xmlns="http://www.w3.org/2000/svg"
          viewBox="0 0 20 20"
        ></svg>
      </div>
    </>
  );
}
