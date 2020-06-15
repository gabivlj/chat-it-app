import React from 'react';
import './App.css';
import Routes from './components/routes';
import { useQuery } from '@apollo/react-hooks';
import { isLogged } from './queries/types/isLogged';
import { IS_LOGED_QUERY } from './queries/user_queries';

function App() {
  const { loading, data } = useQuery<isLogged>(IS_LOGED_QUERY);
  return (
    <div className="App">
      <Routes loged={data} loading={loading} />
    </div>
  );
}

export default App;
