import React, {useState} from 'react';
import './Forgot.css';
import { useNavigate } from 'react-router-dom';
import axios from 'axios';

const ForgotPassword = () => {
    const [email, setEmail] = useState('');
    const [loading, setLoading] = useState(false);
    const [error, setError] = useState('');
    const navigate = useNavigate();

    const handleSendEmail = async (e) => {
        e.preventDefault();
        setError('');

        if(!email){
            alert("Please Enter your email.");
            return;
        }

        try{
            setLoading(true);
            const response = await axios.post('http://localhost:5000/api/forgot-password',
            {email},
            {withCredentials: true}
        );
        alert(response.message || "Reset link sent to your email.");
        navigate('/login');
        }catch (err){
            const errMsg = err.response?.data?.message || 'Error sending email. Please try again.';
        setError(errMsg);
        alert(errMsg);
        } finally {
        setLoading(false);
        }
    };

    return(
        <main className='forgot-password-container'>
            <h2>Forgot Password</h2>
            {error && <p className="error-msg" style={{ color: 'red' }}>{error}</p>}
            <form onSubmit={handleSendEmail} className='forgot-password-form'>
                <label>Email</label>
                <input type='email' placeholder='Enter your registered email' value={email} onChange={(e) => setEmail(e.target.value)} required disabled={loading}/>

                <div className='send-email'>
                    <button type="submit" disabled={loading}>{loading ? 'Sending...' : 'Send Email'}</button>
                </div>
            </form>
        </main>
    );
};

export default ForgotPassword;