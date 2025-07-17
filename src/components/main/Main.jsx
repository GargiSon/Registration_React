import './Main.css';
import React, { useEffect, useState } from 'react';
import { getPosts } from '../../api';
import { useNavigate } from 'react-router-dom';

const Main = () => {
  const [data, setData] = useState([]);
  const [page, setPage] = useState(1);
  const [limit] = useState(5);
  const [total, setTotal] = useState(0);
  const navigate = useNavigate();

  const totalPages = Math.ceil(total / limit);

  useEffect(() => {
    fetchUsers(page);
  }, [page]);

  const fetchUsers = async (page) => {
    try {
      const response = await getPosts(page, limit);
      if (response && response.users && Array.isArray(response.users)) {
        console.log("Fetched Users:", response.users);
        setData(response.users);
        setTotal(response.total || response.users.length);
      } else {
        console.error("Invalid API response structure:", response);
        setData([]);
      }
    } catch (err) {
      console.error("Error Fetching users:", err);
      setData([]);
    }
  };

  const handleAddUser = () => navigate('/add-user');

  const handlePrev = () => page > 1 && setPage(page - 1);
  const handleNext = () => page < totalPages && setPage(page + 1);

  return (
    <main>
      <div className="top-bar">
        <h2>Users List</h2>
        <button onClick={handleAddUser} className="add-user-btn">
          + Add User
        </button>
      </div>
      
      {data.length > 0 ? (
        <>
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
                  <td>{(page - 1) * limit + index + 1}</td>
                  <td>{user.username}</td>
                  <td>{user.email}</td>
                  <td>{user.mobile}</td>
                </tr>
              ))}
            </tbody>
          </table>

          <div className="pagination">
            <button onClick={handlePrev} disabled={page === 1}>Previous</button>
            <span>Page {page} of {totalPages}</span>
            <button onClick={handleNext} disabled={page === totalPages}>Next</button>
          </div>
        </>
      ) : (
        <p>No data</p>
      )}
    </main>
  );
};

export default Main;
