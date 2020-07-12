import { useState, useEffect } from 'react';

export default function useIsBottom() {
  const [isBottom, setIsBottom] = useState(false);
  useEffect(() => {
    const checker = () => {
      const d = document.documentElement;
      const offset = d.scrollTop + window.innerHeight;
      const height = d.offsetHeight;
      return setIsBottom(offset >= height - 1);
    };
    window.addEventListener('scroll', checker);
    return () => window.removeEventListener('scroll', checker);
  }, []);
  return isBottom;
}
