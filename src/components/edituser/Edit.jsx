import React, { useState, useEffect } from 'react';
import { useNavigate, useParams } from 'react-router-dom';
import './Edit.css';
import axios from 'axios';

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

  const [loading, setLoading] = useState(false);
  const [error, setError] = useState('');

  useEffect(() => {
    const fetchUser = async () => {
      try {
        const res = await axios.get(`http://localhost:5000/api/users/${id}`,{
          withCredentials: true,
        });

        const fetchedUser = res.data.user;

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
          image: null,
          removeImage: false,
        }));
        setError('');
      } catch (error) {
        console.error("Error fetching user:", error);
        setError('Failed to load user data. Please try again later.');
      } finally {
        setLoading(false);
      }
    };

    const fetchCountries = async () => {
      try {
        const response = await axios.get('http://localhost:5000/api/countries',{
          withCredentials : true,
        });
        setCountries(response.data || []);
      } catch (error) {
        console.error("Failed to fetch countries:", error);
        setError('Failed to load countries. Please try again later.');
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
      if (checked) {
        setFormData(prev => ({ ...prev, image: null }));
      }
    } else {
      setFormData(prev => ({ ...prev, [name]: value }));
    }
  };

  const handleSubmit = async (e) => {
    e.preventDefault();
    setError('');
    setLoading(true);

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
      const res = await axios.put(`http://localhost:5000/api/users/${id}`, form, {
        headers: {
          'Content-Type' : 'multipart/form-data',
        },
        withCredentials: true,
      });

      alert("User updated successfully!");
      navigate("/");
    } catch (error) {
      const errMsg =
        error.response?.data?.message ||
        error.response?.data ||
        error.message ||
        'Failed to update user.';
      setError(errMsg);
      console.error("Error updating user:", error);
    } finally {
      setLoading(false);
    }
  };

  return (
    <div className="edit-user-container">
      <h2>Edit User</h2>
      {error && <p className="error-msg" style={{ color: 'red' }}>{error}</p>}
      {loading && <p>Loading...</p>}
      <form onSubmit={handleSubmit} encType="multipart/form-data">
        <table>
          <tbody>
            <tr>
              <td>Username: <span className="required-star">*</span></td>
              <td><input name="username" type="text" value={formData.username} onChange={handleChange} required disabled={loading}/></td>
            </tr>
            <tr>
              <td>Email: <span className="required-star">*</span></td>
              <td><input name="email" type="email" value={formData.email} disabled /></td>
            </tr>
            <tr>
              <td>Mobile: <span className="required-star">*</span></td>
              <td><input name="mobile" type="tel" value={formData.mobile} onChange={handleChange} required disabled={loading}/></td>
            </tr>
            <tr>
              <td>Address: <span className="required-star">*</span></td>
              <td><textarea name="address" rows="3" value={formData.address} onChange={handleChange} required disabled={loading}/></td>
            </tr>
            <tr>
              <td>Gender: <span className="required-star">*</span></td>
              <td>
                <label><input type="radio" name="gender" value="Male" checked={formData.gender === "Male"} onChange={handleChange} disabled={loading}/> Male</label>
                <label><input type="radio" name="gender" value="Female" checked={formData.gender === "Female"} onChange={handleChange} disabled={loading}/> Female</label>
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
              <td><input name="image" type="file" accept="image/*" onChange={handleChange} disabled={loading || formData.removeImage}/></td>
            </tr>
            <tr>
              <td>Remove Image:</td>
              <td>
                <label><input type="checkbox" name="removeImage" checked={formData.removeImage} onChange={handleChange} disabled={loading}/> Check to remove current image</label>
              </td>
            </tr>
            <tr>
              <td>Sports:</td>
              <td>
                <label><input type="checkbox" name="sports" value="basketball" checked={formData.sports.includes("basketball")} onChange={handleChange} disabled={loading}/> Basketball</label>
                <label><input type="checkbox" name="sports" value="swimming" checked={formData.sports.includes("swimming")} onChange={handleChange} disabled={loading}/> Swimming</label>
                <label><input type="checkbox" name="sports" value="cricket" checked={formData.sports.includes("cricket")} onChange={handleChange} disabled={loading}/> Cricket</label>
              </td>
            </tr>
            <tr>
              <td>Date of Birth: <span className="required-star">*</span></td>
              <td><input name="dob" type="date" value={formData.dob} onChange={handleChange} required disabled={loading}/></td>
            </tr>
            <tr>
              <td>Country: <span className="required-star">*</span></td>
              <td>
                <select name="country" value={formData.country} onChange={handleChange} required disabled={loading}>
                  <option value="">Select your country</option>
                  {countries.map((country, idx) => (
                    <option key={idx} value={country}>{country}</option>
                  ))}
                </select>
              </td>
            </tr>
            <tr>
              <td colSpan="2" style={{ textAlign: 'center' }}>
                <button type="submit" disabled={loading}>{loading ? 'Updating...' : 'Update User'}</button>
                <button type="button" onClick={() => navigate('/')} disabled={loading}>Cancel</button>
              </td>
            </tr>
          </tbody>
        </table>
      </form>
    </div>
  );
};

export default Edit;
