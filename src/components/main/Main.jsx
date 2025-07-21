import './Main.css';
import React, { useEffect, useState } from 'react';
import { getPosts } from '../../api';
import { useNavigate } from 'react-router-dom';

const Main = () => {
  const [data, setData] = useState([]);
  const [page, setPage] = useState(1);
  const [limit] = useState(5);
  const [total, setTotal] = useState(0);
  const [sortField, setSortField] = useState('_id');
  const [sortOrder, setSortOrder] = useState('desc');
  const navigate = useNavigate();
  const totalPages = Math.ceil(total / limit);

  useEffect(() => {
    fetchUsers();
  }, [page, sortField, sortOrder]);

  const fetchUsers = async () => {
    try {
      const response = await getPosts(page, limit, sortField, sortOrder);
      console.log("Fetched response:", response);
      if (response?.users && Array.isArray(response.users)) {
        setData(response.users);
        setTotal(response.total || 0);
      } else {
        console.error('Invalid API response structure:', response);
        setData([]);
      }
    } catch (err) {
      console.error('Error Fetching users:', err);
      setData([]);
    }
  };

  const handleSort = (field) => {
    if (field === sortField) {
      setSortOrder(sortOrder === "asc" ? "desc" : "asc");
    } else {
      setSortField(field);
      setSortOrder("asc");
    }
  };

  const handleAddUser = () => navigate('/add-user');
  const handlePrev = () => page > 1 && setPage(page - 1);
  const handleNext = () => page < totalPages && setPage(page + 1);

  const getSortArrow = (field) => {
    if (sortField !== field) return '';
    return sortOrder === 'asc' ? ' ↑' : ' ↓';
  };

  const deleteUser = async (id) => {
    if (window.confirm("Are you sure you want to delete this user?")) {
      console.log("Deleting user ID:", id); 
      try {
        const response = await fetch(`http://localhost:5000/api/users/${id}`, {
          method: 'DELETE',
        });

        const text = await response.text();
        if (!response.ok) {
          alert("Delete error: " + text);
          return;
        }

        alert("User deleted successfully!");
        fetchUsers(); //After deleting add refreshing line
      } catch (error) {
        alert("Delete error: " + error.message);
      }
    }
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
        <>
          <table>
            <thead>
              <tr>
                <th onClick={() => handleSort('_id')} style={{ cursor: 'pointer' }}>
                  S. No{getSortArrow('_id')}
                </th>
                <th onClick={() => handleSort('username')} style={{ cursor: 'pointer' }}>
                  Name{getSortArrow('username')}
                </th>
                <th onClick={() => handleSort('email')} style={{ cursor: 'pointer' }}>
                  Email{getSortArrow('email')}
                </th>
                <th onClick={() => handleSort('mobile')} style={{ cursor: 'pointer' }}>
                  Mobile{getSortArrow('mobile')}
                </th>
                <th>Actions</th>
              </tr>
            </thead>
            <tbody>
              {data.map((user, index) => (
                <tr key={user._id}>
                  <td>{(page - 1) * limit + index + 1}</td>
                  <td>{user.username}</td>
                  <td>{user.email}</td>
                  <td>{user.mobile}</td>
                  <td>
                    <button className='edit' onClick={() => navigate(`/edit-user/${user.id}`)}>Edit</button>
                    <button className='delete' onClick={() => deleteUser(user.id)}>Delete</button>
                  </td>
                </tr>
              ))}
            </tbody>
          </table>

          <div className="pagination">
            <button onClick={handlePrev} disabled={page === 1}>
              Previous
            </button>
            <span>
              Page {page} of {totalPages}
            </span>
            <button onClick={handleNext} disabled={page === totalPages}>
              Next
            </button>
          </div>
        </>
      ) : (
        <p>No data</p>
      )}
    </main>
  );
};

export default Main;
