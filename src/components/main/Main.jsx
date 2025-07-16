import './Main.css';
import React, { useEffect, useState } from 'react';
import { getPosts } from '../../api';
import { useNavigate } from 'react-router-dom';

const Main = () => {
  const [data, setData] = useState([]);
  const navigate = useNavigate();

  useEffect(() => {
    getPosts()
      .then((posts) => {
        console.log("Posts from API:", posts);
        setData(posts);
      })
      .catch((err) => console.error("Fetch failed:", err));
  }, []);

  const handleAddUser = () => {
    navigate('/add-user'); 
  };

  return (
    <main>
      <div className="top-bar">
        <h2>Users List</h2>
        <button onClick={handleAddUser} className="add-user-btn">
          + Add User
        </button>
      </div>
      
      {data.length > 0 ? (
        <table>
          <thead>
            <tr>
              <th>S. No.</th>
              <th>Name</th>
              <th>Email</th>
              <th>Mobile</th>
            </tr>
          </thead>
          <tbody>
            {data.map((user, index) => (
              <tr key={user._id || index}>
                <td data-label="S. No.">{index + 1}</td>
                <td data-label="Name">{user.username}</td>
                <td data-label="Email">{user.email}</td>
                <td data-label="Mobile">{user.mobile}</td>
              </tr>
            ))}
          </tbody>
        </table>
      ) : (
        <p>No data</p>
      )}
    </main>
  );
};

export default Main;
