import React, { useState, useEffect } from 'react';
import { useNavigate, useParams } from 'react-router-dom';
import './Edit.css';

const Edit = () => {
  const { id } = useParams();
  const navigate = useNavigate();

  const [countries, setCountries] = useState([]);
  const [formData, setFormData] = useState({
    username: '',
    email: '',
    mobile: '',
    address: '',
    gender: '',
    image: null,
    imageBase64: '',
    removeImage: false,
    sports: [],
    dob: '',
    country: ''
  });

  useEffect(() => {
    const fetchUser = async () => {
      try {
        const res = await fetch(`http://localhost:5000/api/users/${id}`);
        if (!res.ok) throw new Error(`HTTP error! Status: ${res.status}`);
        const data = await res.json();
        const fetchedUser = data.user;

        setFormData(prev => ({
          ...prev,
          username: fetchedUser.username || '',
          email: fetchedUser.email || '',
          mobile: fetchedUser.mobile || '',
          address: fetchedUser.address || '',
          gender: fetchedUser.gender || '',
          dob: fetchedUser.dob || '',
          country: fetchedUser.country || '',
          sports: (fetchedUser.sports || '').split(',').map(s => s.trim()).filter(Boolean),
          imageBase64: fetchedUser.imageBase64 || '',
          image: null
        }));
      } catch (error) {
        console.error("Error fetching user:", error);
      }
    };

    const fetchCountries = async () => {
      try {
        const response = await fetch('http://localhost:5000/api/countries');
        const data = await response.json();
        setCountries(data || []);
      } catch (error) {
        console.error("Failed to fetch countries:", error);
      }
    };

    fetchUser();
    fetchCountries();
  }, [id]);

  const handleChange = (e) => {
    const { name, value, type, files, checked } = e.target;

    if (type === 'file') {
      setFormData(prev => ({ ...prev, [name]: files[0] }));
    } else if (type === 'checkbox' && name === 'sports') {
      setFormData(prev => {
        const updatedSports = checked
          ? [...prev.sports, value]
          : prev.sports.filter(sport => sport !== value);
        return { ...prev, sports: updatedSports };
      });
    } else if (type === 'checkbox' && name === 'removeImage') {
      setFormData(prev => ({ ...prev, removeImage: checked }));
    } else {
      setFormData(prev => ({ ...prev, [name]: value }));
    }
  };

  const handleSubmit = async (e) => {
    e.preventDefault();

    const form = new FormData();
    form.append("username", formData.username);
    form.append("email", formData.email);
    form.append("mobile", formData.mobile);
    form.append("address", formData.address);
    form.append("gender", formData.gender);
    form.append("dob", formData.dob);
    form.append("country", formData.country);
    form.append("remove_image", formData.removeImage ? "1" : "0");

    if (formData.image instanceof File) {
      form.append("image", formData.image);
    }

    if (Array.isArray(formData.sports)) {
      formData.sports.forEach(s => form.append("sports", s));
    }

    try {
      const res = await fetch(`http://localhost:5000/api/users/${id}`, {
        method: "PUT",
        body: form
      });

      const text = await res.text();

      if (res.ok) {
        alert("User updated successfully!");
        navigate("/");
      } else {
        alert("Error updating user: " + text);
      }
    } catch (error) {
      alert("Network or server error occurred.");
    }
  };

  return (
    <div className="edit-user-container">
      <h2>Edit User</h2>
      <form onSubmit={handleSubmit} encType="multipart/form-data">
        <table>
          <tbody>
            <tr>
              <td>Username: <span className="required-star">*</span></td>
              <td><input name="username" type="text" value={formData.username} onChange={handleChange} required /></td>
            </tr>
            <tr>
              <td>Email: <span className="required-star">*</span></td>
              <td><input name="email" type="email" value={formData.email} disabled /></td>
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
                <label><input type="radio" name="gender" value="Male" checked={formData.gender === "Male"} onChange={handleChange} /> Male</label>
                <label><input type="radio" name="gender" value="Female" checked={formData.gender === "Female"} onChange={handleChange} /> Female</label>
              </td>
            </tr>
            <tr>
              <td>Current Image:</td>
              <td>
                {formData.imageBase64 ? (
                  <div>
                    <img src={`data:image/jpeg;base64,${formData.imageBase64}`} alt="Current" style={{ maxWidth: '150px', marginBottom: '10px' }} />
                  </div>
                ) : (
                  <span>No image uploaded</span>
                )}
              </td>
            </tr>
            <tr>
              <td>Upload New Image:</td>
              <td><input name="image" type="file" accept="image/*" onChange={handleChange} /></td>
            </tr>
            <tr>
              <td>Remove Image:</td>
              <td>
                <label><input type="checkbox" name="removeImage" checked={formData.removeImage} onChange={handleChange} /> Check to remove current image</label>
              </td>
            </tr>
            <tr>
              <td>Sports:</td>
              <td>
                <label><input type="checkbox" name="sports" value="basketball" checked={formData.sports.includes("basketball")} onChange={handleChange} /> Basketball</label>
                <label><input type="checkbox" name="sports" value="swimming" checked={formData.sports.includes("swimming")} onChange={handleChange} /> Swimming</label>
                <label><input type="checkbox" name="sports" value="cricket" checked={formData.sports.includes("cricket")} onChange={handleChange} /> Cricket</label>
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
                <button type="submit">Update User</button>
                <button type="button" onClick={() => navigate('/')}>Cancel</button>
              </td>
            </tr>
          </tbody>
        </table>
      </form>
    </div>
  );
};

export default Edit;
