import './App.css';
import { BrowserRouter as Router, Routes, Route, useLocation } from 'react-router-dom';
import Header from './components/header/Header';
import Footer from './components/footer/Footer';
import Main from './components/main/Main';
import AddUser from './components/adduser/AddUser';

const Layout = () => {
  const location = useLocation();
  const hideLayout = location.pathname === '/add-user';

  return (
    <div className="App">
      {!hideLayout && <Header title="Users Listing" subtitle="subtitle2" />}
      
      <Routes>
        <Route path="/" element={<Main />} />
        <Route path="/add-user" element={<AddUser />} />
      </Routes>

      {!hideLayout && <Footer note="Footer note" />}
    </div>
  );
};

function App() {
  return (
    <Router>
      <Layout />
    </Router>
  );
}

export default App;
