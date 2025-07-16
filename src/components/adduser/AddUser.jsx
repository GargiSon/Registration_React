import React, { useState, useEffect } from 'react';
import { useNavigate } from 'react-router-dom';
import './AddUser.css';

const AddUser = () => {
  const navigate = useNavigate();
  const [countries, setCountries] = useState([]);
  const [formData, setFormData] = useState({
    name: '',
    password: '',
    confirm: '',
    email: '',
    mobile: '',
    address: '',
    gender: '',
    image: '',
    sports: '',
    dob: '',
    country: ''
  });

  useEffect(() => {
    const fetchCountries = async () => {
      try {
        const response = await fetch('http://localhost:5000/api/countries');
        const data = await response.json();
        setCountries(data);
      } catch (error) {
        console.error("Failed to fetch countries:", error);
      }
    };

    fetchCountries();
  }, []);

  const handleChange = (e) => {
    const { name, value } = e.target;
    setFormData((prev) => ({
      ...prev,
      [name]: value
    }));
  };

  const handleSubmit = async (e) => {
    e.preventDefault();

    try {
      const res = await fetch('http://localhost:5000/api/users', {
        method: 'POST',
        headers: { 'Content-Type': 'application/json' },
        body: JSON.stringify(formData)
      });

      if (res.ok) {
        alert('User added successfully!');
        navigate('/');
      } else {
        alert('Error adding user');
      }
    } catch (error) {
      console.error('Failed to add user:', error);
    }
  };

  return (
    <div className="add-user-container">
      <h2>Add New User</h2>
      <form onSubmit={handleSubmit}>
        <table>
          <tbody>
            <tr>
              <td>Name: <span className="required-star">*</span></td>
              <td><input name="name" type="text" value={formData.name} onChange={handleChange} required /></td>
            </tr>
            <tr>
              <td>Password: <span className="required-star">*</span></td>
              <td><input name="password" type="password" value={formData.password} onChange={handleChange} required /></td>
            </tr>
            <tr>
              <td>Confirm Password: <span className="required-star">*</span></td>
              <td><input name="confirm" type="password" value={formData.confirm} onChange={handleChange} required /></td>
            </tr>
            <tr>
              <td>Email: <span className="required-star">*</span></td>
              <td><input name="email" type="email" value={formData.email} onChange={handleChange} required /></td>
            </tr>
            <tr>
              <td>Mobile: <span className="required-star">*</span></td>
              <td><input name="mobile" type="tel" value={formData.mobile} onChange={handleChange} required /></td>
            </tr>
            <tr>
              <td>Address: <span className="required-star">*</span></td>
              <td><textarea name="address" rows="3" value={formData.address} onChange={handleChange} required /></td>
            </tr>
            <tr>
              <td>Gender: <span className="required-star">*</span></td>
              <td>
                <select name="gender" value={formData.gender} onChange={handleChange} required>
                  <option value="">Select</option>
                  <option value="Male">Male</option>
                  <option value="Female">Female</option>
                  <option value="Other">Other</option>
                </select>
              </td>
            </tr>
            <tr>
              <td>Upload Image:</td>
              <td><input name="image" type="text" value={formData.image} onChange={handleChange}/></td>
            </tr>
            <tr>
              <td>Sports:</td>
              <td><input name="sports" type="text" value={formData.sports} onChange={handleChange} /></td>
            </tr>
            <tr>
              <td>Date of Birth: <span className="required-star">*</span></td>
              <td><input name="dob" type="date" value={formData.dob} onChange={handleChange} required /></td>
            </tr>
            <tr>
              <td>Country: <span className="required-star">*</span></td>
              <td>
                <select name="country" value={formData.country} onChange={handleChange} required>
                  <option value="">Select your country</option>
                  {countries.map((country, idx) => (
                    <option key={idx} value={country}>{country}</option>
                  ))}
                </select>
              </td>
            </tr>
            <tr>
              <td colSpan="2" style={{ textAlign: 'center' }}>
                <button type="submit">Add User</button>
                <button type="button" onClick={() => navigate('/')}>Cancel</button>
              </td>
            </tr>
          </tbody>
        </table>
      </form>
    </div>
  );
};

export default AddUser;
