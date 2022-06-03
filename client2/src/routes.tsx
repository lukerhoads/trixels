import { BrowserRouter as Router, Route } from 'react-router-dom';
import { Auction } from './pages/Auction'
import { Home } from './pages/Home'

export const Routes = () => {
  return (
    <Router>
      <Route exact path='/'>
        <Home />
      </Route>
      <Route exact path='/auction'>
        <Auction />
      </Route>
    </Router>
  );
};
