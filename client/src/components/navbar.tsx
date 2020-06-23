import React, { useState, useEffect } from 'react';
import { isLogged_loged_user, isLogged } from '../queries/types/isLogged';
import { Link } from 'react-router-dom';
import { CHECK_LOGED_LOCAL, LOG_USER_LOCAL } from '../queries/user_queries';
import { useQuery, useMutation } from '@apollo/react-hooks';

export default function Navbar() {
  const [logUser] = useMutation(LOG_USER_LOCAL, {
    variables: { user: null }
  });
  const { data } = useQuery<isLogged>(CHECK_LOGED_LOCAL);
  const [showing, setShowed] = useState(false);
  useEffect(() => {
    if (document.body.clientWidth >= 1000) {
      setShowed(true);
    }
    let interval = setInterval(() => {
      if (document.body.clientWidth >= 1000 && !showing) setShowed(true);
    }, 500);
    return () => clearInterval(interval);
  }, []);

  return (
    <div>
      <nav className="flex items-center justify-between flex-wrap p-6 bg-teal-900">
        <div className="flex items-center flex-shrink-0 text-white mr-6">
          <svg className="fill-current h-8 w-8 mr-2" viewBox="0 0 24 24">
            <path
              fill="currentColor"
              d="M23 17V19H15V17H23M12 3C17.5 3 22 6.58 22 11C22 11.58 21.92 12.14 21.78 12.68C20.95 12.25 20 12 19 12C15.69 12 13 14.69 13 18L13.08 18.95L12 19C10.76 19 9.57 18.82 8.47 18.5C5.55 21 2 21 2 21C4.33 18.67 4.7 17.1 4.75 16.5C3.05 15.07 2 13.14 2 11C2 6.58 6.5 3 12 3Z"
            />
          </svg>
          <span className="font-semibold text-xl tracking-tight">Chat-it</span>
        </div>
        <div className="block lg:hidden">
          <button
            className="flex items-center px-3 py-2 border rounded text-teal-200 border-teal-400 hover:text-white hover:border-white"
            onClick={() => setShowed(prev => !prev)}
          >
            <svg
              className="fill-current h-3 w-3"
              viewBox="0 0 20 20"
              xmlns="http://www.w3.org/2000/svg"
            >
              <title>Home</title>
              <path d="M0 3h20v2H0V3zm0 6h20v2H0V9zm0 6h20v2H0v-2z" />
            </svg>
          </button>
        </div>
        {showing && (
          <div className="w-full block flex-grow lg:flex lg:items-center lg:w-auto">
            <div className="text-sm lg:flex-grow">
              <Link
                to="/"
                className="block mt-4 lg:inline-block lg:mt-0 text-teal-200 hover:text-white mr-4"
              >
                Home
              </Link>
              <a
                href="https://github.com/gabivlj/chat-it-app"
                className="block mt-4 lg:inline-block lg:mt-0 text-teal-200 hover:text-white mr-4"
              >
                Github
              </a>
              <Link
                to="/about"
                className="block mt-4 lg:inline-block lg:mt-0 text-teal-200 hover:text-white mr-4"
              >
                About
              </Link>
              {data && data.loged.user && (
                <button
                  onClick={() => {
                    localStorage.clear();
                    logUser();
                  }}
                  className="block mt-4 lg:inline-block lg:mt-0 text-teal-200 hover:text-white"
                >
                  Log out
                </button>
              )}
            </div>
            <div>
              <Link
                to={`${
                  data && data.loged.user
                    ? `/user/${data.loged.user.username}`
                    : `/login`
                }`}
                className="inline-block text-sm px-4 py-2 leading-none border rounded text-white border-white hover:border-transparent hover:text-teal-500 hover:bg-white mt-4 lg:mt-0"
              >
                {data && data.loged.user
                  ? `${data.loged.user.username}`
                  : `Login`}
              </Link>
              {(!data || !data.loged.user) && (
                <Link
                  to={`/signup`}
                  className="ml-2 inline-block text-sm px-4 py-2 leading-none border rounded text-white border-white hover:border-transparent hover:text-teal-500 hover:bg-white mt-4 lg:mt-0"
                >
                  Signup
                </Link>
              )}
            </div>
          </div>
        )}
      </nav>
    </div>
  );
}
