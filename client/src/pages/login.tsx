import React from 'react';
import { useQuery } from '@apollo/react-hooks';
import { isLogged } from '../queries/types/isLogged';
import { Redirect } from 'react-router';
import LoginForm from '../components/Auth/LoginForm';
import { CHECK_LOGED_LOCAL } from '../queries/user_queries';

export default function Login() {
  // Check in the client data if he is loged
  const { data } = useQuery<isLogged>(CHECK_LOGED_LOCAL);
  if (data && data.loged && data.loged.user) {
    return <Redirect to={`/user/${data.loged.user.username}`} />;
  }
  return (
    <div>
      <LoginForm />
    </div>
  );
}
