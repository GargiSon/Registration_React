import React, { useState, useEffect } from 'react';
import { useNavigate } from 'react-router-dom';
import './AddUser.css';

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
    const { name, value, type, files } = e.target;

    if (type === 'file') {
      setFormData((prev) => ({
        ...prev,
        [name]: files[0]
      }));
    } else if (type === 'checkbox') {
      const checked = e.target.checked;
      setFormData((prev) => {
        const sports = checked
          ? [...prev.sports, value]
          : prev.sports.filter((sport) => sport !== value);
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

    // File input
    if (formData.image instanceof File) {
      form.append("image", formData.image);
    }

    // Sports (assume it's a list/array in formData.sports)
    if (Array.isArray(formData.sports)) {
      formData.sports.forEach(s => form.append("sports", s));
    } else if (formData.sports) {
      form.append("sports", formData.sports);
    }

    try {
      const res = await fetch("http://localhost:5000/api/users", {
        method: "POST",
        body: form
      });

      const text = await res.text();

      if (res.ok) {
        alert("User added successfully!");
        navigate("/");
      } else {
        alert(`Error adding user: ${text}`);
      }
    } catch (error) {
      console.error("Failed to add user:", error);
      alert("Network or server error occurred.");
    }
  };

  return (
    <div className="add-user-container">
      <h2>Add New User</h2>
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
