import React, { useEffect, useState } from 'react';
import Loading from '../components/Utils/Loading';
import { useQuery } from '@apollo/react-hooks';
import { posts, postsVariables } from '../queries/types/posts';
import { POST_FRONTPAGE } from '../queries/post_queries';
import NotFound from '../components/Utils/NotFound';
import PostsFrontPage from '../components/Posts/PostFrontpage/PostsFrontPage';
import FormNewPost from '../components/Form/FormNewPost';

export default function Home() {
  const [showingForm, setShow] = useState(false);
  const { data, loading, fetchMore, error } = useQuery<posts, postsVariables>(
    POST_FRONTPAGE,
    {
      variables: {
        params: {
          limit: 5
        }
      }
    }
  );
  useEffect(() => {
    if (!data) return;
    const fetchMoreScroll = () => {
      const d = document.documentElement;
      const offset = d.scrollTop + window.innerHeight;
      const height = d.offsetHeight;
      if (offset !== height || loading || !data) {
        return;
      }
      fetchMore({
        variables: {
          params: {
            limit: 5,
            before: data.posts[data.posts.length - 1].id
          }
        },
        updateQuery: (prev: any, { fetchMoreResult }) => {
          if (!fetchMoreResult) return prev;
          return Object.assign({}, prev, {
            posts: [...prev.posts, ...fetchMoreResult.posts]
          });
        }
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
        <button
          className="bg-gray-300 hover:bg-gray-400 text-gray-800 font-bold py-2 px-4 rounded inline-flex items-center"
          onClick={() => setShow(prev => !prev)}
        >
          <svg
            className="fill-current w-4 h-4 mr-2"
            xmlns="http://www.w3.org/2000/svg"
            viewBox="0 0 20 20"
          >
            <path d="M12,8H4A2,2 0 0,0 2,10V14A2,2 0 0,0 4,16H5V20A1,1 0 0,0 6,21H8A1,1 0 0,0 9,20V16H12L17,20V4L12,8M21.5,12C21.5,13.71 20.54,15.26 19,16V8C20.53,8.75 21.5,10.3 21.5,12Z" />
          </svg>
          <span>Create a post</span>
        </button>
      </div>
      {showingForm && <FormNewPost />}
      <PostsFrontPage posts={data.posts} />
    </>
  );
}
