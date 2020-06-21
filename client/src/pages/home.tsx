import React, { useEffect } from 'react';
import Loading from '../components/Utils/Loading';
import { useQuery } from '@apollo/react-hooks';
import { posts, postsVariables } from '../queries/types/posts';
import { POST_FRONTPAGE } from '../queries/post_queries';
import NotFound from '../components/Utils/NotFound';
import PostsFrontPage from '../components/Posts/PostFrontpage/PostsFrontPage';

export default function Home() {
  const { data, loading, fetchMore, error } = useQuery<posts, postsVariables>(
    POST_FRONTPAGE,
    {
      variables: {
        params: {
          limit: 10
        }
      }
    }
  );
  useEffect(() => {
    if (!data) return;
    const fetchMoreScroll = () => {
      if (!data) return;
      // Check if we hitted bottom
      if (
        window.innerHeight + window.pageYOffset <
        document.body.offsetHeight - 2
      ) {
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
  return <PostsFrontPage posts={data.posts} />;
}
