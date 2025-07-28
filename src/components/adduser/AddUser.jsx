import React, { useState, useEffect } from 'react';
import { useNavigate } from 'react-router-dom';
import './AddUser.css';
import axios from 'axios';

const AddUser = () => {
  const navigate = useNavigate();
  const [countries, setCountries] = useState([]);
  const [formData, setFormData] = useState({
    username: '',
    password: '',
    confirm: '',
    email: '',
    mobile: '',
    address: '',
    gender: '',
    image: null,
    sports: [],
    dob: '',
    country: ''
  });

  const [error, setError] = useState('');
  const [loading, setLoading] = useState(false);

  useEffect(() => {
    const fetchCountries = async () => {
      try {
        const response = await axios.get('http://localhost:5000/api/countries',{
          withCredentials : true,
        });
        setCountries(response.data);
      } catch (error) {
        console.error("Failed to fetch countries:", error);
        setError("Failed to load countries. Please try again later.");
      }
    };
    fetchCountries();
  }, []);

  const handleChange = (e) => {
    const { name, value, type, files, checked } = e.target;
    if (type === 'file') {
      setFormData((prev) => ({
        ...prev,
        [name]: files[0] || null
      }));
    } else if (type === 'checkbox') {
      setFormData(prev => {
        let sports = [...prev.sports];
        if (checked) {
          if (!sports.includes(value)) sports.push(value);
        } else {
          sports = sports.filter(sport => sport !== value);
        }
        return { ...prev, sports };
      });
    } else {
      setFormData((prev) => ({
        ...prev,
        [name]: value
      }));
    }
  };

  const handleSubmit = async (e) => {
    e.preventDefault();
    setError('');

    if (formData.password !== formData.confirm) {
      setError("Password and Confirm Password do not match.");
      return;
    }

    const form = new FormData();
    form.append("username", formData.username);
    form.append("password", formData.password);
    form.append("confirm", formData.confirm);
    form.append("email", formData.email);
    form.append("mobile", formData.mobile);
    form.append("address", formData.address);
    form.append("gender", formData.gender);
    form.append("dob", formData.dob);
    form.append("country", formData.country);

    if (formData.image instanceof File) {
      form.append("image", formData.image);
    }

    formData.sports.forEach(s => form.append("sports", s));

    try {
      setLoading(true);
      const res = await axios.post("http://localhost:5000/api/users", form, {
        headers: {
          'Content-Type' : 'multipart/form-data'
        },
        withCredentials: true,
      });
      alert("User added successfully!");
      navigate("/");
    }catch (error) {
      const errMsg = 
        error.response?.data?.message ||
        error.response?.data ||
        error.message ||
        "Server error occurred.";
      setError(errMsg);
      console.error("Failed to add user:", error);
    } finally {
      setLoading(false);
    }
  };

  return (
    <div className="add-user-container">
      <h2>Add New User</h2>
      {error && <p className="error-msg" style={{ color: 'red' }}>{error}</p>}
      <form onSubmit={handleSubmit} encType="multipart/form-data">
        <table>
          <tbody>
            <tr>
              <td>Username: <span className="required-star">*</span></td>
              <td><input name="username" type="text" value={formData.username} onChange={handleChange} required /></td>
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
                <label><input type="radio" name="gender" value="Male" checked={formData.gender === "Male"} onChange={handleChange} required /> Male</label>
                <label><input type="radio" name="gender" value="Female" checked={formData.gender === "Female"} onChange={handleChange} /> Female</label>
              </td>
            </tr>
            <tr>
              <td>Upload Image:</td>
              <td><input name="image" type="file" accept="image/*" onChange={handleChange} /></td>
            </tr>
            <tr>
              <td>Sports:</td>
              <td>
                <label><input type="checkbox" name="sports" value="basketball" onChange={handleChange} /> Basketball</label>
                <label><input type="checkbox" name="sports" value="swimming" onChange={handleChange} /> Swimming</label>
                <label><input type="checkbox" name="sports" value="cricket" onChange={handleChange} /> Cricket</label>
              </td>
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
