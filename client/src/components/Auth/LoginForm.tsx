import React, { useEffect, useState } from 'react';
import { logUserVariables, logUser } from '../../queries/types/logUser';
import { useMutation } from '@apollo/react-hooks';
import {
  LOG_USER_MUTATION,
  LOG_USER_LOCAL,
  LOG_USER_LOCAL_SESSION
} from '../../queries/user_queries';
import '../../App.css';
import { LoginError, ParseError } from '../../utils/error';
import Input from '../Inputs/Input';
import ButtonForm from '../Inputs/ButtonForm';

export default function LoginForm() {
  const [inputState, setInputState] = useState({ username: '', password: '' });
  const [errorState, setErrorState] = useState({
    username: '',
    password: ''
  } as LoginError);
  function changeState(e: React.FormEvent<HTMLInputElement>) {
    setInputState({
      ...inputState,
      [e.currentTarget.name]: e.currentTarget.value
    });
    e.persist();
  }
  const [login, result] = useMutation<logUser, logUserVariables>(
    LOG_USER_MUTATION,
    {
      variables: {
        formParameters: {
          username: inputState.username,
          password: inputState.password
        }
      },
      errorPolicy: 'all',
      onError: err => {
        setErrorState(ParseError.parseLoginError(err));
      }
    }
  );
  // Local
  const [logUserLocally] = useMutation(LOG_USER_LOCAL_SESSION, {
    variables: { user: result.data ? result.data.logUser : null }
  });
  useEffect(() => {
    if (result.loading || !result.data) return;
    logUserLocally();
  }, [result.loading, logUserLocally, result.data]);
  return (
    <div className="container flex justify-center">
      <h3 className="mt-64 text-3xl">
        Log in in chat-it, the reddit inspired realtime post chat! (Long
        advertisement)
      </h3>
      <div className="w-full max-w-xsm mt-32">
        <form className="bg-white shadow-md rounded px-8 pt-6 pb-8 mb-4 ">
          <div className="mb-4">
            <Input
              label="Username"
              className={`${errorState.username && `border-red-500`}`}
              placeholder="Username"
              name={'username'}
              onChange={changeState}
              type="text"
              value={inputState.username}
            />
          </div>
          <div className="mb-6">
            <Input
              label="Password"
              className={`${errorState.password && `border-red-500`}`}
              id="password"
              type="password"
              name="password"
              onChange={changeState}
              value={inputState.password}
              placeholder="******************"
            />
            {errorState.password && (
              <p className="text-red-500 text-xs italic">
                {errorState.password}
              </p>
            )}
          </div>
          <div className="flex items-center justify-between">
            <ButtonForm
              onClick={e => {
                e.preventDefault();
                login();
              }}
            >
              Sign In
            </ButtonForm>
            <a
              className="inline-block align-baseline font-bold text-sm text-blue-500 hover:text-blue-800"
              href="#"
            >
              Forgot Password?
            </a>
          </div>
        </form>
        <p className="text-center text-gray-500 text-xs">
          &copy;2020 Chat-It. All rights reserved.
        </p>
      </div>
    </div>
  );
}
