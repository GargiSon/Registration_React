import React, {useState, useEffect} from 'react';
import './Login.css';
import { useNavigate } from 'react-router-dom';
import {loginUser} from '../../api';
import { useAuth } from "../../hooks/useAuth.jsx";

const Login = () => {
    const [email, setEmail] = useState('');
    const [password, setPassword] = useState('');
    const [loginSuccess, setLoginSuccess] = useState(false);
    const [loginError, setLoginError] = useState('');
    const [loading, setLoading] = useState(false);
    const isAuthenticated = useAuth();
    const navigate = useNavigate();

     useEffect(() => {
        if (isAuthenticated === true) {
        navigate('/', { replace: true });
        }
    }, [isAuthenticated, navigate]);

    const handleLogin = async (e) => {
        e.preventDefault();

        if(!email || !password){
            setLoginError("Please fill in all fields.");
            setLoginSuccess(false);
            setTimeout(() => setLoginError(''), 3000);
            return;
        }

        try{
            setLoading(true);
            const result = await loginUser(email, password);
            localStorage.setItem("token", result.token);
            console.log("Login result:", result);
            setLoginSuccess(true);
            setTimeout(() => navigate('/'), 1500);
        }catch(error){
            console.error("Login error: ", error);
            const errorMsg = error.response?.data?.message || "Login failed. Please try again.";
            setLoginError(errorMsg);
            setLoginSuccess(false);
            setTimeout(() => setLoginError(''), 3000);
        } finally {
            setLoading(false);
        }
    };

    const handleForgotPassword = () => {
        navigate('/forgot-password');
    };

    return (
        <main className='login-container'>
            <h2>Login</h2>
            <form onSubmit={handleLogin} className='login-form'>
                <input type='email' placeholder='Enter your email' value={email} onChange={(e) => setEmail(e.target.value)} disabled = {loading} required/>
                <input type='password' placeholder='Enter your password' value={password} onChange={(e) => setPassword(e.target.value)} disabled={loading} required autoComplete="current-password"/>

                <div className='login-buttons'>
                    <button type='submit' disabled={loading}>{loading ? 'Logging in...' : 'Login'}</button>
                    <button type='button' onClick={handleForgotPassword} disabled={loading}>Forgot Password</button>
                </div>

                {loginSuccess && (<p className='login-success-message'>Login successful!</p>)}
                {loginError && <p className='login-error-message'>{loginError}</p>}
            </form>
        </main>
    );
};

export default Login;