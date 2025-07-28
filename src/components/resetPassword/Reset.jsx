import React, {useState, useEffect } from 'react';
import './Reset.css';
import { useNavigate, useLocation } from 'react-router-dom';
import axios from 'axios';

const ResetPassword = () =>{
    const [newPassword, setNewPassword] = useState('');
    const [confirmPassword, setConfirmPassword] = useState('');
    const [error, setError] = useState('');
    const [successMsg, setSuccessMsg] = useState('');
    const [loading, setLoading] = useState(false);
    const [token, setToken] = useState('');

    const navigate = useNavigate();
    const location = useLocation();

    useEffect(() => {
        const urlParams = new URLSearchParams(location.search);
        const tokenParam = urlParams.get('token');
        if (!tokenParam) {
            setError('Invalid or missing token');
        } else {
            setToken(tokenParam);
        }
    }, [location]);

    const handleSubmit = async (e) => {
        e.preventDefault();
        setError('');
        setSuccessMsg('');

        if(!newPassword || !confirmPassword){
            setError('Both fields are required.');
            return;
        }

        if (newPassword !== confirmPassword) {
            setError('Passwords do not match.');
            return;
        }

        try {
            setLoading(true);
            await axios.post(`http://localhost:5000/api/reset-password?token=${token}`,
                {password: newPassword, confirm: confirmPassword,},
                {headers: {'Content-Type': 'application/json'},
                withCredentials: true,
            });

        setSuccessMsg('Password reset successful! Redirecting to login...');
        setNewPassword('');
        setConfirmPassword('');
        setTimeout(() => navigate('/login'), 2000);
    } catch (err) {
        console.error(err);
        setError(
            err.response?.data?.error || err.response?.data?.message ||
            'Failed to reset password. Please try again.'
        );
        setNewPassword('');
        setConfirmPassword('');
        } finally {
        setLoading(false);
        }
    };

    return (
        <div className='reset-password-container'>
            <h2>Reset your Password</h2>
            <form onSubmit={handleSubmit} className='reset-password-form'>
                <input type='password' placeholder='New Password' value={newPassword} onChange={(e) => setNewPassword(e.target.value)} required disabled={loading}/>
                <input type='password' placeholder='Confirm New Password' value={confirmPassword} onChange={(e) => setConfirmPassword(e.target.value)} required disabled={loading}/>

                {error && <p className='error'>{error}</p>}
                {successMsg && <p className='success'>{successMsg}</p>}

                <button type='submit' disabled={loading}>{loading ? 'Resetting...' : 'Reset password'}</button>
            </form>
        </div>
    );
};

export default ResetPassword;