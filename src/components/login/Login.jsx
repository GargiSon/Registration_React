import React, {useState} from 'react';
import './Login.css';
import { useNavigate } from 'react-router-dom';
import {loginUser} from '../../api'

const Login = () => {
    const [email, setEmail] = useState('');
    const [password, setPassword] = useState('');
    const navigate = useNavigate();

    const handleLogin = async (e) => {
        e.preventDefault();

        if(!email || !password){
            alert("Please fill in all fields.");
            return;
        }

        try{
            const result = await loginUser(email, password);
                console.log("Login result:", result);
                alert("Login Successful!");
                navigate('/');
            }catch(error){
                console.error("Login error: ", error);
                const errorMsg = error.response?.data?.message || error.message || "Login Failed. Try Again.";
                alert(errorMsg);
        }
    };

    const handleForgotPassword = () => {
        navigate('/forgot-password');
    };

    return (
        <main className='login-container'>
            <h2>Login</h2>
            <form onSubmit={handleLogin} className='login-form'>
                <input type='email' placeholder='Enter your email' value={email} onChange={(e) => setEmail(e.target.value)}/>
                <input type='password' placeholder='Enter your password' value={password} onChange={(e) => setPassword(e.target.value)}/>

                <div className='login-buttons'>
                    <button type='submit'>Login</button>
                    <button type='button' onClick={handleForgotPassword}>Forgot Password</button>
                </div>
            </form>
        </main>
    );
};

export default Login;