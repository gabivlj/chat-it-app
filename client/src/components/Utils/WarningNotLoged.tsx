import React from 'react';
import { CHECK_LOGED_LOCAL } from '../../queries/user_queries';
import { isLogged } from '../../queries/types/isLogged';
import { useQuery } from '@apollo/react-hooks';

export default function WarningNotLoged() {
  const resultIsLogged = useQuery<isLogged>(CHECK_LOGED_LOCAL);
  return (
    <div>
      {(!resultIsLogged.data || !resultIsLogged.data.loged.user) && (
        <div
          className={`rounded-lg p-2 bg-indigo-800 items-center text-indigo-100 leading-none  m-3 max-w-lg h-auto float-right clear-both w-lg sm:w-2/3 xl:w-1/3 md:w-1/3 lg:w-2/3 w-2/3`}
          role="alert"
        >
          <span className="rounded-full bg-indigo-500 uppercase px-2 py-1 text-xs font-bold mr-3 w-auto inline-block">
            Warning!
          </span>
          <span
            style={{ wordWrap: 'break-word' }}
            className="font-semibold mt-3 mr-2 text-left max-w-lg"
          >
            You are not logged in! If you want real-time chat you must be logged
            in!
          </span>
          <svg
            className="fill-current opacity-75 h-4 w-4"
            xmlns="http://www.w3.org/2000/svg"
            viewBox="0 0 20 20"
          ></svg>
        </div>
      )}
    </div>
  );
}
