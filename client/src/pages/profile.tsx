import React from 'react';
import { useQuery } from '@apollo/react-hooks';
import { GetUser, GetUserVariables } from '../queries/types/GetUser';
import { GET_USER } from '../queries/user_queries';
import { TODO } from '../utils/todo';
import ProfileCard from '../components/Profile/ProfileCard';

type Props = {
  match: {
    params: {
      username: string;
    };
  };
};

/**
 * This component at the moment is just
 * testing the profile queries and that they work with Apollo.
 */
export default function Profile({ match }: Props) {
  const { data, loading, error } = useQuery<GetUser, GetUserVariables>(
    GET_USER,
    {
      variables: {
        query: {
          username: match.params.username,
          id: null
        }
      },
      onError: err => {
        TODO(`Handle error ${err}`);
      }
    }
  );
  if (!data) {
    return <div className="container">Loading...</div>;
  }
  return (
    <div className="container">
      {loading ? <>Loading...</> : <ProfileCard user={data.user}></ProfileCard>}
    </div>
  );
}
