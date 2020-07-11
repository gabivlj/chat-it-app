import React, { useState, useEffect } from 'react';

export const MediaQueries = {
  Small: '(min-width: 400px)',
  Medium: '(min-width: 800px)',
  Large: '(min-width: 1200px)',
};

export function useMediaQuery(query: string) {
  const [match, updateMatch] = useState(window.matchMedia(query).matches);
  useEffect(() => {
    const updateQuery = () => {
      updateMatch(window.matchMedia(query).matches);
    };
    window.addEventListener('resize', updateQuery);
    window.addEventListener('change', updateQuery);
    return () => {
      window.removeEventListener('resize', updateQuery);
      window.removeEventListener('change', updateQuery);
    };
  }, [query]);
  return match;
}
