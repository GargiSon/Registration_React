import React, {useState} from 'react';
import './Forgot.css';
import { useNavigate } from 'react-router-dom';

const ForgotPassword = () => {
    const [email, setEmail] = useState('');
    const navigate = useNavigate();

    const handleSendEmail = async (e) => {
        e.preventDefault();

        if(!email){
            alert("Please Enter your email.");
            return;
        }

        try{
            const response = await fetch('http://localhost:5000/api/forgot-password',{
                method: 'POST',
                headers: {'Content-Type':'application/json'},
                body: JSON.stringify({email}),
                credentials: 'include',
            });

            const result = await response.json();
            if(response.ok){
                alert(result.message || "Reset link sent to your email.");
                navigate('/login');
            }else{
                alert(result.message || "Error sending email.");
            }
        }catch (error){
            console.error("Forgot password error:", error);
            alert("Something went wrong. Try again.");
        }
    };

    return(
        <main className='forgot-password-container'>
            <h2>Forgot Password</h2>
            <form onSubmit={handleSendEmail} className='forgot-password-form'>
                <label>Email</label>
                <input type='email' placeholder='Enter your registered email' value={email} onChange={(e) => setEmail(e.target.value)}/>

                <div className='send-email'>
                    <button type='submit'>Send Email</button>
                </div>
            </form>
        </main>
    );
};

export default ForgotPassword;