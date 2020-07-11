import React, { useEffect, useState } from 'react';
import Loading from '../components/Utils/Loading';
import { useQuery } from '@apollo/react-hooks';
import { posts, postsVariables } from '../queries/types/posts';
import { POST_FRONTPAGE } from '../queries/post_queries';
import NotFound from '../components/Utils/NotFound';
import PostsFrontPage from '../components/Posts/PostFrontpage/PostsFrontPage';
import FormNewPost from '../components/Form/FormNewPost';
import { isLogged } from '../queries/types/isLogged';
import { CHECK_LOGED_LOCAL } from '../queries/user_queries';
import Button from '../components/Inputs/Button';
import Communication from '../components/Icons/Communication';
import { withRouter, RouteComponentProps } from 'react-router';
import PersonPlus from '../components/Icons/PersonPlus';
import PersonKey from '../components/Icons/PersonKey';
import canFetchMore from '../utils/canFetchMore';

function Home({ history }: RouteComponentProps) {
  const logedData = useQuery<isLogged>(CHECK_LOGED_LOCAL);
  const [showingForm, setShow] = useState(false);
  const { data, loading, fetchMore, error } = useQuery<posts, postsVariables>(
    POST_FRONTPAGE,
    {
      variables: {
        params: {
          limit: 5,
        },
      },
    }
  );
  useEffect(() => {
    if (!data) return;
    const fetchMoreScroll = () => {
      if (!canFetchMore({ loading, data })) {
        return;
      }
      fetchMore({
        variables: {
          params: {
            limit: 5,
            before: data.posts[data.posts.length - 1].id,
          },
        },
        updateQuery: (prev: any, { fetchMoreResult }) => {
          if (!fetchMoreResult) return prev;
          return Object.assign({}, prev, {
            posts: [...prev.posts, ...fetchMoreResult.posts],
          });
        },
      });
    };
    window.addEventListener('scroll', fetchMoreScroll);
    return () => {
      window.removeEventListener('scroll', fetchMoreScroll);
    };
  }, [data]);
  if (loading) return <Loading />;
  if (!data || error) return <NotFound />;
  return (
    <>
      <div className="container mx-auto flex items-center flex-wrap pt-4 pb-12">
        {logedData.data && logedData.data.loged.user ? (
          <Button
            icon={<Communication />}
            text="Create a post"
            onClick={() => setShow((prev) => !prev)}
          />
        ) : (
          <>
            <label className="block text-gray-700 text-sm font-bold mb-2 mr-2">
              You are not loged, you cannot post! Consider the following:
            </label>
            <Button
              icon={<PersonKey />}
              text="Login"
              onClick={() => history.push('/login')}
            />
            <div className="ml-2"></div>
            <Button
              icon={<PersonPlus />}
              text="SignUp"
              onClick={() => history.push('/login')}
            />
          </>
        )}
      </div>
      {showingForm && logedData.data && logedData.data.loged.user ? (
        <FormNewPost />
      ) : (
        ''
      )}
      <PostsFrontPage posts={data.posts} />
    </>
  );
}

export default withRouter(Home);
