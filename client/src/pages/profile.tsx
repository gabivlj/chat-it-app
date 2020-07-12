import React, { useEffect } from 'react';
import { useQuery } from '@apollo/react-hooks';
import { GetUser, GetUserVariables } from '../queries/types/GetUser';
import { GET_USER } from '../queries/user_queries';
import { TODO } from '../utils/todo';
import ProfileCard from '../components/Profile/ProfileCard';
import useIsBottom from '../utils/useIsBottom';
import Loading from '../components/Utils/Loading';

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
  const isBottom = useIsBottom();
  const { data, loading, error, fetchMore } = useQuery<
    GetUser,
    GetUserVariables
  >(GET_USER, {
    variables: {
      query: {
        username: match.params.username,
        id: null,
      },
      paramsPost: {
        limit: 5,
      },
      paramsComments: {
        limit: 5,
      },
    },
    onError: (err) => {
      TODO(`Handle error ${err}`);
    },
  });
  useEffect(() => {
    if (!data) return;
    if (!isBottom) return;
    const { postsUser, commentsUser } = data.user;
    fetchMore({
      variables: {
        paramsPost: {
          limit: 5,
          before: postsUser[postsUser.length - 1].id,
        },
        paramsComments: {
          limit: 5,
          before: commentsUser[commentsUser.length - 1].id,
        },
      },
      updateQuery: (prev: any, { fetchMoreResult }) => {
        if (!fetchMoreResult) return prev;
        return Object.assign({}, prev, {
          user: {
            ...prev.user,
            postsUser: [
              ...prev.user.postsUser,
              ...fetchMoreResult.user.postsUser,
            ],
            commentsUser: [
              ...prev.user.commentsUser,
              ...fetchMoreResult.user.commentsUser,
            ],
          },
        });
      },
    });
  }, [isBottom, data]);
  if (!data) {
    return (
      <div className="container">
        <Loading />
      </div>
    );
  }
  return (
    <div className="container">
      {loading ? (
        <>
          <Loading />
        </>
      ) : (
        <ProfileCard user={data.user}></ProfileCard>
      )}
    </div>
  );
}
