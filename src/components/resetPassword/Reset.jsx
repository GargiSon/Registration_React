import React, {useState} from 'react';
import './Reset.css';
import { useNavigate } from 'react-router-dom';

const ResetPassword = () =>{
    const [newPassword, setNewPassword] = useState('');
    const [confirmPassword, setConfirmPassword] = useState('');
    const [error, setError] = useState('');
    const [successMsg, setSuccessMsg] = useState('');
    const navigate = useNavigate();

    const handleSubmit = (e) => {
        e.preventDefault();

        if(!newPassword || !confirmPassword){
            setError('Both fields are required.');
            setSuccessMsg('');
            return;
        }

        console.log('Password reset successful:', newPassword);
        setError('');
        setSuccessMsg('Password reset Successful! regirecting to login...');
        setTimeout(() => navigate('/login'), 2000);
    };

    return (
        <div className='reset-password-container'>
            <h2>Reset your Password</h2>
            <form onSubmit={handleSubmit}className='reset-password-form'>
                <input type='password' placeholder='New Password' value={newPassword} onChange={(e) => setNewPassword(e.target.value)}/>
                <input type='password' placeholder='Confirm New Password' value={confirmPassword} onChange={(e) => setConfirmPassword(e.target.value)} />

                {error && <p className='error'>{error}</p>}
                {successMsg && <p className='success'>{successMsg}</p>}

                <button type='submit'>Reset password</button>
            </form>
        </div>
    );
};

export default ResetPassword;