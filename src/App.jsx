import './App.css';
import { BrowserRouter as Router, Routes, Route, useLocation } from 'react-router-dom';
import Header from './components/header/Header';
import Footer from './components/footer/Footer';
import Main from './components/main/Main';
import AddUser from './components/adduser/AddUser';
import Edit from './components/edituser/Edit';
import Login from './components/login/Login';
import ForgotPassword from './components/forgotPassword/Forgot';
import ResetPassword from './components/resetPassword/Reset';
import ProtectedRoute from './ProtectedRoute';
import { Navigate } from 'react-router-dom';

const Layout = () => {
  const location = useLocation();

  // Routes which do not show header and Footer
  const noLayoutRoutes = ['/login', '/forgot-password', '/reset-password'];
  const hideLayout = noLayoutRoutes.includes(location.pathname);

  return (
    <div className="App">
      {!hideLayout && <Header title="Users Listing" subtitle="subtitle2" />}
      
      <Routes>
        <Route path='/login' element={<Login/>} />
        <Route path='/forgot-password' element={<ForgotPassword/>} />
        <Route path='/reset-password' element={<ResetPassword/>}/>

        <Route element={<ProtectedRoute />}>
          <Route path="/" element={<Main />} />
          <Route path="/add-user" element={<AddUser />} />
          <Route path="/edit-user/:id" element={<Edit />} />
        </Route>
        <Route path="*" element={<Navigate to="/login" replace />} />
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
