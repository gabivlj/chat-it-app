import React from 'react';
import { useQuery } from '@apollo/react-hooks';
import { GetUser, GetUserVariables } from '../queries/types/GetUser';
import { GET_USER } from '../queries/user_queries';
import { TODO } from '../utils/todo';

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
        // todo: Handle
        TODO(`Handle error ${err}`);
      }
    }
  );
  let posts = data && data.user.posts;
  const user = data && data.user.username;
  return (
    <div className="container">
      <h2 className="text-xl">Username {user}</h2>
      {posts &&
        posts.map(post => (
          <div key={post.title}>
            {post.title} said {post.text}
            <img src={post.image ? post.image.urlXL : ''}></img>
          </div>
        ))}
    </div>
  );
}
